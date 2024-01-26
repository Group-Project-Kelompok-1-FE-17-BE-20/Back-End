package data

import (
	"Laptop/app/database"
	"Laptop/features/user"
	"Laptop/utils/responses"
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: database,
	}
}

// Login implements user.UserDataInterface.
func (r *userQuery) Login(email string) (user.CoreUser, error) {
	var dataUser database.User
	tx := r.db.Where("email = ?", email).First(&dataUser)
	if tx.Error != nil {
		log.Error("Database error:", tx.Error)
		return user.CoreUser{}, errors.New(tx.Error.Error() + ", invalid email")
	}
	if tx.RowsAffected == 0 {
		return user.CoreUser{}, errors.New("login failed, invalid email")
	}

	dataCore := MapUserToCoreUser(dataUser)
	return dataCore, nil
}

// Insert implements user.UserDataInterface.
func (r *userQuery) Insert(inputUser user.CoreUser) (uint, error) {
	NewUser := MapCoreUsertoUser(inputUser)
	NewUser.Password = responses.HashPassword(inputUser.Password)

	//simpan ke db
	tx := r.db.Create(&NewUser)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}
	return NewUser.ID, nil
}

// SelectAll implements user.UserDataInterface.
func (r *userQuery) SelectAll() ([]user.CoreUser, error) {
	var userData []database.User
	tx := r.db.Find(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}

	//mapping from userData -> CoreUser
	coreUserSlice := ListMapUserToCoreUser(userData)
	return coreUserSlice, nil
}

// SelectByID implements user.UserDataInterface.
func (r *userQuery) Select(userId uint) (user.CoreUser, error) {
	var userData database.User
	tx := r.db.First(&userData, userId)
	if tx.Error != nil {
		return user.CoreUser{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.CoreUser{}, errors.New("data not found")
	}
	//Mapping User to CoreUser
	coreUser := MapUserToCoreUser(userData)
	return coreUser, nil
}

// // Update implements user.UserDataInterface.
func (r *userQuery) Update(userId uint, userData user.CoreUser) error {
	var user database.User
	tx := r.db.First(&user, userId)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	if userData.Password != "" {
		userData.Password = responses.HashPassword(userData.Password)

	}
	//Mapping User to CoreUser
	updatedUser := MapCoreUsertoUser(userData)
	tx = r.db.Model(&user).Updates(updatedUser)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update user")
	}
	return nil

}

// Delete implements user.UserDataInterface.
func (r *userQuery) Delete(userId uint) error {
	var user database.User
	tx := r.db.Delete(&user, userId)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

// /-----------------------------
func (r *userQuery) Photo(c echo.Context) *uploader.UploadResult {
	urlCloudinary := "cloudinary://979172954987629:PNgbXcjMn-VOd1AyTlN0yBSvnWU@dv3nso14b"
	fileHeader, _ := c.FormFile("image_profil")

	var user database.User
	_ = c.Bind(&user)

	file, _ := fileHeader.Open()
	//log.Println(fileHeader.Filename)

	ctx := context.Background()
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	//log.Println(resp.SecureURL)

	return resp
}

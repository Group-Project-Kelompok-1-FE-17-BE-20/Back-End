package data

import (
	"Laptop/app/database"
	"Laptop/features/admin"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type adminQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) admin.AdminDataInterface {
	return &adminQuery{
		db: database,
	}
}

// Login implements user.UserDataInterface.
func (r *adminQuery) Login(email string) (admin.CoreAdmin, error) {
	var dataAdmin database.Admin
	tx := r.db.Where("email = ?", email).First(&dataAdmin)
	if tx.Error != nil {
		log.Error("Database error:", tx.Error)
		return admin.CoreAdmin{}, errors.New(tx.Error.Error() + ", invalid email")
	}
	if tx.RowsAffected == 0 {
		return admin.CoreAdmin{}, errors.New("login failed, invalid email")
	}

	dataCore := MapAdminToCoreAdmin(dataAdmin)
	return dataCore, nil
}

// Insert implements user.UserDataInterface.

// func (r *userQuery) Insert(inputUser user.CoreUser) (uint, error) {
// 	NewUser := MapCoreUsertoUser(inputUser)
// 	NewUser.Password = responses.HashPassword(inputUser.Password)

// 	//simpan ke db
// 	tx := r.db.Create(&NewUser)
// 	if tx.Error != nil {
// 		return 0, tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return 0, errors.New("data not found")
// 	}
// 	return NewUser.ID, nil
// }

// SelectAll implements user.UserDataInterface.
func (r *adminQuery) SelectAll() ([]admin.CoreAdmin, error) {
	var adminData []database.Admin
	tx := r.db.Find(&adminData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}

	//mapping from userData -> CoreUser
	coreUserSlice := ListMapAdminToCoreAdmin(adminData)
	return coreUserSlice, nil
}

// SelectByID implements user.UserDataInterface.
// func (r *userQuery) Select(userId uint) (user.CoreUser, error) {
// 	var userData User
// 	tx := r.db.First(&userData, userId)
// 	if tx.Error != nil {
// 		return user.CoreUser{}, tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return user.CoreUser{}, errors.New("data not found")
// 	}
// 	//Mapping User to CoreUser
// 	coreUser := MapUserToCoreUser(userData)
// 	return coreUser, nil
// }

// // Update implements user.UserDataInterface.
// func (r *userQuery) Update(userId uint, userData user.CoreUser) error {
// 	var user User
// 	tx := r.db.First(&user, userId)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New("data not found")
// 	}
// 	if userData.Password != "" {
// 		userData.Password = responses.HashPassword(userData.Password)

// 	}
// 	//Mapping User to CoreUser
// 	updatedUser := MapCoreUsertoUser(userData)
// 	tx = r.db.Model(&user).Updates(updatedUser)
// 	if tx.Error != nil {
// 		return errors.New(tx.Error.Error() + " failed to update user")
// 	}
// 	return nil

// }

// // Delete implements user.UserDataInterface.
// func (r *userQuery) Delete(userId uint) error {
// 	var user User
// 	tx := r.db.Delete(&user, userId)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New("data not found")
// 	}
// 	return nil
// }

// ---------------------------------------------------------------------------------------------------------------------------------------------

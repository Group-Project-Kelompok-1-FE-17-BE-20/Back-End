package service

import (
	"Laptop/app/middlewares"
	"Laptop/features/user"
	"Laptop/utils/responses"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserService struct {
	userRepo user.UserDataInterface
	validate *validator.Validate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &UserService{
		userRepo: repo,
		validate: validator.New(),
	}
}

// Login implements user.UserServiceInterface.
func (s *UserService) Login(email string, password string) (user.CoreUser, string, error) {
	if email == "" {
		return user.CoreUser{}, "", errors.New("email is required")
	} else if password == "" {
		return user.CoreUser{}, "", errors.New("password is required")
	}

	dataLogin, err := s.userRepo.Login(email)
	if err != nil {
		log.Error("Service error:", err)
		return user.CoreUser{}, "", err
	}
	checkPassword := responses.ComparePassword(password, dataLogin.Password)
	if !checkPassword {
		return user.CoreUser{}, "", errors.New("login failed, wrong password")
	}
	token, err := middlewares.CreateToken(dataLogin.ID)
	if err != nil {
		log.Error("Create token error:", err)
		return user.CoreUser{}, "", err
	}
	return dataLogin, token, nil
}

// CreateUser implements user.UserServiceInterface.
func (s *UserService) Create(inputUser user.CoreUser) (uint, error) {
	errValidate := s.validate.Struct(inputUser)
	if errValidate != nil {
		return 0, errors.New("validation error, " + errValidate.Error())
	}
	userID, err := s.userRepo.Insert(inputUser)
	if err != nil {
		return 0, err
	}
	return userID, nil

}

// GetAllUser implements user.UserServiceInterface.
func (s *UserService) GetAll() ([]user.CoreUser, error) {
	result, err := s.userRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserById implements user.UserServiceInterface.
func (s *UserService) GetById(userId uint) (user.CoreUser, error) {
	result, err := s.userRepo.Select(userId)
	if err != nil {
		return user.CoreUser{}, err
	}
	return result, nil
}

// UpdateUserById implements user.UserServiceInterface.
func (s *UserService) UpdateById(userId uint, userData user.CoreUser) error {
	err := s.userRepo.Update(userId, userData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserById implements user.UserServiceInterface.
func (s *UserService) DeleteById(userId uint) error {
	err := s.userRepo.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserService) Photo(c echo.Context) *uploader.UploadResult {
	res := s.userRepo.Photo(c)
	return res
}

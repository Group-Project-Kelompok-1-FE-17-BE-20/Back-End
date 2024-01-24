package service

import (
	//"Laptop/app/middlewares"
	"Laptop/features/admin"
	//"Laptop/utils/responses"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type AdminService struct {
	adminRepo admin.AdminDataInterface
	validate  *validator.Validate
}

func New(repo admin.AdminDataInterface) admin.AdminServiceInterface {
	return &AdminService{
		adminRepo: repo,
		validate:  validator.New(),
	}
}

// Login implements user.UserServiceInterface.
func (s *AdminService) Login(email string, password string) (admin.CoreAdmin, string, error) {
	if email == "" {
		return admin.CoreAdmin{}, "", errors.New("email is required")
	} else if password == "" {
		return admin.CoreAdmin{}, "", errors.New("password is required")
	}

	dataLogin, err := s.adminRepo.Login(email)
	if err != nil {
		log.Error("Service error:", err)
		return admin.CoreAdmin{}, "", err
	}
	// checkPassword := responses.ComparePassword(password, dataLogin.Password)
	// if !checkPassword {
	// 	return admin.CoreAdmin{}, "", errors.New("login failed, wrong password")
	// }
	// token, err := middlewares.CreateToken(dataLogin.ID)
	// if err != nil {
	// 	log.Error("Create token error:", err)
	// 	return admin.CoreAdmin{}, "", err
	// }
	return dataLogin, "", nil // return dataLogin, token, nil
}

// // CreateUser implements user.UserServiceInterface.
// func (s *UserService) Create(inputUser user.CoreUser) (uint, error) {
// 	errValidate := s.validate.Struct(inputUser)
// 	if errValidate != nil {
// 		return 0, errors.New("validation error, " + errValidate.Error())
// 	}
// 	userID, err := s.userRepo.Insert(inputUser)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return userID, nil

// }

// GetAllUser implements user.UserServiceInterface.
func (s *AdminService) GetAll() ([]admin.CoreAdmin, error) {
	result, err := s.adminRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserById implements user.UserServiceInterface.
// func (s *UserService) GetById(userId uint) (user.CoreUser, error) {
// 	result, err := s.userRepo.Select(userId)
// 	if err != nil {
// 		return user.CoreUser{}, err
// 	}
// 	return result, nil
// }

// // UpdateUserById implements user.UserServiceInterface.
// func (s *UserService) UpdateById(userId uint, userData user.CoreUser) error {
// 	err := s.userRepo.Update(userId, userData)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // DeleteUserById implements user.UserServiceInterface.
// func (s *UserService) DeleteById(userId uint) error {
// 	err := s.userRepo.Delete(userId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//--------------------------------------------------------------------------------------------------

package data

import (
	//"Laptop/features/store"

	"Laptop/app/database"
	"Laptop/features/user"

	"gorm.io/gorm"
)

// mapping coreUser to User
func MapCoreUsertoUser(coreUser1 user.CoreUser) database.User {

	var user = &database.User{
		Username:     coreUser1.Username,
		NamaLengkap:  coreUser1.NamaLengkap,
		Email:        coreUser1.Email,
		Password:     coreUser1.Password,
		NomorHP:      coreUser1.NomorHP,
		Alamat:       coreUser1.Alamat,
		JenisKelamin: coreUser1.JenisKelamin,
		ImageProfil:  coreUser1.ImageProfil,
		Model:        gorm.Model{ID: uint(coreUser1.ID)},
	}
	return *user

}

// Mapping User to CoreUser
func MapUserToCoreUser(user1 database.User) user.CoreUser {

	coreUser := &user.CoreUser{
		ID:           user1.ID,
		Username:     user1.Username,
		NamaLengkap:  user1.NamaLengkap,
		Email:        user1.Email,
		Password:     user1.Password,
		NomorHP:      user1.NomorHP,
		Alamat:       user1.Alamat,
		JenisKelamin: user1.JenisKelamin,
		ImageProfil:  user1.ImageProfil,
		CreatedAt:    user1.CreatedAt,
		UpdatedAt:    user1.UpdatedAt,
		DeletedAt:    user1.DeletedAt.Time,
	}
	return *coreUser

}

// Mapping dari []User ke []CoreUser

func ListMapUserToCoreUser(users []database.User) []user.CoreUser {
	coreUsers := make([]user.CoreUser, len(users))

	for i, users := range users {
		coreUser := &user.CoreUser{
			ID:           users.ID,
			Username:     users.Username,
			NamaLengkap:  users.NamaLengkap,
			Email:        users.Email,
			Password:     users.Password,
			NomorHP:      users.NomorHP,
			Alamat:       users.Alamat,
			JenisKelamin: users.JenisKelamin,
			ImageProfil:  users.ImageProfil,
		}
		coreUsers[i] = *coreUser
	}

	return coreUsers
}

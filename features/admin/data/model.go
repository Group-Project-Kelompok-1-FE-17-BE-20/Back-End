package data

import (
	"Laptop/app/database"
	"Laptop/features/admin"
	// "Laptop/controllers/user"
)

// mapping coreUser to User
func MapCoreAdmintoAdmin(coreAdmin admin.CoreAdmin) database.Admin {
	// var users []_userData.User
	// for _, core := range coreAdmin.Users {
	// 	users = append(users, data.MapCoreUserToUser(core))
	// }
	return database.Admin{
		// var user = &User{
		// Username:     coreUser.Username,
		// NamaLengkap: coreUser.NamaLengkap,
		Email:    coreAdmin.Email,
		Password: coreAdmin.Password,

		// NomorHP:      coreUser.NomorHP,
		// Alamat:       coreUser.Alamat,
		// JenisKelamin: coreUser.JenisKelamin,
		// ImageProfil: coreUser.ImageProfil,
		// Stores:       stores,

	}
	//return *user
	// }
}

// Mapping User to CoreUser
func MapAdminToCoreAdmin(user1 database.Admin) admin.CoreAdmin {
	// var coreAdmin []user.CoreUser
	// for _, UserModel := range user1.Users {
	// 	coreAdmin = append(coreAdmin, data.MapUserToCoreUser(UserModel))
	// }
	return admin.CoreAdmin{

		//coreUser := &user.CoreUser{
		ID: user1.ID,
		// Username:     user1.Username,
		// NamaLengkap:  user1.NamaLengkap,
		Email:    user1.Email,
		Password: user1.Password,
		// NomorHP:      user1.NomorHP,
		// Alamat:       user1.Alamat,
		// JenisKelamin: user1.JenisKelamin,
		// ImageProfil:  user1.ImageProfil,
		// Stores: coreStore,
	}
}

func ListMapAdminToCoreAdmin(admins []database.Admin) []admin.CoreAdmin {
	var coreAdmins []admin.CoreAdmin
	for _, adminModel := range admins {
		coreAdmins = append(coreAdmins, MapAdminToCoreAdmin(adminModel))
	}
	return coreAdmins

	// coreUsers := make([]user.CoreUser, len(users))

	// for i, users := range users {
	// 	coreUser := &user.CoreUser{
	// 		Username:     users.Username,
	// 		NamaLengkap:  users.NamaLengkap,
	// 		Email:        users.Email,
	// 		Password:     users.Password,
	// 		NomorHP:      users.NomorHP,
	// 		Alamat:       users.Alamat,
	// 		JenisKelamin: users.JenisKelamin,
	// 		ImageProfil:  users.ImageProfil,
	// 		// NamaToko:     users.NamaToko,
	// 		// ImageToko:    users.ImageToko,
	// 		// AlamatToko:   users.AlamatToko,
	// 		// Lakukan mapping atribut lainnya sesuai kebutuhan
	// 	}
	// 	coreUsers[i] = *coreUser
	// }

	// return coreUsers
}

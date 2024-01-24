package admin

import (
	"Laptop/features/user"
	"time"

	"gorm.io/gorm"
)

type CoreAdmin struct {
	gorm.Model
	ID       uint
	Email    string          `json:"email" form:"email" validate:"required" gorm:"unique"`
	Password string          `json:"password" form:"Password" validate:"required"`
	Users    []user.CoreUser `gorm:"foreignKey:AdminID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type AdminDataInterface interface {
	Login(email string) (CoreAdmin, error)
	// Insert(inputUser CoreAdmin) (uint, error)
	SelectAll() ([]CoreAdmin, error)
	// Select(userId uint) (CoreUser, error)
	// Update(userId uint, userData CoreUser) error
	// Delete(userId uint) error
}

type AdminServiceInterface interface {
	Login(email string, password string) (CoreAdmin, string, error)
	// Create(inputUser CoreAdmin) (uint, error)
	GetAll() ([]CoreAdmin, error)
	// GetById(userId uint) (CoreUser, error)
	// UpdateById(userId uint, userData CoreUser) error
	// DeleteById(userId uint) error
}

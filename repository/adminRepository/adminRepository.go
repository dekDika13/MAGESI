package adminRepository

import (
	"errors"
	adminDto "magesi/dto/admin"
	"magesi/model"
	"time"

	"gorm.io/gorm"
)

type AdminRepository interface {
	// TODO AUTH
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (model.Users, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO ADMIN REPOSITORY HERE

// TODO LOGIN ADMIN
func (u *adminRepository) LoginAdmin(payloads adminDto.LoginDTO) (model.Users, error) {
	var admin model.Users

	query := u.db.Where("username = ?", payloads.Username).First(&admin)
	if query.Error != nil {
		return admin, query.Error
	}

	if query.RowsAffected < 1 {
		return admin, errors.New("username is incorrect")
	}

	return admin, nil
}

// TODO REGISTER ADMIN
func (u *adminRepository) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {
	var user model.Users

	checkUser := u.db.Where("username = ?", payloads.Username).First(&user)

	if checkUser.Error == nil && user.Username == payloads.Username {
		return payloads, errors.New("username used")
	} else if err := u.db.Create(&model.Users{
		RoleId:    payloads.RoleId,
		Username:  payloads.Username,
		Password:  payloads.Password,
		CreatedAT: time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}

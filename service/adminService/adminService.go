package adminService

import (
	"errors"
	adminDto "magesi/dto/admin"
	"magesi/middleware"
	"magesi/repository/adminRepository"
	"magesi/utils"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	// TODO AUTH
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error)
}

type adminService struct {
	adminRepository adminRepository.AdminRepository
}

func NewAdminService(adminRepo adminRepository.AdminRepository) *adminService {
	return &adminService{
		adminRepository: adminRepo,
	}
}

// TODO ADMIN SERVICE HERE

// TODO REGISTER ADMIN
func (s *adminService) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	// encrypt password
	pw, err := utils.HashBcrypt(payloads.Password)

	if err != nil {
		return payloads, err
	}

	payloads.Password = pw

	res, err := s.adminRepository.RegisterAdmin(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

// TODO LOGIN ADMIN
func (s *adminService) LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error) {
	var temp adminDto.LoginJWT

	res, err := s.adminRepository.LoginAdmin(payloads)

	if errh := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(payloads.Password)); errh != nil {
		return temp, errors.New("username or password incorrect")
	}

	token, errt := middleware.CreateToken(res.UserID, res.RoleId, res.Username)

	temp = adminDto.LoginJWT{
		Username: res.Username,
		Token:    token,
	}

	if err != nil {
		return temp, err
	}

	if errt != nil {
		return temp, errt
	}

	return temp, nil
}

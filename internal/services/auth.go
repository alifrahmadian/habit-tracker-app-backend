package services

import (
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/models"
	r "github.com/alifrahmadian/habit-tracker-app-backend/internal/repositories"
	e "github.com/alifrahmadian/habit-tracker-app-backend/pkg/errors"
	"github.com/alifrahmadian/habit-tracker-app-backend/pkg/utils"
)

type AuthService interface {
	Register(user *models.User) (*models.User, error)
	Login(identity, password string) (*models.User, error)
}

type authService struct {
	UserRepo r.UserRepository
}

func NewAuthService(userRepo r.UserRepository) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (s *authService) Register(user *models.User) (*models.User, error) {
	checkUsername, err := s.UserRepo.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	if checkUsername != nil {
		return nil, e.ErrUsernameAlreadyExist
	}

	checkEmail, err := s.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if checkEmail != nil {
		return nil, e.ErrEmailAlreadyExist
	}

	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userModel := &models.User{
		RoleId:    user.RoleId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	usr, err := s.UserRepo.CreateUser(userModel)
	if err != nil {
		return nil, err
	}

	return usr, nil

}

func (s *authService) Login(identity, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByIdentity(identity)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePassword(password, user.Password) {
		return nil, e.ErrUserCredentialsInvalid
	}

	return user, nil
}

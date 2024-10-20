package services

import (
	"errors"
	"time"

	"github.com/nianfouyi/video-user-api/internal/models"
	"github.com/nianfouyi/video-user-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(username, password, securityQuestion, securityAnswer string) error
	Login(username, password string) (*models.User, error)
	UpdateUser(id uint, gender, hobbies string) error
	DeleteUser(id uint) error
	ChangePassword(username, securityAnswer, newPassword string) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	ResetPassword(username, password, newPassword string) error
}

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, password, securityQuestion, securityAnswer string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:         username,
		Password:         string(hashedPassword),
		SecurityQuestion: securityQuestion,
		SecurityAnswer:   securityAnswer,
		RegisterDate:     time.Now(),
	}

	return s.repo.Create(user)
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *UserService) UpdateUser(id uint, gender, hobbies string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.Gender = gender
	user.Hobbies = hobbies

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) ResetPassword(username, password, newPassword string) error {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return err
	}
	oldHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if user.Password != string(oldHashedPassword) {
		return errors.New("the old password you entered is incorrect")
	}
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(newHashedPassword)
	return s.repo.Update(user)
}
func (s *UserService) ChangePassword(username, securityAnswer, newPassword string) error {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return err
	}

	if user.SecurityAnswer != securityAnswer {
		return errors.New("invalid security answer")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.Update(user)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

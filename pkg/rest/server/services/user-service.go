package services

import (
	"gromnew/pkg/rest/server/daos"
	"gromnew/pkg/rest/server/models"
)

type UserService struct {
	userDao *daos.UserDao
}

func NewUserService() (*UserService, error) {
	userDao, err := daos.NewUserDao()
	if err != nil {
		return nil, err
	}
	return &UserService{
		userDao: userDao,
	}, nil
}

func (userService *UserService) CreateUser(user models.User) (*models.User, error) {
	return userService.userDao.CreateUser(user)
}

func (userService *UserService) UpdateUser(id int64, user models.User) (*models.User ,error) {
	return userService.userDao.UpdateUser(id, user)
}

func (userService *UserService) DeleteUser(id int64) error {
	return userService.userDao.DeleteUser(id)
}

func (userService *UserService) ListUsers() ([]models.User, error) {
	return userService.userDao.ListUsers()
}

func (userService *UserService) GetUser(id int64) (*models.User, error) {
	return userService.userDao.GetUser(id)
}

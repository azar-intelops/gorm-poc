package daos

import (
	"errors"
	"gromnew/pkg/rest/server/daos/clients/sqls"
	"gromnew/pkg/rest/server/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() (*UserDao, error) {
	sqlClient, err := sqls.InitSqlDB()
	if err != nil {
		return nil, err
	}

	err = sqlClient.DB.AutoMigrate(models.User{})
	if err != nil {
		return nil, err
	}

	return &UserDao{
		db: sqlClient.DB,
	}, nil
}

func (userDao *UserDao) CreateUser(user models.User) error {
	if err := userDao.db.Create(&user).Error; err != nil {
		return err
	}

	log.Debugf("user created")
	return nil
}

func (userDao *UserDao) ListUsers() ([]models.User, error) {
	var all []models.User
	if err := userDao.db.Find(&all).Error; err != nil {
		return nil, err
	}

	log.Debugf("user listed")
	return all, nil
}

func (userDao *UserDao) UpdateUser(id int64, user models.User) error {
	if id != user.Id || id == 0 {
		return errors.New("id and payload don't match or id can't be 0")
	}

	var m models.User
	if err := userDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		return err
	}
	if id == m.Id {
		userDao.db.Save(&user)
		log.Debugf("user updated")
		return nil
	}
	return errors.New("internal server error")
}

func (userDao *UserDao) DeleteUser(id int64) error {
	var m models.User

	if err := userDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	log.Debugf("user deleted")
	return nil
}

func (userDao *UserDao) GetUser(id int64) (models.User, error) {
	var m models.User
	if err := userDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		return models.User{}, err
	}
	log.Debugf("user retrieved")
	return m, nil
}

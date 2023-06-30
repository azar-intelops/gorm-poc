package daos

import (
	"errors"
	"gormint/pkg/rest/server/daos/clients/sqls"
	user_client "gormint/pkg/rest/server/daos/clients/sqls/user-client"
	"gormint/pkg/rest/server/models"

	log "github.com/sirupsen/logrus"
)

type UserDao struct {
	sqlClient *sqls.SQLiteClient
}

func NewUserDao() (*UserDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = user_client.Migrate(sqlClient, models.User{})
	if err != nil {
		return nil, err
	}
	return &UserDao{
		sqlClient: sqlClient,
	}, nil
}

func (userDao *UserDao) CreateUser(user models.User) error {
	_, err := user_client.Create(userDao.sqlClient, user)
	if err != nil {
		return err
	}
	log.Debugf("user created")
	return nil
}

func (userDao *UserDao) ListUsers() ([]models.User, error) {
	users, err := user_client.All(userDao.sqlClient)
	if err != nil {
		return users, err
	}
	log.Debugf("user listed")
	return users, nil
}

func (userDao *UserDao) UpdateUser(id int64, user models.User) error {
	if id != user.Id {
		return errors.New("id and payload don't match")
	}
	_, err := user_client.Update(userDao.sqlClient, id, user)
	if err != nil {
		return err
	}
	log.Debugf("user updated")
	return nil
}

func (userDao *UserDao) DeleteUser(id int64) error {
	err := user_client.Delete(userDao.sqlClient, id)
	if err != nil {
		return err
	}
	log.Debugf("user deleted")
	return nil
}

func (userDao *UserDao) GetUser(id int64) (models.User, error) {
	user, err := user_client.Get(userDao.sqlClient, id)
	if err != nil {
		return models.User{}, err
	}
	log.Debugf("user retrieved")
	return *user, nil
}

package user_client

import (
	"errors"
	"gormint/pkg/rest/server/daos/clients/sqls"
	"gormint/pkg/rest/server/models"
)

func Migrate(r *sqls.SQLiteClient, m models.User) error {
	db := r.DB.AutoMigrate(m)
	return db.Error
}

func Create(r *sqls.SQLiteClient, m models.User) (*models.User, error) {
	if err := r.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func All(r *sqls.SQLiteClient) ([]models.User, error) {
	var all []models.User

	if err := r.DB.Find(&all).Error; err != nil {
		return nil, err
	}
	return all, nil
}

func Get(r *sqls.SQLiteClient, id int64) (*models.User, error) {
	var m models.User

	if err := r.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func Update(r *sqls.SQLiteClient, id int64, m models.User) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	var user models.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Id == m.Id {
		r.DB.Save(&m)
		return &m, nil
	}
	return nil, errors.New("internal server error")
}

func Delete(r *sqls.SQLiteClient, id int64) error {
	var m models.User

	if err := r.DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

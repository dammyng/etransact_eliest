package implementation

import (
	"eliest/models"
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



type SqlLayer struct {
	Session *gorm.DB
}

func NewSqlLayer(dsn string) *SqlLayer {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	return &SqlLayer{Session: db}
}

func (sql *SqlLayer) CreateAccount(user *models.Account) (string, error) {
	err := sql.Session.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.MSISDN , err
}

func (sql *SqlLayer) FindAccount(arg *models.Account) (*models.Account, error) {
	session := sql.Session
	var dA models.Account
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateUser(old *models.Account, new *models.Account) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}


func (sql *SqlLayer) UpdateUserMap(arg *models.Account, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}

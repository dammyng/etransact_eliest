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


func (sql *SqlLayer) CreateWinning(wins *models.Winnings) (string, error) {
	err := sql.Session.Create(&wins).Error
	if err != nil {
		return "", err
	}
	return wins.Hash , err
}

func (sql *SqlLayer) FindWinning(arg *models.Winnings) (*models.Winnings, error) {
	session := sql.Session
	var dA models.Winnings
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateWinning(old *models.Winnings, new *models.Winnings) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}


func (sql *SqlLayer) UpdateWinningMap(arg *models.Winnings, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}


func (sql *SqlLayer) CreateVoucher(wins *models.Vouchers) (string, error) {
	err := sql.Session.Create(&wins).Error
	if err != nil {
		return "", err
	}
	return wins.Hash , err
}

func (sql *SqlLayer) FindVoucher(arg *models.Vouchers) (*models.Vouchers, error) {
	session := sql.Session
	var dA models.Vouchers
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateVoucher(old *models.Vouchers, new *models.Vouchers) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}

func (sql *SqlLayer) UpdateVoucherMap(arg *models.Vouchers, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}


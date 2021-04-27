package db

import "eliest/models"

type Handler interface {
	FindAccount(*models.Account) (*models.Account ,error)
	CreateAccount(*models.Account) (string, error)
	UpdateUser(*models.Account, *models.Account) error
	UpdateUserMap(*models.Account, map[string]interface{}) error
}

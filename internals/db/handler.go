package db

import "eliest/models"

type Handler interface {
	FindAccount(*models.Account) (*models.Account, error)
	CreateAccount(*models.Account) (string, error)
	UpdateUser(*models.Account, *models.Account) error
	UpdateUserMap(*models.Account, map[string]interface{}) error

	CreateWinning(*models.Winnings) (string, error)
	FindWinning(*models.Winnings) (*models.Winnings, error)
	UpdateWinning(*models.Winnings, *models.Winnings) error
	UpdateWinningMap(*models.Winnings, map[string]interface{}) error
}

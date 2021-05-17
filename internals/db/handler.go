package db

import "eliest/models"

type Handler interface {
	FindAccount(*models.Account) (*models.Account, error)
	FindAgent(*models.Agent) (*models.Agent, error)
	CreateAccount(*models.Account) (string, error)
	UpdateUser(*models.Account, *models.Account) error
	UpdateUserMap(*models.Account, map[string]interface{}) error
	FindAgentWallet(*models.Wallet) (*models.Wallet, error)
	CreateTransaction(*models.Transaction) (string, error)

	UpdateAgent(*models.Agent, *models.Agent) error
	UpdateAgentMap(*models.Agent, map[string]interface{}) error

	CreateWinning(*models.Winnings) (string, error)
	FindWinning(*models.Winnings) (*models.Winnings, error)
	UpdateWinning(*models.Winnings, *models.Winnings) error
	UpdateWinningMap(*models.Winnings, map[string]interface{}) error

	CreateVoucher(*models.Vouchers) (string, error)
	FindVoucher(*models.Vouchers) (*models.Vouchers, error)
	UpdateVoucher(*models.Vouchers, *models.Vouchers) error
	UpdateVoucherMap(*models.Vouchers, map[string]interface{}) error

}

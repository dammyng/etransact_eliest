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

	CreateVoucher(*models.Voucher) (string, error)
	CreateVBatch(*models.VBatch) (string, error)
	FindVoucher(*models.Voucher) (*models.Voucher, error)
	UpdateVoucher(*models.Voucher, *models.Voucher) error
	UpdateVoucherMap(*models.Voucher, map[string]interface{}) error

}

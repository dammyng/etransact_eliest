package helpers

import (
	"crypto/md5"
	"eliest/models"
	"fmt"
	"io"
	"time"
)

func GererateVoucher(amt float64, gen string, ) (models.Winnings, string) {
	pin := RandInt(3)
	serial := RandInt(4)
	code := pin + serial

	return models.Winnings{
		Amount:      amt,
		Status:      "active",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		GeneratedBy: gen,
		Hash:        DemoWHasher(code, serial),
	}, code
}

func VoucherHash(validator string, serial string) models.Vouchers {
	return models.Vouchers{
		Hash: DemoWHasher(validator, serial),
	}
}

func DemoVHasher(validator, serial string) string {
	h := md5.New()
	io.WriteString(h, validator)
	io.WriteString(h, serial)
	return fmt.Sprintf("%v", h.Sum(nil))
}

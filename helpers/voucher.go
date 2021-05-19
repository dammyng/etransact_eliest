package helpers

import (
	"crypto/md5"
	"eliest/models"
	"fmt"
	"io"
	"time"
)

func GererateVoucher(amt float64, gen , batch string ) (models.Voucher) {
	pin := RandInt(3)
	serial := RandInt(4)
	code := pin + serial

	return models.Voucher{
		Amount:      amt,
		Status:      "active",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Code:   code,
		Batch:   batch,
		GeneratedBy: gen,
		Hash:        DemoVHasher(code, serial),
	}
}

func VoucherHash(validator, serial string) models.Voucher {
	return models.Voucher{
		Hash: DemoVHasher(validator, serial),
	}
}

func DemoVHasher(validator, serial string) string {
	h := md5.New()
	io.WriteString(h, validator)
	io.WriteString(h, serial)
	return fmt.Sprintf("%v", h.Sum(nil))
}

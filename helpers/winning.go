package helpers

import (
	"crypto/md5"
	"eliest/models"
	"fmt"
	"io"
	"time"
)

func GererateWinning(amt float64, gen string, ) (models.Winnings, string) {
	pin := RandInt(3)
	serial := RandInt(4)
	code := pin + serial

	return models.Winnings{
		Amount:      amt,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		GeneratedBy: gen,
		Hash:        DemoWHasher(code, serial),
	}, code
}

func GererateUsedWinning(amt float64, gen string, ) (models.Winnings, string) {
	pin := RandInt(3)
	serial := RandInt(4)
	code := pin + serial

	return models.Winnings{
		Amount:      amt,
		Status:      "used",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		GeneratedBy: gen,
		Hash:        DemoWHasher(code, serial),
	}, code
}


func WinningHash(validator string, serial string) models.Winnings {
	return models.Winnings{
		Hash: DemoWHasher(validator, serial),
	}
}

func DemoWHasher(validator, serial string) string {
	h := md5.New()
	io.WriteString(h, validator)
	io.WriteString(h, serial)
	return fmt.Sprintf("%v", h.Sum(nil))
}

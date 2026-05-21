package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	MsgOK            = "success"
	MsgValidasiGagal = "validasi gagal"
)

var (
	ErrTooManyRequests = errors.New("terlalu banyak percobaan login, coba lagi nanti")
	ErrBadCredentials  = errors.New("username atau password salah")
)

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(b), err
}

func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

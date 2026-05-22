package util

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02T15:04:05Z07:00"

	ParamID     = "id"
	ParamTeamID = "teamId"
	FormLogo    = "logo"

	MsgOK              = "success"
	MsgValidasiGagal   = "validasi gagal"
	MsgParamTidakValid = "parameter tidak valid"
	MsgGagalAmbilData  = "gagal mengambil data"
)

var (
	ErrTooManyRequests  = errors.New("terlalu banyak percobaan login, coba lagi nanti")
	ErrBadCredentials   = errors.New("username atau password salah")
	ErrJerseyExists     = errors.New("nomor punggung sudah dipakai pemain lain di tim ini")
	ErrSameTeam         = errors.New("tim home dan away tidak boleh sama")
	ErrMatchFinished    = errors.New("pertandingan yang sudah selesai tidak dapat diubah jadwalnya")
	ErrPlayerNotInMatch = errors.New("pemain bukan anggota tim yang bertanding")
	ErrMatchNotDone     = errors.New("pertandingan belum selesai")
	ErrLogoTooLarge     = errors.New("ukuran logo maksimal 2 MB")
	ErrLogoInvalidType  = errors.New("format logo harus JPEG atau PNG")
	ErrInvalidID        = errors.New("id tidak valid")
	ErrInvalidDate      = errors.New("format tanggal harus YYYY-MM-DD")
)

func ErrNotFound(err error, entity string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s tidak ditemukan", entity)
	}

	return err
}

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(b), err
}

func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func BuildOrderClause(sortBy, sortDir string) string {
	by := "created_at"
	if sortBy != "" {
		by = strings.ToLower(sortBy)
	}

	direction := "DESC"
	if strings.ToUpper(sortDir) == "ASC" {
		direction = "ASC"
	}

	return by + " " + direction
}

func ParseDate(s string) (time.Time, error) {
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return time.Time{}, ErrInvalidDate
	}

	return t, nil
}

func ParseID(c *gin.Context, param string) (int64, error) {
	id, err := strconv.ParseInt(c.Param(param), 10, 64)
	if err != nil {
		Fail(c, http.StatusBadRequest, ErrInvalidID.Error(), nil)
		return 0, err
	}
	return id, nil
}

func ParseOptionalDate(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	return ParseDate(s)
}

func FormatDate(t time.Time) string {
	return t.Format(DateLayout)
}

func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeLayout)
}

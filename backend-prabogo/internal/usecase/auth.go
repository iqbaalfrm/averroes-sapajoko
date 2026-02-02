package usecase

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

const (
	otpExpiryMinutes   = 10
	otpCooldownMinutes = 1
)

type AuthUsecase struct {
	repo      domain.AuthRepository
	jwtSecret string
}

func NewAuthUsecase(repo domain.AuthRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *AuthUsecase) Daftar(ctx context.Context, nama, email, kataSandi, peran string) (*domain.Pengguna, *domain.OTPVerifikasi, error) {
	if peran == "" {
		peran = "user"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(kataSandi), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	pengguna := &domain.Pengguna{
		Nama:            nama,
		Email:           email,
		KataSandiHash:   string(hash),
		Peran:           peran,
		Status:          "aktif",
		SudahVerifikasi: false,
		DibuatPada:      time.Now(),
		DiubahPada:      time.Now(),
	}

	id, err := u.repo.BuatPengguna(ctx, pengguna)
	if err != nil {
		return nil, nil, err
	}
	pengguna.ID = id

	kode := buatKodeOTP()
	otp := &domain.OTPVerifikasi{
		IDPengguna:       id,
		Kode:             kode,
		KadaluarsaPada:   time.Now().Add(otpExpiryMinutes * time.Minute),
		TerakhirKirimPada: time.Now(),
		JumlahKirim:      1,
	}

	if err := u.repo.SimpanOTP(ctx, otp); err != nil {
		return nil, nil, err
	}

	return pengguna, otp, nil
}

func (u *AuthUsecase) VerifikasiOTP(ctx context.Context, email, kode string) error {
	pengguna, err := u.repo.CariPenggunaByEmail(ctx, email)
	if err != nil {
		return err
	}
	if pengguna == nil {
		return errors.New("pengguna tidak ditemukan")
	}

	otp, err := u.repo.AmbilOTPByPengguna(ctx, pengguna.ID)
	if err != nil {
		return err
	}
	if otp == nil {
		return errors.New("otp tidak ditemukan")
	}
	if otp.Kode != kode {
		return errors.New("kode otp tidak sesuai")
	}
	if time.Now().After(otp.KadaluarsaPada) {
		return errors.New("kode otp kadaluarsa")
	}

	return u.repo.TandaiPenggunaTerverifikasi(ctx, pengguna.ID)
}

func (u *AuthUsecase) KirimUlangOTP(ctx context.Context, email string) (*domain.OTPVerifikasi, error) {
	pengguna, err := u.repo.CariPenggunaByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if pengguna == nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}

	otp, err := u.repo.AmbilOTPByPengguna(ctx, pengguna.ID)
	if err != nil {
		return nil, err
	}
	if otp == nil {
		return nil, errors.New("otp tidak ditemukan")
	}

	if time.Since(otp.TerakhirKirimPada) < otpCooldownMinutes*time.Minute {
		return nil, fmt.Errorf("cooldown %d menit belum selesai", otpCooldownMinutes)
	}

	otp.Kode = buatKodeOTP()
	otp.KadaluarsaPada = time.Now().Add(otpExpiryMinutes * time.Minute)
	otp.TerakhirKirimPada = time.Now()
	otp.JumlahKirim += 1

	if err := u.repo.PerbaruiOTP(ctx, otp); err != nil {
		return nil, err
	}

	return otp, nil
}

func (u *AuthUsecase) Masuk(ctx context.Context, email, kataSandi string) (*domain.Pengguna, error) {
	pengguna, err := u.repo.CariPenggunaByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if pengguna == nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}
	if !pengguna.SudahVerifikasi {
		return nil, errors.New("pengguna belum verifikasi")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pengguna.KataSandiHash), []byte(kataSandi)); err != nil {
		return nil, errors.New("kata sandi salah")
	}

	return pengguna, nil
}

func (u *AuthUsecase) Profil(ctx context.Context, id int64) (*domain.Pengguna, error) {
	return u.repo.AmbilPenggunaByID(ctx, id)
}

func buatKodeOTP() string {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "123456"
	}
	n := binary.BigEndian.Uint64(b[:])
	kode := n % 1000000
	return fmt.Sprintf("%06d", kode)
}

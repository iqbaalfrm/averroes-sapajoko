package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) BuatPengguna(ctx context.Context, pengguna *domain.Pengguna) (int64, error) {
	query := `INSERT INTO pengguna (nama, email, kata_sandi_hash, peran, status, sudah_verifikasi, dibuat_pada, diubah_pada)
		VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := r.db.ExecContext(ctx, query, pengguna.Nama, pengguna.Email, pengguna.KataSandiHash, pengguna.Peran, pengguna.Status, pengguna.SudahVerifikasi)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *Repository) CariPenggunaByEmail(ctx context.Context, email string) (*domain.Pengguna, error) {
	query := `SELECT id, nama, email, kata_sandi_hash, peran, status, sudah_verifikasi, dibuat_pada, diubah_pada
		FROM pengguna WHERE email = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, email)
	pengguna := &domain.Pengguna{}
	if err := row.Scan(&pengguna.ID, &pengguna.Nama, &pengguna.Email, &pengguna.KataSandiHash, &pengguna.Peran, &pengguna.Status, &pengguna.SudahVerifikasi, &pengguna.DibuatPada, &pengguna.DiubahPada); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return pengguna, nil
}

func (r *Repository) AmbilPenggunaByID(ctx context.Context, id int64) (*domain.Pengguna, error) {
	query := `SELECT id, nama, email, kata_sandi_hash, peran, status, sudah_verifikasi, dibuat_pada, diubah_pada
		FROM pengguna WHERE id = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, id)
	pengguna := &domain.Pengguna{}
	if err := row.Scan(&pengguna.ID, &pengguna.Nama, &pengguna.Email, &pengguna.KataSandiHash, &pengguna.Peran, &pengguna.Status, &pengguna.SudahVerifikasi, &pengguna.DibuatPada, &pengguna.DiubahPada); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return pengguna, nil
}

func (r *Repository) TandaiPenggunaTerverifikasi(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `UPDATE pengguna SET sudah_verifikasi = 1, diubah_pada = NOW() WHERE id = ?`, id)
	return err
}

func (r *Repository) SimpanOTP(ctx context.Context, otp *domain.OTPVerifikasi) error {
	query := `INSERT INTO otp_verifikasi (id_pengguna, kode, kadaluarsa_pada, terakhir_kirim_pada, jumlah_kirim)
		VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, otp.IDPengguna, otp.Kode, otp.KadaluarsaPada, otp.TerakhirKirimPada, otp.JumlahKirim)
	return err
}

func (r *Repository) AmbilOTPByPengguna(ctx context.Context, idPengguna int64) (*domain.OTPVerifikasi, error) {
	query := `SELECT id, id_pengguna, kode, kadaluarsa_pada, terakhir_kirim_pada, jumlah_kirim
		FROM otp_verifikasi WHERE id_pengguna = ? ORDER BY id DESC LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, idPengguna)
	otp := &domain.OTPVerifikasi{}
	if err := row.Scan(&otp.ID, &otp.IDPengguna, &otp.Kode, &otp.KadaluarsaPada, &otp.TerakhirKirimPada, &otp.JumlahKirim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return otp, nil
}

func (r *Repository) PerbaruiOTP(ctx context.Context, otp *domain.OTPVerifikasi) error {
	query := `UPDATE otp_verifikasi SET kode = ?, kadaluarsa_pada = ?, terakhir_kirim_pada = ?, jumlah_kirim = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, otp.Kode, otp.KadaluarsaPada, otp.TerakhirKirimPada, otp.JumlahKirim, otp.ID)
	return err
}

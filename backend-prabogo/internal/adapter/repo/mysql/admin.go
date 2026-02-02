package mysql

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarPengguna(ctx context.Context) ([]domain.Pengguna, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, nama, email, kata_sandi_hash, peran, status, sudah_verifikasi, dibuat_pada, diubah_pada FROM pengguna ORDER BY dibuat_pada DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Pengguna
	for rows.Next() {
		var item domain.Pengguna
		if err := rows.Scan(&item.ID, &item.Nama, &item.Email, &item.KataSandiHash, &item.Peran, &item.Status, &item.SudahVerifikasi, &item.DibuatPada, &item.DiubahPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) PerbaruiPengguna(ctx context.Context, pengguna *domain.Pengguna) error {
	query := `UPDATE pengguna SET nama = ?, email = ?, peran = ?, status = ?, sudah_verifikasi = ?, diubah_pada = NOW() WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, pengguna.Nama, pengguna.Email, pengguna.Peran, pengguna.Status, pengguna.SudahVerifikasi, pengguna.ID)
	return err
}

func (r *Repository) HapusPengguna(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM pengguna WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatKonfigurasi(ctx context.Context, konfigurasi *domain.Konfigurasi) error {
	query := `INSERT INTO konfigurasi (kunci, nilai, deskripsi) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, konfigurasi.Kunci, konfigurasi.Nilai, konfigurasi.Deskripsi)
	return err
}

func (r *Repository) PerbaruiKonfigurasi(ctx context.Context, konfigurasi *domain.Konfigurasi) error {
	query := `UPDATE konfigurasi SET kunci = ?, nilai = ?, deskripsi = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, konfigurasi.Kunci, konfigurasi.Nilai, konfigurasi.Deskripsi, konfigurasi.ID)
	return err
}

func (r *Repository) HapusKonfigurasi(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM konfigurasi WHERE id = ?`, id)
	return err
}

func (r *Repository) DaftarKonfigurasi(ctx context.Context) ([]domain.Konfigurasi, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, kunci, nilai, deskripsi FROM konfigurasi ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Konfigurasi
	for rows.Next() {
		var item domain.Konfigurasi
		if err := rows.Scan(&item.ID, &item.Kunci, &item.Nilai, &item.Deskripsi); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

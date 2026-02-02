package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarDiskusi(ctx context.Context) ([]domain.Diskusi, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_pengguna, judul, isi, status, dibuat_pada FROM diskusi ORDER BY dibuat_pada DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Diskusi
	for rows.Next() {
		var item domain.Diskusi
		if err := rows.Scan(&item.ID, &item.IDPengguna, &item.Judul, &item.Isi, &item.Status, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DetailDiskusi(ctx context.Context, id int64) (*domain.Diskusi, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, id_pengguna, judul, isi, status, dibuat_pada FROM diskusi WHERE id = ?`, id)
	var item domain.Diskusi
	if err := row.Scan(&item.ID, &item.IDPengguna, &item.Judul, &item.Isi, &item.Status, &item.DibuatPada); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *Repository) DaftarBalasan(ctx context.Context, idDiskusi int64) ([]domain.DiskusiBalas, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_diskusi, id_pengguna, isi, dibuat_pada FROM diskusi_balas WHERE id_diskusi = ? ORDER BY dibuat_pada ASC`, idDiskusi)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.DiskusiBalas
	for rows.Next() {
		var item domain.DiskusiBalas
		if err := rows.Scan(&item.ID, &item.IDDiskusi, &item.IDPengguna, &item.Isi, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) BuatDiskusi(ctx context.Context, diskusi *domain.Diskusi) error {
	query := `INSERT INTO diskusi (id_pengguna, judul, isi, status, dibuat_pada) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, diskusi.IDPengguna, diskusi.Judul, diskusi.Isi, diskusi.Status, diskusi.DibuatPada)
	return err
}

func (r *Repository) BuatBalasan(ctx context.Context, balas *domain.DiskusiBalas) error {
	query := `INSERT INTO diskusi_balas (id_diskusi, id_pengguna, isi, dibuat_pada) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, balas.IDDiskusi, balas.IDPengguna, balas.Isi, balas.DibuatPada)
	return err
}

func (r *Repository) BuatLaporan(ctx context.Context, laporan *domain.DiskusiLaporan) error {
	query := `INSERT INTO diskusi_laporan (id_diskusi, id_pengguna, alasan, dibuat_pada) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, laporan.IDDiskusi, laporan.IDPengguna, laporan.Alasan, laporan.DibuatPada)
	return err
}

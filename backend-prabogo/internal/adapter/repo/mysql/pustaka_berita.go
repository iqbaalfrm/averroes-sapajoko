package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarPustaka(ctx context.Context) ([]domain.Pustaka, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, judul_tampil, judul_asli, penulis, kategori, bahasa, jumlah_halaman, deskripsi, tautan_file FROM pustaka ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Pustaka
	for rows.Next() {
		var item domain.Pustaka
		if err := rows.Scan(&item.ID, &item.JudulTampil, &item.JudulAsli, &item.Penulis, &item.Kategori, &item.Bahasa, &item.JumlahHalaman, &item.Deskripsi, &item.TautanFile); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DetailPustaka(ctx context.Context, id int64) (*domain.Pustaka, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, judul_tampil, judul_asli, penulis, kategori, bahasa, jumlah_halaman, deskripsi, tautan_file FROM pustaka WHERE id = ?`, id)
	var item domain.Pustaka
	if err := row.Scan(&item.ID, &item.JudulTampil, &item.JudulAsli, &item.Penulis, &item.Kategori, &item.Bahasa, &item.JumlahHalaman, &item.Deskripsi, &item.TautanFile); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *Repository) DaftarBerita(ctx context.Context, limit int) ([]domain.Berita, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, judul, ringkasan, isi, kategori, sumber, gambar_url, diterbitkan_pada FROM berita ORDER BY diterbitkan_pada DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Berita
	for rows.Next() {
		var item domain.Berita
		if err := rows.Scan(&item.ID, &item.Judul, &item.Ringkasan, &item.Isi, &item.Kategori, &item.Sumber, &item.GambarURL, &item.DiterbitkanPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarBeritaTerbaru(ctx context.Context, limit int) ([]domain.Berita, error) {
	return r.DaftarBerita(ctx, limit)
}

func (r *Repository) BuatPustaka(ctx context.Context, pustaka *domain.Pustaka) error {
	query := `INSERT INTO pustaka (judul_tampil, judul_asli, penulis, kategori, bahasa, jumlah_halaman, deskripsi, tautan_file) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, pustaka.JudulTampil, pustaka.JudulAsli, pustaka.Penulis, pustaka.Kategori, pustaka.Bahasa, pustaka.JumlahHalaman, pustaka.Deskripsi, pustaka.TautanFile)
	return err
}

func (r *Repository) PerbaruiPustaka(ctx context.Context, pustaka *domain.Pustaka) error {
	query := `UPDATE pustaka SET judul_tampil = ?, judul_asli = ?, penulis = ?, kategori = ?, bahasa = ?, jumlah_halaman = ?, deskripsi = ?, tautan_file = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, pustaka.JudulTampil, pustaka.JudulAsli, pustaka.Penulis, pustaka.Kategori, pustaka.Bahasa, pustaka.JumlahHalaman, pustaka.Deskripsi, pustaka.TautanFile, pustaka.ID)
	return err
}

func (r *Repository) HapusPustaka(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM pustaka WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatBerita(ctx context.Context, berita *domain.Berita) error {
	query := `INSERT INTO berita (judul, ringkasan, isi, kategori, sumber, gambar_url, diterbitkan_pada) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, berita.Judul, berita.Ringkasan, berita.Isi, berita.Kategori, berita.Sumber, berita.GambarURL, berita.DiterbitkanPada)
	return err
}

func (r *Repository) PerbaruiBerita(ctx context.Context, berita *domain.Berita) error {
	query := `UPDATE berita SET judul = ?, ringkasan = ?, isi = ?, kategori = ?, sumber = ?, gambar_url = ?, diterbitkan_pada = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, berita.Judul, berita.Ringkasan, berita.Isi, berita.Kategori, berita.Sumber, berita.GambarURL, berita.DiterbitkanPada, berita.ID)
	return err
}

func (r *Repository) HapusBerita(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM berita WHERE id = ?`, id)
	return err
}

package mysql

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarScreener(ctx context.Context, kategori, cari string) ([]domain.Screener, error) {
	query := `SELECT id, nama_aset, simbol, kategori, skor_syariah, keterangan, harga_terakhir, perubahan_24j, dibuat_pada
		FROM screener WHERE 1=1`
	args := []interface{}{}
	if kategori != "" && kategori != "semua" {
		query += " AND kategori = ?"
		args = append(args, kategori)
	}
	if cari != "" {
		query += " AND (nama_aset LIKE ? OR simbol LIKE ?)"
		like := "%" + cari + "%"
		args = append(args, like, like)
	}
	query += " ORDER BY nama_aset ASC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Screener
	for rows.Next() {
		var item domain.Screener
		if err := rows.Scan(&item.ID, &item.NamaAset, &item.Simbol, &item.Kategori, &item.SkorSyariah, &item.Keterangan, &item.HargaTerakhir, &item.Perubahan24J, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DetailScreener(ctx context.Context, id int64) (*domain.Screener, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, nama_aset, simbol, kategori, skor_syariah, keterangan, harga_terakhir, perubahan_24j, dibuat_pada FROM screener WHERE id = ?`, id)
	var item domain.Screener
	if err := row.Scan(&item.ID, &item.NamaAset, &item.Simbol, &item.Kategori, &item.SkorSyariah, &item.Keterangan, &item.HargaTerakhir, &item.Perubahan24J, &item.DibuatPada); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *Repository) DaftarCatatanScreener(ctx context.Context, idScreener int64) ([]domain.ScreenerCatatan, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_screener, judul, isi, dibuat_pada FROM screener_catatan WHERE id_screener = ? ORDER BY id DESC`, idScreener)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ScreenerCatatan
	for rows.Next() {
		var item domain.ScreenerCatatan
		if err := rows.Scan(&item.ID, &item.IDScreener, &item.Judul, &item.Isi, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarPasar(ctx context.Context) ([]domain.Pasar, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, nama_aset, simbol, harga, volume_24j, perubahan_24j, kapitalisasi_pasar, diperbarui_pada FROM pasar ORDER BY kapitalisasi_pasar DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Pasar
	for rows.Next() {
		var item domain.Pasar
		if err := rows.Scan(&item.ID, &item.NamaAset, &item.Simbol, &item.Harga, &item.Volume24J, &item.Perubahan24J, &item.KapitalisasiPasar, &item.DiperbaruiPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) BuatScreener(ctx context.Context, screener *domain.Screener) error {
	query := `INSERT INTO screener (nama_aset, simbol, kategori, skor_syariah, keterangan, harga_terakhir, perubahan_24j, dibuat_pada) VALUES (?, ?, ?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, screener.NamaAset, screener.Simbol, screener.Kategori, screener.SkorSyariah, screener.Keterangan, screener.HargaTerakhir, screener.Perubahan24J)
	return err
}

func (r *Repository) PerbaruiScreener(ctx context.Context, screener *domain.Screener) error {
	query := `UPDATE screener SET nama_aset = ?, simbol = ?, kategori = ?, skor_syariah = ?, keterangan = ?, harga_terakhir = ?, perubahan_24j = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, screener.NamaAset, screener.Simbol, screener.Kategori, screener.SkorSyariah, screener.Keterangan, screener.HargaTerakhir, screener.Perubahan24J, screener.ID)
	return err
}

func (r *Repository) HapusScreener(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM screener WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatPasar(ctx context.Context, pasar *domain.Pasar) error {
	query := `INSERT INTO pasar (nama_aset, simbol, harga, volume_24j, perubahan_24j, kapitalisasi_pasar, diperbarui_pada) VALUES (?, ?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, pasar.NamaAset, pasar.Simbol, pasar.Harga, pasar.Volume24J, pasar.Perubahan24J, pasar.KapitalisasiPasar)
	return err
}

func (r *Repository) PerbaruiPasar(ctx context.Context, pasar *domain.Pasar) error {
	query := `UPDATE pasar SET nama_aset = ?, simbol = ?, harga = ?, volume_24j = ?, perubahan_24j = ?, kapitalisasi_pasar = ?, diperbarui_pada = NOW() WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, pasar.NamaAset, pasar.Simbol, pasar.Harga, pasar.Volume24J, pasar.Perubahan24J, pasar.KapitalisasiPasar, pasar.ID)
	return err
}

func (r *Repository) HapusPasar(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM pasar WHERE id = ?`, id)
	return err
}

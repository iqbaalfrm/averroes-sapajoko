package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarPortofolio(ctx context.Context, idPengguna int64) ([]domain.Portofolio, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_pengguna, nama_aset, simbol, jumlah, harga_beli, nilai_saat_ini, kategori, dibuat_pada FROM portofolio WHERE id_pengguna = ? ORDER BY dibuat_pada DESC`, idPengguna)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Portofolio
	for rows.Next() {
		var item domain.Portofolio
		if err := rows.Scan(&item.ID, &item.IDPengguna, &item.NamaAset, &item.Simbol, &item.Jumlah, &item.HargaBeli, &item.NilaiSaatIni, &item.Kategori, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) SimpanPortofolio(ctx context.Context, portofolio *domain.Portofolio) error {
	query := `INSERT INTO portofolio (id_pengguna, nama_aset, simbol, jumlah, harga_beli, nilai_saat_ini, kategori, dibuat_pada) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, portofolio.IDPengguna, portofolio.NamaAset, portofolio.Simbol, portofolio.Jumlah, portofolio.HargaBeli, portofolio.NilaiSaatIni, portofolio.Kategori, portofolio.DibuatPada)
	return err
}

func (r *Repository) PerbaruiPortofolio(ctx context.Context, portofolio *domain.Portofolio) error {
	query := `UPDATE portofolio SET nama_aset = ?, simbol = ?, jumlah = ?, harga_beli = ?, nilai_saat_ini = ?, kategori = ? WHERE id = ? AND id_pengguna = ?`
	_, err := r.db.ExecContext(ctx, query, portofolio.NamaAset, portofolio.Simbol, portofolio.Jumlah, portofolio.HargaBeli, portofolio.NilaiSaatIni, portofolio.Kategori, portofolio.ID, portofolio.IDPengguna)
	return err
}

func (r *Repository) HapusPortofolio(ctx context.Context, id int64, idPengguna int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM portofolio WHERE id = ? AND id_pengguna = ?`, id, idPengguna)
	return err
}

func (r *Repository) HargaEmasTerbaru(ctx context.Context) (*domain.HargaEmas, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, tanggal, harga_per_gram FROM harga_emas ORDER BY tanggal DESC LIMIT 1`)
	var item domain.HargaEmas
	if err := row.Scan(&item.ID, &item.Tanggal, &item.HargaPerGram); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *Repository) DaftarRiwayatZakat(ctx context.Context, idPengguna int64) ([]domain.ZakatRiwayat, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_pengguna, total_nilai, nisab, persen_zakat, zakat_terhitung, dibuat_pada FROM zakat_riwayat WHERE id_pengguna = ? ORDER BY dibuat_pada DESC`, idPengguna)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ZakatRiwayat
	for rows.Next() {
		var item domain.ZakatRiwayat
		if err := rows.Scan(&item.ID, &item.IDPengguna, &item.TotalNilai, &item.Nisab, &item.PersenZakat, &item.ZakatTerhitung, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) SimpanRiwayatZakat(ctx context.Context, riwayat *domain.ZakatRiwayat) error {
	query := `INSERT INTO zakat_riwayat (id_pengguna, total_nilai, nisab, persen_zakat, zakat_terhitung, dibuat_pada) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, riwayat.IDPengguna, riwayat.TotalNilai, riwayat.Nisab, riwayat.PersenZakat, riwayat.ZakatTerhitung, riwayat.DibuatPada)
	return err
}

package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarKelas(ctx context.Context) ([]domain.Kelas, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, judul, deskripsi, level, jumlah_modul, durasi_menit, thumbnail_url, status, dibuat_pada FROM kelas ORDER BY dibuat_pada DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Kelas
	for rows.Next() {
		var item domain.Kelas
		if err := rows.Scan(&item.ID, &item.Judul, &item.Deskripsi, &item.Level, &item.JumlahModul, &item.DurasiMenit, &item.ThumbnailURL, &item.Status, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DetailKelas(ctx context.Context, id int64) (*domain.Kelas, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, judul, deskripsi, level, jumlah_modul, durasi_menit, thumbnail_url, status, dibuat_pada FROM kelas WHERE id = ?`, id)
	var item domain.Kelas
	if err := row.Scan(&item.ID, &item.Judul, &item.Deskripsi, &item.Level, &item.JumlahModul, &item.DurasiMenit, &item.ThumbnailURL, &item.Status, &item.DibuatPada); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *Repository) DaftarModulByKelas(ctx context.Context, idKelas int64) ([]domain.Modul, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_kelas, judul, urutan, ringkasan, durasi_menit, dibuat_pada FROM modul WHERE id_kelas = ? ORDER BY urutan ASC`, idKelas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Modul
	for rows.Next() {
		var item domain.Modul
		if err := rows.Scan(&item.ID, &item.IDKelas, &item.Judul, &item.Urutan, &item.Ringkasan, &item.DurasiMenit, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarMateriByModul(ctx context.Context, idModul int64) ([]domain.Materi, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_modul, judul, tipe, konten, url_video, durasi_menit, dibuat_pada FROM materi WHERE id_modul = ? ORDER BY id ASC`, idModul)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Materi
	for rows.Next() {
		var item domain.Materi
		if err := rows.Scan(&item.ID, &item.IDModul, &item.Judul, &item.Tipe, &item.Konten, &item.URLVideo, &item.DurasiMenit, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarUjianByKelas(ctx context.Context, idKelas int64) ([]domain.Ujian, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_kelas, judul, deskripsi, durasi_menit, jumlah_soal, dibuat_pada FROM ujian WHERE id_kelas = ? ORDER BY id ASC`, idKelas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Ujian
	for rows.Next() {
		var item domain.Ujian
		if err := rows.Scan(&item.ID, &item.IDKelas, &item.Judul, &item.Deskripsi, &item.DurasiMenit, &item.JumlahSoal, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarModulSemua(ctx context.Context) ([]domain.Modul, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_kelas, judul, urutan, ringkasan, durasi_menit, dibuat_pada FROM modul ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Modul
	for rows.Next() {
		var item domain.Modul
		if err := rows.Scan(&item.ID, &item.IDKelas, &item.Judul, &item.Urutan, &item.Ringkasan, &item.DurasiMenit, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarMateriSemua(ctx context.Context) ([]domain.Materi, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_modul, judul, tipe, konten, url_video, durasi_menit, dibuat_pada FROM materi ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Materi
	for rows.Next() {
		var item domain.Materi
		if err := rows.Scan(&item.ID, &item.IDModul, &item.Judul, &item.Tipe, &item.Konten, &item.URLVideo, &item.DurasiMenit, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarUjianSemua(ctx context.Context) ([]domain.Ujian, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_kelas, judul, deskripsi, durasi_menit, jumlah_soal, dibuat_pada FROM ujian ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Ujian
	for rows.Next() {
		var item domain.Ujian
		if err := rows.Scan(&item.ID, &item.IDKelas, &item.Judul, &item.Deskripsi, &item.DurasiMenit, &item.JumlahSoal, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DaftarSertifikat(ctx context.Context) ([]domain.Sertifikat, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_pengguna, id_kelas, kode, tanggal_terbit FROM sertifikat ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Sertifikat
	for rows.Next() {
		var item domain.Sertifikat
		if err := rows.Scan(&item.ID, &item.IDPengguna, &item.IDKelas, &item.Kode, &item.TanggalTerbit); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) SimpanProgress(ctx context.Context, progress *domain.ProgressKelas) error {
	query := `INSERT INTO progress_kelas (id_pengguna, id_kelas, persentase, status, terakhir_diakses_pada)
		VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE persentase = VALUES(persentase), status = VALUES(status), terakhir_diakses_pada = VALUES(terakhir_diakses_pada)`
	_, err := r.db.ExecContext(ctx, query, progress.IDPengguna, progress.IDKelas, progress.Persentase, progress.Status, progress.TerakhirDiaksesPada)
	return err
}

func (r *Repository) DaftarProgress(ctx context.Context, idPengguna int64) ([]domain.ProgressKelas, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_pengguna, id_kelas, persentase, status, terakhir_diakses_pada FROM progress_kelas WHERE id_pengguna = ? ORDER BY terakhir_diakses_pada DESC`, idPengguna)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ProgressKelas
	for rows.Next() {
		var item domain.ProgressKelas
		if err := rows.Scan(&item.ID, &item.IDPengguna, &item.IDKelas, &item.Persentase, &item.Status, &item.TerakhirDiaksesPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) BuatKelas(ctx context.Context, kelas *domain.Kelas) error {
	query := `INSERT INTO kelas (judul, deskripsi, level, jumlah_modul, durasi_menit, thumbnail_url, status, dibuat_pada) VALUES (?, ?, ?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, kelas.Judul, kelas.Deskripsi, kelas.Level, kelas.JumlahModul, kelas.DurasiMenit, kelas.ThumbnailURL, kelas.Status)
	return err
}

func (r *Repository) PerbaruiKelas(ctx context.Context, kelas *domain.Kelas) error {
	query := `UPDATE kelas SET judul = ?, deskripsi = ?, level = ?, jumlah_modul = ?, durasi_menit = ?, thumbnail_url = ?, status = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, kelas.Judul, kelas.Deskripsi, kelas.Level, kelas.JumlahModul, kelas.DurasiMenit, kelas.ThumbnailURL, kelas.Status, kelas.ID)
	return err
}

func (r *Repository) HapusKelas(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM kelas WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatModul(ctx context.Context, modul *domain.Modul) error {
	query := `INSERT INTO modul (id_kelas, judul, urutan, ringkasan, durasi_menit, dibuat_pada) VALUES (?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, modul.IDKelas, modul.Judul, modul.Urutan, modul.Ringkasan, modul.DurasiMenit)
	return err
}

func (r *Repository) PerbaruiModul(ctx context.Context, modul *domain.Modul) error {
	query := `UPDATE modul SET id_kelas = ?, judul = ?, urutan = ?, ringkasan = ?, durasi_menit = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, modul.IDKelas, modul.Judul, modul.Urutan, modul.Ringkasan, modul.DurasiMenit, modul.ID)
	return err
}

func (r *Repository) HapusModul(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM modul WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatMateri(ctx context.Context, materi *domain.Materi) error {
	query := `INSERT INTO materi (id_modul, judul, tipe, konten, url_video, durasi_menit, dibuat_pada) VALUES (?, ?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, materi.IDModul, materi.Judul, materi.Tipe, materi.Konten, materi.URLVideo, materi.DurasiMenit)
	return err
}

func (r *Repository) PerbaruiMateri(ctx context.Context, materi *domain.Materi) error {
	query := `UPDATE materi SET id_modul = ?, judul = ?, tipe = ?, konten = ?, url_video = ?, durasi_menit = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, materi.IDModul, materi.Judul, materi.Tipe, materi.Konten, materi.URLVideo, materi.DurasiMenit, materi.ID)
	return err
}

func (r *Repository) HapusMateri(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM materi WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatUjian(ctx context.Context, ujian *domain.Ujian) error {
	query := `INSERT INTO ujian (id_kelas, judul, deskripsi, durasi_menit, jumlah_soal, dibuat_pada) VALUES (?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, ujian.IDKelas, ujian.Judul, ujian.Deskripsi, ujian.DurasiMenit, ujian.JumlahSoal)
	return err
}

func (r *Repository) PerbaruiUjian(ctx context.Context, ujian *domain.Ujian) error {
	query := `UPDATE ujian SET id_kelas = ?, judul = ?, deskripsi = ?, durasi_menit = ?, jumlah_soal = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, ujian.IDKelas, ujian.Judul, ujian.Deskripsi, ujian.DurasiMenit, ujian.JumlahSoal, ujian.ID)
	return err
}

func (r *Repository) HapusUjian(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM ujian WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatSertifikat(ctx context.Context, sertifikat *domain.Sertifikat) error {
	query := `INSERT INTO sertifikat (id_pengguna, id_kelas, kode, tanggal_terbit) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, sertifikat.IDPengguna, sertifikat.IDKelas, sertifikat.Kode, sertifikat.TanggalTerbit)
	return err
}

func (r *Repository) HapusSertifikat(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM sertifikat WHERE id = ?`, id)
	return err
}

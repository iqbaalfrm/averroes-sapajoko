package mysql

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

func (r *Repository) DaftarReels(ctx context.Context, tema string) ([]domain.Reels, error) {
	query := `SELECT id, judul, tema, kutipan, sumber, url_video, thumbnail_url, dibuat_pada FROM reels`
	args := []interface{}{}
	if tema != "" {
		query += " WHERE tema = ?"
		args = append(args, tema)
	}
	query += " ORDER BY dibuat_pada DESC"
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Reels
	for rows.Next() {
		var item domain.Reels
		if err := rows.Scan(&item.ID, &item.Judul, &item.Tema, &item.Kutipan, &item.Sumber, &item.URLVideo, &item.ThumbnailURL, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) DetailReels(ctx context.Context, id int64) (*domain.Reels, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, judul, tema, kutipan, sumber, url_video, thumbnail_url, dibuat_pada FROM reels WHERE id = ?`, id)
	var item domain.Reels
	if err := row.Scan(&item.ID, &item.Judul, &item.Tema, &item.Kutipan, &item.Sumber, &item.URLVideo, &item.ThumbnailURL, &item.DibuatPada); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *Repository) DaftarTadabbur(ctx context.Context, tema string) ([]domain.Tadabbur, error) {
	query := `SELECT id, judul, tema, ringkasan, isi, sumber, dibuat_pada FROM tadabbur`
	args := []interface{}{}
	if tema != "" {
		query += " WHERE tema = ?"
		args = append(args, tema)
	}
	query += " ORDER BY dibuat_pada DESC"
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Tadabbur
	for rows.Next() {
		var item domain.Tadabbur
		if err := rows.Scan(&item.ID, &item.Judul, &item.Tema, &item.Ringkasan, &item.Isi, &item.Sumber, &item.DibuatPada); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) BuatReels(ctx context.Context, reels *domain.Reels) error {
	query := `INSERT INTO reels (judul, tema, kutipan, sumber, url_video, thumbnail_url, dibuat_pada) VALUES (?, ?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, reels.Judul, reels.Tema, reels.Kutipan, reels.Sumber, reels.URLVideo, reels.ThumbnailURL)
	return err
}

func (r *Repository) PerbaruiReels(ctx context.Context, reels *domain.Reels) error {
	query := `UPDATE reels SET judul = ?, tema = ?, kutipan = ?, sumber = ?, url_video = ?, thumbnail_url = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, reels.Judul, reels.Tema, reels.Kutipan, reels.Sumber, reels.URLVideo, reels.ThumbnailURL, reels.ID)
	return err
}

func (r *Repository) HapusReels(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM reels WHERE id = ?`, id)
	return err
}

func (r *Repository) BuatTadabbur(ctx context.Context, tadabbur *domain.Tadabbur) error {
	query := `INSERT INTO tadabbur (judul, tema, ringkasan, isi, sumber, dibuat_pada) VALUES (?, ?, ?, ?, ?, NOW())`
	_, err := r.db.ExecContext(ctx, query, tadabbur.Judul, tadabbur.Tema, tadabbur.Ringkasan, tadabbur.Isi, tadabbur.Sumber)
	return err
}

func (r *Repository) PerbaruiTadabbur(ctx context.Context, tadabbur *domain.Tadabbur) error {
	query := `UPDATE tadabbur SET judul = ?, tema = ?, ringkasan = ?, isi = ?, sumber = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, tadabbur.Judul, tadabbur.Tema, tadabbur.Ringkasan, tadabbur.Isi, tadabbur.Sumber, tadabbur.ID)
	return err
}

func (r *Repository) HapusTadabbur(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM tadabbur WHERE id = ?`, id)
	return err
}

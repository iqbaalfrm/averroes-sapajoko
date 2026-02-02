package usecase

import (
	"context"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type PortofolioUsecase struct {
	repo domain.PortofolioRepository
}

func NewPortofolioUsecase(repo domain.PortofolioRepository) *PortofolioUsecase {
	return &PortofolioUsecase{repo: repo}
}

func (u *PortofolioUsecase) Daftar(ctx context.Context, idPengguna int64) ([]domain.Portofolio, error) {
	return u.repo.DaftarPortofolio(ctx, idPengguna)
}

func (u *PortofolioUsecase) Tambah(ctx context.Context, portofolio *domain.Portofolio) error {
	portofolio.DibuatPada = time.Now()
	return u.repo.SimpanPortofolio(ctx, portofolio)
}

func (u *PortofolioUsecase) Perbarui(ctx context.Context, portofolio *domain.Portofolio) error {
	return u.repo.PerbaruiPortofolio(ctx, portofolio)
}

func (u *PortofolioUsecase) Hapus(ctx context.Context, id int64, idPengguna int64) error {
	return u.repo.HapusPortofolio(ctx, id, idPengguna)
}

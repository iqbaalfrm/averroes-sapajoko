package usecase

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type ScreenerUsecase struct {
	repo domain.ScreenerRepository
}

func NewScreenerUsecase(repo domain.ScreenerRepository) *ScreenerUsecase {
	return &ScreenerUsecase{repo: repo}
}

func (u *ScreenerUsecase) Daftar(ctx context.Context, kategori, cari string) ([]domain.Screener, error) {
	return u.repo.DaftarScreener(ctx, kategori, cari)
}

func (u *ScreenerUsecase) Detail(ctx context.Context, id int64) (*domain.Screener, error) {
	return u.repo.DetailScreener(ctx, id)
}

func (u *ScreenerUsecase) Catatan(ctx context.Context, idScreener int64) ([]domain.ScreenerCatatan, error) {
	return u.repo.DaftarCatatanScreener(ctx, idScreener)
}

func (u *ScreenerUsecase) Pasar(ctx context.Context) ([]domain.Pasar, error) {
	return u.repo.DaftarPasar(ctx)
}

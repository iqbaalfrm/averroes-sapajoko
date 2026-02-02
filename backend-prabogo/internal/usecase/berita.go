package usecase

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type BeritaUsecase struct {
	repo domain.BeritaRepository
}

func NewBeritaUsecase(repo domain.BeritaRepository) *BeritaUsecase {
	return &BeritaUsecase{repo: repo}
}

func (u *BeritaUsecase) Daftar(ctx context.Context, limit int) ([]domain.Berita, error) {
	return u.repo.DaftarBerita(ctx, limit)
}

func (u *BeritaUsecase) Terbaru(ctx context.Context, limit int) ([]domain.Berita, error) {
	return u.repo.DaftarBeritaTerbaru(ctx, limit)
}

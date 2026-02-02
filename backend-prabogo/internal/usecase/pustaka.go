package usecase

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type PustakaUsecase struct {
	repo domain.PustakaRepository
}

func NewPustakaUsecase(repo domain.PustakaRepository) *PustakaUsecase {
	return &PustakaUsecase{repo: repo}
}

func (u *PustakaUsecase) Daftar(ctx context.Context) ([]domain.Pustaka, error) {
	return u.repo.DaftarPustaka(ctx)
}

func (u *PustakaUsecase) Detail(ctx context.Context, id int64) (*domain.Pustaka, error) {
	return u.repo.DetailPustaka(ctx, id)
}

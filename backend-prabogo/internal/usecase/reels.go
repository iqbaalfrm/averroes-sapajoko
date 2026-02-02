package usecase

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type ReelsUsecase struct {
	repo domain.ReelsRepository
}

func NewReelsUsecase(repo domain.ReelsRepository) *ReelsUsecase {
	return &ReelsUsecase{repo: repo}
}

func (u *ReelsUsecase) Daftar(ctx context.Context, tema string) ([]domain.Reels, error) {
	return u.repo.DaftarReels(ctx, tema)
}

func (u *ReelsUsecase) Detail(ctx context.Context, id int64) (*domain.Reels, error) {
	return u.repo.DetailReels(ctx, id)
}

package usecase

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type TadabburUsecase struct {
	repo domain.TadabburRepository
}

func NewTadabburUsecase(repo domain.TadabburRepository) *TadabburUsecase {
	return &TadabburUsecase{repo: repo}
}

func (u *TadabburUsecase) Daftar(ctx context.Context, tema string) ([]domain.Tadabbur, error) {
	return u.repo.DaftarTadabbur(ctx, tema)
}

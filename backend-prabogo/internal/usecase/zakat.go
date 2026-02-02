package usecase

import (
	"context"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type ZakatUsecase struct {
	portofolioRepo domain.PortofolioRepository
	repo           domain.ZakatRepository
}

func NewZakatUsecase(portofolioRepo domain.PortofolioRepository, repo domain.ZakatRepository) *ZakatUsecase {
	return &ZakatUsecase{portofolioRepo: portofolioRepo, repo: repo}
}

func (u *ZakatUsecase) Ringkasan(ctx context.Context, idPengguna int64) (*domain.ZakatRingkasan, error) {
	portofolio, err := u.portofolioRepo.DaftarPortofolio(ctx, idPengguna)
	if err != nil {
		return nil, err
	}

	total := 0.0
	for _, item := range portofolio {
		total += item.NilaiSaatIni
	}

	hargaEmas, err := u.repo.HargaEmasTerbaru(ctx)
	if err != nil {
		return nil, err
	}
	if hargaEmas == nil {
		hargaEmas = &domain.HargaEmas{HargaPerGram: 0}
	}

	nisab := hargaEmas.HargaPerGram * 85
	persen := 0.025
	zakat := total * persen
	wajib := total >= nisab && nisab > 0

	ringkasan := &domain.ZakatRingkasan{
		TotalNilai:    total,
		Nisab:         nisab,
		PersenZakat:   persen * 100,
		ZakatTerhitung: zakat,
		WajibZakat:    wajib,
	}

	if wajib {
		_ = u.repo.SimpanRiwayatZakat(ctx, &domain.ZakatRiwayat{
			IDPengguna:    idPengguna,
			TotalNilai:    total,
			Nisab:         nisab,
			PersenZakat:   persen * 100,
			ZakatTerhitung: zakat,
			DibuatPada:    time.Now(),
		})
	}

	return ringkasan, nil
}

func (u *ZakatUsecase) Riwayat(ctx context.Context, idPengguna int64) ([]domain.ZakatRiwayat, error) {
	return u.repo.DaftarRiwayatZakat(ctx, idPengguna)
}

func (u *ZakatUsecase) HargaEmas(ctx context.Context) (*domain.HargaEmas, error) {
	return u.repo.HargaEmasTerbaru(ctx)
}

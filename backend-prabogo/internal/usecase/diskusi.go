package usecase

import (
	"context"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type DiskusiUsecase struct {
	repo domain.DiskusiRepository
}

func NewDiskusiUsecase(repo domain.DiskusiRepository) *DiskusiUsecase {
	return &DiskusiUsecase{repo: repo}
}

func (u *DiskusiUsecase) Daftar(ctx context.Context) ([]domain.Diskusi, error) {
	return u.repo.DaftarDiskusi(ctx)
}

func (u *DiskusiUsecase) Detail(ctx context.Context, id int64) (*domain.Diskusi, []domain.DiskusiBalas, error) {
	diskusi, err := u.repo.DetailDiskusi(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	balasan, err := u.repo.DaftarBalasan(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	return diskusi, balasan, nil
}

func (u *DiskusiUsecase) Buat(ctx context.Context, idPengguna int64, judul, isi string) (*domain.Diskusi, error) {
	diskusi := &domain.Diskusi{
		IDPengguna: idPengguna,
		Judul:      judul,
		Isi:        isi,
		Status:     "aktif",
		DibuatPada: time.Now(),
	}
	return diskusi, u.repo.BuatDiskusi(ctx, diskusi)
}

func (u *DiskusiUsecase) Balas(ctx context.Context, idPengguna, idDiskusi int64, isi string) (*domain.DiskusiBalas, error) {
	balas := &domain.DiskusiBalas{
		IDDiskusi: idDiskusi,
		IDPengguna: idPengguna,
		Isi:       isi,
		DibuatPada: time.Now(),
	}
	return balas, u.repo.BuatBalasan(ctx, balas)
}

func (u *DiskusiUsecase) Lapor(ctx context.Context, idPengguna, idDiskusi int64, alasan string) error {
	laporan := &domain.DiskusiLaporan{
		IDDiskusi: idDiskusi,
		IDPengguna: idPengguna,
		Alasan:    alasan,
		DibuatPada: time.Now(),
	}
	return u.repo.BuatLaporan(ctx, laporan)
}

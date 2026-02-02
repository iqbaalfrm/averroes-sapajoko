package usecase

import (
	"context"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type EdukasiUsecase struct {
	repo domain.EdukasiRepository
}

func NewEdukasiUsecase(repo domain.EdukasiRepository) *EdukasiUsecase {
	return &EdukasiUsecase{repo: repo}
}

func (u *EdukasiUsecase) DaftarKelas(ctx context.Context) ([]domain.Kelas, error) {
	return u.repo.DaftarKelas(ctx)
}

func (u *EdukasiUsecase) DetailKelas(ctx context.Context, id int64) (*domain.Kelas, error) {
	return u.repo.DetailKelas(ctx, id)
}

func (u *EdukasiUsecase) DaftarModul(ctx context.Context, idKelas int64) ([]domain.Modul, error) {
	return u.repo.DaftarModulByKelas(ctx, idKelas)
}

func (u *EdukasiUsecase) DaftarMateri(ctx context.Context, idModul int64) ([]domain.Materi, error) {
	return u.repo.DaftarMateriByModul(ctx, idModul)
}

func (u *EdukasiUsecase) DaftarUjian(ctx context.Context, idKelas int64) ([]domain.Ujian, error) {
	return u.repo.DaftarUjianByKelas(ctx, idKelas)
}

func (u *EdukasiUsecase) DaftarModulSemua(ctx context.Context) ([]domain.Modul, error) {
	return u.repo.DaftarModulSemua(ctx)
}

func (u *EdukasiUsecase) DaftarMateriSemua(ctx context.Context) ([]domain.Materi, error) {
	return u.repo.DaftarMateriSemua(ctx)
}

func (u *EdukasiUsecase) DaftarUjianSemua(ctx context.Context) ([]domain.Ujian, error) {
	return u.repo.DaftarUjianSemua(ctx)
}

func (u *EdukasiUsecase) DaftarSertifikat(ctx context.Context) ([]domain.Sertifikat, error) {
	return u.repo.DaftarSertifikat(ctx)
}

func (u *EdukasiUsecase) MulaiKelas(ctx context.Context, idPengguna, idKelas int64) error {
	progress := &domain.ProgressKelas{
		IDPengguna:         idPengguna,
		IDKelas:            idKelas,
		Persentase:         0,
		Status:             "berjalan",
		TerakhirDiaksesPada: time.Now(),
	}
	return u.repo.SimpanProgress(ctx, progress)
}

func (u *EdukasiUsecase) DaftarProgress(ctx context.Context, idPengguna int64) ([]domain.ProgressKelas, error) {
	return u.repo.DaftarProgress(ctx, idPengguna)
}

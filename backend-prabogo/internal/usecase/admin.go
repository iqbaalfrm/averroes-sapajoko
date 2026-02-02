package usecase

import (
	"context"

	"github.com/averroes/backend-prabogo/internal/domain"
)

type AdminUsecase struct {
	repo domain.AdminRepository
}

func NewAdminUsecase(repo domain.AdminRepository) *AdminUsecase {
	return &AdminUsecase{repo: repo}
}

func (u *AdminUsecase) DaftarPengguna(ctx context.Context) ([]domain.Pengguna, error) {
	return u.repo.DaftarPengguna(ctx)
}

func (u *AdminUsecase) PerbaruiPengguna(ctx context.Context, pengguna *domain.Pengguna) error {
	return u.repo.PerbaruiPengguna(ctx, pengguna)
}

func (u *AdminUsecase) HapusPengguna(ctx context.Context, id int64) error {
	return u.repo.HapusPengguna(ctx, id)
}

func (u *AdminUsecase) BuatKelas(ctx context.Context, kelas *domain.Kelas) error {
	return u.repo.BuatKelas(ctx, kelas)
}

func (u *AdminUsecase) PerbaruiKelas(ctx context.Context, kelas *domain.Kelas) error {
	return u.repo.PerbaruiKelas(ctx, kelas)
}

func (u *AdminUsecase) HapusKelas(ctx context.Context, id int64) error {
	return u.repo.HapusKelas(ctx, id)
}

func (u *AdminUsecase) BuatModul(ctx context.Context, modul *domain.Modul) error {
	return u.repo.BuatModul(ctx, modul)
}

func (u *AdminUsecase) PerbaruiModul(ctx context.Context, modul *domain.Modul) error {
	return u.repo.PerbaruiModul(ctx, modul)
}

func (u *AdminUsecase) HapusModul(ctx context.Context, id int64) error {
	return u.repo.HapusModul(ctx, id)
}

func (u *AdminUsecase) BuatMateri(ctx context.Context, materi *domain.Materi) error {
	return u.repo.BuatMateri(ctx, materi)
}

func (u *AdminUsecase) PerbaruiMateri(ctx context.Context, materi *domain.Materi) error {
	return u.repo.PerbaruiMateri(ctx, materi)
}

func (u *AdminUsecase) HapusMateri(ctx context.Context, id int64) error {
	return u.repo.HapusMateri(ctx, id)
}

func (u *AdminUsecase) BuatUjian(ctx context.Context, ujian *domain.Ujian) error {
	return u.repo.BuatUjian(ctx, ujian)
}

func (u *AdminUsecase) PerbaruiUjian(ctx context.Context, ujian *domain.Ujian) error {
	return u.repo.PerbaruiUjian(ctx, ujian)
}

func (u *AdminUsecase) HapusUjian(ctx context.Context, id int64) error {
	return u.repo.HapusUjian(ctx, id)
}

func (u *AdminUsecase) BuatSertifikat(ctx context.Context, sertifikat *domain.Sertifikat) error {
	return u.repo.BuatSertifikat(ctx, sertifikat)
}

func (u *AdminUsecase) HapusSertifikat(ctx context.Context, id int64) error {
	return u.repo.HapusSertifikat(ctx, id)
}

func (u *AdminUsecase) BuatPustaka(ctx context.Context, pustaka *domain.Pustaka) error {
	return u.repo.BuatPustaka(ctx, pustaka)
}

func (u *AdminUsecase) PerbaruiPustaka(ctx context.Context, pustaka *domain.Pustaka) error {
	return u.repo.PerbaruiPustaka(ctx, pustaka)
}

func (u *AdminUsecase) HapusPustaka(ctx context.Context, id int64) error {
	return u.repo.HapusPustaka(ctx, id)
}

func (u *AdminUsecase) BuatBerita(ctx context.Context, berita *domain.Berita) error {
	return u.repo.BuatBerita(ctx, berita)
}

func (u *AdminUsecase) PerbaruiBerita(ctx context.Context, berita *domain.Berita) error {
	return u.repo.PerbaruiBerita(ctx, berita)
}

func (u *AdminUsecase) HapusBerita(ctx context.Context, id int64) error {
	return u.repo.HapusBerita(ctx, id)
}

func (u *AdminUsecase) BuatScreener(ctx context.Context, screener *domain.Screener) error {
	return u.repo.BuatScreener(ctx, screener)
}

func (u *AdminUsecase) PerbaruiScreener(ctx context.Context, screener *domain.Screener) error {
	return u.repo.PerbaruiScreener(ctx, screener)
}

func (u *AdminUsecase) HapusScreener(ctx context.Context, id int64) error {
	return u.repo.HapusScreener(ctx, id)
}

func (u *AdminUsecase) BuatPasar(ctx context.Context, pasar *domain.Pasar) error {
	return u.repo.BuatPasar(ctx, pasar)
}

func (u *AdminUsecase) PerbaruiPasar(ctx context.Context, pasar *domain.Pasar) error {
	return u.repo.PerbaruiPasar(ctx, pasar)
}

func (u *AdminUsecase) HapusPasar(ctx context.Context, id int64) error {
	return u.repo.HapusPasar(ctx, id)
}

func (u *AdminUsecase) BuatReels(ctx context.Context, reels *domain.Reels) error {
	return u.repo.BuatReels(ctx, reels)
}

func (u *AdminUsecase) PerbaruiReels(ctx context.Context, reels *domain.Reels) error {
	return u.repo.PerbaruiReels(ctx, reels)
}

func (u *AdminUsecase) HapusReels(ctx context.Context, id int64) error {
	return u.repo.HapusReels(ctx, id)
}

func (u *AdminUsecase) BuatTadabbur(ctx context.Context, tadabbur *domain.Tadabbur) error {
	return u.repo.BuatTadabbur(ctx, tadabbur)
}

func (u *AdminUsecase) PerbaruiTadabbur(ctx context.Context, tadabbur *domain.Tadabbur) error {
	return u.repo.PerbaruiTadabbur(ctx, tadabbur)
}

func (u *AdminUsecase) HapusTadabbur(ctx context.Context, id int64) error {
	return u.repo.HapusTadabbur(ctx, id)
}

func (u *AdminUsecase) BuatKonfigurasi(ctx context.Context, konfigurasi *domain.Konfigurasi) error {
	return u.repo.BuatKonfigurasi(ctx, konfigurasi)
}

func (u *AdminUsecase) PerbaruiKonfigurasi(ctx context.Context, konfigurasi *domain.Konfigurasi) error {
	return u.repo.PerbaruiKonfigurasi(ctx, konfigurasi)
}

func (u *AdminUsecase) HapusKonfigurasi(ctx context.Context, id int64) error {
	return u.repo.HapusKonfigurasi(ctx, id)
}

func (u *AdminUsecase) DaftarKonfigurasi(ctx context.Context) ([]domain.Konfigurasi, error) {
	return u.repo.DaftarKonfigurasi(ctx)
}

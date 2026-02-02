package domain

import "context"

type AuthRepository interface {
	BuatPengguna(ctx context.Context, pengguna *Pengguna) (int64, error)
	CariPenggunaByEmail(ctx context.Context, email string) (*Pengguna, error)
	AmbilPenggunaByID(ctx context.Context, id int64) (*Pengguna, error)
	TandaiPenggunaTerverifikasi(ctx context.Context, id int64) error
	SimpanOTP(ctx context.Context, otp *OTPVerifikasi) error
	AmbilOTPByPengguna(ctx context.Context, idPengguna int64) (*OTPVerifikasi, error)
	PerbaruiOTP(ctx context.Context, otp *OTPVerifikasi) error
}

type ScreenerRepository interface {
	DaftarScreener(ctx context.Context, kategori, cari string) ([]Screener, error)
	DetailScreener(ctx context.Context, id int64) (*Screener, error)
	DaftarCatatanScreener(ctx context.Context, idScreener int64) ([]ScreenerCatatan, error)
	DaftarPasar(ctx context.Context) ([]Pasar, error)
}

type EdukasiRepository interface {
	DaftarKelas(ctx context.Context) ([]Kelas, error)
	DetailKelas(ctx context.Context, id int64) (*Kelas, error)
	DaftarModulByKelas(ctx context.Context, idKelas int64) ([]Modul, error)
	DaftarMateriByModul(ctx context.Context, idModul int64) ([]Materi, error)
	DaftarUjianByKelas(ctx context.Context, idKelas int64) ([]Ujian, error)
	DaftarModulSemua(ctx context.Context) ([]Modul, error)
	DaftarMateriSemua(ctx context.Context) ([]Materi, error)
	DaftarUjianSemua(ctx context.Context) ([]Ujian, error)
	DaftarSertifikat(ctx context.Context) ([]Sertifikat, error)
	SimpanProgress(ctx context.Context, progress *ProgressKelas) error
	DaftarProgress(ctx context.Context, idPengguna int64) ([]ProgressKelas, error)
}

type PustakaRepository interface {
	DaftarPustaka(ctx context.Context) ([]Pustaka, error)
	DetailPustaka(ctx context.Context, id int64) (*Pustaka, error)
}

type BeritaRepository interface {
	DaftarBerita(ctx context.Context, limit int) ([]Berita, error)
	DaftarBeritaTerbaru(ctx context.Context, limit int) ([]Berita, error)
}

type DiskusiRepository interface {
	DaftarDiskusi(ctx context.Context) ([]Diskusi, error)
	DetailDiskusi(ctx context.Context, id int64) (*Diskusi, error)
	DaftarBalasan(ctx context.Context, idDiskusi int64) ([]DiskusiBalas, error)
	BuatDiskusi(ctx context.Context, diskusi *Diskusi) error
	BuatBalasan(ctx context.Context, balas *DiskusiBalas) error
	BuatLaporan(ctx context.Context, laporan *DiskusiLaporan) error
}

type PortofolioRepository interface {
	DaftarPortofolio(ctx context.Context, idPengguna int64) ([]Portofolio, error)
	SimpanPortofolio(ctx context.Context, portofolio *Portofolio) error
	PerbaruiPortofolio(ctx context.Context, portofolio *Portofolio) error
	HapusPortofolio(ctx context.Context, id int64, idPengguna int64) error
}

type ZakatRepository interface {
	HargaEmasTerbaru(ctx context.Context) (*HargaEmas, error)
	DaftarRiwayatZakat(ctx context.Context, idPengguna int64) ([]ZakatRiwayat, error)
	SimpanRiwayatZakat(ctx context.Context, riwayat *ZakatRiwayat) error
}

type ReelsRepository interface {
	DaftarReels(ctx context.Context, tema string) ([]Reels, error)
	DetailReels(ctx context.Context, id int64) (*Reels, error)
}

type TadabburRepository interface {
	DaftarTadabbur(ctx context.Context, tema string) ([]Tadabbur, error)
}

type AdminRepository interface {
	DaftarPengguna(ctx context.Context) ([]Pengguna, error)
	PerbaruiPengguna(ctx context.Context, pengguna *Pengguna) error
	HapusPengguna(ctx context.Context, id int64) error

	BuatKelas(ctx context.Context, kelas *Kelas) error
	PerbaruiKelas(ctx context.Context, kelas *Kelas) error
	HapusKelas(ctx context.Context, id int64) error

	BuatModul(ctx context.Context, modul *Modul) error
	PerbaruiModul(ctx context.Context, modul *Modul) error
	HapusModul(ctx context.Context, id int64) error

	BuatMateri(ctx context.Context, materi *Materi) error
	PerbaruiMateri(ctx context.Context, materi *Materi) error
	HapusMateri(ctx context.Context, id int64) error

	BuatUjian(ctx context.Context, ujian *Ujian) error
	PerbaruiUjian(ctx context.Context, ujian *Ujian) error
	HapusUjian(ctx context.Context, id int64) error

	BuatSertifikat(ctx context.Context, sertifikat *Sertifikat) error
	HapusSertifikat(ctx context.Context, id int64) error

	BuatPustaka(ctx context.Context, pustaka *Pustaka) error
	PerbaruiPustaka(ctx context.Context, pustaka *Pustaka) error
	HapusPustaka(ctx context.Context, id int64) error

	BuatBerita(ctx context.Context, berita *Berita) error
	PerbaruiBerita(ctx context.Context, berita *Berita) error
	HapusBerita(ctx context.Context, id int64) error

	BuatScreener(ctx context.Context, screener *Screener) error
	PerbaruiScreener(ctx context.Context, screener *Screener) error
	HapusScreener(ctx context.Context, id int64) error

	BuatPasar(ctx context.Context, pasar *Pasar) error
	PerbaruiPasar(ctx context.Context, pasar *Pasar) error
	HapusPasar(ctx context.Context, id int64) error

	BuatReels(ctx context.Context, reels *Reels) error
	PerbaruiReels(ctx context.Context, reels *Reels) error
	HapusReels(ctx context.Context, id int64) error

	BuatTadabbur(ctx context.Context, tadabbur *Tadabbur) error
	PerbaruiTadabbur(ctx context.Context, tadabbur *Tadabbur) error
	HapusTadabbur(ctx context.Context, id int64) error

	BuatKonfigurasi(ctx context.Context, konfigurasi *Konfigurasi) error
	PerbaruiKonfigurasi(ctx context.Context, konfigurasi *Konfigurasi) error
	HapusKonfigurasi(ctx context.Context, id int64) error
	DaftarKonfigurasi(ctx context.Context) ([]Konfigurasi, error)
}

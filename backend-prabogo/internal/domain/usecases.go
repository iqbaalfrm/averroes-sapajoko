package domain

import "context"

// UserUseCase defines the interface for user-related business logic
type UserUseCase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email, password string) (*User, error)
	GetProfile(ctx context.Context, userID string) (*User, error)
	UpdateProfile(ctx context.Context, user *User) error
	GetAllUsers(ctx context.Context, offset, limit int) ([]*User, error)
}

// KelasUseCase defines the interface for course-related business logic
type KelasUseCase interface {
	CreateKelas(ctx context.Context, kelas *Kelas) error
	GetKelasByID(ctx context.Context, id string) (*Kelas, error)
	GetAllKelas(ctx context.Context, offset, limit int) ([]*Kelas, error)
	UpdateKelas(ctx context.Context, kelas *Kelas) error
	DeleteKelas(ctx context.Context, id string) error
}

// ModulUseCase defines the interface for module-related business logic
type ModulUseCase interface {
	CreateModul(ctx context.Context, modul *Modul) error
	GetModulByKelasID(ctx context.Context, kelasID string) ([]*Modul, error)
	GetModulByID(ctx context.Context, id string) (*Modul, error)
	UpdateModul(ctx context.Context, modul *Modul) error
	DeleteModul(ctx context.Context, id string) error
}

// MateriUseCase defines the interface for material-related business logic
type MateriUseCase interface {
	CreateMateri(ctx context.Context, materi *Materi) error
	GetMateriByModulID(ctx context.Context, modulID string) ([]*Materi, error)
	GetMateriByID(ctx context.Context, id string) (*Materi, error)
	UpdateMateri(ctx context.Context, materi *Materi) error
	DeleteMateri(ctx context.Context, id string) error
}

// UjianUseCase defines the interface for exam-related business logic
type UjianUseCase interface {
	CreateUjian(ctx context.Context, ujian *Ujian) error
	GetUjianByID(ctx context.Context, id string) (*Ujian, error)
	GetUjianByKelasID(ctx context.Context, kelasID string) ([]*Ujian, error)
	GetUjianByModulID(ctx context.Context, modulID string) ([]*Ujian, error)
	UpdateUjian(ctx context.Context, ujian *Ujian) error
	DeleteUjian(ctx context.Context, id string) error
}

// SoalUseCase defines the interface for question-related business logic
type SoalUseCase interface {
	CreateSoal(ctx context.Context, soal *Soal) error
	GetSoalByUjianID(ctx context.Context, ujianID string) ([]*Soal, error)
	GetSoalByID(ctx context.Context, id string) (*Soal, error)
	UpdateSoal(ctx context.Context, soal *Soal) error
	DeleteSoal(ctx context.Context, id string) error
}

// HasilUjianUseCase defines the interface for exam result-related business logic
type HasilUjianUseCase interface {
	CreateHasilUjian(ctx context.Context, hasil *HasilUjian) error
	GetHasilUjianByUserID(ctx context.Context, userID string) ([]*HasilUjian, error)
	GetHasilUjianByUjianID(ctx context.Context, ujianID string) ([]*HasilUjian, error)
	GetHasilUjianByID(ctx context.Context, id string) (*HasilUjian, error)
	UpdateHasilUjian(ctx context.Context, hasil *HasilUjian) error
	DeleteHasilUjian(ctx context.Context, id string) error
}

// SertifikatUseCase defines the interface for certificate-related business logic
type SertifikatUseCase interface {
	CreateSertifikat(ctx context.Context, sertifikat *Sertifikat) error
	GetSertifikatByUserID(ctx context.Context, userID string) ([]*Sertifikat, error)
	GetSertifikatByKelasID(ctx context.Context, kelasID string) ([]*Sertifikat, error)
	GetSertifikatByID(ctx context.Context, id string) (*Sertifikat, error)
	UpdateSertifikat(ctx context.Context, sertifikat *Sertifikat) error
	DeleteSertifikat(ctx context.Context, id string) error
}

// BukuUseCase defines the interface for book-related business logic
type BukuUseCase interface {
	CreateBuku(ctx context.Context, buku *Buku) error
	GetBukuByID(ctx context.Context, id string) (*Buku, error)
	GetAllBuku(ctx context.Context, offset, limit int) ([]*Buku, error)
	UpdateBuku(ctx context.Context, buku *Buku) error
	DeleteBuku(ctx context.Context, id string) error
}

// BeritaUseCase defines the interface for news-related business logic
type BeritaUseCase interface {
	CreateBerita(ctx context.Context, berita *Berita) error
	GetBeritaByID(ctx context.Context, id string) (*Berita, error)
	GetAllBerita(ctx context.Context, offset, limit int) ([]*Berita, error)
	GetLatestBerita(ctx context.Context, limit int) ([]*Berita, error)
	UpdateBerita(ctx context.Context, berita *Berita) error
	DeleteBerita(ctx context.Context, id string) error
}

// DiskusiThreadUseCase defines the interface for discussion thread-related business logic
type DiskusiThreadUseCase interface {
	CreateThread(ctx context.Context, thread *DiskusiThread) error
	GetThreadByID(ctx context.Context, id string) (*DiskusiThread, error)
	GetAllThreads(ctx context.Context, offset, limit int) ([]*DiskusiThread, error)
	GetThreadByUserID(ctx context.Context, userID string) ([]*DiskusiThread, error)
	UpdateThread(ctx context.Context, thread *DiskusiThread) error
	DeleteThread(ctx context.Context, id string) error
}

// DiskusiKomentarUseCase defines the interface for discussion comment-related business logic
type DiskusiKomentarUseCase interface {
	CreateKomentar(ctx context.Context, komentar *DiskusiKomentar) error
	GetKomentarByThreadID(ctx context.Context, threadID string) ([]*DiskusiKomentar, error)
	GetKomentarByID(ctx context.Context, id string) (*DiskusiKomentar, error)
	UpdateKomentar(ctx context.Context, komentar *DiskusiKomentar) error
	DeleteKomentar(ctx context.Context, id string) error
}

// KonfigurasiAplikasiUseCase defines the interface for application configuration-related business logic
type KonfigurasiAplikasiUseCase interface {
	GetKonfigurasi(ctx context.Context) (*KonfigurasiAplikasi, error)
	UpdateKonfigurasi(ctx context.Context, config *KonfigurasiAplikasi) error
}
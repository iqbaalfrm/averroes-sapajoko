package domain

import "context"

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, offset, limit int) ([]*User, error)
}

// KelasRepository defines the interface for course-related database operations
type KelasRepository interface {
	Create(ctx context.Context, kelas *Kelas) error
	GetByID(ctx context.Context, id string) (*Kelas, error)
	GetAll(ctx context.Context, offset, limit int) ([]*Kelas, error)
	Update(ctx context.Context, kelas *Kelas) error
	Delete(ctx context.Context, id string) error
}

// ModulRepository defines the interface for module-related database operations
type ModulRepository interface {
	Create(ctx context.Context, modul *Modul) error
	GetByKelasID(ctx context.Context, kelasID string) ([]*Modul, error)
	GetByID(ctx context.Context, id string) (*Modul, error)
	Update(ctx context.Context, modul *Modul) error
	Delete(ctx context.Context, id string) error
}

// MateriRepository defines the interface for material-related database operations
type MateriRepository interface {
	Create(ctx context.Context, materi *Materi) error
	GetByModulID(ctx context.Context, modulID string) ([]*Materi, error)
	GetByID(ctx context.Context, id string) (*Materi, error)
	Update(ctx context.Context, materi *Materi) error
	Delete(ctx context.Context, id string) error
}

// UjianRepository defines the interface for exam-related database operations
type UjianRepository interface {
	Create(ctx context.Context, ujian *Ujian) error
	GetByID(ctx context.Context, id string) (*Ujian, error)
	GetByKelasID(ctx context.Context, kelasID string) ([]*Ujian, error)
	GetByModulID(ctx context.Context, modulID string) ([]*Ujian, error)
	Update(ctx context.Context, ujian *Ujian) error
	Delete(ctx context.Context, id string) error
}

// SoalRepository defines the interface for question-related database operations
type SoalRepository interface {
	Create(ctx context.Context, soal *Soal) error
	GetByUjianID(ctx context.Context, ujianID string) ([]*Soal, error)
	GetByID(ctx context.Context, id string) (*Soal, error)
	Update(ctx context.Context, soal *Soal) error
	Delete(ctx context.Context, id string) error
}

// HasilUjianRepository defines the interface for exam result-related database operations
type HasilUjianRepository interface {
	Create(ctx context.Context, hasil *HasilUjian) error
	GetByUserID(ctx context.Context, userID string) ([]*HasilUjian, error)
	GetByUjianID(ctx context.Context, ujianID string) ([]*HasilUjian, error)
	GetByID(ctx context.Context, id string) (*HasilUjian, error)
	Update(ctx context.Context, hasil *HasilUjian) error
	Delete(ctx context.Context, id string) error
}

// SertifikatRepository defines the interface for certificate-related database operations
type SertifikatRepository interface {
	Create(ctx context.Context, sertifikat *Sertifikat) error
	GetByUserID(ctx context.Context, userID string) ([]*Sertifikat, error)
	GetByKelasID(ctx context.Context, kelasID string) ([]*Sertifikat, error)
	GetByID(ctx context.Context, id string) (*Sertifikat, error)
	Update(ctx context.Context, sertifikat *Sertifikat) error
	Delete(ctx context.Context, id string) error
}

// BukuRepository defines the interface for book-related database operations
type BukuRepository interface {
	Create(ctx context.Context, buku *Buku) error
	GetByID(ctx context.Context, id string) (*Buku, error)
	GetAll(ctx context.Context, offset, limit int) ([]*Buku, error)
	Update(ctx context.Context, buku *Buku) error
	Delete(ctx context.Context, id string) error
}

// BeritaRepository defines the interface for news-related database operations
type BeritaRepository interface {
	Create(ctx context.Context, berita *Berita) error
	GetByID(ctx context.Context, id string) (*Berita, error)
	GetAll(ctx context.Context, offset, limit int) ([]*Berita, error)
	GetLatest(ctx context.Context, limit int) ([]*Berita, error)
	Update(ctx context.Context, berita *Berita) error
	Delete(ctx context.Context, id string) error
}

// DiskusiThreadRepository defines the interface for discussion thread-related database operations
type DiskusiThreadRepository interface {
	Create(ctx context.Context, thread *DiskusiThread) error
	GetByID(ctx context.Context, id string) (*DiskusiThread, error)
	GetAll(ctx context.Context, offset, limit int) ([]*DiskusiThread, error)
	GetByUserID(ctx context.Context, userID string) ([]*DiskusiThread, error)
	Update(ctx context.Context, thread *DiskusiThread) error
	Delete(ctx context.Context, id string) error
}

// DiskusiKomentarRepository defines the interface for discussion comment-related database operations
type DiskusiKomentarRepository interface {
	Create(ctx context.Context, komentar *DiskusiKomentar) error
	GetByThreadID(ctx context.Context, threadID string) ([]*DiskusiKomentar, error)
	GetByID(ctx context.Context, id string) (*DiskusiKomentar, error)
	Update(ctx context.Context, komentar *DiskusiKomentar) error
	Delete(ctx context.Context, id string) error
}

// KonfigurasiAplikasiRepository defines the interface for application configuration-related database operations
type KonfigurasiAplikasiRepository interface {
	Get(ctx context.Context) (*KonfigurasiAplikasi, error)
	Update(ctx context.Context, config *KonfigurasiAplikasi) error
}
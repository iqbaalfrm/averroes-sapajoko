package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/averroes/backend-prabogo/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCaseImpl implements the UserUseCase interface
type UserUseCaseImpl struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new user use case instance
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &UserUseCaseImpl{
		userRepo: userRepo,
	}
}

func (uc *UserUseCaseImpl) Register(ctx context.Context, user *domain.User) error {
	// Check if user with email already exists
	existingUser, err := uc.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email sudah terdaftar")
	}
	
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("gagal mengenkripsi kata sandi: %v", err)
	}
	user.KataSandi = string(hashedPassword)
	
	// Set default role
	if user.Peran == "" {
		user.Peran = "user" // default role
	}
	
	// Set status to active
	user.Status = "aktif"
	
	return uc.userRepo.Create(ctx, user)
}

func (uc *UserUseCaseImpl) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("email atau kata sandi salah")
	}
	
	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(password)); err != nil {
		return nil, errors.New("email atau kata sandi salah")
	}
	
	// Check if user is active
	if user.Status != "aktif" {
		return nil, errors.New("akun tidak aktif")
	}
	
	return user, nil
}

func (uc *UserUseCaseImpl) GetProfile(ctx context.Context, userID string) (*domain.User, error) {
	return uc.userRepo.GetByID(ctx, userID)
}

func (uc *UserUseCaseImpl) UpdateProfile(ctx context.Context, user *domain.User) error {
	return uc.userRepo.Update(ctx, user)
}

func (uc *UserUseCaseImpl) GetAllUsers(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	return uc.userRepo.GetAll(ctx, offset, limit)
}

// KelasUseCaseImpl implements the KelasUseCase interface
type KelasUseCaseImpl struct {
	kelasRepo domain.KelasRepository
}

// NewKelasUseCase creates a new kelas use case instance
func NewKelasUseCase(kelasRepo domain.KelasRepository) domain.KelasUseCase {
	return &KelasUseCaseImpl{
		kelasRepo: kelasRepo,
	}
}

func (uc *KelasUseCaseImpl) CreateKelas(ctx context.Context, kelas *domain.Kelas) error {
	// Validate required fields
	if strings.TrimSpace(kelas.Judul) == "" {
		return errors.New("judul kelas wajib diisi")
	}
	
	if kelas.Harga < 0 {
		return errors.New("harga tidak boleh negatif")
	}
	
	// Set default status to draft
	if kelas.Status == "" {
		kelas.Status = "draft"
	}
	
	return uc.kelasRepo.Create(ctx, kelas)
}

func (uc *KelasUseCaseImpl) GetKelasByID(ctx context.Context, id string) (*domain.Kelas, error) {
	return uc.kelasRepo.GetByID(ctx, id)
}

func (uc *KelasUseCaseImpl) GetAllKelas(ctx context.Context, offset, limit int) ([]*domain.Kelas, error) {
	return uc.kelasRepo.GetAll(ctx, offset, limit)
}

func (uc *KelasUseCaseImpl) UpdateKelas(ctx context.Context, kelas *domain.Kelas) error {
	// Validate required fields
	if strings.TrimSpace(kelas.Judul) == "" {
		return errors.New("judul kelas wajib diisi")
	}
	
	if kelas.Harga < 0 {
		return errors.New("harga tidak boleh negatif")
	}
	
	return uc.kelasRepo.Update(ctx, kelas)
}

func (uc *KelasUseCaseImpl) DeleteKelas(ctx context.Context, id string) error {
	// In a real implementation, we might want to check if the class has associated modules, etc.
	return uc.kelasRepo.Delete(ctx, id)
}

// BukuUseCaseImpl implements the BukuUseCase interface
type BukuUseCaseImpl struct {
	bukuRepo domain.BukuRepository
}

// NewBukuUseCase creates a new buku use case instance
func NewBukuUseCase(bukuRepo domain.BukuRepository) domain.BukuUseCase {
	return &BukuUseCaseImpl{
		bukuRepo: bukuRepo,
	}
}

func (uc *BukuUseCaseImpl) CreateBuku(ctx context.Context, buku *domain.Buku) error {
	// Validate required fields
	if strings.TrimSpace(buku.JudulTampil) == "" {
		return errors.New("judul tampil wajib diisi")
	}
	
	if strings.TrimSpace(buku.JudulAsli) == "" {
		return errors.New("judul asli wajib diisi")
	}
	
	if strings.TrimSpace(buku.Penulis) == "" {
		return errors.New("penulis wajib diisi")
	}
	
	return uc.bukuRepo.Create(ctx, buku)
}

func (uc *BukuUseCaseImpl) GetBukuByID(ctx context.Context, id string) (*domain.Buku, error) {
	return uc.bukuRepo.GetByID(ctx, id)
}

func (uc *BukuUseCaseImpl) GetAllBuku(ctx context.Context, offset, limit int) ([]*domain.Buku, error) {
	return uc.bukuRepo.GetAll(ctx, offset, limit)
}

func (uc *BukuUseCaseImpl) UpdateBuku(ctx context.Context, buku *domain.Buku) error {
	// Validate required fields
	if strings.TrimSpace(buku.JudulTampil) == "" {
		return errors.New("judul tampil wajib diisi")
	}
	
	if strings.TrimSpace(buku.JudulAsli) == "" {
		return errors.New("judul asli wajib diisi")
	}
	
	if strings.TrimSpace(buku.Penulis) == "" {
		return errors.New("penulis wajib diisi")
	}
	
	return uc.bukuRepo.Update(ctx, buku)
}

func (uc *BukuUseCaseImpl) DeleteBuku(ctx context.Context, id string) error {
	return uc.bukuRepo.Delete(ctx, id)
}

// BeritaUseCaseImpl implements the BeritaUseCase interface
type BeritaUseCaseImpl struct {
	beritaRepo domain.BeritaRepository
}

// NewBeritaUseCase creates a new berita use case instance
func NewBeritaUseCase(beritaRepo domain.BeritaRepository) domain.BeritaUseCase {
	return &BeritaUseCaseImpl{
		beritaRepo: beritaRepo,
	}
}

func (uc *BeritaUseCaseImpl) CreateBerita(ctx context.Context, berita *domain.Berita) error {
	// Validate required fields
	if strings.TrimSpace(berita.Judul) == "" {
		return errors.New("judul berita wajib diisi")
	}
	
	if strings.TrimSpace(berita.Isi) == "" {
		return errors.New("isi berita wajib diisi")
	}
	
	// Set default status to draft
	if berita.Status == "" {
		berita.Status = "draft"
	}
	
	return uc.beritaRepo.Create(ctx, berita)
}

func (uc *BeritaUseCaseImpl) GetBeritaByID(ctx context.Context, id string) (*domain.Berita, error) {
	return uc.beritaRepo.GetByID(ctx, id)
}

func (uc *BeritaUseCaseImpl) GetAllBerita(ctx context.Context, offset, limit int) ([]*domain.Berita, error) {
	return uc.beritaRepo.GetAll(ctx, offset, limit)
}

func (uc *BeritaUseCaseImpl) GetLatestBerita(ctx context.Context, limit int) ([]*domain.Berita, error) {
	// For demo purposes, we'll just get all and sort by date
	// In a real implementation, the repository would handle this
	allBerita, err := uc.beritaRepo.GetAll(ctx, 0, 100) // Get max 100 for demo
	if err != nil {
		return nil, err
	}
	
	// Sort by date descending (most recent first)
	// For simplicity in this mock, we'll just return the first 'limit' items
	if len(allBerita) > limit {
		allBerita = allBerita[:limit]
	}
	
	return allBerita, nil
}

func (uc *BeritaUseCaseImpl) UpdateBerita(ctx context.Context, berita *domain.Berita) error {
	// Validate required fields
	if strings.TrimSpace(berita.Judul) == "" {
		return errors.New("judul berita wajib diisi")
	}
	
	if strings.TrimSpace(berita.Isi) == "" {
		return errors.New("isi berita wajib diisi")
	}
	
	return uc.beritaRepo.Update(ctx, berita)
}

func (uc *BeritaUseCaseImpl) DeleteBerita(ctx context.Context, id string) error {
	return uc.beritaRepo.Delete(ctx, id)
}

// KonfigurasiAplikasiUseCaseImpl implements the KonfigurasiAplikasiUseCase interface
type KonfigurasiAplikasiUseCaseImpl struct {
	configRepo domain.KonfigurasiAplikasiRepository
}

// NewKonfigurasiAplikasiUseCase creates a new konfigurasi aplikasi use case instance
func NewKonfigurasiAplikasiUseCase(configRepo domain.KonfigurasiAplikasiRepository) domain.KonfigurasiAplikasiUseCase {
	return &KonfigurasiAplikasiUseCaseImpl{
		configRepo: configRepo,
	}
}

func (uc *KonfigurasiAplikasiUseCaseImpl) GetKonfigurasi(ctx context.Context) (*domain.KonfigurasiAplikasi, error) {
	return uc.configRepo.Get(ctx)
}

func (uc *KonfigurasiAplikasiUseCaseImpl) UpdateKonfigurasi(ctx context.Context, config *domain.KonfigurasiAplikasi) error {
	// Validate required fields
	if strings.TrimSpace(config.NamaAplikasi) == "" {
		return errors.New("nama aplikasi wajib diisi")
	}
	
	return uc.configRepo.Update(ctx, config)
}
package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/averroes/backend-prabogo/internal/domain"
)

// ModulUseCaseImpl implements the ModulUseCase interface
type ModulUseCaseImpl struct {
	modulRepo domain.ModulRepository
	kelasRepo domain.KelasRepository
}

// NewModulUseCase creates a new modul use case instance
func NewModulUseCase(modulRepo domain.ModulRepository, kelasRepo domain.KelasRepository) domain.ModulUseCase {
	return &ModulUseCaseImpl{
		modulRepo: modulRepo,
		kelasRepo: kelasRepo,
	}
}

func (uc *ModulUseCaseImpl) CreateModul(ctx context.Context, modul *domain.Modul) error {
	// Validate required fields
	if strings.TrimSpace(modul.Judul) == "" {
		return errors.New("judul modul wajib diisi")
	}
	
	if strings.TrimSpace(modul.IDKelas) == "" {
		return errors.New("id kelas wajib diisi")
	}
	
	// Check if kelas exists
	_, err := uc.kelasRepo.GetByID(ctx, modul.IDKelas)
	if err != nil {
		return errors.New("kelas tidak ditemukan")
	}
	
	return uc.modulRepo.Create(ctx, modul)
}

func (uc *ModulUseCaseImpl) GetModulByKelasID(ctx context.Context, kelasID string) ([]*domain.Modul, error) {
	return uc.modulRepo.GetByKelasID(ctx, kelasID)
}

func (uc *ModulUseCaseImpl) GetModulByID(ctx context.Context, id string) (*domain.Modul, error) {
	return uc.modulRepo.GetByID(ctx, id)
}

func (uc *ModulUseCaseImpl) UpdateModul(ctx context.Context, modul *domain.Modul) error {
	// Validate required fields
	if strings.TrimSpace(modul.Judul) == "" {
		return errors.New("judul modul wajib diisi")
	}
	
	return uc.modulRepo.Update(ctx, modul)
}

func (uc *ModulUseCaseImpl) DeleteModul(ctx context.Context, id string) error {
	return uc.modulRepo.Delete(ctx, id)
}

// MateriUseCaseImpl implements the MateriUseCase interface
type MateriUseCaseImpl struct {
	materiRepo domain.MateriRepository
	modulRepo  domain.ModulRepository
}

// NewMateriUseCase creates a new materi use case instance
func NewMateriUseCase(materiRepo domain.MateriRepository, modulRepo domain.ModulRepository) domain.MateriUseCase {
	return &MateriUseCaseImpl{
		materiRepo: materiRepo,
		modulRepo:  modulRepo,
	}
}

func (uc *MateriUseCaseImpl) CreateMateri(ctx context.Context, materi *domain.Materi) error {
	// Validate required fields
	if strings.TrimSpace(materi.Judul) == "" {
		return errors.New("judul materi wajib diisi")
	}
	
	if strings.TrimSpace(materi.IDModul) == "" {
		return errors.New("id modul wajib diisi")
	}
	
	// Check if modul exists
	_, err := uc.modulRepo.GetByID(ctx, materi.IDModul)
	if err != nil {
		return errors.New("modul tidak ditemukan")
	}
	
	return uc.materiRepo.Create(ctx, materi)
}

func (uc *MateriUseCaseImpl) GetMateriByModulID(ctx context.Context, modulID string) ([]*domain.Materi, error) {
	return uc.materiRepo.GetByModulID(ctx, modulID)
}

func (uc *MateriUseCaseImpl) GetMateriByID(ctx context.Context, id string) (*domain.Materi, error) {
	return uc.materiRepo.GetByID(ctx, id)
}

func (uc *MateriUseCaseImpl) UpdateMateri(ctx context.Context, materi *domain.Materi) error {
	// Validate required fields
	if strings.TrimSpace(materi.Judul) == "" {
		return errors.New("judul materi wajib diisi")
	}
	
	return uc.materiRepo.Update(ctx, materi)
}

func (uc *MateriUseCaseImpl) DeleteMateri(ctx context.Context, id string) error {
	return uc.materiRepo.Delete(ctx, id)
}

// UjianUseCaseImpl implements the UjianUseCase interface
type UjianUseCaseImpl struct {
	ujianRepo domain.UjianRepository
	kelasRepo domain.KelasRepository
	modulRepo domain.ModulRepository
}

// NewUjianUseCase creates a new ujian use case instance
func NewUjianUseCase(ujianRepo domain.UjianRepository, kelasRepo domain.KelasRepository, modulRepo domain.ModulRepository) domain.UjianUseCase {
	return &UjianUseCaseImpl{
		ujianRepo: ujianRepo,
		kelasRepo: kelasRepo,
		modulRepo: modulRepo,
	}
}

func (uc *UjianUseCaseImpl) CreateUjian(ctx context.Context, ujian *domain.Ujian) error {
	// Validate required fields
	if strings.TrimSpace(ujian.Judul) == "" {
		return errors.New("judul ujian wajib diisi")
	}
	
	if ujian.IDKelas == "" && ujian.IDModul == "" {
		return errors.New("ujian harus terkait dengan kelas atau modul")
	}
	
	// Check if kelas exists (if provided)
	if ujian.IDKelas != "" {
		_, err := uc.kelasRepo.GetByID(ctx, ujian.IDKelas)
		if err != nil {
			return errors.New("kelas tidak ditemukan")
		}
	}
	
	// Check if modul exists (if provided)
	if ujian.IDModul != "" {
		_, err := uc.modulRepo.GetByID(ctx, ujian.IDModul)
		if err != nil {
			return errors.New("modul tidak ditemukan")
		}
	}
	
	return uc.ujianRepo.Create(ctx, ujian)
}

func (uc *UjianUseCaseImpl) GetUjianByID(ctx context.Context, id string) (*domain.Ujian, error) {
	return uc.ujianRepo.GetByID(ctx, id)
}

func (uc *UjianUseCaseImpl) GetUjianByKelasID(ctx context.Context, kelasID string) ([]*domain.Ujian, error) {
	return uc.ujianRepo.GetByKelasID(ctx, kelasID)
}

func (uc *UjianUseCaseImpl) GetUjianByModulID(ctx context.Context, modulID string) ([]*domain.Ujian, error) {
	return uc.ujianRepo.GetByModulID(ctx, modulID)
}

func (uc *UjianUseCaseImpl) UpdateUjian(ctx context.Context, ujian *domain.Ujian) error {
	// Validate required fields
	if strings.TrimSpace(ujian.Judul) == "" {
		return errors.New("judul ujian wajib diisi")
	}
	
	return uc.ujianRepo.Update(ctx, ujian)
}

func (uc *UjianUseCaseImpl) DeleteUjian(ctx context.Context, id string) error {
	return uc.ujianRepo.Delete(ctx, id)
}

// SoalUseCaseImpl implements the SoalUseCase interface
type SoalUseCaseImpl struct {
	soalRepo domain.SoalRepository
	ujianRepo domain.UjianRepository
}

// NewSoalUseCase creates a new soal use case instance
func NewSoalUseCase(soalRepo domain.SoalRepository, ujianRepo domain.UjianRepository) domain.SoalUseCase {
	return &SoalUseCaseImpl{
		soalRepo:  soalRepo,
		ujianRepo: ujianRepo,
	}
}

func (uc *SoalUseCaseImpl) CreateSoal(ctx context.Context, soal *domain.Soal) error {
	// Validate required fields
	if strings.TrimSpace(soal.Soal) == "" {
		return errors.New("soal wajib diisi")
	}
	
	if strings.TrimSpace(soal.IDUjian) == "" {
		return errors.New("id ujian wajib diisi")
	}
	
	// Check if ujian exists
	_, err := uc.ujianRepo.GetByID(ctx, soal.IDUjian)
	if err != nil {
		return errors.New("ujian tidak ditemukan")
	}
	
	return uc.soalRepo.Create(ctx, soal)
}

func (uc *SoalUseCaseImpl) GetSoalByUjianID(ctx context.Context, ujianID string) ([]*domain.Soal, error) {
	return uc.soalRepo.GetByUjianID(ctx, ujianID)
}

func (uc *SoalUseCaseImpl) GetSoalByID(ctx context.Context, id string) (*domain.Soal, error) {
	return uc.soalRepo.GetByID(ctx, id)
}

func (uc *SoalUseCaseImpl) UpdateSoal(ctx context.Context, soal *domain.Soal) error {
	// Validate required fields
	if strings.TrimSpace(soal.Soal) == "" {
		return errors.New("soal wajib diisi")
	}
	
	return uc.soalRepo.Update(ctx, soal)
}

func (uc *SoalUseCaseImpl) DeleteSoal(ctx context.Context, id string) error {
	return uc.soalRepo.Delete(ctx, id)
}

// HasilUjianUseCaseImpl implements the HasilUjianUseCase interface
type HasilUjianUseCaseImpl struct {
	hasilUjianRepo domain.HasilUjianRepository
	userRepo       domain.UserRepository
	ujianRepo      domain.UjianRepository
}

// NewHasilUjianUseCase creates a new hasil ujian use case instance
func NewHasilUjianUseCase(hasilUjianRepo domain.HasilUjianRepository, userRepo domain.UserRepository, ujianRepo domain.UjianRepository) domain.HasilUjianUseCase {
	return &HasilUjianUseCaseImpl{
		hasilUjianRepo: hasilUjianRepo,
		userRepo:       userRepo,
		ujianRepo:      ujianRepo,
	}
}

func (uc *HasilUjianUseCaseImpl) CreateHasilUjian(ctx context.Context, hasil *domain.HasilUjian) error {
	// Validate required fields
	if strings.TrimSpace(hasil.IDUser) == "" {
		return errors.New("id user wajib diisi")
	}
	
	if strings.TrimSpace(hasil.IDUjian) == "" {
		return errors.New("id ujian wajib diisi")
	}
	
	// Check if user exists
	_, err := uc.userRepo.GetByID(ctx, hasil.IDUser)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	
	// Check if ujian exists
	_, err = uc.ujianRepo.GetByID(ctx, hasil.IDUjian)
	if err != nil {
		return errors.New("ujian tidak ditemukan")
	}
	
	return uc.hasilUjianRepo.Create(ctx, hasil)
}

func (uc *HasilUjianUseCaseImpl) GetHasilUjianByUserID(ctx context.Context, userID string) ([]*domain.HasilUjian, error) {
	return uc.hasilUjianRepo.GetByUserID(ctx, userID)
}

func (uc *HasilUjianUseCaseImpl) GetHasilUjianByUjianID(ctx context.Context, ujianID string) ([]*domain.HasilUjian, error) {
	return uc.hasilUjianRepo.GetByUjianID(ctx, ujianID)
}

func (uc *HasilUjianUseCaseImpl) GetHasilUjianByID(ctx context.Context, id string) (*domain.HasilUjian, error) {
	return uc.hasilUjianRepo.GetByID(ctx, id)
}

func (uc *HasilUjianUseCaseImpl) UpdateHasilUjian(ctx context.Context, hasil *domain.HasilUjian) error {
	return uc.hasilUjianRepo.Update(ctx, hasil)
}

func (uc *HasilUjianUseCaseImpl) DeleteHasilUjian(ctx context.Context, id string) error {
	return uc.hasilUjianRepo.Delete(ctx, id)
}

// SertifikatUseCaseImpl implements the SertifikatUseCase interface
type SertifikatUseCaseImpl struct {
	sertifikatRepo domain.SertifikatRepository
	userRepo       domain.UserRepository
	kelasRepo      domain.KelasRepository
}

// NewSertifikatUseCase creates a new sertifikat use case instance
func NewSertifikatUseCase(sertifikatRepo domain.SertifikatRepository, userRepo domain.UserRepository, kelasRepo domain.KelasRepository) domain.SertifikatUseCase {
	return &SertifikatUseCaseImpl{
		sertifikatRepo: sertifikatRepo,
		userRepo:       userRepo,
		kelasRepo:      kelasRepo,
	}
}

func (uc *SertifikatUseCaseImpl) CreateSertifikat(ctx context.Context, sertifikat *domain.Sertifikat) error {
	// Validate required fields
	if strings.TrimSpace(sertifikat.IDUser) == "" {
		return errors.New("id user wajib diisi")
	}
	
	if strings.TrimSpace(sertifikat.IDKelas) == "" {
		return errors.New("id kelas wajib diisi")
	}
	
	// Check if user exists
	_, err := uc.userRepo.GetByID(ctx, sertifikat.IDUser)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	
	// Check if kelas exists
	_, err = uc.kelasRepo.GetByID(ctx, sertifikat.IDKelas)
	if err != nil {
		return errors.New("kelas tidak ditemukan")
	}
	
	// Generate certificate number if not provided
	if sertifikat.NomorSertifikat == "" {
		sertifikat.NomorSertifikat = fmt.Sprintf("CERT-%d", sertifikat.TanggalTerbit.Unix())
	}
	
	return uc.sertifikatRepo.Create(ctx, sertifikat)
}

func (uc *SertifikatUseCaseImpl) GetSertifikatByUserID(ctx context.Context, userID string) ([]*domain.Sertifikat, error) {
	return uc.sertifikatRepo.GetByUserID(ctx, userID)
}

func (uc *SertifikatUseCaseImpl) GetSertifikatByKelasID(ctx context.Context, kelasID string) ([]*domain.Sertifikat, error) {
	return uc.sertifikatRepo.GetByKelasID(ctx, kelasID)
}

func (uc *SertifikatUseCaseImpl) GetSertifikatByID(ctx context.Context, id string) (*domain.Sertifikat, error) {
	return uc.sertifikatRepo.GetByID(ctx, id)
}

func (uc *SertifikatUseCaseImpl) UpdateSertifikat(ctx context.Context, sertifikat *domain.Sertifikat) error {
	return uc.sertifikatRepo.Update(ctx, sertifikat)
}

func (uc *SertifikatUseCaseImpl) DeleteSertifikat(ctx context.Context, id string) error {
	return uc.sertifikatRepo.Delete(ctx, id)
}

// DiskusiThreadUseCaseImpl implements the DiskusiThreadUseCase interface
type DiskusiThreadUseCaseImpl struct {
	threadRepo domain.DiskusiThreadRepository
	userRepo   domain.UserRepository
}

// NewDiskusiThreadUseCase creates a new diskusi thread use case instance
func NewDiskusiThreadUseCase(threadRepo domain.DiskusiThreadRepository, userRepo domain.UserRepository) domain.DiskusiThreadUseCase {
	return &DiskusiThreadUseCaseImpl{
		threadRepo: threadRepo,
		userRepo:   userRepo,
	}
}

func (uc *DiskusiThreadUseCaseImpl) CreateThread(ctx context.Context, thread *domain.DiskusiThread) error {
	// Validate required fields
	if strings.TrimSpace(thread.Judul) == "" {
		return errors.New("judul thread wajib diisi")
	}
	
	if strings.TrimSpace(thread.Isi) == "" {
		return errors.New("isi thread wajib diisi")
	}
	
	if strings.TrimSpace(thread.IDUser) == "" {
		return errors.New("id user wajib diisi")
	}
	
	// Check if user exists
	_, err := uc.userRepo.GetByID(ctx, thread.IDUser)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	
	return uc.threadRepo.Create(ctx, thread)
}

func (uc *DiskusiThreadUseCaseImpl) GetThreadByID(ctx context.Context, id string) (*domain.DiskusiThread, error) {
	return uc.threadRepo.GetByID(ctx, id)
}

func (uc *DiskusiThreadUseCaseImpl) GetAllThreads(ctx context.Context, offset, limit int) ([]*domain.DiskusiThread, error) {
	return uc.threadRepo.GetAll(ctx, offset, limit)
}

func (uc *DiskusiThreadUseCaseImpl) GetThreadByUserID(ctx context.Context, userID string) ([]*domain.DiskusiThread, error) {
	return uc.threadRepo.GetByUserID(ctx, userID)
}

func (uc *DiskusiThreadUseCaseImpl) UpdateThread(ctx context.Context, thread *domain.DiskusiThread) error {
	// Validate required fields
	if strings.TrimSpace(thread.Judul) == "" {
		return errors.New("judul thread wajib diisi")
	}
	
	if strings.TrimSpace(thread.Isi) == "" {
		return errors.New("isi thread wajib diisi")
	}
	
	return uc.threadRepo.Update(ctx, thread)
}

func (uc *DiskusiThreadUseCaseImpl) DeleteThread(ctx context.Context, id string) error {
	return uc.threadRepo.Delete(ctx, id)
}

// DiskusiKomentarUseCaseImpl implements the DiskusiKomentarUseCase interface
type DiskusiKomentarUseCaseImpl struct {
	komentarRepo domain.DiskusiKomentarRepository
	threadRepo   domain.DiskusiThreadRepository
	userRepo     domain.UserRepository
}

// NewDiskusiKomentarUseCase creates a new diskusi komentar use case instance
func NewDiskusiKomentarUseCase(komentarRepo domain.DiskusiKomentarRepository, threadRepo domain.DiskusiThreadRepository, userRepo domain.UserRepository) domain.DiskusiKomentarUseCase {
	return &DiskusiKomentarUseCaseImpl{
		komentarRepo: komentarRepo,
		threadRepo:   threadRepo,
		userRepo:     userRepo,
	}
}

func (uc *DiskusiKomentarUseCaseImpl) CreateKomentar(ctx context.Context, komentar *domain.DiskusiKomentar) error {
	// Validate required fields
	if strings.TrimSpace(komentar.Isi) == "" {
		return errors.New("isi komentar wajib diisi")
	}
	
	if strings.TrimSpace(komentar.IDThread) == "" {
		return errors.New("id thread wajib diisi")
	}
	
	if strings.TrimSpace(komentar.IDUser) == "" {
		return errors.New("id user wajib diisi")
	}
	
	// Check if thread exists
	_, err := uc.threadRepo.GetByID(ctx, komentar.IDThread)
	if err != nil {
		return errors.New("thread tidak ditemukan")
	}
	
	// Check if user exists
	_, err = uc.userRepo.GetByID(ctx, komentar.IDUser)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	
	return uc.komentarRepo.Create(ctx, komentar)
}

func (uc *DiskusiKomentarUseCaseImpl) GetKomentarByThreadID(ctx context.Context, threadID string) ([]*domain.DiskusiKomentar, error) {
	return uc.komentarRepo.GetByThreadID(ctx, threadID)
}

func (uc *DiskusiKomentarUseCaseImpl) GetKomentarByID(ctx context.Context, id string) (*domain.DiskusiKomentar, error) {
	return uc.komentarRepo.GetByID(ctx, id)
}

func (uc *DiskusiKomentarUseCaseImpl) UpdateKomentar(ctx context.Context, komentar *domain.DiskusiKomentar) error {
	// Validate required fields
	if strings.TrimSpace(komentar.Isi) == "" {
		return errors.New("isi komentar wajib diisi")
	}
	
	return uc.komentarRepo.Update(ctx, komentar)
}

func (uc *DiskusiKomentarUseCaseImpl) DeleteKomentar(ctx context.Context, id string) error {
	return uc.komentarRepo.Delete(ctx, id)
}
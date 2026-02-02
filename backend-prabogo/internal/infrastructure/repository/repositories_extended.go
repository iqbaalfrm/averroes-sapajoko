package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
)

// ModulRepositoryImpl implements the ModulRepository interface
type ModulRepositoryImpl struct {
	db *MockSawitDB
}

// NewModulRepository creates a new modul repository instance
func NewModulRepository(db *MockSawitDB) *ModulRepositoryImpl {
	return &ModulRepositoryImpl{db: db}
}

func (r *ModulRepositoryImpl) Create(ctx context.Context, modul *domain.Modul) error {
	modul.ID = fmt.Sprintf("mdl_%d", time.Now().Unix())
	modul.TanggalBuat = time.Now()
	modul.TanggalUbah = time.Now()
	return r.db.Save("modul", modul.ID, modul)
}

func (r *ModulRepositoryImpl) GetByKelasID(ctx context.Context, kelasID string) ([]*domain.Modul, error) {
	allModulData, err := r.db.FindAll("modul", 0, 1000) // Assuming max 1000 modules for demo
	if err != nil {
		return nil, err
	}
	
	var modulList []*domain.Modul
	for _, modulData := range allModulData {
		modulMap, ok := modulData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if modulKelasID, ok := modulMap["id_kelas"].(string); ok && modulKelasID == kelasID {
			modul := &domain.Modul{}
			dataBytes, err := json.Marshal(modulData)
			if err != nil {
				log.Printf("Gagal mengkonversi data modul: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, modul); err != nil {
				log.Printf("Gagal menguraikan data modul: %v", err)
				continue
			}
			
			modulList = append(modulList, modul)
		}
	}
	
	return modulList, nil
}

func (r *ModulRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Modul, error) {
	data, err := r.db.FindByID("modul", id)
	if err != nil {
		return nil, err
	}
	
	modul := &domain.Modul{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, modul); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data modul: %v", err)
	}
	
	return modul, nil
}

func (r *ModulRepositoryImpl) Update(ctx context.Context, modul *domain.Modul) error {
	modul.TanggalUbah = time.Now()
	return r.db.Save("modul", modul.ID, modul)
}

func (r *ModulRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("modul", id)
}

// MateriRepositoryImpl implements the MateriRepository interface
type MateriRepositoryImpl struct {
	db *MockSawitDB
}

// NewMateriRepository creates a new materi repository instance
func NewMateriRepository(db *MockSawitDB) *MateriRepositoryImpl {
	return &MateriRepositoryImpl{db: db}
}

func (r *MateriRepositoryImpl) Create(ctx context.Context, materi *domain.Materi) error {
	materi.ID = fmt.Sprintf("mtr_%d", time.Now().Unix())
	materi.TanggalBuat = time.Now()
	materi.TanggalUbah = time.Now()
	return r.db.Save("materi", materi.ID, materi)
}

func (r *MateriRepositoryImpl) GetByModulID(ctx context.Context, modulID string) ([]*domain.Materi, error) {
	allMateriData, err := r.db.FindAll("materi", 0, 1000) // Assuming max 1000 materials for demo
	if err != nil {
		return nil, err
	}
	
	var materiList []*domain.Materi
	for _, materiData := range allMateriData {
		materiMap, ok := materiData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if materiModulID, ok := materiMap["id_modul"].(string); ok && materiModulID == modulID {
			materi := &domain.Materi{}
			dataBytes, err := json.Marshal(materiData)
			if err != nil {
				log.Printf("Gagal mengkonversi data materi: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, materi); err != nil {
				log.Printf("Gagal menguraikan data materi: %v", err)
				continue
			}
			
			materiList = append(materiList, materi)
		}
	}
	
	return materiList, nil
}

func (r *MateriRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Materi, error) {
	data, err := r.db.FindByID("materi", id)
	if err != nil {
		return nil, err
	}
	
	materi := &domain.Materi{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, materi); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data materi: %v", err)
	}
	
	return materi, nil
}

func (r *MateriRepositoryImpl) Update(ctx context.Context, materi *domain.Materi) error {
	materi.TanggalUbah = time.Now()
	return r.db.Save("materi", materi.ID, materi)
}

func (r *MateriRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("materi", id)
}

// UjianRepositoryImpl implements the UjianRepository interface
type UjianRepositoryImpl struct {
	db *MockSawitDB
}

// NewUjianRepository creates a new ujian repository instance
func NewUjianRepository(db *MockSawitDB) *UjianRepositoryImpl {
	return &UjianRepositoryImpl{db: db}
}

func (r *UjianRepositoryImpl) Create(ctx context.Context, ujian *domain.Ujian) error {
	ujian.ID = fmt.Sprintf("ujn_%d", time.Now().Unix())
	ujian.TanggalBuat = time.Now()
	ujian.TanggalUbah = time.Now()
	return r.db.Save("ujian", ujian.ID, ujian)
}

func (r *UjianRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Ujian, error) {
	data, err := r.db.FindByID("ujian", id)
	if err != nil {
		return nil, err
	}
	
	ujian := &domain.Ujian{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, ujian); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data ujian: %v", err)
	}
	
	return ujian, nil
}

func (r *UjianRepositoryImpl) GetByKelasID(ctx context.Context, kelasID string) ([]*domain.Ujian, error) {
	allUjianData, err := r.db.FindAll("ujian", 0, 1000) // Assuming max 1000 exams for demo
	if err != nil {
		return nil, err
	}
	
	var ujianList []*domain.Ujian
	for _, ujianData := range allUjianData {
		ujianMap, ok := ujianData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if ujianKelasID, ok := ujianMap["id_kelas"].(string); ok && ujianKelasID == kelasID {
			ujian := &domain.Ujian{}
			dataBytes, err := json.Marshal(ujianData)
			if err != nil {
				log.Printf("Gagal mengkonversi data ujian: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, ujian); err != nil {
				log.Printf("Gagal menguraikan data ujian: %v", err)
				continue
			}
			
			ujianList = append(ujianList, ujian)
		}
	}
	
	return ujianList, nil
}

func (r *UjianRepositoryImpl) GetByModulID(ctx context.Context, modulID string) ([]*domain.Ujian, error) {
	allUjianData, err := r.db.FindAll("ujian", 0, 1000) // Assuming max 1000 exams for demo
	if err != nil {
		return nil, err
	}
	
	var ujianList []*domain.Ujian
	for _, ujianData := range allUjianData {
		ujianMap, ok := ujianData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if ujianModulID, ok := ujianMap["id_modul"].(string); ok && ujianModulID == modulID {
			ujian := &domain.Ujian{}
			dataBytes, err := json.Marshal(ujianData)
			if err != nil {
				log.Printf("Gagal mengkonversi data ujian: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, ujian); err != nil {
				log.Printf("Gagal menguraikan data ujian: %v", err)
				continue
			}
			
			ujianList = append(ujianList, ujian)
		}
	}
	
	return ujianList, nil
}

func (r *UjianRepositoryImpl) Update(ctx context.Context, ujian *domain.Ujian) error {
	ujian.TanggalUbah = time.Now()
	return r.db.Save("ujian", ujian.ID, ujian)
}

func (r *UjianRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("ujian", id)
}

// SoalRepositoryImpl implements the SoalRepository interface
type SoalRepositoryImpl struct {
	db *MockSawitDB
}

// NewSoalRepository creates a new soal repository instance
func NewSoalRepository(db *MockSawitDB) *SoalRepositoryImpl {
	return &SoalRepositoryImpl{db: db}
}

func (r *SoalRepositoryImpl) Create(ctx context.Context, soal *domain.Soal) error {
	soal.ID = fmt.Sprintf("sql_%d", time.Now().Unix())
	return r.db.Save("soal", soal.ID, soal)
}

func (r *SoalRepositoryImpl) GetByUjianID(ctx context.Context, ujianID string) ([]*domain.Soal, error) {
	allSoalData, err := r.db.FindAll("soal", 0, 1000) // Assuming max 1000 questions for demo
	if err != nil {
		return nil, err
	}
	
	var soalList []*domain.Soal
	for _, soalData := range allSoalData {
		soalMap, ok := soalData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if soalUjianID, ok := soalMap["id_ujian"].(string); ok && soalUjianID == ujianID {
			soal := &domain.Soal{}
			dataBytes, err := json.Marshal(soalData)
			if err != nil {
				log.Printf("Gagal mengkonversi data soal: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, soal); err != nil {
				log.Printf("Gagal menguraikan data soal: %v", err)
				continue
			}
			
			soalList = append(soalList, soal)
		}
	}
	
	return soalList, nil
}

func (r *SoalRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Soal, error) {
	data, err := r.db.FindByID("soal", id)
	if err != nil {
		return nil, err
	}
	
	soal := &domain.Soal{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, soal); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data soal: %v", err)
	}
	
	return soal, nil
}

func (r *SoalRepositoryImpl) Update(ctx context.Context, soal *domain.Soal) error {
	return r.db.Save("soal", soal.ID, soal)
}

func (r *SoalRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("soal", id)
}

// HasilUjianRepositoryImpl implements the HasilUjianRepository interface
type HasilUjianRepositoryImpl struct {
	db *MockSawitDB
}

// NewHasilUjianRepository creates a new hasil ujian repository instance
func NewHasilUjianRepository(db *MockSawitDB) *HasilUjianRepositoryImpl {
	return &HasilUjianRepositoryImpl{db: db}
}

func (r *HasilUjianRepositoryImpl) Create(ctx context.Context, hasil *domain.HasilUjian) error {
	hasil.ID = fmt.Sprintf("hsu_%d", time.Now().Unix())
	hasil.TanggalUjian = time.Now()
	return r.db.Save("hasil_ujian", hasil.ID, hasil)
}

func (r *HasilUjianRepositoryImpl) GetByUserID(ctx context.Context, userID string) ([]*domain.HasilUjian, error) {
	allHasilData, err := r.db.FindAll("hasil_ujian", 0, 1000) // Assuming max 1000 results for demo
	if err != nil {
		return nil, err
	}
	
	var hasilList []*domain.HasilUjian
	for _, hasilData := range allHasilData {
		hasilMap, ok := hasilData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if hasilUserID, ok := hasilMap["id_user"].(string); ok && hasilUserID == userID {
			hasil := &domain.HasilUjian{}
			dataBytes, err := json.Marshal(hasilData)
			if err != nil {
				log.Printf("Gagal mengkonversi data hasil ujian: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, hasil); err != nil {
				log.Printf("Gagal menguraikan data hasil ujian: %v", err)
				continue
			}
			
			hasilList = append(hasilList, hasil)
		}
	}
	
	return hasilList, nil
}

func (r *HasilUjianRepositoryImpl) GetByUjianID(ctx context.Context, ujianID string) ([]*domain.HasilUjian, error) {
	allHasilData, err := r.db.FindAll("hasil_ujian", 0, 1000) // Assuming max 1000 results for demo
	if err != nil {
		return nil, err
	}
	
	var hasilList []*domain.HasilUjian
	for _, hasilData := range allHasilData {
		hasilMap, ok := hasilData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if hasilUjianID, ok := hasilMap["id_ujian"].(string); ok && hasilUjianID == ujianID {
			hasil := &domain.HasilUjian{}
			dataBytes, err := json.Marshal(hasilData)
			if err != nil {
				log.Printf("Gagal mengkonversi data hasil ujian: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, hasil); err != nil {
				log.Printf("Gagal menguraikan data hasil ujian: %v", err)
				continue
			}
			
			hasilList = append(hasilList, hasil)
		}
	}
	
	return hasilList, nil
}

func (r *HasilUjianRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.HasilUjian, error) {
	data, err := r.db.FindByID("hasil_ujian", id)
	if err != nil {
		return nil, err
	}
	
	hasil := &domain.HasilUjian{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, hasil); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data hasil ujian: %v", err)
	}
	
	return hasil, nil
}

func (r *HasilUjianRepositoryImpl) Update(ctx context.Context, hasil *domain.HasilUjian) error {
	hasil.TanggalUjian = time.Now()
	return r.db.Save("hasil_ujian", hasil.ID, hasil)
}

func (r *HasilUjianRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("hasil_ujian", id)
}

// SertifikatRepositoryImpl implements the SertifikatRepository interface
type SertifikatRepositoryImpl struct {
	db *MockSawitDB
}

// NewSertifikatRepository creates a new sertifikat repository instance
func NewSertifikatRepository(db *MockSawitDB) *SertifikatRepositoryImpl {
	return &SertifikatRepositoryImpl{db: db}
}

func (r *SertifikatRepositoryImpl) Create(ctx context.Context, sertifikat *domain.Sertifikat) error {
	sertifikat.ID = fmt.Sprintf("srt_%d", time.Now().Unix())
	sertifikat.TanggalTerbit = time.Now()
	sertifikat.TanggalBerlaku = time.Now().AddDate(1, 0, 0) // Valid for 1 year
	return r.db.Save("sertifikat", sertifikat.ID, sertifikat)
}

func (r *SertifikatRepositoryImpl) GetByUserID(ctx context.Context, userID string) ([]*domain.Sertifikat, error) {
	allSertifikatData, err := r.db.FindAll("sertifikat", 0, 1000) // Assuming max 1000 certificates for demo
	if err != nil {
		return nil, err
	}
	
	var sertifikatList []*domain.Sertifikat
	for _, sertifikatData := range allSertifikatData {
		sertifikatMap, ok := sertifikatData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if sertifikatUserID, ok := sertifikatMap["id_user"].(string); ok && sertifikatUserID == userID {
			sertifikat := &domain.Sertifikat{}
			dataBytes, err := json.Marshal(sertifikatData)
			if err != nil {
				log.Printf("Gagal mengkonversi data sertifikat: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, sertifikat); err != nil {
				log.Printf("Gagal menguraikan data sertifikat: %v", err)
				continue
			}
			
			sertifikatList = append(sertifikatList, sertifikat)
		}
	}
	
	return sertifikatList, nil
}

func (r *SertifikatRepositoryImpl) GetByKelasID(ctx context.Context, kelasID string) ([]*domain.Sertifikat, error) {
	allSertifikatData, err := r.db.FindAll("sertifikat", 0, 1000) // Assuming max 1000 certificates for demo
	if err != nil {
		return nil, err
	}
	
	var sertifikatList []*domain.Sertifikat
	for _, sertifikatData := range allSertifikatData {
		sertifikatMap, ok := sertifikatData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if sertifikatKelasID, ok := sertifikatMap["id_kelas"].(string); ok && sertifikatKelasID == kelasID {
			sertifikat := &domain.Sertifikat{}
			dataBytes, err := json.Marshal(sertifikatData)
			if err != nil {
				log.Printf("Gagal mengkonversi data sertifikat: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, sertifikat); err != nil {
				log.Printf("Gagal menguraikan data sertifikat: %v", err)
				continue
			}
			
			sertifikatList = append(sertifikatList, sertifikat)
		}
	}
	
	return sertifikatList, nil
}

func (r *SertifikatRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Sertifikat, error) {
	data, err := r.db.FindByID("sertifikat", id)
	if err != nil {
		return nil, err
	}
	
	sertifikat := &domain.Sertifikat{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, sertifikat); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data sertifikat: %v", err)
	}
	
	return sertifikat, nil
}

func (r *SertifikatRepositoryImpl) Update(ctx context.Context, sertifikat *domain.Sertifikat) error {
	sertifikat.TanggalBerlaku = time.Now().AddDate(1, 0, 0) // Reset validity to 1 year from now
	return r.db.Save("sertifikat", sertifikat.ID, sertifikat)
}

func (r *SertifikatRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("sertifikat", id)
}

// BeritaRepositoryImpl implements the BeritaRepository interface
type BeritaRepositoryImpl struct {
	db *MockSawitDB
}

// NewBeritaRepository creates a new berita repository instance
func NewBeritaRepository(db *MockSawitDB) *BeritaRepositoryImpl {
	return &BeritaRepositoryImpl{db: db}
}

func (r *BeritaRepositoryImpl) Create(ctx context.Context, berita *domain.Berita) error {
	berita.ID = fmt.Sprintf("brt_%d", time.Now().Unix())
	berita.TanggalBuat = time.Now()
	berita.TanggalUbah = time.Now()
	return r.db.Save("berita", berita.ID, berita)
}

func (r *BeritaRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Berita, error) {
	data, err := r.db.FindByID("berita", id)
	if err != nil {
		return nil, err
	}
	
	berita := &domain.Berita{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, berita); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data berita: %v", err)
	}
	
	return berita, nil
}

func (r *BeritaRepositoryImpl) GetAll(ctx context.Context, offset, limit int) ([]*domain.Berita, error) {
	beritaData, err := r.db.FindAll("berita", offset, limit)
	if err != nil {
		return nil, err
	}
	
	beritaList := make([]*domain.Berita, 0, len(beritaData))
	for _, beritaItem := range beritaData {
		berita := &domain.Berita{}
		dataBytes, err := json.Marshal(beritaItem)
		if err != nil {
			log.Printf("Gagal mengkonversi data berita: %v", err)
			continue
		}
		
		if err := json.Unmarshal(dataBytes, berita); err != nil {
			log.Printf("Gagal menguraikan data berita: %v", err)
			continue
		}
		
		beritaList = append(beritaList, berita)
	}
	
	return beritaList, nil
}

func (r *BeritaRepositoryImpl) GetLatest(ctx context.Context, limit int) ([]*domain.Berita, error) {
	allBerita, err := r.GetAll(ctx, 0, 1000) // Get all for sorting
	if err != nil {
		return nil, err
	}
	
	// Sort by date descending (most recent first)
	// For simplicity in this mock, we'll just return the first 'limit' items
	// In a real implementation, the database would handle this
	if len(allBerita) > limit {
		allBerita = allBerita[:limit]
	}
	
	return allBerita, nil
}

func (r *BeritaRepositoryImpl) Update(ctx context.Context, berita *domain.Berita) error {
	berita.TanggalUbah = time.Now()
	return r.db.Save("berita", berita.ID, berita)
}

func (r *BeritaRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("berita", id)
}

// DiskusiThreadRepositoryImpl implements the DiskusiThreadRepository interface
type DiskusiThreadRepositoryImpl struct {
	db *MockSawitDB
}

// NewDiskusiThreadRepository creates a new diskusi thread repository instance
func NewDiskusiThreadRepository(db *MockSawitDB) *DiskusiThreadRepositoryImpl {
	return &DiskusiThreadRepositoryImpl{db: db}
}

func (r *DiskusiThreadRepositoryImpl) Create(ctx context.Context, thread *domain.DiskusiThread) error {
	thread.ID = fmt.Sprintf("dst_%d", time.Now().Unix())
	thread.TanggalBuat = time.Now()
	thread.TanggalUbah = time.Now()
	return r.db.Save("diskusi_thread", thread.ID, thread)
}

func (r *DiskusiThreadRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.DiskusiThread, error) {
	data, err := r.db.FindByID("diskusi_thread", id)
	if err != nil {
		return nil, err
	}
	
	thread := &domain.DiskusiThread{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, thread); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data thread: %v", err)
	}
	
	return thread, nil
}

func (r *DiskusiThreadRepositoryImpl) GetAll(ctx context.Context, offset, limit int) ([]*domain.DiskusiThread, error) {
	threadData, err := r.db.FindAll("diskusi_thread", offset, limit)
	if err != nil {
		return nil, err
	}
	
	threadList := make([]*domain.DiskusiThread, 0, len(threadData))
	for _, threadItem := range threadData {
		thread := &domain.DiskusiThread{}
		dataBytes, err := json.Marshal(threadItem)
		if err != nil {
			log.Printf("Gagal mengkonversi data thread: %v", err)
			continue
		}
		
		if err := json.Unmarshal(dataBytes, thread); err != nil {
			log.Printf("Gagal menguraikan data thread: %v", err)
			continue
		}
		
		threadList = append(threadList, thread)
	}
	
	return threadList, nil
}

func (r *DiskusiThreadRepositoryImpl) GetByUserID(ctx context.Context, userID string) ([]*domain.DiskusiThread, error) {
	allThreadData, err := r.db.FindAll("diskusi_thread", 0, 1000) // Assuming max 1000 threads for demo
	if err != nil {
		return nil, err
	}
	
	var threadList []*domain.DiskusiThread
	for _, threadData := range allThreadData {
		threadMap, ok := threadData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if threadUserID, ok := threadMap["id_user"].(string); ok && threadUserID == userID {
			thread := &domain.DiskusiThread{}
			dataBytes, err := json.Marshal(threadData)
			if err != nil {
				log.Printf("Gagal mengkonversi data thread: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, thread); err != nil {
				log.Printf("Gagal menguraikan data thread: %v", err)
				continue
			}
			
			threadList = append(threadList, thread)
		}
	}
	
	return threadList, nil
}

func (r *DiskusiThreadRepositoryImpl) Update(ctx context.Context, thread *domain.DiskusiThread) error {
	thread.TanggalUbah = time.Now()
	return r.db.Save("diskusi_thread", thread.ID, thread)
}

func (r *DiskusiThreadRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("diskusi_thread", id)
}

// DiskusiKomentarRepositoryImpl implements the DiskusiKomentarRepository interface
type DiskusiKomentarRepositoryImpl struct {
	db *MockSawitDB
}

// NewDiskusiKomentarRepository creates a new diskusi komentar repository instance
func NewDiskusiKomentarRepository(db *MockSawitDB) *DiskusiKomentarRepositoryImpl {
	return &DiskusiKomentarRepositoryImpl{db: db}
}

func (r *DiskusiKomentarRepositoryImpl) Create(ctx context.Context, komentar *domain.DiskusiKomentar) error {
	komentar.ID = fmt.Sprintf("dkm_%d", time.Now().Unix())
	komentar.TanggalBuat = time.Now()
	komentar.TanggalUbah = time.Now()
	return r.db.Save("diskusi_komentar", komentar.ID, komentar)
}

func (r *DiskusiKomentarRepositoryImpl) GetByThreadID(ctx context.Context, threadID string) ([]*domain.DiskusiKomentar, error) {
	allKomentarData, err := r.db.FindAll("diskusi_komentar", 0, 1000) // Assuming max 1000 comments for demo
	if err != nil {
		return nil, err
	}
	
	var komentarList []*domain.DiskusiKomentar
	for _, komentarData := range allKomentarData {
		komentarMap, ok := komentarData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if komentarThreadID, ok := komentarMap["id_thread"].(string); ok && komentarThreadID == threadID {
			komentar := &domain.DiskusiKomentar{}
			dataBytes, err := json.Marshal(komentarData)
			if err != nil {
				log.Printf("Gagal mengkonversi data komentar: %v", err)
				continue
			}
			
			if err := json.Unmarshal(dataBytes, komentar); err != nil {
				log.Printf("Gagal menguraikan data komentar: %v", err)
				continue
			}
			
			komentarList = append(komentarList, komentar)
		}
	}
	
	return komentarList, nil
}

func (r *DiskusiKomentarRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.DiskusiKomentar, error) {
	data, err := r.db.FindByID("diskusi_komentar", id)
	if err != nil {
		return nil, err
	}
	
	komentar := &domain.DiskusiKomentar{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, komentar); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data komentar: %v", err)
	}
	
	return komentar, nil
}

func (r *DiskusiKomentarRepositoryImpl) Update(ctx context.Context, komentar *domain.DiskusiKomentar) error {
	komentar.TanggalUbah = time.Now()
	return r.db.Save("diskusi_komentar", komentar.ID, komentar)
}

func (r *DiskusiKomentarRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("diskusi_komentar", id)
}

// KonfigurasiAplikasiRepositoryImpl implements the KonfigurasiAplikasiRepository interface
type KonfigurasiAplikasiRepositoryImpl struct {
	db *MockSawitDB
}

// NewKonfigurasiAplikasiRepository creates a new konfigurasi aplikasi repository instance
func NewKonfigurasiAplikasiRepository(db *MockSawitDB) *KonfigurasiAplikasiRepositoryImpl {
	return &KonfigurasiAplikasiRepositoryImpl{db: db}
}

func (r *KonfigurasiAplikasiRepositoryImpl) Get(ctx context.Context) (*domain.KonfigurasiAplikasi, error) {
	// For demo purposes, we'll return a default configuration
	// In a real implementation, we would fetch from the database
	defaultConfig := &domain.KonfigurasiAplikasi{
		ID:          "config_1",
		NamaAplikasi: "Averroes",
		WarnaUtama:  "#22c55e", // Emerald green
		LinkSosial: map[string]string{
			"facebook": "",
			"twitter":  "",
			"instagram": "",
			"youtube":  "",
		},
		TanggalUbah: time.Now(),
	}
	
	return defaultConfig, nil
}

func (r *KonfigurasiAplikasiRepositoryImpl) Update(ctx context.Context, config *domain.KonfigurasiAplikasi) error {
	config.TanggalUbah = time.Now()
	return r.db.Save("konfigurasi_aplikasi", config.ID, config)
}
package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
)

// MockSawitDB simulates the SawitDB file-based storage
type MockSawitDB struct {
	data map[string][]byte
	mu   sync.RWMutex
	path string
}

// NewMockSawitDB creates a new mock database instance
func NewMockSawitDB(path string) *MockSawitDB {
	db := &MockSawitDB{
		data: make(map[string][]byte),
		path: path,
	}
	
	// Initialize with empty collections
	db.initCollections()
	
	return db
}

// initCollections initializes the database with empty collections
func (db *MockSawitDB) initCollections() {
	collections := []string{
		"pengguna", "kelas", "modul", "materi", "ujian", "soal", 
		"hasil_ujian", "sertifikat", "pustaka", "berita", 
		"diskusi_thread", "diskusi_komentar", "konfigurasi_aplikasi",
	}
	
	for _, collection := range collections {
		db.data[collection] = []byte("[]") // Empty array as initial data
	}
}

// Save stores data in the mock database
func (db *MockSawitDB) Save(collection, id string, data interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	// Get existing data for the collection
	existingData := []interface{}{}
	if val, exists := db.data[collection]; exists {
		if err := json.Unmarshal(val, &existingData); err != nil {
			// If unmarshaling fails, start with empty array
			existingData = []interface{}{}
		}
	}
	
	// Convert data to map for easier manipulation
	dataMap := make(map[string]interface{})
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
		return fmt.Errorf("gagal mengkonversi data ke map: %v", err)
	}
	
	// Add or update the record
	found := false
	for i, item := range existingData {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		
		if itemMap["id"] == id {
			existingData[i] = dataMap
			found = true
			break
		}
	}
	
	if !found {
		existingData = append(existingData, dataMap)
	}
	
	// Save back to database
	updatedData, err := json.Marshal(existingData)
	if err != nil {
		return fmt.Errorf("gagal menyimpan data: %v", err)
	}
	
	db.data[collection] = updatedData
	
	return nil
}

// FindByID finds a record by ID in the specified collection
func (db *MockSawitDB) FindByID(collection, id string) (interface{}, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	existingData := []interface{}{}
	if val, exists := db.data[collection]; exists {
		if err := json.Unmarshal(val, &existingData); err != nil {
			return nil, fmt.Errorf("gagal membaca data koleksi: %v", err)
		}
	}
	
	for _, item := range existingData {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		
		if itemMap["id"] == id {
			return item, nil
		}
	}
	
	return nil, errors.New("data tidak ditemukan")
}

// FindAll returns all records from the specified collection with pagination
func (db *MockSawitDB) FindAll(collection string, offset, limit int) ([]interface{}, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	existingData := []interface{}{}
	if val, exists := db.data[collection]; exists {
		if err := json.Unmarshal(val, &existingData); err != nil {
			return nil, fmt.Errorf("gagal membaca data koleksi: %v", err)
		}
	}
	
	start := offset
	if start >= len(existingData) {
		return []interface{}{}, nil
	}
	
	end := start + limit
	if end > len(existingData) {
		end = len(existingData)
	}
	
	return existingData[start:end], nil
}

// Delete removes a record by ID from the specified collection
func (db *MockSawitDB) Delete(collection, id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	existingData := []interface{}{}
	if val, exists := db.data[collection]; exists {
		if err := json.Unmarshal(val, &existingData); err != nil {
			return fmt.Errorf("gagal membaca data koleksi: %v", err)
		}
	}
	
	for i, item := range existingData {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		
		if itemMap["id"] == id {
			// Remove the item
			existingData = append(existingData[:i], existingData[i+1:]...)
			
			// Save back to database
			updatedData, err := json.Marshal(existingData)
			if err != nil {
				return fmt.Errorf("gagal menyimpan data setelah penghapusan: %v", err)
			}
			
			db.data[collection] = updatedData
			return nil
		}
	}
	
	return errors.New("data tidak ditemukan")
}

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *MockSawitDB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *MockSawitDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	user.ID = fmt.Sprintf("usr_%d", time.Now().Unix())
	user.TanggalBuat = time.Now()
	user.TanggalUbah = time.Now()
	return r.db.Save("pengguna", user.ID, user)
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.User, error) {
	data, err := r.db.FindByID("pengguna", id)
	if err != nil {
		return nil, err
	}
	
	user := &domain.User{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, user); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data pengguna: %v", err)
	}
	
	return user, nil
}

func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	// Find all users and filter by email
	allUsersData, err := r.db.FindAll("pengguna", 0, 1000) // Assuming max 1000 users for demo
	if err != nil {
		return nil, err
	}
	
	for _, userData := range allUsersData {
		userMap, ok := userData.(map[string]interface{})
		if !ok {
			continue
		}
		
		if userEmail, ok := userMap["email"].(string); ok && userEmail == email {
			user := &domain.User{}
			dataBytes, err := json.Marshal(userData)
			if err != nil {
				continue
			}
			
			if err := json.Unmarshal(dataBytes, user); err != nil {
				continue
			}
			
			return user, nil
		}
	}
	
	return nil, errors.New("pengguna tidak ditemukan")
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	user.TanggalUbah = time.Now()
	return r.db.Save("pengguna", user.ID, user)
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("pengguna", id)
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	usersData, err := r.db.FindAll("pengguna", offset, limit)
	if err != nil {
		return nil, err
	}
	
	users := make([]*domain.User, 0, len(usersData))
	for _, userData := range usersData {
		user := &domain.User{}
		dataBytes, err := json.Marshal(userData)
		if err != nil {
			log.Printf("Gagal mengkonversi data pengguna: %v", err)
			continue
		}
		
		if err := json.Unmarshal(dataBytes, user); err != nil {
			log.Printf("Gagal menguraikan data pengguna: %v", err)
			continue
		}
		
		users = append(users, user)
	}
	
	return users, nil
}

// KelasRepositoryImpl implements the KelasRepository interface
type KelasRepositoryImpl struct {
	db *MockSawitDB
}

// NewKelasRepository creates a new kelas repository instance
func NewKelasRepository(db *MockSawitDB) *KelasRepositoryImpl {
	return &KelasRepositoryImpl{db: db}
}

func (r *KelasRepositoryImpl) Create(ctx context.Context, kelas *domain.Kelas) error {
	kelas.ID = fmt.Sprintf("kls_%d", time.Now().Unix())
	kelas.TanggalBuat = time.Now()
	kelas.TanggalUbah = time.Now()
	return r.db.Save("kelas", kelas.ID, kelas)
}

func (r *KelasRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Kelas, error) {
	data, err := r.db.FindByID("kelas", id)
	if err != nil {
		return nil, err
	}
	
	kelas := &domain.Kelas{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, kelas); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data kelas: %v", err)
	}
	
	return kelas, nil
}

func (r *KelasRepositoryImpl) GetAll(ctx context.Context, offset, limit int) ([]*domain.Kelas, error) {
	kelasData, err := r.db.FindAll("kelas", offset, limit)
	if err != nil {
		return nil, err
	}
	
	kelasList := make([]*domain.Kelas, 0, len(kelasData))
	for _, kelasItem := range kelasData {
		kelas := &domain.Kelas{}
		dataBytes, err := json.Marshal(kelasItem)
		if err != nil {
			log.Printf("Gagal mengkonversi data kelas: %v", err)
			continue
		}
		
		if err := json.Unmarshal(dataBytes, kelas); err != nil {
			log.Printf("Gagal menguraikan data kelas: %v", err)
			continue
		}
		
		kelasList = append(kelasList, kelas)
	}
	
	return kelasList, nil
}

func (r *KelasRepositoryImpl) Update(ctx context.Context, kelas *domain.Kelas) error {
	kelas.TanggalUbah = time.Now()
	return r.db.Save("kelas", kelas.ID, kelas)
}

func (r *KelasRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("kelas", id)
}

// BukuRepositoryImpl implements the BukuRepository interface
type BukuRepositoryImpl struct {
	db *MockSawitDB
}

// NewBukuRepository creates a new buku repository instance
func NewBukuRepository(db *MockSawitDB) *BukuRepositoryImpl {
	return &BukuRepositoryImpl{db: db}
}

func (r *BukuRepositoryImpl) Create(ctx context.Context, buku *domain.Buku) error {
	buku.ID = fmt.Sprintf("bku_%d", time.Now().Unix())
	buku.TanggalBuat = time.Now()
	buku.TanggalUbah = time.Now()
	return r.db.Save("pustaka", buku.ID, buku)
}

func (r *BukuRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.Buku, error) {
	data, err := r.db.FindByID("pustaka", id)
	if err != nil {
		return nil, err
	}
	
	buku := &domain.Buku{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("gagal mengkonversi data: %v", err)
	}
	
	if err := json.Unmarshal(dataBytes, buku); err != nil {
		return nil, fmt.Errorf("gagal menguraikan data buku: %v", err)
	}
	
	return buku, nil
}

func (r *BukuRepositoryImpl) GetAll(ctx context.Context, offset, limit int) ([]*domain.Buku, error) {
	bukuData, err := r.db.FindAll("pustaka", offset, limit)
	if err != nil {
		return nil, err
	}
	
	bukuList := make([]*domain.Buku, 0, len(bukuData))
	for _, bukuItem := range bukuData {
		buku := &domain.Buku{}
		dataBytes, err := json.Marshal(bukuItem)
		if err != nil {
			log.Printf("Gagal mengkonversi data buku: %v", err)
			continue
		}
		
		if err := json.Unmarshal(dataBytes, buku); err != nil {
			log.Printf("Gagal menguraikan data buku: %v", err)
			continue
		}
		
		bukuList = append(bukuList, buku)
	}
	
	return bukuList, nil
}

func (r *BukuRepositoryImpl) Update(ctx context.Context, buku *domain.Buku) error {
	buku.TanggalUbah = time.Now()
	return r.db.Save("pustaka", buku.ID, buku)
}

func (r *BukuRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.Delete("pustaka", id)
}
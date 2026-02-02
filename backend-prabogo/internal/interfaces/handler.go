package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/averroes/backend-prabogo/internal/domain"
	"github.com/averroes/backend-prabogo/internal/infrastructure/repository"
	"github.com/averroes/backend-prabogo/internal/usecase"
	"github.com/averroes/backend-prabogo/pkg/config"
	"github.com/gorilla/mux"
)

// Handler holds all the use cases
type Handler struct {
	Config                 *config.Config
	UserUseCase            domain.UserUseCase
	KelasUseCase           domain.KelasUseCase
	ModulUseCase           domain.ModulUseCase
	MateriUseCase          domain.MateriUseCase
	UjianUseCase           domain.UjianUseCase
	SoalUseCase            domain.SoalUseCase
	HasilUjianUseCase      domain.HasilUjianUseCase
	SertifikatUseCase      domain.SertifikatUseCase
	BukuUseCase            domain.BukuUseCase
	BeritaUseCase          domain.BeritaUseCase
	DiskusiThreadUseCase   domain.DiskusiThreadUseCase
	DiskusiKomentarUseCase domain.DiskusiKomentarUseCase
	KonfigurasiUseCase     domain.KonfigurasiAplikasiUseCase
}

// NewHandler creates a new handler instance
func NewHandler(config *config.Config) *Handler {
	// Initialize mock database
	db := repository.NewMockSawitDB(config.DB.SawitDBPath)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	kelasRepo := repository.NewKelasRepository(db)
	modulRepo := repository.NewModulRepository(db)
	materiRepo := repository.NewMateriRepository(db)
	ujianRepo := repository.NewUjianRepository(db)
	soalRepo := repository.NewSoalRepository(db)
	hasilUjianRepo := repository.NewHasilUjianRepository(db)
	sertifikatRepo := repository.NewSertifikatRepository(db)
	bukuRepo := repository.NewBukuRepository(db)
	beritaRepo := repository.NewBeritaRepository(db)
	diskusiThreadRepo := repository.NewDiskusiThreadRepository(db)
	diskusiKomentarRepo := repository.NewDiskusiKomentarRepository(db)
	konfigurasiRepo := repository.NewKonfigurasiAplikasiRepository(db)

	// Initialize use cases
	userUC := usecase.NewUserUseCase(userRepo)
	kelasUC := usecase.NewKelasUseCase(kelasRepo)
	modulUC := usecase.NewModulUseCase(modulRepo, kelasRepo)
	materiUC := usecase.NewMateriUseCase(materiRepo, modulRepo)
	ujianUC := usecase.NewUjianUseCase(ujianRepo, kelasRepo, modulRepo)
	soalUC := usecase.NewSoalUseCase(soalRepo, ujianRepo)
	hasilUjianUC := usecase.NewHasilUjianUseCase(hasilUjianRepo, userRepo, ujianRepo)
	sertifikatUC := usecase.NewSertifikatUseCase(sertifikatRepo, userRepo, kelasRepo)
	bukuUC := usecase.NewBukuUseCase(bukuRepo)
	beritaUC := usecase.NewBeritaUseCase(beritaRepo)
	diskusiThreadUC := usecase.NewDiskusiThreadUseCase(diskusiThreadRepo, userRepo)
	diskusiKomentarUC := usecase.NewDiskusiKomentarUseCase(diskusiKomentarRepo, diskusiThreadRepo, userRepo)
	konfigurasiUC := usecase.NewKonfigurasiAplikasiUseCase(konfigurasiRepo)

	return &Handler{
		Config:                config,
		UserUseCase:           userUC,
		KelasUseCase:          kelasUC,
		ModulUseCase:          modulUC,
		MateriUseCase:         materiUC,
		UjianUseCase:          ujianUC,
		SoalUseCase:           soalUC,
		HasilUjianUseCase:     hasilUjianUC,
		SertifikatUseCase:     sertifikatUC,
		BukuUseCase:           bukuUC,
		BeritaUseCase:         beritaUC,
		DiskusiThreadUseCase:  diskusiThreadUC,
		DiskusiKomentarUseCase: diskusiKomentarUC,
		KonfigurasiUseCase:    konfigurasiUC,
	}
}

// RegisterRoutes registers all the routes for the application
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Authentication routes
	router.HandleFunc("/api/masuk", h.Login).Methods("POST")
	router.HandleFunc("/api/daftar", h.Register).Methods("POST")
	router.HandleFunc("/api/keluar", h.Logout).Methods("POST")
	router.HandleFunc("/api/profil", h.GetProfile).Methods("GET")

	// LMS routes
	router.HandleFunc("/api/kelas", h.GetAllKelas).Methods("GET")
	router.HandleFunc("/api/kelas/{id}", h.GetKelasByID).Methods("GET")
	router.HandleFunc("/api/kelas/{id}/modul", h.GetModulByKelasID).Methods("GET")
	router.HandleFunc("/api/modul/{id}/materi", h.GetMateriByModulID).Methods("GET")
	router.HandleFunc("/api/kelas/{id}/ujian", h.GetUjianByKelasID).Methods("GET")

	// Pustaka routes
	router.HandleFunc("/api/pustaka", h.GetAllBuku).Methods("GET")
	router.HandleFunc("/api/pustaka/{id}", h.GetBukuByID).Methods("GET")

	// Berita routes
	router.HandleFunc("/api/berita", h.GetBerita).Methods("GET")              // with query param limit
	router.HandleFunc("/api/berita/terbaru", h.GetLatestBerita).Methods("GET") // with query param limit

	// Diskusi routes
	router.HandleFunc("/api/diskusi", h.GetAllDiskusiThread).Methods("GET")
	router.HandleFunc("/api/diskusi", h.CreateDiskusiThread).Methods("POST")
	router.HandleFunc("/api/diskusi/{id}/balas", h.CreateDiskusiKomentar).Methods("POST")
}

// Helper function to send JSON response
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Helper function to send error response in Indonesian
func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	sendJSONResponse(w, status, map[string]interface{}{
		"error":   true,
		"message": message,
	})
}

// Helper function to send success response
func sendSuccessResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	response := map[string]interface{}{
		"error":   false,
		"message": message,
	}

	if data != nil {
		response["data"] = data
	}

	sendJSONResponse(w, status, response)
}

// Login handles user login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Data permintaan tidak valid")
		return
	}

	email := strings.TrimSpace(reqBody["email"])
	password := strings.TrimSpace(reqBody["password"])

	if email == "" || password == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Email dan kata sandi wajib diisi")
		return
	}

	user, err := h.UserUseCase.Login(r.Context(), email, password)
	if err != nil {
		sendErrorResponse(w, http.StatusUnauthorized, "Email atau kata sandi salah")
		return
	}

	// In a real implementation, you would generate a JWT token here
	token := "mock_token_for_demo" // Replace with actual JWT token generation

	sendSuccessResponse(w, http.StatusOK, "Login berhasil", map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

// Register handles user registration
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var reqBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Data permintaan tidak valid")
		return
	}

	nama := strings.TrimSpace(reqBody["nama"])
	email := strings.TrimSpace(reqBody["email"])
	password := strings.TrimSpace(reqBody["password"])

	if nama == "" || email == "" || password == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Nama, email, dan kata sandi wajib diisi")
		return
	}

	user := &domain.User{
		Nama:      nama,
		Email:     email,
		KataSandi: password,
		Peran:     "user", // default role
		Status:    "aktif", // default status
	}

	err = h.UserUseCase.Register(r.Context(), user)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mendaftarkan pengguna: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusCreated, "Pendaftaran berhasil", nil)
}

// Logout handles user logout
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Berhasil keluar",
	})
}

// GetProfile returns user profile
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, you would extract the user ID from the JWT token
	// For demo purposes, we'll use a mock user ID
	userID := "1" // This should come from the JWT token in a real implementation

	user, err := h.UserUseCase.GetProfile(r.Context(), userID)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil profil pengguna: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Profil pengguna berhasil diambil", user)
}

// GetAllKelas returns all classes
func (h *Handler) GetAllKelas(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	var offset, limit int
	var err error

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}
	} else {
		offset = 0
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit > 100 {
			limit = 10
		}
	} else {
		limit = 10
	}

	kelasList, err := h.KelasUseCase.GetAllKelas(r.Context(), offset, limit)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil daftar kelas: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar kelas berhasil diambil", map[string]interface{}{
		"data": kelasList,
		"meta": map[string]int{
			"offset": offset,
			"limit":  limit,
			"total":  len(kelasList),
		},
	})
}

// GetKelasByID returns a specific class by ID
func (h *Handler) GetKelasByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	kelas, err := h.KelasUseCase.GetKelasByID(r.Context(), id)
	if err != nil {
		sendErrorResponse(w, http.StatusNotFound, "Kelas tidak ditemukan: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Detail kelas berhasil diambil", kelas)
}

// GetModulByKelasID returns modules for a specific class
func (h *Handler) GetModulByKelasID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kelasID := vars["id"]

	modulList, err := h.ModulUseCase.GetModulByKelasID(r.Context(), kelasID)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil modul: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar modul berhasil diambil", modulList)
}

// GetMateriByModulID returns materials for a specific module
func (h *Handler) GetMateriByModulID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modulID := vars["id"]

	materiList, err := h.MateriUseCase.GetMateriByModulID(r.Context(), modulID)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil materi: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar materi berhasil diambil", materiList)
}

// GetUjianByKelasID returns exams for a specific class
func (h *Handler) GetUjianByKelasID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kelasID := vars["id"]

	ujianList, err := h.UjianUseCase.GetUjianByKelasID(r.Context(), kelasID)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil ujian: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar ujian berhasil diambil", ujianList)
}

// GetAllBuku returns all books in the library
func (h *Handler) GetAllBuku(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	var offset, limit int
	var err error

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}
	} else {
		offset = 0
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit > 100 {
			limit = 10
		}
	} else {
		limit = 10
	}

	bukuList, err := h.BukuUseCase.GetAllBuku(r.Context(), offset, limit)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil daftar buku: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar buku berhasil diambil", map[string]interface{}{
		"data": bukuList,
		"meta": map[string]int{
			"offset": offset,
			"limit":  limit,
			"total":  len(bukuList),
		},
	})
}

// GetBukuByID returns a specific book by ID
func (h *Handler) GetBukuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	buku, err := h.BukuUseCase.GetBukuByID(r.Context(), id)
	if err != nil {
		sendErrorResponse(w, http.StatusNotFound, "Buku tidak ditemukan: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Detail buku berhasil diambil", buku)
}

// GetBerita returns news articles
func (h *Handler) GetBerita(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")

	var limit int
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit > 20 {
			limit = 5
		}
	} else {
		limit = 5
	}

	beritaList, err := h.BeritaUseCase.GetAllBerita(r.Context(), 0, limit)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil berita: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar berita berhasil diambil", beritaList)
}

// GetLatestBerita returns latest news articles
func (h *Handler) GetLatestBerita(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")

	var limit int
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit > 20 {
			limit = 20
		}
	} else {
		limit = 20
	}

	beritaList, err := h.BeritaUseCase.GetLatestBerita(r.Context(), limit)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil berita terbaru: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar berita terbaru berhasil diambil", beritaList)
}

// GetAllDiskusiThread returns all discussion threads
func (h *Handler) GetAllDiskusiThread(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	var offset, limit int
	var err error

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			offset = 0
		}
	} else {
		offset = 0
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit > 100 {
			limit = 10
		}
	} else {
		limit = 10
	}

	threadList, err := h.DiskusiThreadUseCase.GetAllThreads(r.Context(), offset, limit)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil daftar diskusi: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusOK, "Daftar diskusi berhasil diambil", map[string]interface{}{
		"data": threadList,
		"meta": map[string]int{
			"offset": offset,
			"limit":  limit,
			"total":  len(threadList),
		},
	})
}

// CreateDiskusiThread creates a new discussion thread
func (h *Handler) CreateDiskusiThread(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Data permintaan tidak valid")
		return
	}

	// Extract required fields
	judul := strings.TrimSpace(reqBody["judul"])
	isi := strings.TrimSpace(reqBody["isi"])

	if judul == "" || isi == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Judul dan isi wajib diisi")
		return
	}

	// In a real implementation, we would extract user ID from JWT token
	// For demo purposes, we'll use a mock user ID
	userID := "1" // This should come from the JWT token in a real implementation

	thread := &domain.DiskusiThread{
		IDUser: userID,
		Judul:  judul,
		Isi:    isi,
		Status: "aktif", // default status
	}

	err = h.DiskusiThreadUseCase.CreateThread(r.Context(), thread)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal membuat diskusi: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusCreated, "Diskusi berhasil dibuat", thread)
}

// CreateDiskusiKomentar creates a new comment in a discussion thread
func (h *Handler) CreateDiskusiKomentar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["id"]

	// Parse request body
	var reqBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Data permintaan tidak valid")
		return
	}

	// Extract required field
	isi := strings.TrimSpace(reqBody["isi"])

	if isi == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Isi komentar wajib diisi")
		return
	}

	// In a real implementation, we would extract user ID from JWT token
	// For demo purposes, we'll use a mock user ID
	userID := "1" // This should come from the JWT token in a real implementation

	komentar := &domain.DiskusiKomentar{
		IDThread: threadID,
		IDUser:   userID,
		Isi:      isi,
		Status:   "aktif", // default status
	}

	err = h.DiskusiKomentarUseCase.CreateKomentar(r.Context(), komentar)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Gagal membuat komentar: "+err.Error())
		return
	}

	sendSuccessResponse(w, http.StatusCreated, "Komentar berhasil ditambahkan", komentar)
}
package http

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
	"github.com/averroes/backend-prabogo/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type Handler struct {
	AuthUsecase       *usecase.AuthUsecase
	ScreenerUsecase   *usecase.ScreenerUsecase
	EdukasiUsecase    *usecase.EdukasiUsecase
	PustakaUsecase    *usecase.PustakaUsecase
	BeritaUsecase     *usecase.BeritaUsecase
	DiskusiUsecase    *usecase.DiskusiUsecase
	PortofolioUsecase *usecase.PortofolioUsecase
	ZakatUsecase      *usecase.ZakatUsecase
	ReelsUsecase      *usecase.ReelsUsecase
	TadabburUsecase   *usecase.TadabburUsecase
	AdminUsecase      *usecase.AdminUsecase
	JWTSecret         string
	Versi             string
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Health check at root level for Railway
	router.HandleFunc("/health", h.HealthCheck).Methods("GET")
	router.HandleFunc("/", h.HealthCheck).Methods("GET")

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/status", h.Status).Methods("GET")

	api.HandleFunc("/daftar", h.Daftar).Methods("POST")
	api.HandleFunc("/verifikasi-otp", h.VerifikasiOTP).Methods("POST")
	api.HandleFunc("/kirim-ulang-otp", h.KirimUlangOTP).Methods("POST")
	api.HandleFunc("/masuk", h.Masuk).Methods("POST")
	api.Handle("/keluar", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.Keluar))).Methods("POST")
	api.Handle("/profil", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.Profil))).Methods("GET")

	api.HandleFunc("/screener", h.DaftarScreener).Methods("GET")
	api.HandleFunc("/screener/{id}", h.DetailScreener).Methods("GET")
	api.HandleFunc("/screener/{id}/catatan", h.CatatanScreener).Methods("GET")
	api.HandleFunc("/pasar", h.DaftarPasar).Methods("GET")

	api.HandleFunc("/kelas", h.DaftarKelas).Methods("GET")
	api.HandleFunc("/kelas/{id}", h.DetailKelas).Methods("GET")
	api.HandleFunc("/kelas/{id}/modul", h.DaftarModul).Methods("GET")
	api.HandleFunc("/modul/{id}/materi", h.DaftarMateri).Methods("GET")
	api.HandleFunc("/kelas/{id}/ujian", h.DaftarUjian).Methods("GET")
	api.Handle("/kelas/{id}/mulai", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.MulaiKelas))).Methods("POST")
	api.Handle("/progress", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.DaftarProgress))).Methods("GET")

	api.HandleFunc("/pustaka", h.DaftarPustaka).Methods("GET")
	api.HandleFunc("/pustaka/{id}", h.DetailPustaka).Methods("GET")

	api.HandleFunc("/berita", h.DaftarBerita).Methods("GET")
	api.HandleFunc("/berita/terbaru", h.DaftarBeritaTerbaru).Methods("GET")

	api.HandleFunc("/diskusi", h.DaftarDiskusi).Methods("GET")
	api.HandleFunc("/diskusi/{id}", h.DetailDiskusi).Methods("GET")
	api.Handle("/diskusi", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.BuatDiskusi))).Methods("POST")
	api.Handle("/diskusi/{id}/balas", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.BalasDiskusi))).Methods("POST")
	api.Handle("/diskusi/{id}/lapor", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.LaporDiskusi))).Methods("POST")

	api.Handle("/portofolio", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.DaftarPortofolio))).Methods("GET")
	api.Handle("/portofolio", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.TambahPortofolio))).Methods("POST")
	api.Handle("/portofolio/{id}", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.PerbaruiPortofolio))).Methods("PUT")
	api.Handle("/portofolio/{id}", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.HapusPortofolio))).Methods("DELETE")

	api.Handle("/zakat/ringkasan", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.RingkasanZakat))).Methods("GET")
	api.Handle("/zakat/riwayat", AuthMiddleware(h.JWTSecret)(http.HandlerFunc(h.RiwayatZakat))).Methods("GET")
	api.HandleFunc("/harga-emas", h.HargaEmas).Methods("GET")

	api.HandleFunc("/reels", h.DaftarReels).Methods("GET")
	api.HandleFunc("/reels/{id}", h.DetailReels).Methods("GET")

	api.HandleFunc("/tadabbur", h.DaftarTadabbur).Methods("GET")

	admin := api.PathPrefix("/admin").Subrouter()
	if strings.ToLower(os.Getenv("ADMIN_NO_AUTH")) != "true" {
		admin.Use(AuthMiddleware(h.JWTSecret))
	}
	admin.HandleFunc("/pengguna", h.AdminDaftarPengguna).Methods("GET")
	admin.HandleFunc("/pengguna", h.AdminPerbaruiPengguna).Methods("PUT")
	admin.HandleFunc("/pengguna/{id}", h.AdminHapusPengguna).Methods("DELETE")

	admin.HandleFunc("/kelas", h.AdminDaftarKelas).Methods("GET")
	admin.HandleFunc("/kelas", h.AdminBuatKelas).Methods("POST")
	admin.HandleFunc("/kelas/{id}", h.AdminPerbaruiKelas).Methods("PUT")
	admin.HandleFunc("/kelas/{id}", h.AdminHapusKelas).Methods("DELETE")

	admin.HandleFunc("/modul", h.AdminDaftarModul).Methods("GET")
	admin.HandleFunc("/modul", h.AdminBuatModul).Methods("POST")
	admin.HandleFunc("/modul/{id}", h.AdminPerbaruiModul).Methods("PUT")
	admin.HandleFunc("/modul/{id}", h.AdminHapusModul).Methods("DELETE")

	admin.HandleFunc("/materi", h.AdminDaftarMateri).Methods("GET")
	admin.HandleFunc("/materi", h.AdminBuatMateri).Methods("POST")
	admin.HandleFunc("/materi/{id}", h.AdminPerbaruiMateri).Methods("PUT")
	admin.HandleFunc("/materi/{id}", h.AdminHapusMateri).Methods("DELETE")

	admin.HandleFunc("/ujian", h.AdminDaftarUjian).Methods("GET")
	admin.HandleFunc("/ujian", h.AdminBuatUjian).Methods("POST")
	admin.HandleFunc("/ujian/{id}", h.AdminPerbaruiUjian).Methods("PUT")
	admin.HandleFunc("/ujian/{id}", h.AdminHapusUjian).Methods("DELETE")

	admin.HandleFunc("/sertifikat", h.AdminDaftarSertifikat).Methods("GET")
	admin.HandleFunc("/sertifikat", h.AdminBuatSertifikat).Methods("POST")
	admin.HandleFunc("/sertifikat/{id}", h.AdminHapusSertifikat).Methods("DELETE")

	admin.HandleFunc("/pustaka", h.AdminDaftarPustaka).Methods("GET")
	admin.HandleFunc("/pustaka", h.AdminBuatPustaka).Methods("POST")
	admin.HandleFunc("/pustaka/{id}", h.AdminPerbaruiPustaka).Methods("PUT")
	admin.HandleFunc("/pustaka/{id}", h.AdminHapusPustaka).Methods("DELETE")

	admin.HandleFunc("/berita", h.AdminDaftarBerita).Methods("GET")
	admin.HandleFunc("/berita", h.AdminBuatBerita).Methods("POST")
	admin.HandleFunc("/berita/{id}", h.AdminPerbaruiBerita).Methods("PUT")
	admin.HandleFunc("/berita/{id}", h.AdminHapusBerita).Methods("DELETE")

	admin.HandleFunc("/diskusi", h.AdminDaftarDiskusi).Methods("GET")

	admin.HandleFunc("/screener", h.AdminDaftarScreener).Methods("GET")
	admin.HandleFunc("/screener", h.AdminBuatScreener).Methods("POST")
	admin.HandleFunc("/screener/{id}", h.AdminPerbaruiScreener).Methods("PUT")
	admin.HandleFunc("/screener/{id}", h.AdminHapusScreener).Methods("DELETE")

	admin.HandleFunc("/pasar", h.AdminDaftarPasar).Methods("GET")
	admin.HandleFunc("/pasar", h.AdminBuatPasar).Methods("POST")
	admin.HandleFunc("/pasar/{id}", h.AdminPerbaruiPasar).Methods("PUT")
	admin.HandleFunc("/pasar/{id}", h.AdminHapusPasar).Methods("DELETE")

	admin.HandleFunc("/reels", h.AdminDaftarReels).Methods("GET")
	admin.HandleFunc("/reels", h.AdminBuatReels).Methods("POST")
	admin.HandleFunc("/reels/{id}", h.AdminPerbaruiReels).Methods("PUT")
	admin.HandleFunc("/reels/{id}", h.AdminHapusReels).Methods("DELETE")

	admin.HandleFunc("/tadabbur", h.AdminDaftarTadabbur).Methods("GET")
	admin.HandleFunc("/tadabbur", h.AdminBuatTadabbur).Methods("POST")
	admin.HandleFunc("/tadabbur/{id}", h.AdminPerbaruiTadabbur).Methods("PUT")
	admin.HandleFunc("/tadabbur/{id}", h.AdminHapusTadabbur).Methods("DELETE")

	admin.HandleFunc("/pengaturan", h.AdminDaftarKonfigurasi).Methods("GET")
	admin.HandleFunc("/pengaturan", h.AdminBuatKonfigurasi).Methods("POST")
	admin.HandleFunc("/pengaturan/{id}", h.AdminPerbaruiKonfigurasi).Methods("PUT")
	admin.HandleFunc("/pengaturan/{id}", h.AdminHapusKonfigurasi).Methods("DELETE")
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	ResponSukses(w, http.StatusOK, "Status layanan aktif", map[string]interface{}{
		"status":       "ok",
		"versi":        h.Versi,
		"waktu_server": time.Now().Format(time.RFC3339),
	})
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "healthy",
		"app":    "Averroes API",
		"versi":  h.Versi,
	})
}

type daftarRequest struct {
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	KataSandi string `json:"kata_sandi"`
	Peran     string `json:"peran"`
}

func (h *Handler) Daftar(w http.ResponseWriter, r *http.Request) {
	var req daftarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}

	if strings.TrimSpace(req.Nama) == "" || strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.KataSandi) == "" {
		ResponGagal(w, http.StatusBadRequest, "Nama, email, dan kata sandi wajib diisi", nil)
		return
	}

	pengguna, otp, err := h.AuthUsecase.Daftar(r.Context(), strings.TrimSpace(req.Nama), strings.TrimSpace(req.Email), req.KataSandi, req.Peran)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mendaftar", err.Error())
		return
	}

	ResponSukses(w, http.StatusCreated, "Pendaftaran berhasil, silakan verifikasi OTP", map[string]interface{}{
		"pengguna":        pengguna,
		"otp":             otp.Kode,
		"kadaluarsa_pada": otp.KadaluarsaPada,
	})
}

type verifikasiRequest struct {
	Email string `json:"email"`
	Kode  string `json:"kode"`
}

func (h *Handler) VerifikasiOTP(w http.ResponseWriter, r *http.Request) {
	var req verifikasiRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Kode) == "" {
		ResponGagal(w, http.StatusBadRequest, "Email dan kode OTP wajib diisi", nil)
		return
	}
	if err := h.AuthUsecase.VerifikasiOTP(r.Context(), strings.TrimSpace(req.Email), strings.TrimSpace(req.Kode)); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Verifikasi OTP gagal", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Verifikasi OTP berhasil", nil)
}

func (h *Handler) KirimUlangOTP(w http.ResponseWriter, r *http.Request) {
	var req verifikasiRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Email) == "" {
		ResponGagal(w, http.StatusBadRequest, "Email wajib diisi", nil)
		return
	}
	otp, err := h.AuthUsecase.KirimUlangOTP(r.Context(), strings.TrimSpace(req.Email))
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "Gagal kirim ulang OTP", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "OTP berhasil dikirim ulang", map[string]interface{}{
		"otp":             otp.Kode,
		"kadaluarsa_pada": otp.KadaluarsaPada,
	})
}

type masukRequest struct {
	Email     string `json:"email"`
	KataSandi string `json:"kata_sandi"`
}

func (h *Handler) Masuk(w http.ResponseWriter, r *http.Request) {
	var req masukRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.KataSandi) == "" {
		ResponGagal(w, http.StatusBadRequest, "Email dan kata sandi wajib diisi", nil)
		return
	}

	pengguna, err := h.AuthUsecase.Masuk(r.Context(), strings.TrimSpace(req.Email), req.KataSandi)
	if err != nil {
		ResponGagal(w, http.StatusUnauthorized, "Gagal masuk", err.Error())
		return
	}

	claims := jwt.MapClaims{
		"id_pengguna": pengguna.ID,
		"peran":       pengguna.Peran,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat token", err.Error())
		return
	}

	ResponSukses(w, http.StatusOK, "Berhasil masuk", map[string]interface{}{
		"token":    tokenStr,
		"pengguna": pengguna,
	})
}

func (h *Handler) Keluar(w http.ResponseWriter, r *http.Request) {
	ResponSukses(w, http.StatusOK, "Berhasil keluar", nil)
}

func (h *Handler) Profil(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	pengguna, err := h.AuthUsecase.Profil(r.Context(), idPengguna)
	if err != nil || pengguna == nil {
		ResponGagal(w, http.StatusNotFound, "Profil tidak ditemukan", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Profil berhasil diambil", pengguna)
}

func (h *Handler) DaftarScreener(w http.ResponseWriter, r *http.Request) {
	kategori := r.URL.Query().Get("kategori")
	cari := r.URL.Query().Get("cari")
	data, err := h.ScreenerUsecase.Daftar(r.Context(), kategori, cari)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil screener", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar screener berhasil diambil", data)
}

func (h *Handler) DetailScreener(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID screener tidak valid", nil)
		return
	}
	data, err := h.ScreenerUsecase.Detail(r.Context(), id)
	if err != nil || data == nil {
		ResponGagal(w, http.StatusNotFound, "Screener tidak ditemukan", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Detail screener berhasil diambil", data)
}

func (h *Handler) CatatanScreener(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID screener tidak valid", nil)
		return
	}
	data, err := h.ScreenerUsecase.Catatan(r.Context(), id)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil catatan screener", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Catatan screener berhasil diambil", data)
}

func (h *Handler) DaftarPasar(w http.ResponseWriter, r *http.Request) {
	data, err := h.ScreenerUsecase.Pasar(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil data pasar", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar pasar berhasil diambil", data)
}

func (h *Handler) DaftarKelas(w http.ResponseWriter, r *http.Request) {
	data, err := h.EdukasiUsecase.DaftarKelas(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil daftar kelas", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar kelas berhasil diambil", data)
}

func (h *Handler) DetailKelas(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID kelas tidak valid", nil)
		return
	}
	data, err := h.EdukasiUsecase.DetailKelas(r.Context(), id)
	if err != nil || data == nil {
		ResponGagal(w, http.StatusNotFound, "Kelas tidak ditemukan", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Detail kelas berhasil diambil", data)
}

func (h *Handler) DaftarModul(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID kelas tidak valid", nil)
		return
	}
	data, err := h.EdukasiUsecase.DaftarModul(r.Context(), id)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil modul", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar modul berhasil diambil", data)
}

func (h *Handler) DaftarMateri(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID modul tidak valid", nil)
		return
	}
	data, err := h.EdukasiUsecase.DaftarMateri(r.Context(), id)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil materi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar materi berhasil diambil", data)
}

func (h *Handler) DaftarUjian(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID kelas tidak valid", nil)
		return
	}
	data, err := h.EdukasiUsecase.DaftarUjian(r.Context(), id)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil ujian", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar ujian berhasil diambil", data)
}

func (h *Handler) MulaiKelas(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	idKelas, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID kelas tidak valid", nil)
		return
	}
	if err := h.EdukasiUsecase.MulaiKelas(r.Context(), idPengguna, idKelas); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menyimpan progress", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Progress kelas disimpan", nil)
}

func (h *Handler) DaftarProgress(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	data, err := h.EdukasiUsecase.DaftarProgress(r.Context(), idPengguna)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil progress", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Progress berhasil diambil", data)
}

func (h *Handler) DaftarPustaka(w http.ResponseWriter, r *http.Request) {
	data, err := h.PustakaUsecase.Daftar(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil pustaka", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar pustaka berhasil diambil", data)
}

func (h *Handler) DetailPustaka(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID pustaka tidak valid", nil)
		return
	}
	data, err := h.PustakaUsecase.Detail(r.Context(), id)
	if err != nil || data == nil {
		ResponGagal(w, http.StatusNotFound, "Pustaka tidak ditemukan", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Detail pustaka berhasil diambil", data)
}

func (h *Handler) DaftarBerita(w http.ResponseWriter, r *http.Request) {
	limit := parseLimit(r.URL.Query().Get("limit"), 5)
	data, err := h.BeritaUsecase.Daftar(r.Context(), limit)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil berita", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Berita berhasil diambil", data)
}

func (h *Handler) DaftarBeritaTerbaru(w http.ResponseWriter, r *http.Request) {
	limit := parseLimit(r.URL.Query().Get("limit"), 20)
	data, err := h.BeritaUsecase.Terbaru(r.Context(), limit)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil berita terbaru", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Berita terbaru berhasil diambil", data)
}

func (h *Handler) DaftarDiskusi(w http.ResponseWriter, r *http.Request) {
	data, err := h.DiskusiUsecase.Daftar(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil diskusi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar diskusi berhasil diambil", data)
}

func (h *Handler) DetailDiskusi(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID diskusi tidak valid", nil)
		return
	}
	diskusi, balasan, err := h.DiskusiUsecase.Detail(r.Context(), id)
	if err != nil || diskusi == nil {
		ResponGagal(w, http.StatusNotFound, "Diskusi tidak ditemukan", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Detail diskusi berhasil diambil", map[string]interface{}{
		"diskusi": diskusi,
		"balasan": balasan,
	})
}

func (h *Handler) BuatDiskusi(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Judul string `json:"judul"`
		Isi   string `json:"isi"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Judul) == "" || strings.TrimSpace(req.Isi) == "" {
		ResponGagal(w, http.StatusBadRequest, "Judul dan isi wajib diisi", nil)
		return
	}
	idPengguna := r.Context().Value(ContextUserID).(int64)
	data, err := h.DiskusiUsecase.Buat(r.Context(), idPengguna, req.Judul, req.Isi)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat diskusi", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Diskusi berhasil dibuat", data)
}

func (h *Handler) BalasDiskusi(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Isi string `json:"isi"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Isi) == "" {
		ResponGagal(w, http.StatusBadRequest, "Isi balasan wajib diisi", nil)
		return
	}
	idDiskusi, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID diskusi tidak valid", nil)
		return
	}
	idPengguna := r.Context().Value(ContextUserID).(int64)
	data, err := h.DiskusiUsecase.Balas(r.Context(), idPengguna, idDiskusi, req.Isi)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat balasan", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Balasan berhasil ditambahkan", data)
}

func (h *Handler) LaporDiskusi(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Alasan string `json:"alasan"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Alasan) == "" {
		ResponGagal(w, http.StatusBadRequest, "Alasan laporan wajib diisi", nil)
		return
	}
	idDiskusi, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID diskusi tidak valid", nil)
		return
	}
	idPengguna := r.Context().Value(ContextUserID).(int64)
	if err := h.DiskusiUsecase.Lapor(r.Context(), idPengguna, idDiskusi, req.Alasan); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengirim laporan", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Laporan berhasil dikirim", nil)
}

func (h *Handler) DaftarPortofolio(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	data, err := h.PortofolioUsecase.Daftar(r.Context(), idPengguna)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil portofolio", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Portofolio berhasil diambil", data)
}

func (h *Handler) TambahPortofolio(w http.ResponseWriter, r *http.Request) {
	var req domain.Portofolio
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.NamaAset) == "" || strings.TrimSpace(req.Simbol) == "" {
		ResponGagal(w, http.StatusBadRequest, "Nama aset dan simbol wajib diisi", nil)
		return
	}
	idPengguna := r.Context().Value(ContextUserID).(int64)
	req.IDPengguna = idPengguna
	if err := h.PortofolioUsecase.Tambah(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menambah portofolio", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Portofolio berhasil ditambahkan", req)
}

func (h *Handler) PerbaruiPortofolio(w http.ResponseWriter, r *http.Request) {
	var req domain.Portofolio
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	idPengguna := r.Context().Value(ContextUserID).(int64)
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID portofolio tidak valid", nil)
		return
	}
	req.ID = id
	req.IDPengguna = idPengguna
	if err := h.PortofolioUsecase.Perbarui(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui portofolio", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Portofolio berhasil diperbarui", req)
}

func (h *Handler) HapusPortofolio(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID portofolio tidak valid", nil)
		return
	}
	if err := h.PortofolioUsecase.Hapus(r.Context(), id, idPengguna); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus portofolio", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Portofolio berhasil dihapus", nil)
}

func (h *Handler) RingkasanZakat(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	data, err := h.ZakatUsecase.Ringkasan(r.Context(), idPengguna)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghitung zakat", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Ringkasan zakat berhasil dihitung", data)
}

func (h *Handler) RiwayatZakat(w http.ResponseWriter, r *http.Request) {
	idPengguna := r.Context().Value(ContextUserID).(int64)
	data, err := h.ZakatUsecase.Riwayat(r.Context(), idPengguna)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil riwayat zakat", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Riwayat zakat berhasil diambil", data)
}

func (h *Handler) HargaEmas(w http.ResponseWriter, r *http.Request) {
	data, err := h.ZakatUsecase.HargaEmas(r.Context())
	if err != nil || data == nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil harga emas", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Harga emas berhasil diambil", data)
}

func (h *Handler) DaftarReels(w http.ResponseWriter, r *http.Request) {
	tema := r.URL.Query().Get("tema")
	data, err := h.ReelsUsecase.Daftar(r.Context(), tema)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil reels", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar reels berhasil diambil", data)
}

func (h *Handler) DetailReels(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID reels tidak valid", nil)
		return
	}
	data, err := h.ReelsUsecase.Detail(r.Context(), id)
	if err != nil || data == nil {
		ResponGagal(w, http.StatusNotFound, "Reels tidak ditemukan", nil)
		return
	}
	ResponSukses(w, http.StatusOK, "Detail reels berhasil diambil", data)
}

func (h *Handler) DaftarTadabbur(w http.ResponseWriter, r *http.Request) {
	tema := r.URL.Query().Get("tema")
	data, err := h.TadabburUsecase.Daftar(r.Context(), tema)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil tadabbur", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar tadabbur berhasil diambil", data)
}

func (h *Handler) AdminDaftarPengguna(w http.ResponseWriter, r *http.Request) {
	data, err := h.AdminUsecase.DaftarPengguna(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil pengguna", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar pengguna berhasil diambil", data)
}

func (h *Handler) AdminDaftarKelas(w http.ResponseWriter, r *http.Request) {
	data, err := h.EdukasiUsecase.DaftarKelas(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil kelas", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar kelas berhasil diambil", data)
}

func (h *Handler) AdminDaftarModul(w http.ResponseWriter, r *http.Request) {
	data, err := h.EdukasiUsecase.DaftarModulSemua(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil modul", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar modul berhasil diambil", data)
}

func (h *Handler) AdminDaftarMateri(w http.ResponseWriter, r *http.Request) {
	data, err := h.EdukasiUsecase.DaftarMateriSemua(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil materi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar materi berhasil diambil", data)
}

func (h *Handler) AdminDaftarUjian(w http.ResponseWriter, r *http.Request) {
	data, err := h.EdukasiUsecase.DaftarUjianSemua(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil ujian", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar ujian berhasil diambil", data)
}

func (h *Handler) AdminDaftarSertifikat(w http.ResponseWriter, r *http.Request) {
	data, err := h.EdukasiUsecase.DaftarSertifikat(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil sertifikat", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar sertifikat berhasil diambil", data)
}

func (h *Handler) AdminDaftarPustaka(w http.ResponseWriter, r *http.Request) {
	data, err := h.PustakaUsecase.Daftar(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil pustaka", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar pustaka berhasil diambil", data)
}

func (h *Handler) AdminDaftarBerita(w http.ResponseWriter, r *http.Request) {
	data, err := h.BeritaUsecase.Terbaru(r.Context(), 100)
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil berita", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar berita berhasil diambil", data)
}

func (h *Handler) AdminDaftarDiskusi(w http.ResponseWriter, r *http.Request) {
	data, err := h.DiskusiUsecase.Daftar(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil diskusi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar diskusi berhasil diambil", data)
}

func (h *Handler) AdminDaftarScreener(w http.ResponseWriter, r *http.Request) {
	data, err := h.ScreenerUsecase.Daftar(r.Context(), "semua", "")
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil screener", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar screener berhasil diambil", data)
}

func (h *Handler) AdminDaftarPasar(w http.ResponseWriter, r *http.Request) {
	data, err := h.ScreenerUsecase.Pasar(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil pasar", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar pasar berhasil diambil", data)
}

func (h *Handler) AdminDaftarReels(w http.ResponseWriter, r *http.Request) {
	data, err := h.ReelsUsecase.Daftar(r.Context(), "")
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil reels", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar reels berhasil diambil", data)
}

func (h *Handler) AdminDaftarTadabbur(w http.ResponseWriter, r *http.Request) {
	data, err := h.TadabburUsecase.Daftar(r.Context(), "")
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil tadabbur", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar tadabbur berhasil diambil", data)
}

func (h *Handler) AdminDaftarKonfigurasi(w http.ResponseWriter, r *http.Request) {
	data, err := h.AdminUsecase.DaftarKonfigurasi(r.Context())
	if err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal mengambil konfigurasi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Daftar konfigurasi berhasil diambil", data)
}

func (h *Handler) AdminPerbaruiPengguna(w http.ResponseWriter, r *http.Request) {
	var req domain.Pengguna
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if req.ID == 0 {
		ResponGagal(w, http.StatusBadRequest, "ID pengguna wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.PerbaruiPengguna(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui pengguna", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Pengguna berhasil diperbarui", req)
}

func (h *Handler) AdminHapusPengguna(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID pengguna tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusPengguna(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus pengguna", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Pengguna berhasil dihapus", nil)
}

func (h *Handler) AdminBuatKelas(w http.ResponseWriter, r *http.Request) {
	var req domain.Kelas
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Judul) == "" {
		ResponGagal(w, http.StatusBadRequest, "Judul kelas wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatKelas(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat kelas", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Kelas berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiKelas(w http.ResponseWriter, r *http.Request) {
	var req domain.Kelas
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID kelas tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiKelas(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui kelas", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Kelas berhasil diperbarui", req)
}

func (h *Handler) AdminHapusKelas(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID kelas tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusKelas(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus kelas", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Kelas berhasil dihapus", nil)
}

func (h *Handler) AdminBuatModul(w http.ResponseWriter, r *http.Request) {
	var req domain.Modul
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if req.IDKelas == 0 || strings.TrimSpace(req.Judul) == "" {
		ResponGagal(w, http.StatusBadRequest, "ID kelas dan judul modul wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatModul(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat modul", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Modul berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiModul(w http.ResponseWriter, r *http.Request) {
	var req domain.Modul
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID modul tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiModul(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui modul", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Modul berhasil diperbarui", req)
}

func (h *Handler) AdminHapusModul(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID modul tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusModul(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus modul", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Modul berhasil dihapus", nil)
}

func (h *Handler) AdminBuatMateri(w http.ResponseWriter, r *http.Request) {
	var req domain.Materi
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if req.IDModul == 0 || strings.TrimSpace(req.Judul) == "" {
		ResponGagal(w, http.StatusBadRequest, "ID modul dan judul materi wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatMateri(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat materi", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Materi berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiMateri(w http.ResponseWriter, r *http.Request) {
	var req domain.Materi
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID materi tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiMateri(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui materi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Materi berhasil diperbarui", req)
}

func (h *Handler) AdminHapusMateri(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID materi tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusMateri(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus materi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Materi berhasil dihapus", nil)
}

func (h *Handler) AdminBuatUjian(w http.ResponseWriter, r *http.Request) {
	var req domain.Ujian
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if req.IDKelas == 0 || strings.TrimSpace(req.Judul) == "" {
		ResponGagal(w, http.StatusBadRequest, "ID kelas dan judul ujian wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatUjian(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat ujian", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Ujian berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiUjian(w http.ResponseWriter, r *http.Request) {
	var req domain.Ujian
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID ujian tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiUjian(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui ujian", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Ujian berhasil diperbarui", req)
}

func (h *Handler) AdminHapusUjian(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID ujian tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusUjian(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus ujian", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Ujian berhasil dihapus", nil)
}

func (h *Handler) AdminBuatSertifikat(w http.ResponseWriter, r *http.Request) {
	var req domain.Sertifikat
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if req.IDPengguna == 0 || req.IDKelas == 0 || strings.TrimSpace(req.Kode) == "" {
		ResponGagal(w, http.StatusBadRequest, "ID pengguna, ID kelas, dan kode sertifikat wajib diisi", nil)
		return
	}
	if req.TanggalTerbit.IsZero() {
		req.TanggalTerbit = time.Now()
	}
	if err := h.AdminUsecase.BuatSertifikat(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat sertifikat", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Sertifikat berhasil dibuat", req)
}

func (h *Handler) AdminHapusSertifikat(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID sertifikat tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusSertifikat(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus sertifikat", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Sertifikat berhasil dihapus", nil)
}

func (h *Handler) AdminBuatPustaka(w http.ResponseWriter, r *http.Request) {
	var req domain.Pustaka
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.JudulTampil) == "" {
		ResponGagal(w, http.StatusBadRequest, "Judul pustaka wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatPustaka(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat pustaka", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Pustaka berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiPustaka(w http.ResponseWriter, r *http.Request) {
	var req domain.Pustaka
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID pustaka tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiPustaka(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui pustaka", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Pustaka berhasil diperbarui", req)
}

func (h *Handler) AdminHapusPustaka(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID pustaka tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusPustaka(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus pustaka", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Pustaka berhasil dihapus", nil)
}

func (h *Handler) AdminBuatBerita(w http.ResponseWriter, r *http.Request) {
	var req domain.Berita
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Judul) == "" {
		ResponGagal(w, http.StatusBadRequest, "Judul berita wajib diisi", nil)
		return
	}
	if req.DiterbitkanPada.IsZero() {
		req.DiterbitkanPada = time.Now()
	}
	if err := h.AdminUsecase.BuatBerita(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat berita", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Berita berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiBerita(w http.ResponseWriter, r *http.Request) {
	var req domain.Berita
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID berita tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiBerita(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui berita", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Berita berhasil diperbarui", req)
}

func (h *Handler) AdminHapusBerita(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID berita tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusBerita(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus berita", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Berita berhasil dihapus", nil)
}

func (h *Handler) AdminBuatScreener(w http.ResponseWriter, r *http.Request) {
	var req domain.Screener
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.NamaAset) == "" || strings.TrimSpace(req.Simbol) == "" {
		ResponGagal(w, http.StatusBadRequest, "Nama aset dan simbol wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatScreener(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat screener", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Screener berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiScreener(w http.ResponseWriter, r *http.Request) {
	var req domain.Screener
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID screener tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiScreener(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui screener", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Screener berhasil diperbarui", req)
}

func (h *Handler) AdminHapusScreener(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID screener tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusScreener(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus screener", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Screener berhasil dihapus", nil)
}

func (h *Handler) AdminBuatPasar(w http.ResponseWriter, r *http.Request) {
	var req domain.Pasar
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.NamaAset) == "" || strings.TrimSpace(req.Simbol) == "" {
		ResponGagal(w, http.StatusBadRequest, "Nama aset dan simbol wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatPasar(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat pasar", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Data pasar berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiPasar(w http.ResponseWriter, r *http.Request) {
	var req domain.Pasar
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID pasar tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiPasar(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui pasar", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Data pasar berhasil diperbarui", req)
}

func (h *Handler) AdminHapusPasar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID pasar tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusPasar(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus pasar", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Data pasar berhasil dihapus", nil)
}

func (h *Handler) AdminBuatReels(w http.ResponseWriter, r *http.Request) {
	var req domain.Reels
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Judul) == "" || strings.TrimSpace(req.Tema) == "" {
		ResponGagal(w, http.StatusBadRequest, "Judul dan tema reels wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatReels(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat reels", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Reels berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiReels(w http.ResponseWriter, r *http.Request) {
	var req domain.Reels
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID reels tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiReels(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui reels", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Reels berhasil diperbarui", req)
}

func (h *Handler) AdminHapusReels(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID reels tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusReels(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus reels", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Reels berhasil dihapus", nil)
}

func (h *Handler) AdminBuatTadabbur(w http.ResponseWriter, r *http.Request) {
	var req domain.Tadabbur
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Judul) == "" || strings.TrimSpace(req.Tema) == "" {
		ResponGagal(w, http.StatusBadRequest, "Judul dan tema tadabbur wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatTadabbur(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat tadabbur", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Tadabbur berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiTadabbur(w http.ResponseWriter, r *http.Request) {
	var req domain.Tadabbur
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID tadabbur tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiTadabbur(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui tadabbur", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Tadabbur berhasil diperbarui", req)
}

func (h *Handler) AdminHapusTadabbur(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID tadabbur tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusTadabbur(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus tadabbur", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Tadabbur berhasil dihapus", nil)
}

func (h *Handler) AdminBuatKonfigurasi(w http.ResponseWriter, r *http.Request) {
	var req domain.Konfigurasi
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	if strings.TrimSpace(req.Kunci) == "" {
		ResponGagal(w, http.StatusBadRequest, "Kunci konfigurasi wajib diisi", nil)
		return
	}
	if err := h.AdminUsecase.BuatKonfigurasi(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal membuat konfigurasi", err.Error())
		return
	}
	ResponSukses(w, http.StatusCreated, "Konfigurasi berhasil dibuat", req)
}

func (h *Handler) AdminPerbaruiKonfigurasi(w http.ResponseWriter, r *http.Request) {
	var req domain.Konfigurasi
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponGagal(w, http.StatusBadRequest, "Data tidak valid", err.Error())
		return
	}
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID konfigurasi tidak valid", nil)
		return
	}
	req.ID = id
	if err := h.AdminUsecase.PerbaruiKonfigurasi(r.Context(), &req); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal memperbarui konfigurasi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Konfigurasi berhasil diperbarui", req)
}

func (h *Handler) AdminHapusKonfigurasi(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		ResponGagal(w, http.StatusBadRequest, "ID konfigurasi tidak valid", nil)
		return
	}
	if err := h.AdminUsecase.HapusKonfigurasi(r.Context(), id); err != nil {
		ResponGagal(w, http.StatusInternalServerError, "Gagal menghapus konfigurasi", err.Error())
		return
	}
	ResponSukses(w, http.StatusOK, "Konfigurasi berhasil dihapus", nil)
}

func parseID(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func parseLimit(value string, fallback int) int {
	if value == "" {
		return fallback
	}
	limit, err := strconv.Atoi(value)
	if err != nil || limit <= 0 {
		return fallback
	}
	return limit
}

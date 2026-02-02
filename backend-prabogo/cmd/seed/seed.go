package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/averroes/backend-prabogo/internal/domain"
	"github.com/averroes/backend-prabogo/internal/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
)

// SeedData represents the structure of our seed data
type SeedData struct {
	Users     []domain.User                 `json:"users"`
	Classes   []domain.Kelas                `json:"classes"`
	Modules   []domain.Modul                `json:"modules"`
	Materials []domain.Materi               `json:"materials"`
	Books     []domain.Buku                 `json:"books"`
	News      []domain.Berita               `json:"news"`
	Threads   []domain.DiskusiThread        `json:"threads"`
	Config    domain.KonfigurasiAplikasi    `json:"config"`
}

func main() {
	// Initialize mock database
	db := repository.NewMockSawitDB("./data.sawit")
	
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	kelasRepo := repository.NewKelasRepository(db)
	modulRepo := repository.NewModulRepository(db)
	materiRepo := repository.NewMateriRepository(db)
	bukuRepo := repository.NewBukuRepository(db)
	beritaRepo := repository.NewBeritaRepository(db)
	diskusiThreadRepo := repository.NewDiskusiThreadRepository(db)
	konfigurasiRepo := repository.NewKonfigurasiAplikasiRepository(db)
	
	fmt.Println("Memulai seeding data...")
	
	// Create default user with hashed password
	adminUser := &domain.User{
		ID:          "usr_admin_001",
		Nama:        "Administrator Averroes",
		Email:       "admin@averroes.id",
		KataSandi:   hashPassword("admin123"),
		Peran:       "admin",
		Status:      "aktif",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	editorUser := &domain.User{
		ID:          "usr_editor_001",
		Nama:        "Editor Konten",
		Email:       "editor@averroes.id",
		KataSandi:   hashPassword("editor123"),
		Peran:       "editor",
		Status:      "aktif",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	user1 := &domain.User{
		ID:          "usr_user_001",
		Nama:        "Ahmad Rahman",
		Email:       "ahmad.rahman@example.com",
		KataSandi:   hashPassword("user123"),
		Peran:       "user",
		Status:      "aktif",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	user2 := &domain.User{
		ID:          "usr_user_002",
		Nama:        "Siti Nurhaliza",
		Email:       "siti.nur@example.com",
		KataSandi:   hashPassword("user123"),
		Peran:       "user",
		Status:      "aktif",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	// Save users
	users := []*domain.User{adminUser, editorUser, user1, user2}
	for _, user := range users {
		if err := userRepo.Save(nil, user.ID, user); err != nil {
			log.Printf("Gagal menyimpan user %s: %v", user.Nama, err)
		} else {
			fmt.Printf("Berhasil menyimpan user: %s (%s)\n", user.Nama, user.Email)
		}
	}
	
	// Create classes
	class1 := &domain.Kelas{
		ID:           "kls_001",
		Judul:        "Dasar-Dasar Hukum Islam",
		Ringkasan:    "Pelajari dasar-dasar hukum Islam dengan pendekatan yang mudah dipahami",
		Level:        "pemula",
		Harga:        0,
		Status:       "publik",
		GambarSampul: "/images/kelas1.jpg",
		TanggalBuat:  time.Now(),
		TanggalUbah:  time.Now(),
	}
	
	class2 := &domain.Kelas{
		ID:           "kls_002",
		Judul:        "Fiqh Muamalah Kontemporer",
		Ringkasan:    "Memahami hukum-hukum transaksi dalam Islam di era digital",
		Level:        "menengah",
		Harga:        150000,
		Status:       "publik",
		GambarSampul: "/images/kelas2.jpg",
		TanggalBuat:  time.Now(),
		TanggalUbah:  time.Now(),
	}
	
	classes := []*domain.Kelas{class1, class2}
	for _, class := range classes {
		if err := kelasRepo.Save(nil, class.ID, class); err != nil {
			log.Printf("Gagal menyimpan kelas %s: %v", class.Judul, err)
		} else {
			fmt.Printf("Berhasil menyimpan kelas: %s\n", class.Judul)
		}
	}
	
	// Create modules
	module1 := &domain.Modul{
		ID:        "mdl_001",
		IDKelas:   "kls_001",
		Judul:     "Pengantar Hukum Islam",
		Urutan:    1,
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	module2 := &domain.Modul{
		ID:        "mdl_002",
		IDKelas:   "kls_001",
		Judul:     "Sumber-Sumber Hukum",
		Urutan:    2,
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	modules := []*domain.Modul{module1, module2}
	for _, mod := range modules {
		if err := modulRepo.Save(nil, mod.ID, mod); err != nil {
			log.Printf("Gagal menyimpan modul %s: %v", mod.Judul, err)
		} else {
			fmt.Printf("Berhasil menyimpan modul: %s\n", mod.Judul)
		}
	}
	
	// Create materials
	material1 := &domain.Materi{
		ID:        "mtr_001",
		IDModul:   "mdl_001",
		Jenis:     "teks",
		Judul:     "Pengertian Hukum Islam",
		Konten:    "Hukum Islam adalah sistem hukum yang bersumber dari Al-Quran dan Hadits...",
		Durasi:    15,
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	material2 := &domain.Materi{
		ID:        "mtr_002",
		IDModul:   "mdl_001",
		Jenis:     "video",
		Judul:     "Sejarah Perkembangan Hukum Islam",
		Konten:    "Video ini menjelaskan sejarah perkembangan hukum Islam dari masa ke masa...",
		Durasi:    25,
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	materials := []*domain.Materi{material1, material2}
	for _, mat := range materials {
		if err := materiRepo.Save(nil, mat.ID, mat); err != nil {
			log.Printf("Gagal menyimpan materi %s: %v", mat.Judul, err)
		} else {
			fmt.Printf("Berhasil menyimpan materi: %s\n", mat.Judul)
		}
	}
	
	// Create books
	book1 := &domain.Buku{
		ID:             "bku_001",
		JudulTampil:    "Hukum Fikih tentang Uang Kertas (Fiat)",
		JudulAsli:      "Hukum Fiqih terhadap Uang Kertas(fiat)",
		Penulis:        "Penulis Terkenal",
		Kategori:       "Ekonomi Syariah",
		Bahasa:         "Indonesia",
		JumlahHalaman:  250,
		Deskripsi:      "Buku ini membahas secara mendalam tentang hukum fikih terkait uang kertas dalam perspektif Islam.",
		LinkFile:       "https://drive.google.com/file/d/1234567890abcdefg/view?usp=sharing",
		TanggalBuat:    time.Now(),
		TanggalUbah:    time.Now(),
	}
	
	book2 := &domain.Buku{
		ID:             "bku_002",
		JudulTampil:    "Al-Ahkam Al-Fiqhiyyah Terkait Mata Uang Elektronik",
		JudulAsli:      "a l - a h k ā m a l - f i q h i y y a h a l - m u t a ‘ a l l i q a b i l - ‘ u m a l a a t a l - i l i k t i r ū n i y y a h",
		Penulis:        "Ahmad Al-Buhari",
		Kategori:       "Digital Currency",
		Bahasa:         "Arab-Indonesia",
		JumlahHalaman:  180,
		Deskripsi:      "Kajian komprehensif tentang hukum fiqih terkait mata uang elektronik dan cryptocurrency.",
		LinkFile:       "https://drive.google.com/file/d/0987654321fedcba/view?usp=sharing",
		TanggalBuat:    time.Now(),
		TanggalUbah:    time.Now(),
	}
	
	books := []*domain.Buku{book1, book2}
	for _, book := range books {
		if err := bukuRepo.Save(nil, book.ID, book); err != nil {
			log.Printf("Gagal menyimpan buku %s: %v", book.JudulTampil, err)
		} else {
			fmt.Printf("Berhasil menyimpan buku: %s\n", book.JudulTampil)
		}
	}
	
	// Create news
	news1 := &domain.Berita{
		ID:          "brt_001",
		Judul:       "Pentingnya Pendidikan Hukum Islam di Era Digital",
		Ringkasan:   "Mengapa pendidikan hukum Islam tetap relevan di zaman sekarang",
		Isi:         "Di era digital seperti sekarang, pemahaman hukum Islam menjadi semakin penting...",
		Thumbnail:   "/images/berita1.jpg",
		Status:      "publik",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	news2 := &domain.Berita{
		ID:          "brt_002",
		Judul:       "Perkembangan Fiqh Muamalah di Dunia FinTech",
		Ringkasan:   "Bagaimana hukum Islam mengatur transaksi digital",
		Isi:         "Dengan kemajuan teknologi finansial, muncul berbagai pertanyaan hukum baru...",
		Thumbnail:   "/images/berita2.jpg",
		Status:      "publik",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	news := []*domain.Berita{news1, news2}
	for _, n := range news {
		if err := beritaRepo.Save(nil, n.ID, n); err != nil {
			log.Printf("Gagal menyimpan berita %s: %v", n.Judul, err)
		} else {
			fmt.Printf("Berhasil menyimpan berita: %s\n", n.Judul)
		}
	}
	
	// Create discussion threads
	thread1 := &domain.DiskusiThread{
		ID:          "dst_001",
		IDUser:      "usr_user_001",
		Judul:       "Diskusi tentang Hukum Jual Beli Online",
		Isi:         "Bagaimana hukum jual beli melalui platform digital menurut pandangan ulama?",
		Status:      "aktif",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	thread2 := &domain.DiskusiThread{
		ID:          "dst_002",
		IDUser:      "usr_user_002",
		Judul:       "Perbedaan Pendapat tentang Zakat Profesi",
		Isi:         "Apakah zakat profesi wajib bagi semua pekerja?",
		Status:      "aktif",
		TanggalBuat: time.Now(),
		TanggalUbah: time.Now(),
	}
	
	threads := []*domain.DiskusiThread{thread1, thread2}
	for _, thread := range threads {
		if err := diskusiThreadRepo.Save(nil, thread.ID, thread); err != nil {
			log.Printf("Gagal menyimpan thread %s: %v", thread.Judul, err)
		} else {
			fmt.Printf("Berhasil menyimpan thread: %s\n", thread.Judul)
		}
	}
	
	// Create configuration
	config := &domain.KonfigurasiAplikasi{
		ID:          "cfg_001",
		NamaAplikasi: "Averroes",
		WarnaUtama:  "#22c55e",
		LinkSosial: map[string]string{
			"facebook":  "https://facebook.com/averroes",
			"twitter":   "https://twitter.com/averroes",
			"instagram": "https://instagram.com/averroes",
			"youtube":   "https://youtube.com/averroes",
		},
		TanggalUbah: time.Now(),
	}
	
	if err := konfigurasiRepo.Save(nil, config.ID, config); err != nil {
		log.Printf("Gagal menyimpan konfigurasi: %v", err)
	} else {
		fmt.Printf("Berhasil menyimpan konfigurasi aplikasi\n")
	}
	
	fmt.Println("\nProses seeding data selesai!")
}

// hashPassword hashes a password using bcrypt
func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
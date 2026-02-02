package domain

import "time"

// User represents a user in the system
type User struct {
	ID           string    `json:"id"`
	Nama         string    `json:"nama"`
	Email        string    `json:"email"`
	KataSandi    string    `json:"kata_sandi"`
	Peran        string    `json:"peran"` // admin, editor, moderator, user
	Status       string    `json:"status"` // aktif, tidak_aktif, diblokir
	TanggalBuat  time.Time `json:"tanggal_buat"`
	TanggalUbah  time.Time `json:"tanggal_ubah"`
}

// Kelas represents a course in the LMS
type Kelas struct {
	ID          string    `json:"id"`
	Judul       string    `json:"judul"`
	Ringkasan   string    `json:"ringkasan"`
	Level       string    `json:"level"` // pemula, menengah, mahir
	Harga       float64   `json:"harga"`
	Status      string    `json:"status"` // draft, publik
	GambarSampul string   `json:"gambar_sampul"`
	TanggalBuat time.Time `json:"tanggal_buat"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}

// Modul represents a module within a course
type Modul struct {
	ID        string    `json:"id"`
	IDKelas   string    `json:"id_kelas"`
	Judul     string    `json:"judul"`
	Urutan    int       `json:"urutan"`
	TanggalBuat time.Time `json:"tanggal_buat"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}

// Materi represents learning material within a module
type Materi struct {
	ID        string    `json:"id"`
	IDModul   string    `json:"id_modul"`
	Jenis     string    `json:"jenis"` // video, teks, kuis
	Judul     string    `json:"judul"`
	Konten    string    `json:"konten"`
	Durasi    int       `json:"durasi"` // dalam menit
	Lampiran  string    `json:"lampiran"` // path to attachment
	TanggalBuat time.Time `json:"tanggal_buat"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}

// Ujian represents an exam for a course or module
type Ujian struct {
	ID            string    `json:"id"`
	IDKelas       string    `json:"id_kelas"`
	IDModul       string    `json:"id_modul,omitempty"`
	Judul         string    `json:"judul"`
	BankSoal      []Soal    `json:"bank_soal"`
	PassingGrade  float64   `json:"passing_grade"`
	TanggalBuat   time.Time `json:"tanggal_buat"`
	TanggalUbah   time.Time `json:"tanggal_ubah"`
}

// Soal represents a question in an exam
type Soal struct {
	ID        string   `json:"id"`
	IDUjian   string   `json:"id_ujian"`
	Soal      string   `json:"soal"`
	Jawaban   []string `json:"jawaban"`
	JawabanBenar int    `json:"jawaban_benar"` // index of correct answer
	Tipe      string   `json:"tipe"` // pilihan_ganda, esai
}

// HasilUjian represents exam results
type HasilUjian struct {
	ID        string    `json:"id"`
	IDUser    string    `json:"id_user"`
	IDUjian   string    `json:"id_ujian"`
	Skor      float64   `json:"skor"`
	Status    string    `json:"status"` // lulus, gagal
	TanggalUjian time.Time `json:"tanggal_ujian"`
}

// Sertifikat represents a certificate issued to a user
type Sertifikat struct {
	ID              string    `json:"id"`
	IDUser          string    `json:"id_user"`
	IDKelas         string    `json:"id_kelas"`
	NomorSertifikat string    `json:"nomor_sertifikat"`
	TanggalTerbit   time.Time `json:"tanggal_terbit"`
	TanggalBerlaku  time.Time `json:"tanggal_berlaku"`
}

// Buku represents a book in the library
type Buku struct {
	ID             string    `json:"id"`
	JudulTampil    string    `json:"judul_tampil"` // readable title
	JudulAsli      string    `json:"judul_asli"`   // metadata title
	Penulis        string    `json:"penulis"`
	Kategori       string    `json:"kategori"`
	Bahasa         string    `json:"bahasa"`
	JumlahHalaman  int       `json:"jumlah_halaman"`
	Deskripsi      string    `json:"deskripsi"`
	LinkFile       string    `json:"link_file"` // google drive direct link
	TanggalBuat    time.Time `json:"tanggal_buat"`
	TanggalUbah    time.Time `json:"tanggal_ubah"`
}

// Berita represents news articles
type Berita struct {
	ID          string    `json:"id"`
	Judul       string    `json:"judul"`
	Ringkasan   string    `json:"ringkasan"`
	Isi         string    `json:"isi"`
	Thumbnail   string    `json:"thumbnail"`
	Status      string    `json:"status"` // publik, draft, arsip
	TanggalBuat time.Time `json:"tanggal_buat"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}

// DiskusiThread represents a discussion thread
type DiskusiThread struct {
	ID          string    `json:"id"`
	IDUser      string    `json:"id_user"`
	Judul       string    `json:"judul"`
	Isi         string    `json:"isi"`
	Status      string    `json:"status"` // aktif, diarsipkan, disembunyikan
	TanggalBuat time.Time `json:"tanggal_buat"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}

// DiskusiKomentar represents a comment in a discussion thread
type DiskusiKomentar struct {
	ID          string    `json:"id"`
	IDThread    string    `json:"id_thread"`
	IDUser      string    `json:"id_user"`
	Isi         string    `json:"isi"`
	Status      string    `json:"status"` // aktif, dihapus, disembunyikan
	TanggalBuat time.Time `json:"tanggal_buat"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}

// KonfigurasiAplikasi represents application configuration
type KonfigurasiAplikasi struct {
	ID          string    `json:"id"`
	NamaAplikasi string   `json:"nama_aplikasi"`
	WarnaUtama  string    `json:"warna_utama"`
	LinkSosial  map[string]string `json:"link_sosial"`
	TanggalUbah time.Time `json:"tanggal_ubah"`
}
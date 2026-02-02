package domain

import "time"

type Pengguna struct {
	ID              int64     `json:"id"`
	Nama            string    `json:"nama"`
	Email           string    `json:"email"`
	KataSandiHash   string    `json:"-"`
	Peran           string    `json:"peran"`
	Status          string    `json:"status"`
	SudahVerifikasi bool      `json:"sudah_verifikasi"`
	DibuatPada      time.Time `json:"dibuat_pada"`
	DiubahPada      time.Time `json:"diubah_pada"`
}

type OTPVerifikasi struct {
	ID               int64     `json:"id"`
	IDPengguna       int64     `json:"id_pengguna"`
	Kode             string    `json:"kode"`
	KadaluarsaPada   time.Time `json:"kadaluarsa_pada"`
	TerakhirKirimPada time.Time `json:"terakhir_kirim_pada"`
	JumlahKirim      int       `json:"jumlah_kirim"`
}

type Screener struct {
	ID            int64     `json:"id"`
	NamaAset      string    `json:"nama_aset"`
	Simbol        string    `json:"simbol"`
	Kategori      string    `json:"kategori"`
	SkorSyariah   float64   `json:"skor_syariah"`
	Keterangan    string    `json:"keterangan"`
	HargaTerakhir float64   `json:"harga_terakhir"`
	Perubahan24J  float64   `json:"perubahan_24j"`
	DibuatPada    time.Time `json:"dibuat_pada"`
}

type ScreenerCatatan struct {
	ID        int64     `json:"id"`
	IDScreener int64    `json:"id_screener"`
	Judul     string    `json:"judul"`
	Isi       string    `json:"isi"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

type Pasar struct {
	ID                int64     `json:"id"`
	NamaAset          string    `json:"nama_aset"`
	Simbol            string    `json:"simbol"`
	Harga             float64   `json:"harga"`
	Volume24J         float64   `json:"volume_24j"`
	Perubahan24J      float64   `json:"perubahan_24j"`
	KapitalisasiPasar float64   `json:"kapitalisasi_pasar"`
	DiperbaruiPada    time.Time `json:"diperbarui_pada"`
}

type Kelas struct {
	ID           int64     `json:"id"`
	Judul        string    `json:"judul"`
	Deskripsi    string    `json:"deskripsi"`
	Level        string    `json:"level"`
	JumlahModul  int       `json:"jumlah_modul"`
	DurasiMenit  int       `json:"durasi_menit"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Status       string    `json:"status"`
	DibuatPada   time.Time `json:"dibuat_pada"`
}

type Modul struct {
	ID          int64     `json:"id"`
	IDKelas     int64     `json:"id_kelas"`
	Judul       string    `json:"judul"`
	Urutan      int       `json:"urutan"`
	Ringkasan   string    `json:"ringkasan"`
	DurasiMenit int       `json:"durasi_menit"`
	DibuatPada  time.Time `json:"dibuat_pada"`
}

type Materi struct {
	ID          int64     `json:"id"`
	IDModul     int64     `json:"id_modul"`
	Judul       string    `json:"judul"`
	Tipe        string    `json:"tipe"`
	Konten      string    `json:"konten"`
	URLVideo    string    `json:"url_video"`
	DurasiMenit int       `json:"durasi_menit"`
	DibuatPada  time.Time `json:"dibuat_pada"`
}

type Ujian struct {
	ID          int64     `json:"id"`
	IDKelas     int64     `json:"id_kelas"`
	Judul       string    `json:"judul"`
	Deskripsi   string    `json:"deskripsi"`
	DurasiMenit int       `json:"durasi_menit"`
	JumlahSoal  int       `json:"jumlah_soal"`
	DibuatPada  time.Time `json:"dibuat_pada"`
}

type ProgressKelas struct {
	ID                 int64     `json:"id"`
	IDPengguna         int64     `json:"id_pengguna"`
	IDKelas            int64     `json:"id_kelas"`
	Persentase         float64   `json:"persentase"`
	Status             string    `json:"status"`
	TerakhirDiaksesPada time.Time `json:"terakhir_diakses_pada"`
}

type Sertifikat struct {
	ID            int64     `json:"id"`
	IDPengguna    int64     `json:"id_pengguna"`
	IDKelas       int64     `json:"id_kelas"`
	Kode          string    `json:"kode"`
	TanggalTerbit time.Time `json:"tanggal_terbit"`
}

type Pustaka struct {
	ID            int64     `json:"id"`
	JudulTampil   string    `json:"judul_tampil"`
	JudulAsli     string    `json:"judul_asli"`
	Penulis       string    `json:"penulis"`
	Kategori      string    `json:"kategori"`
	Bahasa        string    `json:"bahasa"`
	JumlahHalaman int       `json:"jumlah_halaman"`
	Deskripsi     string    `json:"deskripsi"`
	TautanFile    string    `json:"tautan_file"`
}

type Berita struct {
	ID             int64     `json:"id"`
	Judul          string    `json:"judul"`
	Ringkasan      string    `json:"ringkasan"`
	Isi            string    `json:"isi"`
	Kategori       string    `json:"kategori"`
	Sumber         string    `json:"sumber"`
	GambarURL      string    `json:"gambar_url"`
	DiterbitkanPada time.Time `json:"diterbitkan_pada"`
}

type Diskusi struct {
	ID        int64     `json:"id"`
	IDPengguna int64    `json:"id_pengguna"`
	Judul     string    `json:"judul"`
	Isi       string    `json:"isi"`
	Status    string    `json:"status"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

type DiskusiBalas struct {
	ID        int64     `json:"id"`
	IDDiskusi int64     `json:"id_diskusi"`
	IDPengguna int64    `json:"id_pengguna"`
	Isi       string    `json:"isi"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

type DiskusiLaporan struct {
	ID        int64     `json:"id"`
	IDDiskusi int64     `json:"id_diskusi"`
	IDPengguna int64    `json:"id_pengguna"`
	Alasan    string    `json:"alasan"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

type Portofolio struct {
	ID          int64     `json:"id"`
	IDPengguna  int64     `json:"id_pengguna"`
	NamaAset    string    `json:"nama_aset"`
	Simbol      string    `json:"simbol"`
	Jumlah      float64   `json:"jumlah"`
	HargaBeli   float64   `json:"harga_beli"`
	NilaiSaatIni float64  `json:"nilai_saat_ini"`
	Kategori    string    `json:"kategori"`
	DibuatPada  time.Time `json:"dibuat_pada"`
}

type ZakatRingkasan struct {
	TotalNilai    float64 `json:"total_nilai"`
	Nisab         float64 `json:"nisab"`
	PersenZakat   float64 `json:"persen_zakat"`
	ZakatTerhitung float64 `json:"zakat_terhitung"`
	WajibZakat    bool    `json:"wajib_zakat"`
}

type ZakatRiwayat struct {
	ID            int64     `json:"id"`
	IDPengguna    int64     `json:"id_pengguna"`
	TotalNilai    float64   `json:"total_nilai"`
	Nisab         float64   `json:"nisab"`
	PersenZakat   float64   `json:"persen_zakat"`
	ZakatTerhitung float64  `json:"zakat_terhitung"`
	DibuatPada    time.Time `json:"dibuat_pada"`
}

type HargaEmas struct {
	ID         int64     `json:"id"`
	Tanggal    time.Time `json:"tanggal"`
	HargaPerGram float64 `json:"harga_per_gram"`
}

type Reels struct {
	ID          int64     `json:"id"`
	Judul       string    `json:"judul"`
	Tema        string    `json:"tema"`
	Kutipan     string    `json:"kutipan"`
	Sumber      string    `json:"sumber"`
	URLVideo    string    `json:"url_video"`
	ThumbnailURL string   `json:"thumbnail_url"`
	DibuatPada  time.Time `json:"dibuat_pada"`
}

type Tadabbur struct {
	ID         int64     `json:"id"`
	Judul      string    `json:"judul"`
	Tema       string    `json:"tema"`
	Ringkasan  string    `json:"ringkasan"`
	Isi        string    `json:"isi"`
	Sumber     string    `json:"sumber"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

type Konfigurasi struct {
	ID       int64  `json:"id"`
	Kunci    string `json:"kunci"`
	Nilai    string `json:"nilai"`
	Deskripsi string `json:"deskripsi"`
}

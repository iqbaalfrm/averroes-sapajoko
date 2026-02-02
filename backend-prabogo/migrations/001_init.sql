CREATE TABLE IF NOT EXISTS pengguna (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nama VARCHAR(150) NOT NULL,
  email VARCHAR(150) NOT NULL UNIQUE,
  kata_sandi_hash VARCHAR(255) NOT NULL,
  peran VARCHAR(50) NOT NULL,
  status VARCHAR(50) NOT NULL,
  sudah_verifikasi TINYINT(1) NOT NULL DEFAULT 0,
  dibuat_pada DATETIME NOT NULL,
  diubah_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS otp_verifikasi (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_pengguna BIGINT NOT NULL,
  kode VARCHAR(10) NOT NULL,
  kadaluarsa_pada DATETIME NOT NULL,
  terakhir_kirim_pada DATETIME NOT NULL,
  jumlah_kirim INT NOT NULL DEFAULT 0,
  INDEX idx_otp_pengguna (id_pengguna),
  CONSTRAINT fk_otp_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS screener (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nama_aset VARCHAR(150) NOT NULL,
  simbol VARCHAR(20) NOT NULL,
  kategori VARCHAR(50) NOT NULL,
  skor_syariah DECIMAL(5,2) NOT NULL,
  keterangan TEXT NOT NULL,
  harga_terakhir DECIMAL(20,4) NOT NULL,
  perubahan_24j DECIMAL(8,2) NOT NULL,
  dibuat_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS screener_catatan (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_screener BIGINT NOT NULL,
  judul VARCHAR(150) NOT NULL,
  isi TEXT NOT NULL,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_catatan_screener (id_screener),
  CONSTRAINT fk_catatan_screener FOREIGN KEY (id_screener) REFERENCES screener(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS pasar (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nama_aset VARCHAR(150) NOT NULL,
  simbol VARCHAR(20) NOT NULL,
  harga DECIMAL(20,4) NOT NULL,
  volume_24j DECIMAL(20,4) NOT NULL,
  perubahan_24j DECIMAL(8,2) NOT NULL,
  kapitalisasi_pasar DECIMAL(20,4) NOT NULL,
  diperbarui_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS kelas (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  judul VARCHAR(150) NOT NULL,
  deskripsi TEXT NOT NULL,
  level VARCHAR(50) NOT NULL,
  jumlah_modul INT NOT NULL DEFAULT 0,
  durasi_menit INT NOT NULL DEFAULT 0,
  thumbnail_url TEXT NOT NULL,
  status VARCHAR(50) NOT NULL,
  dibuat_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS modul (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_kelas BIGINT NOT NULL,
  judul VARCHAR(150) NOT NULL,
  urutan INT NOT NULL,
  ringkasan TEXT NOT NULL,
  durasi_menit INT NOT NULL DEFAULT 0,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_modul_kelas (id_kelas),
  CONSTRAINT fk_modul_kelas FOREIGN KEY (id_kelas) REFERENCES kelas(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS materi (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_modul BIGINT NOT NULL,
  judul VARCHAR(150) NOT NULL,
  tipe VARCHAR(50) NOT NULL,
  konten TEXT NOT NULL,
  url_video TEXT NOT NULL,
  durasi_menit INT NOT NULL DEFAULT 0,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_materi_modul (id_modul),
  CONSTRAINT fk_materi_modul FOREIGN KEY (id_modul) REFERENCES modul(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS ujian (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_kelas BIGINT NOT NULL,
  judul VARCHAR(150) NOT NULL,
  deskripsi TEXT NOT NULL,
  durasi_menit INT NOT NULL DEFAULT 0,
  jumlah_soal INT NOT NULL DEFAULT 0,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_ujian_kelas (id_kelas),
  CONSTRAINT fk_ujian_kelas FOREIGN KEY (id_kelas) REFERENCES kelas(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS progress_kelas (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_pengguna BIGINT NOT NULL,
  id_kelas BIGINT NOT NULL,
  persentase DECIMAL(5,2) NOT NULL DEFAULT 0,
  status VARCHAR(50) NOT NULL,
  terakhir_diakses_pada DATETIME NOT NULL,
  UNIQUE KEY uk_progress (id_pengguna, id_kelas),
  CONSTRAINT fk_progress_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE,
  CONSTRAINT fk_progress_kelas FOREIGN KEY (id_kelas) REFERENCES kelas(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS sertifikat (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_pengguna BIGINT NOT NULL,
  id_kelas BIGINT NOT NULL,
  kode VARCHAR(100) NOT NULL,
  tanggal_terbit DATETIME NOT NULL,
  INDEX idx_sertifikat_pengguna (id_pengguna),
  CONSTRAINT fk_sertifikat_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE,
  CONSTRAINT fk_sertifikat_kelas FOREIGN KEY (id_kelas) REFERENCES kelas(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS pustaka (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  judul_tampil VARCHAR(200) NOT NULL,
  judul_asli VARCHAR(200) NOT NULL,
  penulis VARCHAR(150) NOT NULL,
  kategori VARCHAR(100) NOT NULL,
  bahasa VARCHAR(100) NOT NULL,
  jumlah_halaman INT NOT NULL,
  deskripsi TEXT NOT NULL,
  tautan_file TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS berita (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  judul VARCHAR(200) NOT NULL,
  ringkasan TEXT NOT NULL,
  isi TEXT NOT NULL,
  kategori VARCHAR(100) NOT NULL,
  sumber VARCHAR(150) NOT NULL,
  gambar_url TEXT NOT NULL,
  diterbitkan_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS diskusi (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_pengguna BIGINT NOT NULL,
  judul VARCHAR(200) NOT NULL,
  isi TEXT NOT NULL,
  status VARCHAR(50) NOT NULL,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_diskusi_pengguna (id_pengguna),
  CONSTRAINT fk_diskusi_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS diskusi_balas (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_diskusi BIGINT NOT NULL,
  id_pengguna BIGINT NOT NULL,
  isi TEXT NOT NULL,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_diskusi_balas (id_diskusi),
  CONSTRAINT fk_balas_diskusi FOREIGN KEY (id_diskusi) REFERENCES diskusi(id) ON DELETE CASCADE,
  CONSTRAINT fk_balas_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS diskusi_laporan (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_diskusi BIGINT NOT NULL,
  id_pengguna BIGINT NOT NULL,
  alasan TEXT NOT NULL,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_laporan_diskusi (id_diskusi),
  CONSTRAINT fk_laporan_diskusi FOREIGN KEY (id_diskusi) REFERENCES diskusi(id) ON DELETE CASCADE,
  CONSTRAINT fk_laporan_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS portofolio (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_pengguna BIGINT NOT NULL,
  nama_aset VARCHAR(150) NOT NULL,
  simbol VARCHAR(20) NOT NULL,
  jumlah DECIMAL(20,8) NOT NULL,
  harga_beli DECIMAL(20,4) NOT NULL,
  nilai_saat_ini DECIMAL(20,4) NOT NULL,
  kategori VARCHAR(50) NOT NULL,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_portofolio_pengguna (id_pengguna),
  CONSTRAINT fk_portofolio_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS zakat_riwayat (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  id_pengguna BIGINT NOT NULL,
  total_nilai DECIMAL(20,4) NOT NULL,
  nisab DECIMAL(20,4) NOT NULL,
  persen_zakat DECIMAL(6,2) NOT NULL,
  zakat_terhitung DECIMAL(20,4) NOT NULL,
  dibuat_pada DATETIME NOT NULL,
  INDEX idx_zakat_pengguna (id_pengguna),
  CONSTRAINT fk_zakat_pengguna FOREIGN KEY (id_pengguna) REFERENCES pengguna(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS harga_emas (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  tanggal DATE NOT NULL,
  harga_per_gram DECIMAL(20,4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS reels (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  judul VARCHAR(200) NOT NULL,
  tema VARCHAR(100) NOT NULL,
  kutipan TEXT NOT NULL,
  sumber VARCHAR(150) NOT NULL,
  url_video TEXT NOT NULL,
  thumbnail_url TEXT NOT NULL,
  dibuat_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS tadabbur (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  judul VARCHAR(200) NOT NULL,
  tema VARCHAR(100) NOT NULL,
  ringkasan TEXT NOT NULL,
  isi TEXT NOT NULL,
  sumber VARCHAR(150) NOT NULL,
  dibuat_pada DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS konfigurasi (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  kunci VARCHAR(100) NOT NULL,
  nilai TEXT NOT NULL,
  deskripsi TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

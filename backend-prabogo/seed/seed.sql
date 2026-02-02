INSERT INTO pengguna (nama, email, kata_sandi_hash, peran, status, sudah_verifikasi, dibuat_pada, diubah_pada) VALUES
('Admin Averroes', 'admin@averroes.id', '$2a$10$KMg7u19uNKufsfpOJCjEO.z7NhzfT5m7g3pyvEPpY1hdZgPwW/3UW', 'admin', 'aktif', 1, NOW(), NOW()),
('Editor Konten', 'editor@averroes.id', '$2a$10$/6ChuXqZtgNSEgNC6nJrveOJvHuvMSHx/iPzyZ8Cmm7KhihgYz5qO', 'editor', 'aktif', 1, NOW(), NOW()),
('Moderator Forum', 'moderator@averroes.id', '$2a$10$EnXpqPlOl0FrAKp6byW6OurHv5G.rWp5Wm8MJhsqCEIt2RxL0lTaG', 'moderator', 'aktif', 1, NOW(), NOW()),
('Ahmad Fahri', 'ahmad@averroes.id', '$2a$10$17w8CCEGflAGDTO.uLryNeazD37ULpLDdaz3cEN7XEdS8C2SnzHBa', 'user', 'aktif', 1, NOW(), NOW());

INSERT INTO screener (nama_aset, simbol, kategori, skor_syariah, keterangan, harga_terakhir, perubahan_24j, dibuat_pada) VALUES
('Amanah Coin', 'AMAN', 'halal', 88.50, 'Struktur token utilitas dan transparansi syariah terverifikasi', 1.2450, 2.10, NOW()),
('Mizan Token', 'MZN', 'proses', 72.30, 'Sedang proses kajian dewan pengawas syariah', 0.8420, -1.25, NOW()),
('RibaX', 'RBX', 'tidak_direkomendasikan', 35.10, 'Model bisnis mengandung unsur spekulatif tinggi', 0.1200, -4.80, NOW());

INSERT INTO screener_catatan (id_screener, judul, isi, dibuat_pada) VALUES
(1, 'Catatan Kepatuhan', 'Audit internal menyatakan kepatuhan muamalah pada lapis transaksi utama.', NOW()),
(2, 'Status Kajian', 'Menunggu publikasi ringkasan keputusan dewan syariah.', NOW());

INSERT INTO pasar (nama_aset, simbol, harga, volume_24j, perubahan_24j, kapitalisasi_pasar, diperbarui_pada) VALUES
('Amanah Coin', 'AMAN', 1.2450, 1250000, 2.10, 55000000, NOW()),
('Mizan Token', 'MZN', 0.8420, 980000, -1.25, 32000000, NOW()),
('Sukuk Chain', 'SKC', 2.1500, 1870000, 3.45, 76000000, NOW());

INSERT INTO kelas (judul, deskripsi, level, jumlah_modul, durasi_menit, thumbnail_url, status, dibuat_pada) VALUES
('Fiqh Muamalah Aset Digital', 'Memahami prinsip muamalah dalam aset digital dan kripto syariah.', 'pemula', 4, 180, 'https://picsum.photos/seed/kelas1/600/400', 'publik', NOW()),
('Analisis Risiko Syariah', 'Membedah risiko dan mitigasi syariah pada investasi aset digital.', 'menengah', 3, 150, 'https://picsum.photos/seed/kelas2/600/400', 'publik', NOW());

INSERT INTO modul (id_kelas, judul, urutan, ringkasan, durasi_menit, dibuat_pada) VALUES
(1, 'Pengantar Aset Digital Syariah', 1, 'Dasar konsep aset digital dalam perspektif muamalah.', 45, NOW()),
(1, 'Kaidah Muamalah Terapan', 2, 'Prinsip halal-haram dan gharar dalam aset digital.', 45, NOW()),
(1, 'Studi Kasus Proyek Kripto', 3, 'Menganalisis proyek nyata dari sisi kepatuhan.', 50, NOW()),
(2, 'Kerangka Risiko Syariah', 1, 'Model identifikasi risiko syariah pada investasi.', 50, NOW());

INSERT INTO materi (id_modul, judul, tipe, konten, url_video, durasi_menit, dibuat_pada) VALUES
(1, 'Definisi Aset Digital', 'teks', 'Aset digital adalah representasi nilai berbasis teknologi yang harus memenuhi prinsip muamalah.', '', 10, NOW()),
(1, 'Video Pengantar', 'video', 'Ringkasan prinsip utama aset digital syariah.', 'https://www.example.com/video1', 15, NOW()),
(2, 'Kaidah Dasar', 'teks', 'Larangan riba, maysir, dan gharar menjadi pagar utama.', '', 15, NOW());

INSERT INTO ujian (id_kelas, judul, deskripsi, durasi_menit, jumlah_soal, dibuat_pada) VALUES
(1, 'Ujian Dasar Muamalah', 'Ujian pemahaman dasar aset digital syariah.', 30, 10, NOW()),
(2, 'Ujian Risiko Syariah', 'Evaluasi risiko dan mitigasi syariah.', 25, 8, NOW());

INSERT INTO pustaka (judul_tampil, judul_asli, penulis, kategori, bahasa, jumlah_halaman, deskripsi, tautan_file) VALUES
('Hukum Fikih tentang Uang Kertas (Fiat)', 'Hukum Fiqih terhadap Uang Kertas (Fiat)', 'Penulis Terkenal', 'Ekonomi Syariah', 'Indonesia', 250, 'Kajian fikih mengenai uang kertas dan muamalah modern.', 'https://drive.google.com/file/d/1234567890abcdefg/view?usp=sharing'),
('Al-Ahkam Al-Fiqhiyyah Mata Uang Elektronik', 'Al-Ahkam Al-Fiqhiyyah Al-Mutaaliqah bil-Umalaat Al-Iliktruniyyah', 'Ahmad Al-Buhari', 'Digital Currency', 'Arab-Indonesia', 180, 'Kajian komprehensif mata uang elektronik dalam fiqh muamalah.', 'https://drive.google.com/file/d/0987654321fedcba/view?usp=sharing'),
('Manajemen Portofolio Syariah', 'Manajemen Portofolio Syariah', 'Hana Syafira', 'Investasi', 'Indonesia', 210, 'Panduan mengelola portofolio syariah untuk aset digital.', 'https://drive.google.com/file/d/portofolio123/view?usp=sharing');

INSERT INTO berita (judul, ringkasan, isi, kategori, sumber, gambar_url, diterbitkan_pada) VALUES
('Regulator Dorong Standar Crypto Syariah', 'Standar kepatuhan syariah untuk crypto semakin diperkuat.', 'Regulator dan akademisi menyepakati kerangka standar kepatuhan...', 'Regulasi', 'Averroes News', 'https://picsum.photos/seed/berita1/600/400', NOW()),
('Dewan Syariah Rilis Panduan Screener', 'Panduan screener untuk aset digital syariah dirilis.', 'Panduan ini menekankan prinsip transparansi dan manfaat...', 'Panduan', 'Averroes News', 'https://picsum.photos/seed/berita2/600/400', NOW()),
('Portofolio Muamalah 2026', 'Tren portofolio syariah aset digital tahun 2026.', 'Tren terbaru menunjukkan peningkatan minat pada aset utilitas...', 'Analisis', 'Averroes News', 'https://picsum.photos/seed/berita3/600/400', NOW());

INSERT INTO diskusi (id_pengguna, judul, isi, status, dibuat_pada) VALUES
(4, 'Bagaimana menilai proyek halal?', 'Apa indikator utama proyek kripto yang halal menurut muamalah?', 'aktif', NOW()),
(4, 'Strategi zakat portofolio', 'Bagaimana menghitung zakat dari aset digital yang volatile?', 'aktif', NOW());

INSERT INTO diskusi_balas (id_diskusi, id_pengguna, isi, dibuat_pada) VALUES
(1, 3, 'Fokus pada utility, transparansi, dan tidak ada skema riba.', NOW()),
(2, 2, 'Gunakan nilai rata-rata tahunan dan bandingkan dengan nisab.', NOW());

INSERT INTO portofolio (id_pengguna, nama_aset, simbol, jumlah, harga_beli, nilai_saat_ini, kategori, dibuat_pada) VALUES
(4, 'Amanah Coin', 'AMAN', 1000, 1.0000, 1245.00, 'halal', NOW()),
(4, 'Mizan Token', 'MZN', 500, 0.9000, 421.00, 'proses', NOW());

INSERT INTO zakat_riwayat (id_pengguna, total_nilai, nisab, persen_zakat, zakat_terhitung, dibuat_pada) VALUES
(4, 1666.00, 85000000.00, 2.50, 41.65, NOW());

INSERT INTO harga_emas (tanggal, harga_per_gram) VALUES
(CURDATE(), 1000000.00);

INSERT INTO reels (judul, tema, kutipan, sumber, url_video, thumbnail_url, dibuat_pada) VALUES
('Transaksi Amanah', 'muamalah', 'Amanah adalah pondasi muamalah.', 'QS Al-Muminun', 'https://www.example.com/reels1', 'https://picsum.photos/seed/reels1/400/600', NOW()),
('Sabar dalam Usaha', 'sabar_usaha', 'Kesabaran memperkuat ikhtiar.', 'Hadis Riwayat Ahmad', 'https://www.example.com/reels2', 'https://picsum.photos/seed/reels2/400/600', NOW()),
('Rezeki dan Takdir', 'takdir_rezeki', 'Rezeki ditakar namun usaha wajib.', 'QS Hud', 'https://www.example.com/reels3', 'https://picsum.photos/seed/reels3/400/600', NOW());

INSERT INTO tadabbur (judul, tema, ringkasan, isi, sumber, dibuat_pada) VALUES
('Amanah dalam Aset Digital', 'amanah', 'Menjaga amanah pada transaksi digital.', 'Renungkan bagaimana amanah menjaga keberkahan transaksi.', 'QS Al-Anfal', NOW()),
('Muamalah Berkeadilan', 'muamalah', 'Keadilan dalam transaksi membawa keberkahan.', 'Muamalah menuntut kejelasan, ridha, dan manfaat bersama.', 'QS An-Nisa', NOW()),
('Rezeki dan Ikhtiar', 'rezeki', 'Ikhtiar dan tawakal dalam mencari rezeki.', 'Upayakan halal dan hindari unsur gharar dalam investasi.', 'Hadis Shahih', NOW());

INSERT INTO konfigurasi (kunci, nilai, deskripsi) VALUES
('tema_utama', '#16a34a', 'Warna utama aplikasi Averroes'),
('versi_aplikasi', '1.0.0', 'Versi aplikasi saat ini');

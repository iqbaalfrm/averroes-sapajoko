package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/averroes/backend-prabogo/internal/adapter/repo/mysql"
	"github.com/averroes/backend-prabogo/pkg/config"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_ = godotenv.Load()

	cfg := config.NewConfig()

	migrationsDir := filepath.Join("..", "..", "migrations")
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal("Gagal membaca folder migrasi: ", err)
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		files = append(files, entry.Name())
	}
	sort.Strings(files)

	db, err := mysql.Open(cfg.DB.MySQLDSN())
	if err != nil {
		log.Fatal("Gagal koneksi database: ", err)
	}
	defer db.Close()

	for _, name := range files {
		path := filepath.Join(migrationsDir, name)
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatal("Gagal membaca file migrasi: ", err)
		}
		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("Gagal menjalankan migrasi %s: %v", name, err)
		}
		fmt.Println("Migrasi berhasil:", name)
	}
}

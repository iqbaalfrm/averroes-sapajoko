package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/averroes/backend-prabogo/internal/adapter/repo/mysql"
	"github.com/averroes/backend-prabogo/pkg/config"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_ = godotenv.Load()

	cfg := config.NewConfig()
	seedPath := filepath.Join("..", "..", "seed", "seed.sql")

	content, err := os.ReadFile(seedPath)
	if err != nil {
		log.Fatal("Gagal membaca seed: ", err)
	}

	db, err := mysql.Open(cfg.DB.MySQLDSN())
	if err != nil {
		log.Fatal("Gagal koneksi database: ", err)
	}
	defer db.Close()

	if _, err := db.Exec(string(content)); err != nil {
		log.Fatal("Gagal menjalankan seed: ", err)
	}

	fmt.Println("Seed berhasil dijalankan")
}

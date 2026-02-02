package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	httphandler "github.com/averroes/backend-prabogo/internal/adapter/http"
	"github.com/averroes/backend-prabogo/internal/adapter/repo/mysql"
	"github.com/averroes/backend-prabogo/internal/adapter/repo/postgres"
	"github.com/averroes/backend-prabogo/internal/usecase"
	"github.com/averroes/backend-prabogo/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.NewConfig()

	var db *sql.DB
	var err error

	// Select database driver
	if cfg.DB.DBDriver == "postgres" {
		log.Println("Menggunakan PostgreSQL/Supabase...")
		db, err = postgres.Open(cfg.DB.PostgresDSN())
		if err != nil {
			log.Fatal("Gagal koneksi database PostgreSQL: ", err)
		}
	} else {
		log.Println("Menggunakan MySQL...")
		db, err = mysql.Open(cfg.DB.MySQLDSN())
		if err != nil {
			log.Fatal("Gagal koneksi database MySQL: ", err)
		}
	}

	// Create repository based on driver
	var repo interface {
		// Define common repository interface methods here if needed
	}

	if cfg.DB.DBDriver == "postgres" {
		repo = postgres.NewRepository(db)
	} else {
		repo = mysql.NewRepository(db)
	}

	// Type assert to the MySQL repository for now (we'll create a common interface later)
	mysqlRepo := mysql.NewRepository(db)

	authUC := usecase.NewAuthUsecase(mysqlRepo, cfg.JWT.Secret)
	screenerUC := usecase.NewScreenerUsecase(mysqlRepo)
	edukasiUC := usecase.NewEdukasiUsecase(mysqlRepo)
	pustakaUC := usecase.NewPustakaUsecase(mysqlRepo)
	beritaUC := usecase.NewBeritaUsecase(mysqlRepo)
	diskusiUC := usecase.NewDiskusiUsecase(mysqlRepo)
	portofolioUC := usecase.NewPortofolioUsecase(mysqlRepo)
	zakatUC := usecase.NewZakatUsecase(mysqlRepo, mysqlRepo)
	reelsUC := usecase.NewReelsUsecase(mysqlRepo)
	tadabburUC := usecase.NewTadabburUsecase(mysqlRepo)
	adminUC := usecase.NewAdminUsecase(mysqlRepo)

	// Suppress unused variable warning
	_ = repo

	handler := &httphandler.Handler{
		AuthUsecase:       authUC,
		ScreenerUsecase:   screenerUC,
		EdukasiUsecase:    edukasiUC,
		PustakaUsecase:    pustakaUC,
		BeritaUsecase:     beritaUC,
		DiskusiUsecase:    diskusiUC,
		PortofolioUsecase: portofolioUC,
		ZakatUsecase:      zakatUC,
		ReelsUsecase:      reelsUC,
		TadabburUsecase:   tadabburUC,
		AdminUsecase:      adminUC,
		JWTSecret:         cfg.JWT.Secret,
		Versi:             "1.0.0",
	}

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
	}

	log.Printf("Server berjalan di port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

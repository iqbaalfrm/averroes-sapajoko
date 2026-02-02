package config

import (
	"os"
	"time"
)

// Config represents the application configuration
type Config struct {
	Server ServerConfig
	DB     DBConfig
	JWT    JWTConfig
}

// ServerConfig holds server-related configurations
type ServerConfig struct {
	Port       string
	Host       string
	TimeFormat string
}

// DBConfig holds database-related configurations
type DBConfig struct {
	SawitDBPath string
	// MySQL config (legacy)
	MySQLHost   string
	MySQLPort   string
	MySQLUser   string
	MySQLPass   string
	MySQLName   string
	MySQLParams string
	// PostgreSQL/Supabase config
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresSSLMode  string
	// Database driver selection
	DBDriver string // "mysql" or "postgres"
}

// JWTConfig holds JWT-related configurations
type JWTConfig struct {
	Secret string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:       getEnvOrDefault("SERVER_PORT", "8080"),
			Host:       getEnvOrDefault("SERVER_HOST", "localhost"),
			TimeFormat: time.Now().Format(time.RFC3339),
		},
		DB: DBConfig{
			SawitDBPath: getEnvOrDefault("SAWIT_DB_PATH", "./data.sawit"),
			// MySQL config (legacy)
			MySQLHost:   getEnvOrDefault("MYSQL_HOST", "127.0.0.1"),
			MySQLPort:   getEnvOrDefault("MYSQL_PORT", "3306"),
			MySQLUser:   getEnvOrDefault("MYSQL_USER", "averroes"),
			MySQLPass:   getEnvOrDefault("MYSQL_PASS", "averroes123"),
			MySQLName:   getEnvOrDefault("MYSQL_NAME", "averroes_db"),
			MySQLParams: getEnvOrDefault("MYSQL_PARAMS", "parseTime=true&charset=utf8mb4&loc=Local&multiStatements=true"),
			// PostgreSQL/Supabase config
			PostgresHost:     getEnvOrDefault("POSTGRES_HOST", ""),
			PostgresPort:     getEnvOrDefault("POSTGRES_PORT", "5432"),
			PostgresUser:     getEnvOrDefault("POSTGRES_USER", "postgres"),
			PostgresPassword: getEnvOrDefault("POSTGRES_PASSWORD", ""),
			PostgresDBName:   getEnvOrDefault("POSTGRES_DBNAME", "postgres"),
			PostgresSSLMode:  getEnvOrDefault("POSTGRES_SSLMODE", "require"),
			// Database driver selection
			DBDriver: getEnvOrDefault("DB_DRIVER", "postgres"),
		},
		JWT: JWTConfig{
			Secret: getEnvOrDefault("JWT_SECRET", "default_secret_key_for_development"),
		},
	}
}

func (db DBConfig) MySQLDSN() string {
	return db.MySQLUser + ":" + db.MySQLPass + "@tcp(" + db.MySQLHost + ":" + db.MySQLPort + ")/" + db.MySQLName + "?" + db.MySQLParams
}

func (db DBConfig) PostgresDSN() string {
	return "host=" + db.PostgresHost +
		" port=" + db.PostgresPort +
		" user=" + db.PostgresUser +
		" password=" + db.PostgresPassword +
		" dbname=" + db.PostgresDBName +
		" sslmode=" + db.PostgresSSLMode
}

// getEnvOrDefault retrieves environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

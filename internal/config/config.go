package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort       string
	MysqlDSN       string
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioBucket    string
	DataDir        string
	VersionDataDir string
	ApkDir         string
}

var Cfg *Config

func Load() {
	_ = godotenv.Load()

	dataDir := getEnv("DATA_DIR", "./data")

	Cfg = &Config{
		HTTPPort:       getEnv("APP_PORT", "8080"),
		MysqlDSN:       getEnv("MYSQL_DSN", ""),
		MinioEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinioAccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinioSecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin123"),
		MinioBucket:    getEnv("MINIO_BUCKET", "otaku-backups"),
		DataDir:        dataDir,
		VersionDataDir: dataDir + "/version/android", // Assuming this structure
		ApkDir:         dataDir + "/apks",            // Assuming this structure
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

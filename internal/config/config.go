package config

import (
	"os"
)

type Config struct {
	HTTPPort       string
	BaseURL        string
	VersionDataDir string
	ApkDir         string
}

var Cfg Config

func Load() {
	Cfg = Config{
		HTTPPort:       getEnv("HTTP_PORT", "8080"),
		BaseURL:        getEnv("BASE_URL", "http://localhost:8080"),
		VersionDataDir: getEnv("VERSION_DATA_DIR", "/data/version"),
		ApkDir:         getEnv("APK_DIR", "/data/apks"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

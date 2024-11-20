package db

import (
	"github.com/saintbyte/postgresURItoDSN"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

var Database *gorm.DB

func GetPGUrl() string {
	value, ok := os.LookupEnv("DATABASE_URL")
	if ok {
		return value
	}
	return ""
}

func ConnectPG() error {
	slog.Error("GetPGUrl:", GetPGUrl())
	dsn, err := postgresURItoDSN.UriToDSN(GetPGUrl())
	if err != nil {
		slog.Error("Database url error:", err)
		return err
	}
	Database, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		slog.Error("Database connect error:", err)
		return err
	}
	return nil
}

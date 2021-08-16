package config

import (
	"os"

	"github.com/abdulsalam01/go_gorm/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Name     string
	Host     string
	Username string
	Password string
}

func InitEnvironment() {
	godotenv.Load()
}

func InitConnection() (*gorm.DB, error) {
	dbConfig := DatabaseConfig{
		Name:     os.Getenv("DATABASE_NAME"),
		Host:     os.Getenv("DATABASE_HOST"),
		Username: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
	}

	return gorm.Open(sqlite.Open(dbConfig.Name), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func InitSchema(db *gorm.DB) {
	db.Logger.LogMode(logger.Info)

	db.AutoMigrate(&models.User{}, &models.Contact{})
	db.Set("gorm:table_options", "ENGINE=InnoDB")
}

func InitPort() string {
	return os.Getenv("PORT")
}

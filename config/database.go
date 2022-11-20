package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Secure   string
}

var Db *gorm.DB

func InitDb() {
	config := DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "complexpassword",
		Name:     "postgres",
		Secure:   "disable",
	}

	_ = setup(config)
}

func InitTestDb() {
	config := DatabaseConfig{
		Host:     "localhost",
		Port:     "5433",
		User:     "postgres",
		Password: "VerySecurePassword",
		Name:     "postgres",
		Secure:   "disable",
	}

	_ = setup(config)
}

func GetDb() *gorm.DB {
	return Db
}

func ResetDb() {
	Db = nil
}

func setup(config DatabaseConfig) error {
	var err error

	creds := generateCreds(config)
	Db, err = gorm.Open("postgres", creds)

	if err != nil {
		return err
	}

	Db.DB().SetMaxIdleConns(5)

	return nil
}

func generateCreds(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Name, config.Password, config.Secure)
}

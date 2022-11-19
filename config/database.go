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

var db *gorm.DB

func InitDb() {
	// TODO: implement the function
}

func InitTestDb() {
	config := DatabaseConfig{
		Host:     "localhost",
		Port:     "5433",
		User:     "postgres",
		Password: "VerySecurePassword",
		Name:     "postgres",
		Secure:   "false",
	}

	_ = setup(config)
}

func setup(config DatabaseConfig) error {
	creds := generateCreds(config)
	db, err := gorm.Open("postgres", creds)

	if err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(5)

	return nil
}

func generateCreds(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Name, config.Password, config.Secure)
}

func GetDb() *gorm.DB {
	return db
}

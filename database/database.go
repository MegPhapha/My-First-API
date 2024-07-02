package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Driver   string
}

type Migrations struct {
	DB     *gorm.DB
	Models []interface{}
}

func InitDb() *gorm.DB {

	dbConfig := DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "test.db",
		Driver:   "sqlite",
	}
	db := NewDatabaseConnection(dbConfig)
	return db
}

func NewDatabaseConnection(config DatabaseConfig) *gorm.DB {
	var dsn string
	switch config.Driver {

	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.DbName, config.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {

			fmt.Println("Failed to connect database")
			return nil
		}
		fmt.Println("Database connected")
		return db

	case "sqlite":
		dsn = config.DbName
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println("Database connected")
		return db

	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.DbName, config.Driver)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println("Database connected")
		return db

	default:
		fmt.Println("Unsupported database")

		return nil

	}
}

func RunMigrations(migrations Migrations) {
	if migrations.DB == nil {
		fmt.Println("Database connection is nil, cannot run migration")
		return
	}
	for _, model := range migrations.Models {
		err := migrations.DB.AutoMigrate(model)
		if err != nil {
			fmt.Println("error migrating models")
		}
	}
}

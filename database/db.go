package database

import (
	"dts/learn_middleware/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5433"
	dbname   = "universal"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := "root:Secret123@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	//db.Debug().Migrator().DropTable(models.User{}, models.Product{})
	db.Debug().AutoMigrate(models.User{}, models.Product{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error :", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}

func GetDB() *gorm.DB {
	return db
}

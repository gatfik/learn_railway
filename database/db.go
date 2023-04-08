package database

import (
	"dts/learn_middleware/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	host     = os.Getenv("MYSQLHOST")
	user     = os.Getenv("MYSQLUSER")
	password = os.Getenv("MYSQLPASSWORD")
	dbPort   = os.Getenv("MYSQLPORT")
	dbname   = os.Getenv("MYSQLDATABASE")
	db       *gorm.DB
	err      error
)

func StartDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbname)
	fmt.Println("dsn : ", dsn)
	//dsn := "root:Secret123@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
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

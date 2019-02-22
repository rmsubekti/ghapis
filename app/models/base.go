package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //impot gorm dialect postgres
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")
	dbType := os.Getenv("DB_TYPE")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println("Using ENV : " + dbURI)

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(dbType, dbURI) // gorm checks Ping on Open
		if err == nil {
			fmt.Println("Database Connected ...")
			break
		}
		fmt.Println("Retrying to connect to database ... ...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Account{}, &Role{}, &Product{})

	//initial role rows
	db.FirstOrCreate(&Role{}, Role{RoleName: RoleUser})
	db.FirstOrCreate(&Role{}, Role{RoleName: RoleAdmin})
	db.FirstOrCreate(&Role{}, Role{RoleName: RolePM})
}

// GetDB export
func GetDB() *gorm.DB {
	return db
}

package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // here
)

var DB *gorm.DB

func Init() {
	dbString := getDatabaseString()
	connect(dbString, false)
}

func InitTest() {
	dbString := getTestDatabaseString()
	connect(dbString, false)
}

func connect(dbString string, logMode bool) {
	var err error

	DB, err = gorm.Open("postgres", dbString)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&User{})
	DB.LogMode(logMode)
}

func getDatabaseString() (dbString string) {
	dbString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("ETH_RBAC_DB_HOST"),
		os.Getenv("ETH_RBAC_DB_PORT"),
		os.Getenv("ETH_RBAC_DB_USERNAME"),
		os.Getenv("ETH_RBAC_DB_NAME"),
		os.Getenv("ETH_RBAC_DB_PASSWORD"))

	return
}

func getTestDatabaseString() (dbString string) {
	dbString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("ETH_RBAC_DB_HOST"),
		os.Getenv("ETH_RBAC_DB_PORT"),
		os.Getenv("ETH_RBAC_DB_TEST_USERNAME"),
		os.Getenv("ETH_RBAC_DB_TEST_NAME"),
		os.Getenv("ETH_RBAC_DB_TEST_PASSWORD"))

	return
}

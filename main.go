package main

import (
	"fmt"
	"log"
	"os" //Package os provides a platform-independent interface to operating system functionality.

	"github.com/jinzhu/gorm" // this is gorm stable version v1.9.16 (you can even use gorm-version 2 as well)

	//blank import used for postgres db connection string
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// we declare the db and err variables global so that we can use them out of this package as well when needed.
var (
	db  *gorm.DB
	err error
)

// DBConfig stores database related configurations
type DBConfig struct {
	UserName string
	Password string
	DbName   string
	SslMode  string
}

func main() {

	var config DBConfig
	// os.Getenv("environment_variable") you will place your env-variable in it os.Getenv("") and it will return it's value
	dbUser := os.Getenv("YOUR_DB_USER_NAME")
	dbUserPwd := os.Getenv("YOUR_DB_USER_PASSWORD")
	dbName := os.Getenv("YOUR_DB_NAME")
	sslmode := os.Getenv("YOUR_DB_SSL_MODE")

	// we are now storing all environment variables values in DBConfig struct
	config = DBConfig{UserName: dbUser, Password: dbUserPwd, DbName: dbName, SslMode: sslmode}

	//Build connection string
	dbURI := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s", config.UserName, config.DbName, config.SslMode, config.Password)

	// open db connection here
	db, err = gorm.Open("postgres", dbURI)

	// verify connection and catch the error if occurs
	if err != nil {
		log.Println("Connection Failed to Open due to following error => ", err)
	} else {
		log.Println("Connection Established")
	}
}

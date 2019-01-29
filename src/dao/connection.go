package dao

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	mgo "gopkg.in/mgo.v2"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var postgreDb *gorm.DB
var mariaDb *gorm.DB
var mongoDb *mgo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("get string\n - ok: val: %v\n", os.Getenv("MYSQL_URL"))
	mariaDb = InitMaria()
	mongoDb = InitMongo()
}

// var mongoDB *mongo.Client

// InitPostgre creates and migrates the database
func InitPostgre() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	postgreDb, err := gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME")+" password="+os.Getenv("DB_PASS"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("he error")
		return nil
	}
	postgreDb.LogMode(true)
	return postgreDb
}

// InitMaria creates and migrates the database
func InitMaria() *gorm.DB {
	mariaDb, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	mariaDb.LogMode(true)

	return mariaDb
}

// InitMongo func
func InitMongo() *mgo.Database {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("connected!")
	mongoDb = session.DB(os.Getenv("MONGO_DBNAME"))

	return mongoDb
}

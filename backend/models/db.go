package models

import (
	"fmt"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(host string, username string, password string, name string, port string) {
	// PRODUCTION
	cleanup, err := pgxv5.RegisterDriver("cloudsql-postgres")
	if err != nil {
			panic(err)
	}
	defer cleanup()
	fmt.Println(host,username,password,name,port)
	dsnGcloud := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",host,username,password,name)
	cloudSqlDb,err:=gorm.Open(postgres.New(
		postgres.Config{
			DriverName: "cloudsql-postgres",
			DSN: dsnGcloud,
		}))
	if err!=nil{
		panic("Failed to connect to database!")
	}
	err = cloudSqlDb.AutoMigrate(&User{}, &Job{},&JobDB{})
	if err != nil {
		return
	}
	DB=cloudSqlDb
	fmt.Println(DB)
	fmt.Println("Database up and running.........")

	// LOCAL
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, name, port)
	// database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to database!")
	// }
	// err = database.AutoMigrate(&User{}, &Job{}, &JobDB{})
	// if err != nil {
	// 	return
	// }
	// DB = database
	// fmt.Println(DB)
	// fmt.Println("Database up and running.........")
}

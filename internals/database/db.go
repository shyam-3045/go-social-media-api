package database

import (
	"fmt"

	"socio/models/friendships"
	"socio/models/posts"
	"socio/models/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Config() {
	dsn := "host=localhost port=5432 user=postgres password=shyam123 dbname=socio sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Unable to open  database , error :%v", err)
		panic(err)
	}

	sql, err := db.DB()

	if err != nil {
		fmt.Printf("\nUnable to get database from gorm , error : %v", err)
		panic(err)
	}

	if err := sql.Ping(); err != nil {
		fmt.Printf("Unable to connect to database , error :%v", err)
		panic(err)
	}

	DB = db 

	DB.AutoMigrate(&users.Users{} , &friendships.Friendships{} ,&posts.Posts{})

}

package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbPostgre  *gorm.DB
	errPostgre error
)

func InitPostgre() {
	var (
		host     = "127.1.1.1"
		dbname   = "test_kenda"
		port     = 5432
		user     = "postgres"
		password = "1980000"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	dbPostgre, errPostgre = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errPostgre != nil {
		fmt.Println(errPostgre.Error())
		panic("Can't connect to Postgre Database.")
	}
}

func GetPostgre() *gorm.DB {
	return dbPostgre
}

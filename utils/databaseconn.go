package utils

import(
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)


func Connectdb()(*gorm.DB){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	fmt.Print("\n connected to the database sucessfully \n")

	if err != nil {
		panic("failed to connect database")
	}

	return  db

}
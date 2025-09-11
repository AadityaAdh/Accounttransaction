package repository

import (
  "gorm.io/driver/sqlite" 
  "gorm.io/gorm"
)



func Createdatabaseconnection(dbname string)(*gorm.DB,error){
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})

	if err!=nil{
		return db,err
	}
	return db,nil

}
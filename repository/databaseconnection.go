package repository

import (
  "gorm.io/driver/sqlite" 
  "gorm.io/gorm"
)


type Database interface{
	Connectdb() (*gorm.DB,error)
}

type Sqlitedb struct{
	Dbname string
}

func (s Sqlitedb) Connectdb()(*gorm.DB,error){
	db, err := gorm.Open(sqlite.Open(s.Dbname), &gorm.Config{})

	if err!=nil{
		return db,err
	}
	return db,nil

}


func Getdatabase(db Database)(*gorm.DB,error){
	dbconn,err:=db.Connectdb()
	if err!=nil{
		return dbconn,err
	}
	return dbconn,err




}




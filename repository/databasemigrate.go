package repository

import (
	"gorm.io/gorm"
)



func Createtableifnotexists(db *gorm.DB)error{

	accounttemplate:=Account{}

	transactiontemplate:=Transaction{}

	if err := db.AutoMigrate(&accounttemplate, &transactiontemplate); err != nil {
        return err
    }

    return nil

	
}
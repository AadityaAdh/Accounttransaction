package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)


func Addaccount(account Account,db *gorm.DB)error{

	result:=db.Create(&account)

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return errors.New("your row was not inserted")
	}


	return nil


}

func Updateaccount(account Account,db *gorm.DB)(error){
	existingaccount,err:=Getaccount(account.Accountid,db)

	if err!=nil{
		return errors.New("could not even fetch the single row with this id let along update")

	}

	if account.Balance!=0{
		existingaccount.Balance=account.Balance
	}

	if account.Owner!=""{
		existingaccount.Owner=account.Owner
	}

	if account.Status!=""{
		existingaccount.Status=account.Status
	}

	//naya xa vanae tyo natra jasta ko testai

	fmt.Print("\n Existingaccount\n")

	fmt.Print(existingaccount)

	result:=db.Model(&existingaccount).Where("accountid=?",existingaccount.Accountid).Updates(existingaccount)


	


	if result.Error!=nil{
		return result.Error
	}
	return nil




}



func Getallaccount(db *gorm.DB)([]Account,error){
	var accounts []Account

	result:=db.Find(&accounts)

	if result.Error!=nil{
		return accounts,result.Error

	}
	return accounts,nil

}

func Getaccount(accountid int32,db *gorm.DB)(Account,error){

	var account=Account{}
	result:=db.First(&account,"accountid=?",accountid)

	if result.Error!=nil{
		return  account,result.Error
	}
	return account,nil




}


func Deleteaccount(accountid int32,db *gorm.DB)error{
	var account Account
	result:=db.Where("accountid=?",accountid).Delete(&account)


	if result.Error!=nil{
		return result.Error
	}
	return nil



}


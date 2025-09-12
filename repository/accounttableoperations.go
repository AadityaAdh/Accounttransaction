package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Accountrepo struct{
	Db *gorm.DB
}


func (ar Accountrepo) Addaccount(account Account)error{

	result:=ar.Db.Create(&account)

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return errors.New("your row was not inserted")
	}


	return nil


}

func (ar Accountrepo) Updateaccount(account Account)(error){
	existingaccount,err:=ar.Getaccount(account.Accountid)

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

	result:=ar.Db.Model(&existingaccount).Where("accountid=?",existingaccount.Accountid).Updates(existingaccount)


	


	if result.Error!=nil{
		return result.Error
	}
	return nil




}



func (ar Accountrepo) Getallaccount()([]Account,error){
	var accounts []Account

	result:=ar.Db.Find(&accounts)

	if result.Error!=nil{
		return accounts,result.Error

	}
	return accounts,nil

}

func (ar Accountrepo) Getaccount(accountid int32)(Account,error){

	var account=Account{}
	result:=ar.Db.First(&account,"accountid=?",accountid)

	if result.Error!=nil{
		return  account,result.Error
	}
	return account,nil




}


func (ar Accountrepo) Deleteaccount(accountid int32)error{
	var account Account
	result:=ar.Db.Where("accountid=?",accountid).Delete(&account)


	if result.Error!=nil{
		return result.Error
	}
	return nil



}


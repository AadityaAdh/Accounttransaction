package service

import (
	"acctraproject/protogen"
	"acctraproject/repository"
	"acctraproject/utils"

	"gorm.io/gorm"
)



func Addaccountservice(account repository.Account,db *gorm.DB)(string,error){

	err:=repository.Addaccount(account,db)

	if err!=nil{
		return "",err
	}
	tokenstring,err:=utils.Generatetoken(account.Owner)
	return tokenstring,err

}


func Updateaccountservice(account repository.Account,db *gorm.DB)error{
	err:=repository.Updateaccount(account,db)
	return err
	//done
}


func Getallaccountservice(db *gorm.DB)(*protogen.Getallaccountsmsg,error){
	accounts,err:=repository.Getallaccount(db)

	if err!=nil{
		return &protogen.Getallaccountsmsg{},err
	}

	//but we need every element of type *protogen.Getallaccountsmsg to return from that getallaccountservice handler
	//so lets do that conversion

	//from list of accounts to *protogen.Getallaccountsmsg

	//convering accounts to accountsmsg
	accountsmsg:=make([]*protogen.Accountmsg,len(accounts))

	for _,acc:=range accounts{
		accountsmsg=append(accountsmsg, &protogen.Accountmsg{
			Status: acc.Status,
			Owner: acc.Owner,
			Balance: acc.Balance,
			AccountId: acc.Accountid,
		})
	}

	return &protogen.Getallaccountsmsg{Listaccs: accountsmsg},nil


	//done
}

func Getaccountservice(account repository.Account,db *gorm.DB)(repository.Account,error){
	account,err:=repository.Getaccount(account.Accountid,db)
	return account,err
	//done
}

func Deleteaccountservice(accountid int32 ,db *gorm.DB)error{
	err:=repository.Deleteaccount(accountid,db)
	return err
	//done
}
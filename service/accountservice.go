package service

import (
	"acctraproject/protogen"
	"acctraproject/repository"
	"acctraproject/utils"
	// "gorm.io/gorm"
)


type Accountservice struct{
	Repo repository.Accountrepo
}



func(s Accountservice) Addaccountservice(account repository.Account)(string,error){

	err:=s.Repo.Addaccount(account)

	if err!=nil{
		return "",err
	}
	tokenstring,err:=utils.Generatetoken(account.Owner)
	return tokenstring,err

}


func(s Accountservice) Updateaccountservice(account repository.Account)error{
	err:=s.Repo.Updateaccount(account)
	return err
	//done
}


func (s Accountservice) Getallaccountservice()(*protogen.Getallaccountsmsg,error){
	accounts,err:=s.Repo.Getallaccount()

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

func (s Accountservice) Getaccountservice(account repository.Account)(repository.Account,error){
	account,err:=s.Repo.Getaccount(account.Accountid)
	return account,err
	//done
}

func (s Accountservice) Deleteaccountservice(accountid int32)error{
	err:=s.Repo.Deleteaccount(accountid)
	return err
	//done
}


func Returnnewaccountservice(repo repository.Accountrepo)Accountservice{
	return Accountservice{Repo: repo}



}
package handler

import (
	"acctraproject/protogen"
	"acctraproject/repository"
	"acctraproject/service"
	"context"
	"fmt"

	"gorm.io/gorm"
)



type Server struct{
	protogen.UnimplementedAccountserviceServer
}

var db *gorm.DB

func init(){
	var err error
	db,err=repository.Createdatabaseconnection("mydb1.db")
	if err!=nil{
		fmt.Print("Could not create db connection")
	}
	err=repository.Createtableifnotexists(db)

	if err!=nil{
		fmt.Print("could not create the table")
	}


}

func converter(req *protogen.Accountmsg)repository.Account{
	
	
	account:=repository.Account{
		Accountid: req.AccountId,
		Owner: req.Owner,
		Balance: req.Balance,
		Status: req.Status,
	}
	return account

}

func optionconverter(req *protogen.Updateaccountmsg)repository.Account{

	var owner string
	var balance int32
	var status string
	


	if req.Balance!=nil{
		balance=*req.Balance
	}

	if req.Owner!=nil{
		owner=*req.Owner
	}
	if req.Status!=nil{
		status=*req.Status
	}

	return  repository.Account{
		Accountid: req.AccountId,
		Owner: owner,
		Balance: balance,
		Status: status,
	}
}




func (s *Server) Addaccount(ctx context.Context,req *protogen.Accountmsg) (*protogen.Accountaddedmsg, error){

	fmt.Print("Hit Add account")

	account:=converter(req)

	
	tokenstring,err:=service.Addaccountservice(account,db)

	if err!=nil{

		return &protogen.Accountaddedmsg{Message: "error in generating token",Tokenstring: tokenstring},nil

	}

	fmt.Print("\n account added \n")

	return &protogen.Accountaddedmsg{Message: "done",Tokenstring: tokenstring},nil

}


func (s *Server) Updateaccount(ctx context.Context,req *protogen.Updateaccountmsg) (*protogen.Statusmsg, error){

	fmt.Print("\n update accounthandler hit\n")

	fmt.Print(req.Balance)

	account:=optionconverter(req)

	err:=service.Updateaccountservice(account,db)

	if err!=nil{
		return &protogen.Statusmsg{Status: 400,Message: "Update Unsucessful"},err

	}


	return &protogen.Statusmsg{Status: 200,Message: "Update done"},nil




}


func (s *Server) Getallaccounts(ctx context.Context,req *protogen.Getallaccountsmsg) (*protogen.Getallaccountsmsg, error){

	getaccountmsg,err:=service.Getallaccountservice(db)

	if err!=nil{
		return &protogen.Getallaccountsmsg{},err
	}

	

	



	return getaccountmsg,nil







}




func (s *Server) Deleteaccount(ctx context.Context,req *protogen.Deleteaccountmsg) (*protogen.Statusmsg, error){
	err:=service.Deleteaccountservice(req.AccountId,db)

	if err!=nil{
		return &protogen.Statusmsg{Status: 400,Message: "delete unsucess"},err

	}


	return &protogen.Statusmsg{Status: 200,Message: "deleted sucessfully"},nil
}
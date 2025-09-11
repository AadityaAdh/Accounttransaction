package handler

import (
	"fmt"
	tp "acctraproject/transactionprotogen"
	"context"
	"acctraproject/repository"
	"acctraproject/service"
)



func tmtotconverter(req *tp.Transactionmsg)repository.Transaction{
	
	
	transaction:=repository.Transaction{
		Accountid: req.AccountId,
		Transactionid: req.Transactionid,
		Transactiontype: req.Type,
		Amount: req.Amount,
	}
	return transaction

}

func umtotconverter(req *tp.Updatetransactionmsg)repository.Transaction{

	var mtype string
	var amount int32
	


	if req.Type!=nil{
		mtype=*req.Type
	}

	if req.Amount!=nil{
		amount=*req.Amount
	}
	

	return  repository.Transaction{
		Transactionid: req.Transactionid,
		Amount: amount,
		Transactiontype: mtype,
	}
}











func (s *Servertran) Addtransaction(ctx context.Context,req *tp.Transactionmsg) (*tp.Transactionaddedmsg, error){

	fmt.Print("Hit Add Transactions")

	transaction:=tmtotconverter(req)

	
	tokenstring,err:=service.Addtransactionservice(transaction,db)

	if err!=nil{

		return &tp.Transactionaddedmsg{Message: "error in generating token",Tokenstring: tokenstring},nil

	}

	fmt.Print("\n Transaction added \n")

	return &tp.Transactionaddedmsg{Message: "done",Tokenstring: tokenstring},nil




}



func (s *Servertran) Updatetransaction(ctx context.Context,req *tp.Updatetransactionmsg) (*tp.Transactionstatusmsg, error){

	fmt.Print("\n update transactionhandler hit\n")

	

	transaction:=umtotconverter(req)

	err:=service.Updatetransactionservice(transaction,db)

	if err!=nil{
		return &tp.Transactionstatusmsg{Status: 400,Message: "Update Unsucessful"},err

	}


	return &tp.Transactionstatusmsg{Status: 200,Message: "Update done"},nil




}


func (s *Servertran) Getalltransaction(ctx context.Context,req *tp.Getalltransactionsmsg) (*tp.Getalltransactionsmsg, error){

	gettransmsg,err:=service.Getalltransactionservice(db)

	if err!=nil{
		return &tp.Getalltransactionsmsg{},err
	}

	

	



	return gettransmsg,nil







}


func (s *Servertran) Deletetransaction(ctx context.Context,req *tp.Deletetransactionmsg) (*tp.Transactionstatusmsg, error){
	err:=service.Deletetransactionservice(req.Transactionid,db)

	if err!=nil{
		return &tp.Transactionstatusmsg{Status: 400,Message: "delete unsucess"},err

	}


	return &tp.Transactionstatusmsg{Status: 200,Message: "deleted sucessfully"},nil
}



















type Servertran struct{
	tp.UnimplementedTransactionserviceServer
}
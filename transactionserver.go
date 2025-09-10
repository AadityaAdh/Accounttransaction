package main


import (
	"fmt"
	"net"
	tp "acctraproject/transactionprotogen"
	"google.golang.org/grpc"
	"context"
	u "acctraproject/utils"
)











func (s *server) Addtransaction(ctx context.Context,req *tp.Transactionmsg) (*tp.Transactionaddedmsg, error){

	fmt.Print("Hit Add Transaction")

	db:=u.Connectdb() //yeutai package ma xadai xa ni yo ta so use garnamili halyo import na gari


	//creating table if not exists from the Accountmsg struct

	transactiontemplate:=tp.Transactionmsg{}


	db.AutoMigrate(&transactiontemplate)


	db.Create(&req)

	tokenstring,err:=u.Generatetoken(string(req.AccountId))

	if err!=nil{

		return &tp.Transactionaddedmsg{Message: "error in generating token",Tokenstring: tokenstring},nil



	}

	fmt.Print("\n transaction added \n")




	return &tp.Transactionaddedmsg{Message: "done",Tokenstring: tokenstring},nil




}

func tranauthenticationInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error){

	fmt.Print(info.FullMethod)

	if info.FullMethod != "/accounttransaction.Transactionservice/Addtransaction" && info.FullMethod != "/accounttransaction.Transactionservice/Getalltransaction" {
		fmt.Print("\n token step taken \n")
		tokenstring:=req.(*tp.Updatetransactionmsg).Tokenstring //type assertion vando raixa yo chai 
		tokenverified:=u.Verifytoken(tokenstring)
		if tokenverified{

			res,err:=handler(ctx,req)
			return res,err
		}
		return &tp.Transactionstatusmsg{Status: 400,Message: "Token could not be verified"},nil
	}
	
	res,err:=handler(ctx,req)
	return res,err

	

}

func (s *server) Updatetransaction(ctx context.Context,req *tp.Updatetransactionmsg) (*tp.Transactionstatusmsg, error){

	var transaction tp.Transactionmsg //note yesko type le chai kun table ma operation vai raa xa vanni hunxa
	//yo struct ko name aausar variable ko chai haina 

	//first fetch that row
	//if you didnt do this this account is empty struct so empty value will be placed

	db:=u.Connectdb()

	db.First(&transaction,"transactionid=?",req.Transactionid)

	

	

	if req.Amount!=nil{
		transaction.Amount=*req.Amount
	}
	if req.Type!=nil{
		transaction.Type=*req.Type
	}

	

	


	// db.Save(&account)

	db.Model(&transaction).Where("transactionid=?",req.Transactionid).Updates(transaction) //lock sock aaru k k field ni hola tesko warning aako

	fmt.Print("\n update done \n")


	return &tp.Transactionstatusmsg{Status: 200,Message: "Update done"},nil




}


func (s *server) Getalltransaction(ctx context.Context,req *tp.Getalltransactionsmsg) (*tp.Getalltransactionsmsg, error){

	var transaction []*tp.Transactionmsg

	db:=u.Connectdb()

	db.Find(&transaction)

	fmt.Print("found transaction are \n")
	fmt .Print(transaction)


	return &tp.Getalltransactionsmsg{Listtran: transaction},nil







}


func (s *server) Deletetransaction(ctx context.Context,req *tp.Deletetransactionmsg) (*tp.Transactionstatusmsg, error){
	var transaction tp.Transactionmsg //table sanga match garai rana yesko banako hai

	db:=u.Connectdb()

	db.Where("transactionid=?",req.Transactionid).Delete(&transaction)

	//that &account means
	//find the type name of that struct/interface
	//then puralize that
	//that will be the table we are querying on


	return &tp.Transactionstatusmsg{Status: 200,Message: "deleted sucessfully"},nil
}



















type server struct{
	tp.UnimplementedTransactionserviceServer
}

func main(){
	lis, err := net.Listen("tcp", ":50068")

	if err!=nil{
		fmt.Print(err)
	}

	grpcserver:=grpc.NewServer(grpc.ChainUnaryInterceptor(tranauthenticationInterceptor))

	tp.RegisterTransactionserviceServer(grpcserver,&server{})


	servererr:=grpcserver.Serve(lis);

	if servererr!=nil{
		fmt.Print(servererr)
	}

}
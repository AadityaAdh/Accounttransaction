package main


import (
	"fmt"
	"net"
	"acctraproject/protogen"
	"google.golang.org/grpc"
	"context"
	u "acctraproject/utils"
	
	
)






func (s *server) Addaccount(ctx context.Context,req *protogen.Accountmsg) (*protogen.Accountaddedmsg, error){

	fmt.Print("Hit Add account")

	db:=u.Connectdb()


	//creating table if not exists from the Accountmsg struct

	accounttemplate:=protogen.Accountmsg{}

	db.AutoMigrate(&accounttemplate)


	db.Create(&req)

	tokenstring,err:=u.Generatetoken(req.Owner)

	if err!=nil{

		return &protogen.Accountaddedmsg{Message: "error in generating token",Tokenstring: tokenstring},nil



	}

	fmt.Print("\n account added \n")




	return &protogen.Accountaddedmsg{Message: "done",Tokenstring: tokenstring},nil




}

func authenticationInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error){

	fmt.Print(info.FullMethod)

	if info.FullMethod != "/accounttransaction.Accountservice/Addaccount" && info.FullMethod != "/accounttransaction.Accountservice/Getallaccounts" {
		fmt.Print("\n token step taken \n")
		tokenstring:=req.(*protogen.Updateaccountmsg).Tokenstring //type assertion vando raixa yo chai 
		tokenverified:=u.Verifytoken(tokenstring)
		if tokenverified{

			res,err:=handler(ctx,req)
			return res,err
		}
		return &protogen.Statusmsg{Status: 400,Message: "Token could not be verified"},nil
	}
	
	res,err:=handler(ctx,req)
	return res,err

	

}

func (s *server) Updateaccount(ctx context.Context,req *protogen.Updateaccountmsg) (*protogen.Statusmsg, error){

	var account protogen.Accountmsg //note yesko type le chai kun table ma operation vai raa xa vanni hunxa
	//yo struct ko name aausar variable ko chai haina 

	//first fetch that row
	//if you didnt do this this account is empty struct so empty value will be placed

	db:=u.Connectdb()

	db.First(&account,"account_id=?",req.AccountId)

	

	

	if req.Balance!=nil{
		account.Balance=*req.Balance
	}
	if req.Owner!=nil{
		account.Owner=*req.Owner
	}

	if req.Status!=nil{
		account.Status=*req.Status
	}

	


	// db.Save(&account)

	db.Model(&account).Where("account_id=?",req.AccountId).Updates(account) //lock sock aaru k k field ni hola tesko warning aako

	fmt.Print("\n update done \n")


	return &protogen.Statusmsg{Status: 200,Message: "Update done"},nil




}


func (s *server) Getallaccounts(ctx context.Context,req *protogen.Getallaccountsmsg) (*protogen.Getallaccountsmsg, error){

	var accounts []*protogen.Accountmsg

	db:=u.Connectdb()

	db.Find(&accounts)

	fmt.Print("found accounts are \n")
	fmt .Print(accounts)


	return &protogen.Getallaccountsmsg{Listaccs: accounts},nil







}


func (s *server) Deleteaccount(ctx context.Context,req *protogen.Deleteaccountmsg) (*protogen.Statusmsg, error){
	var account protogen.Accountmsg //table sanga match garai rana yesko banako hai

	db:=u.Connectdb()

	db.Where("account_id=?",req.AccountId).Delete(&account)

	//that &account means
	//find the type name of that struct/interface
	//then puralize that
	//that will be the table we are querying on


	return &protogen.Statusmsg{Status: 200,Message: "deleted sucessfully"},nil
}



















type server struct{
	protogen.UnimplementedAccountserviceServer
}

func main(){
	lis, err := net.Listen("tcp", ":50063")

	if err!=nil{
		fmt.Print(err)
	}

	grpcserver:=grpc.NewServer(grpc.ChainUnaryInterceptor(authenticationInterceptor))

	protogen.RegisterAccountserviceServer(grpcserver,&server{})

	servererr:=grpcserver.Serve(lis);

	if servererr!=nil{
		fmt.Print(servererr)
	}

}
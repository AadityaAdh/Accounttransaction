package main

import(
	"fmt"
	tp "acctraproject/transactionprotogen"
	"google.golang.org/grpc"
	"context"
	"time"
)





func main(){

	fmt.Println("Client has started \n")

	conn, err := grpc.Dial("localhost:50068", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v \n", err)
	}
	defer conn.Close()

	


	// client:=protogen.NewAccountserviceClient(conn)

	client:=tp.NewTransactionserviceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()


	//now start calling from here

	mrespose,err:=client.Addtransaction(ctx,&tp.Transactionmsg{
		AccountId: 14,
		Transactionid: 15,
		Type: "cash",
		Amount: 1000,
		


	})

	if err!=nil{
		fmt.Print("\n error in add account response\n")
		fmt.Print(err)
	}
	fmt.Print(mrespose)

	mytokenstring:=mrespose.Tokenstring

	//optional field chai address pass garnu pardo raixa

	new_balance:=int32(30000) //this int 32 is also required


	updaterespose,err:=client.Updatetransaction(ctx,&tp.Updatetransactionmsg{
		Transactionid: 15,
		Tokenstring: mytokenstring,
		Amount: &new_balance,
		
		


	})

	if err==nil{
		fmt.Print(updaterespose)
	}
	fmt.Print("\n")

	// var x []*protogen.Accountmsg


	getallaccountresponse,err:=client.Getalltransaction(ctx,&tp.Getalltransactionsmsg{

	})

	if err==nil{
		fmt.Print(getallaccountresponse)
	}

	





}
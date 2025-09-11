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

	conn, err := grpc.Dial("localhost:50080", grpc.WithInsecure())
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
		Transactionid: 1,
		Type: "cash",
		Amount: 1000,
		


	})

	if err!=nil{
		fmt.Println("\n error in add account response\n")
		fmt.Println(err)
	}
	fmt.Print(mrespose)


	getalltransactionresponse,err:=client.Getalltransaction(ctx,&tp.Getalltransactionsmsg{

	})

	if err==nil{

		fmt.Println("after adding transaction")
		fmt.Println(getalltransactionresponse)
	}




	mytokenstring:=mrespose.Tokenstring

	//optional field chai address pass garnu pardo raixa as pointer type banauxa optional filed le

	new_balance:=int32(30000) //this int 32 is also required


	updaterespose,err:=client.Updatetransaction(ctx,&tp.Updatetransactionmsg{
		Transactionid: 1,
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
		fmt.Println("After Update ")
		fmt.Println(getallaccountresponse)
	}


	deleteres,err:=client.Deletetransaction(ctx,&tp.Deletetransactionmsg{
		Transactionid: 1,
		Tokenstring: mytokenstring,
	})

	if err==nil{
		println(deleteres)
	}

	getallaccountresponse,err=client.Getalltransaction(ctx,&tp.Getalltransactionsmsg{

	})

	if err==nil{
		fmt.Println("After Detete ")
		fmt.Println(getallaccountresponse)
	}

	

	





}
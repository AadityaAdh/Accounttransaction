package main

import(
	"fmt"
	"acctraproject/protogen"
	"google.golang.org/grpc"
	"context"
	"time"
)





func main(){

	fmt.Println("Client has started \n")

	conn, err := grpc.Dial("localhost:50063", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v \n", err)
	}
	defer conn.Close()

	


	client:=protogen.NewAccountserviceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()


	//now start calling from here

	mrespose,err:=client.Addaccount(ctx,&protogen.Accountmsg{
		AccountId: 14,
		Owner: "ram",
		Status: "active",
		Balance: 10000,
		


	})

	if err!=nil{
		fmt.Print("\n error in add account response\n")
		fmt.Print(err)
	}
	fmt.Print(mrespose)

	mytokenstring:=mrespose.Tokenstring

	//optional field chai address pass garnu pardo raixa

	new_balance:=int32(30000) //this int 32 is also required


	updaterespose,err:=client.Updateaccount(ctx,&protogen.Updateaccountmsg{
		AccountId: 14,
		Tokenstring: mytokenstring,
		Balance: &new_balance,
		
		


	})

	if err==nil{
		fmt.Print(updaterespose)
	}
	fmt.Print("\n")

	// var x []*protogen.Accountmsg


	getallaccountresponse,err:=client.Getallaccounts(ctx,&protogen.Getallaccountsmsg{

	})

	if err==nil{
		fmt.Print(getallaccountresponse)
	}

	





}
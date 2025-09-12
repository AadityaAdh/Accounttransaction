package main

import (
	"acctraproject/protogen"
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
)





func main(){

	fmt.Println("Client has started ")

	conn, err := grpc.Dial("localhost:50090", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v \n", err)
	}
	defer conn.Close()

	


	client:=protogen.NewAccountserviceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()


	//now start calling from here

	mrespose,err:=client.Addaccount(ctx,&protogen.Accountmsg{
		AccountId: 1,
		Owner: "ram",
		Status: "active",
		Balance: 10000,
		

	})

	if err!=nil{
		fmt.Print("\n error in add account response\n")
		fmt.Print(err)
	}
	fmt.Print(mrespose)


	getallaccountresponse,err:=client.Getallaccounts(ctx,&protogen.Getallaccountsmsg{

	})

	if err==nil{

		fmt.Println("after adding all the accounts are")
		fmt.Println(getallaccountresponse)
	}

	mytokenstring:=mrespose.Tokenstring

	//optional field chai address pass garnu pardo raixa

	//because the Balance , and all the remaining fields will be pointer

	// Balance       *int32                 `protobuf:"varint,2,opt,name=balance,proto3,oneof" json:"balance,omitempty"`
	// Owner         *string                `protobuf:"bytes,3,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
	// Status        *string 

	//so if you didnt pass anything it will pass in nil as default value of pointer
	//note this nil for pointer and errors 

	//string haru ma nil garnai didaina

	// var s string=nil
	//you cant do this

	// var s *string=nil

	//you can so this

	//simply optional field jati chai pointers banxan so nil hua sakxa tesko value


	new_balance:=int32(30000) //this int 32 is also required


	updaterespose,err:=client.Updateaccount(ctx,&protogen.Updateaccountmsg{
		AccountId: 1,
		Tokenstring: mytokenstring,
		Balance: &new_balance,
		
		


	})

	if err==nil{
		fmt.Print(updaterespose)
	}else{
		fmt.Print(err)
	

		

	}
	fmt.Print("\n")

	// var x []*protogen.Accountmsg


	getallaccountresponse,err=client.Getallaccounts(ctx,&protogen.Getallaccountsmsg{

	})

	if err==nil{
		fmt.Println("After updating all the accounts are")
		fmt.Println(getallaccountresponse)
	}

	//lets try to pass fake token 

	tamperedtoken:=strings.Replace(mytokenstring,"a","e",1)

	updaterespose,err=client.Updateaccount(ctx,&protogen.Updateaccountmsg{
		AccountId: 1,
		Tokenstring: tamperedtoken,
		Balance: &new_balance,
		
		


	})

	if err==nil{
		fmt.Println(updaterespose)
	}else{
		fmt.Println(err)
	

		

	}




	//delete account

	deleteres,err:=client.Deleteaccount(ctx,&protogen.Deleteaccountmsg{
		AccountId: 1,
		Tokenstring: mytokenstring,
	})

	if err==nil{
		println(deleteres)
	}

	getallaccountresponse,err=client.Getallaccounts(ctx,&protogen.Getallaccountsmsg{

	})

	if err==nil{
		fmt.Println("After deleting all the accounts are")
		fmt.Println(getallaccountresponse)
	}



	





}
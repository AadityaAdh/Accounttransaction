package main

import (
	"acctraproject/protogen"
	"acctraproject/rest/handlers"
	"fmt"

	"github.com/labstack/echo"
	"google.golang.org/grpc"
)


func main(){

	fmt.Println("Client has started ")

	conn, err := grpc.Dial("localhost:50093", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect: %v \n", err)
	}
	defer conn.Close()

	
	

	client:=protogen.NewAccountserviceClient(conn)


	transactionhandler:=handlers.Accounthandler{Client: client}

	e:=echo.New()

	e.POST("/addaccount",transactionhandler.Addaccount)

	e.GET("/getallaccount",transactionhandler.Getallaccount)

	e.PUT("/updateaccount",transactionhandler.Updateaccount)

	

	e.DELETE("/deleteaccount/:id",transactionhandler.Deleteaccount)


	e.Start(":1338")





	
	
}
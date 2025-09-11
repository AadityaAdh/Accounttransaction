package main


import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"acctraproject/interceptors"
	tp "acctraproject/transactionprotogen"
	"acctraproject/handler"
)

func main(){
	lis, err := net.Listen("tcp", ":50080")

	if err!=nil{
		fmt.Print(err)
	}

	grpcserver:=grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.TranauthenticationInterceptor))

	tp.RegisterTransactionserviceServer(grpcserver,&handler.Servertran{})


	servererr:=grpcserver.Serve(lis);

	if servererr!=nil{
		fmt.Print(servererr)
	}

}
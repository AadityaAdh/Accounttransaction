package main


import (
	"fmt"
	"net"
	"acctraproject/protogen"
	"google.golang.org/grpc"
	"acctraproject/interceptors"
	"acctraproject/handler"
	
	
)




func main(){
	lis, err := net.Listen("tcp", ":50076")

	if err!=nil{
		fmt.Print(err)
	}

	grpcserver:=grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.AuthenticationInterceptor))

	protogen.RegisterAccountserviceServer(grpcserver,&handler.Server{})

	servererr:=grpcserver.Serve(lis);

	if servererr!=nil{
		fmt.Print(servererr)
	}

}
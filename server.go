package main

import (
	"acctraproject/handler"
	"acctraproject/interceptors"
	"acctraproject/protogen"
	"acctraproject/repository"
	"acctraproject/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)




func main(){
	lis, err := net.Listen("tcp", ":50093")

	fmt.Println("Server started on 50093")

	if err!=nil{
		fmt.Print(err)
	}

	sqlvar:=repository.Sqlitedb{Dbname: "mydb2.db"}

	db,err:=repository.Getdatabase(sqlvar)

	

	if err!=nil{
		fmt.Println("error in getting database conn")
		fmt.Println(err)
	}

	err=repository.Createtableifnotexists(db)


	ar:=repository.Accountrepo{Db: db}

	as:=service.Accountservice{Repo: ar}

	server:=handler.Server{Service: as}
	

	grpcserver:=grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.AuthenticationInterceptor))

	protogen.RegisterAccountserviceServer(grpcserver,&server)

	servererr:=grpcserver.Serve(lis);

	if servererr!=nil{
		fmt.Print(servererr)
	}

}
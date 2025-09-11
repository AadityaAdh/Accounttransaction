package interceptors

import (
	"fmt"

	"acctraproject/protogen"
	"google.golang.org/grpc"
	"context"
	"acctraproject/utils"
	
	
)


func AuthenticationInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error){

	fmt.Print(info.FullMethod)

	if info.FullMethod != "/accounttransaction.Accountservice/Addaccount" && info.FullMethod != "/accounttransaction.Accountservice/Getallaccounts" {
		fmt.Print("\n token step taken \n")
		updatemsg,ok:=req.(*protogen.Updateaccountmsg) //type assertion vando raixa yo chai 

		var tokenstring string

		if ok{
			tokenstring=updatemsg.Tokenstring
		}else{
			deletemsg,ok:=req.(*protogen.Deleteaccountmsg)
			if ok{

				tokenstring=deletemsg.Tokenstring
			}else{

				return &protogen.Statusmsg{Status: 400,Message: "What kind of message are you sending msg should be either update or delete "},nil

			}
			
		}
		tokenverified:=utils.Verifytoken(tokenstring)
		if tokenverified{

			res,err:=handler(ctx,req)
			return res,err
		}
		return &protogen.Statusmsg{Status: 400,Message: "Token could not be verified"},nil
	}
	
	res,err:=handler(ctx,req)
	return res,err

	

}
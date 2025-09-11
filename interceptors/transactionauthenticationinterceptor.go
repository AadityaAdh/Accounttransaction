package interceptors

import (
	"fmt"
	tp "acctraproject/transactionprotogen"
	"google.golang.org/grpc"
	"context"
	"acctraproject/utils"

)




func TranauthenticationInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error){

	fmt.Print(info.FullMethod)

	if info.FullMethod != "/accounttransaction.Transactionservice/Addtransaction" && info.FullMethod != "/accounttransaction.Transactionservice/Getalltransaction" {
		fmt.Print("\n token step taken \n")
		updatemsg,ok:=req.(*tp.Updatetransactionmsg) //type assertion vando raixa yo chai 
		var tokenstring string
		if ok{
			tokenstring=updatemsg.Tokenstring
			
		}else{
			deletemsg,ok:=req.(*tp.Deletetransactionmsg)
			if ok{
				tokenstring=deletemsg.Tokenstring
			}else{
				return &tp.Transactionstatusmsg{Status: 400,Message: "Neede either delete msg or update msg"},nil

			}


		}

		tokenverified:=utils.Verifytoken(tokenstring)
		if tokenverified{

			res,err:=handler(ctx,req)
			return res,err
		}
		return &tp.Transactionstatusmsg{Status: 400,Message: "Token could not be verified"},nil
	}
	
	res,err:=handler(ctx,req)
	return res,err

	

}
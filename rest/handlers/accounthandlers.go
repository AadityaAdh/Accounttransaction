package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"acctraproject/protogen"

	"github.com/labstack/echo"
)


type Accounthandler struct{
	Client protogen.AccountserviceClient
	
}






func (ah Accounthandler) Addaccount(c echo.Context)error{
	var account protogen.Accountmsg
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	

	err:=c.Bind(&account)

	fmt.Println("Accounte after Binding")

	fmt.Println(account)
	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":err.Error()})

	}



	addresponse,err:=ah.Client.Addaccount(ctx,&account)

	if err!=nil{

		return c.JSON(http.StatusBadGateway,err)

	}

	fmt.Println(addresponse)

	err=c.JSON(http.StatusOK,addresponse)

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]any{"message":err})
	}
	return err


}



func (ah Accounthandler) Deleteaccount(c echo.Context)error{
	fmt.Print("Added account")

	err:=c.JSON(http.StatusOK,map[string]string{"message":"Added sucessfully"})

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":"Could not add"})
	}
	return err


}



func (ah Accounthandler) Updateaccount(c echo.Context)error{
	

	var body map[string]interface{}
    if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid JSON"})
    }
	var account_id int32
	var tokenstring string
	var balance *int32
	var owner *int32
	var status *string

	accountidfloat,_:=body["AccountId"].(float64)
	account_id=int32(accountidfloat)

	tokenstring,_=body["Tokenstring"].(string)

	_,err:=body["Balance"].(float64)
	if !err{
		balancefloat:=body["Balance"].(float64)
		balance=&(int32(balancefloat))


	}





	account:=protogen.Updateaccountmsg{AccountId: int32(account_id)}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err:=c.Bind(&account)

	fmt.Println(account)

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":err.Error()})

	}

	updateresponse,err:=ah.Client.Updateaccount(ctx,&account)

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":err.Error()})

	}



	err=c.JSON(http.StatusOK,updateresponse)

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":"Could not update"})
	}
	return err


}



func (ah Accounthandler) Getaccount(c echo.Context)error{
	fmt.Print("Added account")

	err:=c.JSON(http.StatusOK,map[string]string{"message":"Added sucessfully"})

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":"Could not add"})
	}
	return err


}


func (ah Accounthandler) Getallaccount(c echo.Context)error{

	var getallaccountmsg=protogen.Getallaccountsmsg{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	getallaccountresponse,err:=ah.Client.Getallaccounts(ctx,&getallaccountmsg)
	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":"Could not get response"})

	}


	err=c.JSON(http.StatusOK,getallaccountresponse)

	if err!=nil{
		return c.JSON(http.StatusBadGateway,map[string]string{"message":"Could not add"})
	}
	return err


}
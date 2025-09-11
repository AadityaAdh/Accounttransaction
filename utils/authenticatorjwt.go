package utils


import(
	"github.com/golang-jwt/jwt/v5"
	"fmt"
	"os"
)



var secretkey = []byte(os.Getenv("secret_key"))


func Generatetoken(name string)(string,error){
	claims:=jwt.MapClaims{
		"username":name,
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenstring,err:=token.SignedString(secretkey)
	if err!=nil{
		return "",nil

	}
	return tokenstring,nil

}


func Verifytoken(tokenstring string)bool{
	token,err:=jwt.Parse(tokenstring,func (token *jwt.Token) (interface{}, error){
		return secretkey,nil
	})
	if err!=nil{
		fmt.Print("some error occured while parsing", err)
		return false
	}
	if token.Valid{
		return true
	}
	return false


}
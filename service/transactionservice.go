package service

import (
	tp "acctraproject/transactionprotogen"
	"acctraproject/repository"
	"acctraproject/utils"

	"gorm.io/gorm"
)



func Addtransactionservice(transaction repository.Transaction,db *gorm.DB)(string,error){

	err:=repository.Addtransaction(transaction,db)

	if err!=nil{
		return "",err
	}
	tokenstring,err:=utils.Generatetoken(transaction.Transactiontype)//yei yeutya ta string xa
	return tokenstring,err

}


func Updatetransactionservice(transaction repository.Transaction,db *gorm.DB)error{
	err:=repository.Updatetransaction(transaction,db)
	return err
	//done
}


func Getalltransactionservice(db *gorm.DB)(*tp.Getalltransactionsmsg,error){
	transactions,err:=repository.Getalltransaction(db)

	if err!=nil{
		return &tp.Getalltransactionsmsg{},err
	}

	//but we need every element of type *protogen.Getallaccountsmsg to return from that getallaccountservice handler
	//so lets do that conversion

	//from list of accounts to *protogen.Getallaccountsmsg

	//convering accounts to accountsmsg
	transactionsmsg:=make([]*tp.Transactionmsg,len(transactions))

	for _,tra:=range transactions{
		transactionsmsg=append(transactionsmsg, &tp.Transactionmsg{
			AccountId: tra.Accountid,
			Type: tra.Transactiontype,
			Transactionid: tra.Transactionid,
			Amount: tra.Amount,
		})
	}

	return &tp.Getalltransactionsmsg{Listtran: transactionsmsg},nil


	//done
}

func Gettransactionservice(transaction repository.Transaction,db *gorm.DB)(repository.Transaction,error){
	transaction,err:=repository.Gettransaction(transaction.Transactionid,db)
	return transaction,err
	//done
}

func Deletetransactionservice(transactionid int32 ,db *gorm.DB)error{
	err:=repository.Deletetransaction(transactionid,db)
	return err
	//done
}
package repository

import (
	"gorm.io/gorm"
	"errors"
	"fmt"
)


func Addtransaction(transaction Transaction,db *gorm.DB)error{
	result:=db.Create(&transaction)

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return errors.New("your row was not inserted")
	}


	return nil

}

func Updatetransaction(transaction Transaction,db *gorm.DB)error{

	existingtransaction,err:=Gettransaction(transaction.Transactionid,db)

	if err!=nil{
		return errors.New("could not even fetch the single row with this id let along update")

	}

	if transaction.Amount!=0{
		existingtransaction.Amount=transaction.Amount
	}

	if transaction.Transactiontype!=""{
		existingtransaction.Transactiontype=transaction.Transactiontype
	}

	

	//naya xa vanae tyo natra jasta ko testai

	fmt.Print("\n Existing transaction\n")

	fmt.Print(existingtransaction)

	result:=db.Model(&existingtransaction).Where("transactionid=?",existingtransaction.Transactionid).Updates(existingtransaction)


	


	if result.Error!=nil{
		return result.Error
	}
	return nil

}



func Getalltransaction(db *gorm.DB)([]Transaction,error){

	var transaction []Transaction

	result:=db.Find(&transaction)

	if result.Error!=nil{
		return transaction,result.Error

	}
	return transaction,nil

}

func Gettransaction(transactionid int32,db *gorm.DB)(Transaction,error){

	var transaction=Transaction{}
	result:=db.First(&transaction,"transactionid=?",transactionid)

	if result.Error!=nil{
		return  transaction,result.Error
	}
	return transaction,nil

}


func Deletetransaction(transactionid int32,db *gorm.DB)error{

	var transaction Transaction
	result:=db.Where("transactionid=?",transactionid).Delete(&transaction)


	if result.Error!=nil{
		return result.Error
	}
	return nil

}


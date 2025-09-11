package repository

type Transaction struct {
	Transactionid   int32 `gorm:"primarykey"`
	Accountid       int32 
	Amount          int32
	Transactiontype string

	Account *Account `gorm:"foreignKey:Accountid;references:Accountid"`
}




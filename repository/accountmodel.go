package repository



type Account struct{
	
	Accountid int32 `gorm:"primarykey"`
	Balance int32
	Owner string
	Status string
	Transaction []Transaction `gorm:"foreignKey:Accountid"`

}
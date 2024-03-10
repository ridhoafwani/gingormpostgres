package models

type Item struct {
	ItemId      int    `gorm:"type:int;primary_key;autoIncrement:true"`
	ItemCode    string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Quantity    int    `gorm:"type:int"`
	OrderId     int    `gorm:"type:int"`
}

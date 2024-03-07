package models

type Items struct {
	ItemId      int    `gorm:"type:int;primary_key"`
	ItemCode    string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Quantity    int    `gorm:"type:int"`
	OrderId     int    `gorm:"type:int"`
}

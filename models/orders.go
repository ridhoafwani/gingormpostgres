package models

import "time"

type Orders struct {
	OrderId      int    `gorm:"type:int;primary_key;autoIncrement:true"`
	CustomerName string `gorm:"type: varchar(255)"`
	OrderedAt    time.Time
	Items        []Item `gorm:"foreignKey:OrderId"`
}

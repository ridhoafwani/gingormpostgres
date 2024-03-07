package models

type Orders struct {
	OrderId      int    `gorm:"type: int"; primary_key`
	CustomerName string `gorm:"type: varchar(255)"`
	OrderedAt    string `gorm:"type: varchar(255)"`
}

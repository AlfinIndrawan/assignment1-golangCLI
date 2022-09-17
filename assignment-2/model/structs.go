package model

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	// on delete cascade
	Items []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ItemCode    uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint
}

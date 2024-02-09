package pkg

import "time"

type Order struct {
	ID           uint `json: "id" gorm:"primaryKey"`
	CratedAt     time.Time
	ProductRefer int `json:"product_id"`
	//Relacion con producto: foreign key
	Product   Product `gorm:"foreignKey: ProductRefer"`
	UserRefer int     `json:"user_id"`
	User      User    `gorm:"foreignKey:UserRefer"`
}

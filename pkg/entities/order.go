package entities

import "time"

type Order struct {
	ID           uint `json: "id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int `json:"product_id"`
	//Relacion con producto: foreign key
	Product   Product `gorm:"foreignKey: ProductRefer"`
	UserRefer int     `json:"user_id"`
	//Relacion con Usuario
	User User `gorm:"foreignKey:UserRefer"`
}

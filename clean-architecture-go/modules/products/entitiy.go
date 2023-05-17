package products

import "relasi-go/modules/categories"

type Product struct {
	ProductId   int    `gorm:"primaryKey" json:"product_id"`
	Name        string `gorm:"varchar(300)" json:"name_product"`
	Description string `gorm:"varchar(300)" json:"description"`
	Price       int    `gorm:"int(200)" json:"price"`
	CategoryId  int    `gorm:"foreignKey:CategoryId" json:"category_id"`
	Category    categories.Category
}

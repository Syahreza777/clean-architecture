package categories

type Category struct {
	CategoryId   int    `gorm:"primaryKey" json:"category_id"`
	Name string `gorm:"varchar(300)" json:"name_category"`
	Description string `gorm:"varchar(300)" json:"description"`
}
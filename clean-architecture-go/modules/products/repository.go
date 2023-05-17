package products

import "gorm.io/gorm"

type Repository struct{
	DB *gorm.DB
}

func (repo Repository) GetAllProducts() ([]Product, error){
	var products []Product
	result := repo.DB.Preload("Category").Find(&products)

	return products, result.Error
}

func (repo Repository) GetProductById(id string) (*Product, error){
	var product *Product
	result := repo.DB.Preload("Category").First(&product, id)

	return product, result.Error
}

func (repo Repository) AddProduct(product Product) error{
	result := repo.DB.Create(&product)

	return result.Error
}

func (repo Repository) EditProduct(id string, product Product) error{
	result := repo.DB.Where(id).Updates(product)

	return result.Error
}

func (repo Repository) DeleteProduct(id string) error{
	var product Product
	result := repo.DB.Delete(product, id)

	return result.Error
}
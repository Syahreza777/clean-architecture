package products

type Usecase struct {
	Repo Repository
}

func (uc Usecase) UcGetAllProducts() ([]Product, error) {
	products, err := uc.Repo.GetAllProducts()
	return products, err
}

func (uc Usecase) UcGetProductById(id string) (*Product, error) {
	product, err := uc.Repo.GetProductById(id)
	return product, err
}

func (uc Usecase) UcAddProduct(product Product) error {
	err := uc.Repo.AddProduct(product)
	return err
}

func (uc Usecase) UcEditProduct(id string, product Product) error {
	err := uc.Repo.EditProduct(id, product)
	return err
}

func (uc Usecase) UcDeleteProduct(id string) error {
	err := uc.Repo.DeleteProduct(id)
	return err
}

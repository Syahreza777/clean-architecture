package categories

type Usecase struct {
	Repo Repository
}

func (uc Usecase) GetAllCategories() ([]Category, error) {
	categories, err := uc.Repo.GetAllCategories()
	return categories, err
}

func (uc Usecase) GetCategoryById(id string) (*Category, error) {
	category, err := uc.Repo.GetCategoryById(id)
	return category, err
}

func (uc Usecase) AddCategory(category Category) error {
	err := uc.Repo.AddCategory(category)
	return err
}

func (uc Usecase) EditCategory(id string, category Category) error {
	err := uc.Repo.EditCategory(id, category)
	return err
}

func (uc Usecase) DeleteCategory(id string) error {
	err := uc.Repo.DeleteCategory(id)
	return err
}

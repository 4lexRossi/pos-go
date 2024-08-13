package product

type ProductUseCase struct {
	repository ProductRepository
}

func NewProductUseCase(repository ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository}
}

// on real example shoud create a DTO, this is just for example purposes
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}

package category

type Category struct {
	ID   string
	Name string
}

func NewCategory(
	id string,
	name string,
) *Category {
	return &Category{
		ID:   id,
		Name: name,
	}
}

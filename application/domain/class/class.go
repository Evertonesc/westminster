package class

type Class struct {
	ID   string
	Name string
}

func New() Class {
	return Class{}
}

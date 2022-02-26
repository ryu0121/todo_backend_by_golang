package domain

type Category struct {
	ID   int
	Name string
	Base
}

type Categories []Category

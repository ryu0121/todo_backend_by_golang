package domain

type Todo struct {
	ID          int `json:"id" form:"id" query:"id"`
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
	Expiration  string `json:"expiration" form:"expiration" query:"expiration"`
	CategoryId  int `json:"categoryId" form:"categoryId" query:"categoryId"`
	Category    Category
	Base
}

type Todos []Todo

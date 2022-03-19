package domain

type Todo struct {
	ID      int    `json:"id" form:"id" query:"id"`
	Content string `json:"content" form:"content" query:"content"`
	Checked bool   `json:"checked" form:"checked" query:"checked"`
	Removed bool   `json:"removed" form:"removed" query:"removed"`
}

type Todos []Todo

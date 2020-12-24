package domain

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=32"`
}

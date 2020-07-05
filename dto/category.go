package dto

type Category struct {
	Id         int      `validate:"number,min=0,max=9999999999"`
	Name       string   `validate:"required,alphanum,min=1,max=30"`
	IsSelected bool
}
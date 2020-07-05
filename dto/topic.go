package dto

type Topic struct {
	Id           int    `validate:"number,min=0,max=9999999999"`
	Title        string `validate:"required,alphanum,min=1,max=100"`
	Label        string `validate:"required,alphanum,min=1,max=100"`
	CategoryId   int    `validate:"number,min=0,max=99"`
	CategoryName string
	Content      string `validate:"required,min=8,max=100"`
	Time         *Time
}

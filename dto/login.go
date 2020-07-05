package dto

type Login struct {
	Email    string `validate:"required,email,min=6,max=100"`
	Password string `validate:"required,alphanum,min=8,max=50"`
	Check    string `validate:"min=0,max=4"`
}

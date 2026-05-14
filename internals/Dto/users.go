package dto 

type UserCreate struct {
	Name string `json:"name" validate:"required,max=30"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}


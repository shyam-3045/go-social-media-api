package users

import (
	"context"
	dto "socio/internals/Dto"
	"socio/models/users"
	models "socio/models/users"

	"github.com/google/uuid"
)


func Create(ctx context.Context , body dto.UserCreate) (*models.Users , error){
	
	user := users.New()

	user.Name = body.Name
	user.Email = body.Email
	user.Password = body.Password

	if err := user.Create(ctx) ; err != nil{
		return nil ,err
	}

	return user,nil
}

func GetById(ctx context.Context , id uuid.UUID) (*models.Users , error){
	return users.GetByID(ctx,id)
}
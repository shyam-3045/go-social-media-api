package users

import (
	"context"
	"socio/internals/database"
	"time"
	"github.com/google/uuid"
)

type Users struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` 
}

func New() *Users {
	return &Users{}
}

func (u *Users) Create(ctx context.Context) error {
	return database.Client().
		WithContext(ctx).
		Create(u).Error
}

func GetByID(ctx context.Context , id uuid.UUID) (*Users , error){
	var user Users

	err := database.Client().WithContext(ctx).First(&user,id).Error

	if err != nil {
		return nil ,err
	}

	return &user,nil
}
package friendships

import (
	"context"
	"socio/internals/database"
	"socio/models/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Friendships struct {
	gorm.Model

	UserID   uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"user_id"`
	FriendID uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"friend_id"`

	User   users.Users `gorm:"foreignKey:UserID;references:ID" json:"-" `
	Friend users.Users `gorm:"foreignKey:FriendID;references:ID" json:"-"`
}

func New() *Friendships {
	return &Friendships{}
}

func (f *Friendships) Add(ctx context.Context) error {
	return database.Client().WithContext(ctx).Create(f).Error
} 

package dto

import "github.com/google/uuid"

type Friends struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid"`
	FriendID uuid.UUID `json:"friend_id" validate:"required,uuid"`

}
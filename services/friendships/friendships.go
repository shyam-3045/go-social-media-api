package friendships

import (
	"context"
	dto "socio/internals/Dto"
	"socio/models/friendships"
)

func Add(ctx context.Context , body dto.Friends) (*friendships.Friendships , error){
	frnd := friendships.New()

	frnd.UserID = body.UserID
	frnd.FriendID = body.FriendID

	if err := frnd.Add(ctx) ; err != nil {
		return nil , err
	}

	return frnd,nil
}
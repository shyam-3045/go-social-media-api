package config

import (
	"socio/internals/database"
	"socio/models/friendships"
	"socio/models/posts"
	"socio/models/users"
)

func AutoMigrate() {
	 database.DB.AutoMigrate(&users.Users{} , &friendships.Friendships{} ,&posts.Posts{})
}
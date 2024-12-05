package structs

import "github.com/google/uuid"

type UserInput struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreatePost struct {
	Post string `json:"post" validate:"required"`
}

type CreateComment struct {
	Comment string    `json:"comment" validate:"required"`
	PostId  uuid.UUID `json:"post_id" validate:"required"`
}

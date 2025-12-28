package params

import "github.com/google/uuid"

type UserCreate struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`

	RoleIds []uuid.UUID `json:"role_ids" validate:"min=1,dive,uuid"`
}

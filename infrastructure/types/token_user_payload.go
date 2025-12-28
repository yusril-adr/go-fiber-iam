package types

import "github.com/google/uuid"

type TUserPayload struct {
	Id      uuid.UUID   `json:"id"`
	RoleIds []uuid.UUID `json:"role_ids"`
}

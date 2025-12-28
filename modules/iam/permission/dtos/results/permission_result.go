package results

import (
	"github.com/google/uuid"

	iamModel "iam-service/models/maindb/iam"
)

type PermissionResult struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Key  string    `json:"key"`
}

func (p *PermissionResult) MapModel(permission iamModel.Permission) {
	p.Id = permission.Id
	p.Name = permission.Name
	p.Key = permission.Key
}

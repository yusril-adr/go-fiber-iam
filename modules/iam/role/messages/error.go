package messages

import (
	"fmt"

	"github.com/google/uuid"
)

func ErrRolesNotFound(roles []uuid.UUID) string {
	return fmt.Sprintf("Roles not found: %v", roles)
}

var (
	ErrRoleNotFound        = "Role not found"
	ErrRoleAssignedToUsers = "Cannot delete role that is assigned to users"
)

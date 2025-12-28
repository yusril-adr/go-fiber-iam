package messages

import (
	"fmt"

	"github.com/google/uuid"
)

func ErrPermissionsNotFound(permissions []uuid.UUID) string {
	return fmt.Sprintf("Permissions not found: %v", permissions)
}

var (
	ErrPermissionNotFound        = "Permission not found"
	ErrPermissionAssignedToRoles = "Cannot delete permission that is assigned to roles"
)

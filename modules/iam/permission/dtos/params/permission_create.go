package params

type PermissionCreate struct {
	Name string `json:"name" validate:"required"`
}

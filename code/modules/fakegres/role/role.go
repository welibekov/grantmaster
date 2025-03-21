package role

import "github.com/welibekov/grantmaster/modules/role/base"

type FGRole struct {
	*base.Role
}

func New(cfg map[string]string) (*FGRole, error) {
	baseRole := base.NewRole(cfg)

	return &FGRole{
		Role: baseRole,
	}, nil
}

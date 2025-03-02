package template

import (
	"context"
	"text/template"

	"github.com/welibekov/grantmaster/modules/template/funcs"
)

func (t *Template) NewFuncs(ctx context.Context) template.FuncMap {
	fn := funcs.New(ctx, t.pool)

	return template.FuncMap{
		"isRoleExist": fn.IsRoleExist,
	}
}

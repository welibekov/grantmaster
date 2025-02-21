package greenplum

import (
	"github.com/welibekov/grantmaster/modules/database/base"
)

type Greenplum struct {
	*base.Database
}

func New(config map[string]string) (*Greenplum, error) {
	return &Greenplum{
		base.NewDatabase(),
	}, nil

}

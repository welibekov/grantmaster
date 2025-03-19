package databaser

type RunTesterer interface {
	Prepare() (func() error, error)
	Execute() error
}

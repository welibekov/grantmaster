package databaser

type Testerer interface {
	Prepare() (func() error, error)
	Execute() error
}

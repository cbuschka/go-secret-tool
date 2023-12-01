package backend

type Secret interface {
	Name() string
}

type Backend interface {
	ListSecrets() ([]Secret, error)
	Close() error
}

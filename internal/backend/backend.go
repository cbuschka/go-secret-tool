package backend

type Secret interface {
	Name() string
	Value() string
}

type Backend interface {
	ListSecrets() ([]Secret, error)
	Close() error
}

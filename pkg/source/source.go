package source

// Source is interface for defining source
type Source interface {
	With() error
	To() error
}

package source

// Source is interface for defining source
type Source interface {
	Out() <-chan interface{}
}

package receiver

// Receiver defines main interface for receiving
type Receiver interface {
	In(interface{})
}

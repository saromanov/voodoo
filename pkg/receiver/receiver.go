package receiver

type Receiver interface {
	In(<-chan interface{})
}

// Package transform contains actions for data transform
package transform

// Transform is interface for data manipulation/transformation
type Transform interface {
	In(elem interface{})
	Out() <-chan interface{}
}

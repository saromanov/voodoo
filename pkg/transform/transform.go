// Package transform contains actions for data transform
package transform

// Transform is interface for data manipulation/transformation
type Transform interface {
	Do() error
}

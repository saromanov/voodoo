package voodoo

// Voodoo defines main app structure
type Voodoo struct {

}

// New creates new instance of the project
func (v *Voodoo) New() *Voodoo {
	return &Voodoo {}
}
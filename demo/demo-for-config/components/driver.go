package components

// Driver class
type Driver struct {
	ID       string
	Name     string
	Birthday string
	Sex      string
	Car      *Car
	MyCars   []*Car
}

package elements

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// struct ComExample1

type ComExample1 struct {
	Name string
}

////////////////////////////////////////////////////////////////////////////////
// struct ComExample2

type ComExample2 struct {
	Com1ref *ComExample1
}

func (inst *ComExample2) Open() error {
	fmt.Println("Com1.Open")
	return nil
}

func (inst *ComExample2) Close() error {
	fmt.Println("Com1.Close")
	return nil
}

func (inst *ComExample2) Loop() error {
	fmt.Println("Com1.Loop")
	return nil
}

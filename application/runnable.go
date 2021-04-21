package application

type Runnable interface {
	Run(context RuntimeContext) error
}

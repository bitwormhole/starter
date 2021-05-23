package application

type Runnable interface {
	Run(context Context) error
}

package collection

type Arguments interface {
	GetReader(flag string) (ArgumentReader, bool)
	Import([]string)
	Export() []string
}

type ArgumentReader interface {
	SetEnding(ending string)
	Ending() string
	Read() (string, bool)
}

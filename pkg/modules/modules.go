package modules

type Module [S any, C any] interface {
	Name() string
	Settings() S
	Config() C
	Exec(filepath, filename string) error
}

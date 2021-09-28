package bolter

import "github.com/go-logr/logr"

type option func(*Bolter)

func SetFilePath(path string) option {
	return func(b *Bolter) {
		b.Path = path
	}
}

func SetLogger(logger logr.Logger) option {
	return func(b *Bolter) {
		b.logger = logger
	}
}

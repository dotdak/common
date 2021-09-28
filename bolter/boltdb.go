package bolter

import (
	"io/fs"

	"github.com/boltdb/bolt"
	"github.com/dotdak/common/logger"
	"github.com/go-logr/logr"
)

type Bolter struct {
	*bolt.DB
	Path     string
	FileMode uint32
	logger   logr.Logger
}

func NewBolter(opts ...option) *Bolter {
	bolter := &Bolter{
		Path:     "default.db",
		logger:   logger.LOG(),
		FileMode: 0600,
	}

	for _, o := range opts {
		o(bolter)
	}

	return bolter
}

func (b *Bolter) Open() error {
	db, err := bolt.Open(b.Path, fs.FileMode(b.FileMode), nil)
	if err != nil {
		b.DB = nil
		return err
	}

	b.DB = db

	return nil
}
func (b *Bolter) Close() error {
	return b.DB.Close()
}

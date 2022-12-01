package db

type Store interface {
	Add()
	Get()
	GetDir() string
}

type FileStore struct {
	dir string
}

func NewFileStore(dir string) *FileStore {
	return &FileStore{
		dir: dir,
	}
}

func (fs *FileStore) GetDir() string {
	return fs.dir
}

func (fs *FileStore) Add() {
}

func (fs *FileStore) Get() {
}

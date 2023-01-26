package models

type StorageAccess interface {
	Browse() ([]Storage, error)
	Read(name string) (Storage, error)
}

type Storage struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

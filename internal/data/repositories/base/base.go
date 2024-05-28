package base

type BaseRepository[T any, Id any] interface {
	Create(*T) (*T, error)
	Update(*T) (*T, error)
	FindById(Id) (*T, error)
	FindAll() ([]*T, error)
	DeleteById(Id) error
}

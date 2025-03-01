package entity

type IPersist[T any] interface {
	Get(ID string) T
	Add(ID string, entity T) error
}

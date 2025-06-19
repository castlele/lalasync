package storage

type DB[K any, V any] interface {
	GetAll() []*V
	Get(K) *V
	Set(K, *V) error
}

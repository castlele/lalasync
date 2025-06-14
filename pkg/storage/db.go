package storage

type DB[K any, V any] interface {
	Get(K) *V
	Set(K, *V) error
}

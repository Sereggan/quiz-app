package store

type Store interface {
	New(address string) *Store
}

package structures

// Index is used for storing data in an efficient manner
// for reading and writing data quickly instead of 
// manually going through each block
type Index[K, V any] interface {
	// Insert creates an entry into this index that maps
	// key K to value V
	Insert(key K, value V) error

	// Get finds an entry by key
	Get(key K) (K, V, error)

	// LargerThanOrEqualTo finds all entries that contain a key
	// larger than or equal to specified key
	LargerThanOrEqualTo(key K) Iterator[K, V]

	// LargerThan finds all entries that contain a key larger
	// than specified key
	LargerThan(key K) Iterator[K, V]

	// LessThanOrEqualTo finds all entries that contain a key 
	// less than or equal specified key
	LessThanOrEqualTo(key K) Iterator[K, V]

	// LessThan finds all entries that contain a key 
	// less than specified key 
	LessThan(key K) Iterator[K, V]

	// Delete deletes an entry from this index using a comparer function
	Delete(key K, value V, valueComparer func(V, V) bool) bool

	// DeleteAll deletes all entries of given key
	DeleteAll(key K) bool
}

type Iterator[K, V any] interface {
	Next() bool
	Key() K
	Value() V
	Close() error
}
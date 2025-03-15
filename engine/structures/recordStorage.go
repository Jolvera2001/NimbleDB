package structures

// RecordStorage provides methods for storing and retrieving variable
// sized records. It abstracts the underlying block storage, allowing the
// database to work with logical records rather than physical blocks
type RecordStorage interface {
	// Update updates a record
	Update(recordId uint, data []byte) error

	// Find grabs a record's data
	Find(recordId uint) ([]byte, error)

	// Create makes a new record
	CreateEmpty() (uint32, error)

	// Creates a new record with given data
	CreateWithData(data []byte) (uint32, error)

	// Creates a new record similar to CreateWithData 
	// but with a generator which makes data after a 
	// record is allocated
	CreateWithGenerator(dataGenerator func(uint32) []byte) (uint32, error)

	// Delete deletes a record by its Id
	Delete(recordId uint32) error
}
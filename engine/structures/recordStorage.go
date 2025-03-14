package structures

type IRecordStorage interface {
	Update(recordId uint, data []byte) error
	Find(recordId uint) ([]byte, error)
	CreateEmpty() (uint32, error)
	CreateWithData(data []byte) (uint32, error)
	CreateWithGenerator(dataGenerator func(uint32) []byte) (uint32, error)
	Delete(recordId uint32) error
}
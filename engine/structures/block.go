package structures

// Block represents a storage container for data
type Block interface {
	// ID returns the unique identifier of the block
	ID() uint32

	// GetHeader retrieves the value of a header field
	GetHeader(field int) int64

	// SetHeader changes the value of specified header
	// data must not be written on disk until the block is disposed
	SetHeader(field int, value int64)

	// Read reads the content of this block (src) into given buffer (dst)
	Read(dst []byte, dstOffset int, srcOffset int, count int) error

	// Write writes the content of given buffer (src) into this (dst)
	Write(src []byte, srcOffset int, dstOffset int, count int) error
}

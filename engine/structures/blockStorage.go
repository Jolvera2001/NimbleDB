package structures

// IBlockStorage represents a container that manages blocks
type IBlockStorage interface {
	// BlockContentSize returns the number of bytes
	// of custom data per block that this storage can handle
	BlockContentSize() int

	// BlockHeaderSize returns the total number of bytes in a header
	BlockHeaderSize() int

	// BlockSize returns the total block size
	// equal to content size + header size
	// should be a multiple of 128B
	BlockSize() int

	// Find finds a block by its id
	Find(blockId uint32) (IBlock, error)

	// CreateNew allocates a new block
	// extends the length of underlying storage
	CreateNew() (IBlock, error)
}
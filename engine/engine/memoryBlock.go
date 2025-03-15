package engine

import "os"

type MemoryBlock struct {
	firstSector       []byte
	cachedHeaderValue []int64
	file              *os.File
	storage           *MemoryBlockStorage
	id                uint32

	isDirty    bool
	isDisposed bool
}

func NewMemoryBlock() (*MemoryBlock, error) {
	
}

func (b *MemoryBlock) ID() uint32 {
	return b.id
}
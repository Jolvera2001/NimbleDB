package engine

type MemoryBlock struct {
	firstSector []byte
	cachedHeaderValue []int64
}
package graphics

const (
	maxBlocks = 32
)

type Cache struct {
	BlockCount     int
	BlockPositions []float32
	BlockSizes     []float32
}

func NewCache() *Cache {
	return &Cache{
		BlockCount:     0,
		BlockPositions: make([]float32, maxBlocks*3),
		BlockSizes:     make([]float32, maxBlocks*2),
	}
}

func (c *Cache) Reset() {
	c.BlockCount = 0
	for i := range c.BlockPositions {
		c.BlockPositions[i] = 0
	}
	for i := range c.BlockSizes {
		c.BlockSizes[i] = 0
	}
}

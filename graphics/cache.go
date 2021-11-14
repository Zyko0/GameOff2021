package graphics

const (
	maxBlocks = 32
)

type Cache struct {
	BlockCount     int
	BlockPositions []float32
	BlockSizes     []float32
	BlockKinds     []float32
	BlockSeeds     []float32
}

func NewCache() *Cache {
	return &Cache{
		BlockCount:     0,
		BlockPositions: make([]float32, maxBlocks*3),
		BlockSizes:     make([]float32, maxBlocks*2),
		BlockKinds:     make([]float32, maxBlocks),
		BlockSeeds:     nil,
	}
}

func (c *Cache) Reset() {
	c.BlockCount = 0
	c.BlockSeeds = nil
	for i := range c.BlockPositions {
		c.BlockPositions[i] = 0
	}
	for i := range c.BlockSizes {
		c.BlockSizes[i] = 0
	}
	for i := range c.BlockKinds {
		c.BlockKinds[i] = 0
	}
}

package core

const (
	BlockWidth0 = 0.2
	BlockWidth1 = 0.4

	BlockHeight0 = 0.2
	BlockHeight1 = 0.4

	BlockPosY0 = 0.
	BlockPosY1 = 0.4
)

type BlockKind float32

const (
	BlockKindRegular      BlockKind = 4
	BlockKindHarder       BlockKind = 5
	BlockKindHarder2      BlockKind = 6
	BlockKindHeart        BlockKind = 7
	BlockKindGoldenHeart  BlockKind = 8
)

type Block struct {
	width, height, x, y, z float64
	kind                   BlockKind
}

func newBlock(x, y, z, width, height float64, kind BlockKind) *Block {
	return &Block{
		width:  width,
		height: height,
		x:      x + width/2,
		y:      y + height/2,
		z:      z,
		kind:   kind,
	}
}

func (b *Block) GetX() float64 {
	return b.x
}

func (b *Block) GetY() float64 {
	return b.y
}

func (b *Block) GetZ() float64 {
	return b.z
}

func (b *Block) GetWidth() float64 {
	return b.width
}

func (b *Block) GetHeight() float64 {
	return b.height
}

func (b *Block) GetKind() BlockKind {
	return b.kind
}

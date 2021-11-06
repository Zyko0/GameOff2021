package graphics

import (
	"github.com/Zyko0/GameOff2021/drafts/01/automata"
	"github.com/Zyko0/GameOff2021/drafts/01/graphics/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	boxIndices = [6]uint16{0, 1, 2, 1, 2, 3}
)

type Graphics struct {
	palette      []float32
	outlineColor []float32
}

func New() *Graphics {
	outlineColor := []float32{
		1.0, 1.0, 1.0, 1.0, // TODO: make a funnier color
	}
	return &Graphics{
		palette:      NewRandomPalette(),
		outlineColor: outlineColor,
	}
}

func (g *Graphics) SetPalette(palette []float32) {
	g.palette = palette
}

func (g *Graphics) DrawParticles(screen *ebiten.Image, cellSize float32, particles []*automata.Particle) {
	vertices := make([]ebiten.Vertex, 0, len(particles)*4)
	indices := make([]uint16, 0, len(particles)*6)
	triangleIndex := 0
	for _, p := range particles {
		pos := p.GetPosition()
		v := float32(p.GetValue())
		vertices = append(vertices, []ebiten.Vertex{
			{
				DstX:   float32(pos[0]) * cellSize,
				DstY:   float32(pos[1]) * cellSize,
				SrcX:   0,
				SrcY:   0,
				ColorR: v,
				ColorG: 0,
				ColorB: 0,
			},
			{
				DstX:   float32(pos[0]+1) * cellSize,
				DstY:   float32(pos[1]) * cellSize,
				SrcX:   0,
				SrcY:   0,
				ColorR: v,
				ColorG: 0,
				ColorB: 0,
			},
			{
				DstX:   float32(pos[0]) * cellSize,
				DstY:   float32(pos[1]+1) * cellSize,
				SrcX:   0,
				SrcY:   0,
				ColorR: v,
				ColorG: 0,
				ColorB: 0,
			},
			{
				DstX:   float32(pos[0]+1) * cellSize,
				DstY:   float32(pos[1]+1) * cellSize,
				SrcX:   0,
				SrcY:   0,
				ColorR: v,
				ColorG: 0,
				ColorB: 0,
			},
		}...)

		indiceCursor := uint16(triangleIndex * 4)
		indices = append(indices, []uint16{
			boxIndices[0] + indiceCursor,
			boxIndices[1] + indiceCursor,
			boxIndices[2] + indiceCursor,
			boxIndices[3] + indiceCursor,
			boxIndices[4] + indiceCursor,
			boxIndices[5] + indiceCursor,
		}...)
		triangleIndex++
	}
	screen.DrawTrianglesShader(
		vertices, indices, shaders.ParticleShader, &ebiten.DrawTrianglesShaderOptions{
			Uniforms: map[string]interface{}{
				"Palette": g.palette,
			},
			CompositeMode: ebiten.CompositeModeSourceOver, // TODO: Hmm.. does Multiply results in a intuitive effect ?
		},
	)
}

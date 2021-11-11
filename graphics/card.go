package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	brushImage       = ebiten.NewImage(1, 1)
	boxIndices       = []uint16{0, 1, 2, 1, 2, 3}
	borderBoxIndices = []uint16{0, 2, 4, 2, 4, 6, 1, 3, 5, 3, 5, 7}
)

func init() {
	brushImage.Fill(color.White)
}

func DrawRect(dst *ebiten.Image, x, y, width, height float32, clr []float32) {
	dst.DrawTriangles(
		[]ebiten.Vertex{
			{
				DstX:   x,
				DstY:   y,
				SrcX:   0,
				SrcY:   0,
				ColorR: clr[0],
				ColorG: clr[1],
				ColorB: clr[2],
				ColorA: clr[3],
			},
			{
				DstX:   x + width,
				DstY:   y,
				SrcX:   1,
				SrcY:   0,
				ColorR: clr[0],
				ColorG: clr[1],
				ColorB: clr[2],
				ColorA: clr[3],
			},
			{
				DstX:   x,
				DstY:   y + height,
				SrcX:   0,
				SrcY:   1,
				ColorR: clr[0],
				ColorG: clr[1],
				ColorB: clr[2],
				ColorA: clr[3],
			},
			{
				DstX:   x + width,
				DstY:   y + height,
				SrcX:   1,
				SrcY:   1,
				ColorR: clr[0],
				ColorG: clr[1],
				ColorB: clr[2],
				ColorA: clr[3],
			},
		}, boxIndices, brushImage, &ebiten.DrawTrianglesOptions{},
	)
}

func DrawRectBorder(dst *ebiten.Image, x, y, width, height, borderWidth float32, clr []float32) {
	dst.DrawTriangles([]ebiten.Vertex{
		{
			DstX:   x,
			DstY:   y,
			SrcX:   0,
			SrcY:   0,
			ColorR: clr[0],
			ColorG: clr[1],
			ColorB: clr[2],
			ColorA: clr[3],
		},
		{
			DstX: x + borderWidth,
			DstY: y + borderWidth,
			SrcX: 0,
			SrcY: 0,
		},
		{
			DstX:   x + width,
			DstY:   y,
			SrcX:   1,
			SrcY:   0,
			ColorR: clr[0],
			ColorG: clr[1],
			ColorB: clr[2],
			ColorA: clr[3],
		},
		{
			DstX: x + width - borderWidth,
			DstY: y + borderWidth,
			SrcX: 1,
			SrcY: 0,
		},
		{
			DstX:   x,
			DstY:   y + height,
			SrcX:   0,
			SrcY:   1,
			ColorR: clr[0],
			ColorG: clr[1],
			ColorB: clr[2],
			ColorA: clr[3],
		},
		{
			DstX: x + borderWidth,
			DstY: y + height - borderWidth,
			SrcX: 0,
			SrcY: 1,
		},
		{
			DstX:   x + width,
			DstY:   y + height,
			SrcX:   1,
			SrcY:   1,
			ColorR: clr[0],
			ColorG: clr[1],
			ColorB: clr[2],
			ColorA: clr[3],
		},
		{
			DstX: x + width - borderWidth,
			DstY: y + height - borderWidth,
			SrcX: 1,
			SrcY: 1,
		},
	}, borderBoxIndices, brushImage, &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	})
}

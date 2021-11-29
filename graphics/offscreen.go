package graphics

import (
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	QualityHigh = 1 + iota
	QualityMedium
	QualityLow
	QualityVeryLow
)

var (
	quality = QualityHigh

	offscreenVeryLow = ebiten.NewImage(
		logic.GameSquareDim/QualityVeryLow,
		logic.GameSquareDim/QualityVeryLow,
	)
	offscreenLow = ebiten.NewImage(
		logic.GameSquareDim/QualityLow,
		logic.GameSquareDim/QualityLow,
	)
	offscreenMedium = ebiten.NewImage(
		logic.GameSquareDim/QualityMedium,
		logic.GameSquareDim/QualityMedium,
	)
	offscreenHigh = ebiten.NewImage(
		logic.GameSquareDim/QualityHigh,
		logic.GameSquareDim/QualityHigh,
	)

	optsVeryLow = &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
	optsLow = &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
	optsMedium = &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
	optsHigh = &ebiten.DrawImageOptions{
		Filter: ebiten.FilterNearest,
	}
)

func init() {
	optsVeryLow.GeoM.Scale(float64(QualityVeryLow), float64(QualityVeryLow))
	optsVeryLow.GeoM.Translate(float64(logic.ScreenWidth-logic.GameSquareDim)/2, 0)

	optsLow.GeoM.Scale(float64(QualityLow), float64(QualityLow))
	optsLow.GeoM.Translate(float64(logic.ScreenWidth-logic.GameSquareDim)/2, 0)

	optsMedium.GeoM.Scale(float64(QualityMedium), float64(QualityMedium))
	optsMedium.GeoM.Translate(float64(logic.ScreenWidth-logic.GameSquareDim)/2, 0)

	optsHigh.GeoM.Scale(float64(QualityHigh), float64(QualityHigh))
	optsHigh.GeoM.Translate(float64(logic.ScreenWidth-logic.GameSquareDim)/2, 0)
}

func GetOffscreenImage() *ebiten.Image {
	switch quality {
	case QualityVeryLow:
		return offscreenVeryLow
	case QualityLow:
		return offscreenLow
	case QualityMedium:
		return offscreenMedium
	case QualityHigh:
		return offscreenHigh
	default:
		return offscreenMedium
	}
}

func GetOffscreenOpts() *ebiten.DrawImageOptions {
	switch quality {
	case QualityVeryLow:
		return optsVeryLow
	case QualityLow:
		return optsLow
	case QualityMedium:
		return optsMedium
	case QualityHigh:
		return optsHigh
	default:
		return optsMedium
	}
}

func GetQuality() int {
	switch quality {
	case QualityVeryLow:
		return QualityVeryLow
	case QualityLow:
		return QualityLow
	case QualityMedium:
		return QualityMedium
	case QualityHigh:
		return QualityHigh
	default:
		return QualityMedium
	}
}

func SetQuality(q int) {
	quality = q
}

func UpdateQualitySettings() bool {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.Key1):
		quality = QualityVeryLow
	case inpututil.IsKeyJustPressed(ebiten.Key2):
		quality = QualityLow
	case inpututil.IsKeyJustPressed(ebiten.Key3):
		quality = QualityMedium
	case inpututil.IsKeyJustPressed(ebiten.Key4):
		quality = QualityHigh
	default:
		return false
	}

	return true
}

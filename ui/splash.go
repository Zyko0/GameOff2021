package ui

import (
	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	splashDisplayDuration = logic.TPS * 3
)

type SplashView struct {
	ticks  uint
	active bool

	geom   ebiten.GeoM
	colorm ebiten.ColorM
	image  *ebiten.Image
}

func NewSplashView() *SplashView {
	splashImg := assets.SplashScreen()

	geom := ebiten.GeoM{}
	geom.Translate(0, 0)
	geom.Scale(
		float64(logic.ScreenWidth)/float64(splashImg.Bounds().Max.X),
		float64(logic.ScreenHeight)/float64(splashImg.Bounds().Max.Y),
	)
	colorm := ebiten.ColorM{}

	return &SplashView{
		ticks:  0,
		active: true,

		geom:   geom,
		colorm: colorm,
		image:  splashImg,
	}
}

func (sv *SplashView) Active() bool {
	return sv.active
}

func (sv *SplashView) Update() {
	sv.ticks++

	if sv.ticks > splashDisplayDuration {
		sv.active = false
		return
	}
	if len(inpututil.AppendPressedKeys(nil)) > 0 || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		sv.active = false
	}

	d := float64(sv.ticks) / float64(splashDisplayDuration)
	sc := (-(d * d) + d) * 4
	sv.colorm.Reset()
	sv.colorm.Scale(sc, sc, sc, 1.)
}

func (sv *SplashView) Draw(screen *ebiten.Image) {
	screen.DrawImage(sv.image, &ebiten.DrawImageOptions{
		GeoM:   sv.geom,
		Filter: ebiten.FilterLinear,
		ColorM: sv.colorm,
	})
}

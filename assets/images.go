package assets

import (
	"bytes"
	"image/png"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed images/splash_black.png
	splashScreenBytes []byte
	splashScreenImage *ebiten.Image
)

func init() {
	img, err := png.Decode(bytes.NewReader(splashScreenBytes))
	if err != nil {
		log.Fatal(err)
	}
	splashScreenImage = ebiten.NewImageFromImage(img)
}

func SplashScreen() *ebiten.Image {
	return splashScreenImage
}

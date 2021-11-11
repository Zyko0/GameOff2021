package assets

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
)

var (
	CardTitleFontFace font.Face
	CardBodyText   font.Face
)

func init() {
	pfont, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		log.Fatal(err)
	}
	CardTitleFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 24,
	})

	pfont, err = truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}
	CardBodyText = truetype.NewFace(pfont, &truetype.Options{
		Size: 16,
	})
}

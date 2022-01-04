package shaders

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed internal/raymarching.kage
	raymarchingShaderSrc []byte

	RaymarchingShader *ebiten.Shader
)

func init() {
	var err error

	RaymarchingShader, err = ebiten.NewShader(raymarchingShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}

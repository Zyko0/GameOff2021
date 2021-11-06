package shaders

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	particleShaderSrc = []byte(
		`
package main

var Palette [4]vec4
var Value float

func palette(t float, a, b, c, d vec4) vec3 {
	var clr vec3
	if t >= a.a && t <= b.a {
		clr = mix(a.rgb*a.rgb, b.rgb*b.rgb, vec3((t-a.a)/(b.a-a.a)))
	} else if t >= b.a && t <= c.a {
		clr = mix(b.rgb*b.rgb, c.rgb*c.rgb, vec3((t-b.a)/(c.a-b.a)))
	} else {
		clr = mix(c.rgb*c.rgb, d.rgb*d.rgb, vec3((t-c.a)/(d.a-c.a)))
	}
	clr = clr*clr*(3.0-2.0*clr)
	return sqrt(clr)
}

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	v := color.r
	c := vec4(palette(v, Palette[0], Palette[1], Palette[2], Palette[3]), 1)
	
	return c
}
`)

	ParticleShader *ebiten.Shader
)

func init() {
	var err error

	ParticleShader, err = ebiten.NewShader(particleShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}

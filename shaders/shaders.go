package shaders

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	raymarchShaderSrc = []byte(
		`
package main

const (
	PlayerIndex = 0.
	RoadIndex = 1.
	PlaneIndex = 2.
	BlockIndex = 4.
	BlockHarderIndex = 5.
	BlockHarder2Index = 6.
	BlockHeartIndex = 7.
	
	MaxBlocks = 32.
	MaxDepth = 30.
)

var ScreenSize vec2
var PlayerPosition vec3
var PlayerRadius float
var Camera vec3
var Distance float

var BlockCount float
var BlockPositions [32]vec3
var BlockSizes [32]vec2
var BlockKinds [32]float
var BlockSeeds [7]float

// TODO: If this is too much uniform for a specific platform, can always hardcode these
var PalettePlayer [4]vec4
var PaletteBlock [4]vec4
var PaletteRoad [4]vec4
var PaletteBlockHarder [4]vec4
var PaletteBlockHarder2 [4]vec4
var PaletteHeart [4]vec4

func hash(p vec2, seed float) float { 
	return fract(sin(dot(p, vec2(12.9898, 4.1414))) * 43758.5453 * seed)
}

// https://gist.github.com/patriciogonzalezvivo/670c22f3966e662d2f83#generic-123-noise
func noise(p vec2, seed float) float {
	ip := floor(p)
	u := fract(p)
	u = u*u*(3.0-2.0*u)
	
	res := mix(
		mix(hash(ip, seed), hash(ip+vec2(1.,0.), seed), u.x),
		mix(hash(ip+vec2(0.,1.), seed), hash(ip+vec2(1.,1.), seed), u.x), u.y)
	
	return res*res
}

func background(p vec3) vec3 {
	space := vec3(0.01, 0.01, 0.01)

	oriY := p.y
	// p.x = abs(p.x) // TODO: do not want symmetry
	p.x -= sign(p.x) * p.z * 0.005
	p.x = abs(p.x)*0.1
	p.y += p.z * 0.002
	n := fract(noise(p.xy*500., BlockSeeds[0]))
	if oriY < 0.03 && n < 0.002 {
		return vec3(0.5)
	}
	return space
}

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

func colorize(p vec3, t, index, seed float) vec3 {
	var pal [4]vec4

	if index == PlayerIndex {
		const scale = 16.

		// X Rotation
		s := sin(PlayerPosition.x*8.)
		c := cos(PlayerPosition.x*8.)
		t = noise(p.xy*mat2(c, -s, s, c)*scale, seed)
		pal = PalettePlayer
	} else if index == BlockIndex {
		t = noise(p.xy*8., seed)
		pal = PaletteBlock
	} else if index == RoadIndex {
		p.z -= Distance
		t = noise(p.xz*2., seed)
		t *= (1+noise(p.xz*4., seed))
		t *= (1+noise(p.xz*6., seed))
		pal = PaletteRoad
	} else if index == PlaneIndex {
		// TODO: make the plane more fancy maybe ?
		return vec3(0., 0., 0.247)
	} else if index == BlockHarderIndex {
		t = noise(p.xy*8., seed)
		pal = PaletteBlockHarder
	} else if index == BlockHarder2Index {
		t = noise(p.xy*8., seed)
		pal = PaletteBlockHarder2
	} else if index == BlockHeartIndex {
		p = normalize(p)
		y := -p.y-abs(p.x)
		t = abs(sqrt(p.x*p.x+y*y) - 1.0)
		// t = noise(p.xy*64., seed) // length(p.xy)-0.1 // noise(p.xy*8./5.)
		pal = PaletteHeart
	}
	
	return palette(t, pal[0], pal[1], pal[2], pal[3])
}

func translate(p, offset vec3) vec3 {
	return p - offset
}

func sdSphere(p vec3, r float, offset vec3, index float) mat3 {
	p = translate(p, offset)
	d := length(p) - r

	return mat3(
		vec3(d, index, 0.),
		offset,
		vec3(0.),
	)
}

func sdBox(p, b, offset vec3, index float) mat3 {
	p = translate(p, offset)
	q := abs(p) - b
  	d := length(max(q,0.0)) + min(max(q.x,max(q.y,q.z)),0.0)

	return mat3(
		vec3(d, index, 0.),
		offset,
		vec3(0.),
	)
}

func sdPlane(p, n vec3, h float, index float) mat3 {
	// n must be normalized
	d := dot(p,n) + h

	return mat3(
		vec3(d, index, 0.),
		vec3(0.),
		vec3(0.),
	)
}

func sdHeart(p, b, offset vec3, index float) mat3 {
	p = translate(p, offset)

	x := p.x
	y := -p.y*1.2-abs(p.x)
	z := p.z
	d := sqrt(x*x+y*y+z*z) - b.x/5.*4. // Dirty fix
	offset.y += p.y*0.2

	return mat3(
		vec3(d, index, 0.),
		offset,
		vec3(0.),
	)
}

func sdBlock(p vec3, i float) mat3 {
	bi := int(i)

	bs := BlockSizes[bi]
	blockOffset := vec3(0., 0.999-bs.y/2, 0.)
	blockOffset = translate(blockOffset, BlockPositions[bi])

	kind := BlockKinds[bi]
	if kind == BlockIndex {
		obj := sdBox(p, vec3(bs.x, bs.y, bs.x), blockOffset, BlockIndex)
		p = translate(p, obj[1].xyz)
		c := 20.
		obj[0].x -= noise(p.xy*c, BlockSeeds[bi])*p.z*0.2
		return obj
	} else if kind == BlockHarderIndex {
		obj := sdBox(p, vec3(bs.x, bs.y, bs.x), blockOffset, BlockHarderIndex)
		p = translate(p, obj[1].xyz)
		c := 30.
		obj[0].x -= noise(p.xy*c, BlockSeeds[bi])*p.z*0.2
		return obj
	} else if kind == BlockHarder2Index {
		obj := sdBox(p, vec3(bs.x, bs.y, bs.x), blockOffset, BlockHarder2Index)
		p = translate(p, obj[1].xyz)
		c := 40.
		obj[0].x -= noise(p.xy*c, BlockSeeds[bi])*p.z*0.2
		return obj
	} else if kind == BlockHeartIndex {
		return sdHeart(p, vec3(bs.x, bs.y, bs.x), blockOffset, BlockHeartIndex)
	}

	return mat3(0.)
}

func minWithData(obj1, obj2 mat3) mat3 {
	if obj2[0].x < obj1[0].x {
		return obj2
	}

	return obj1
}

func colorFromObj(p vec3, obj mat3) vec3 {
	p = translate(p, obj[1])
	seed := BlockSeeds[int(obj[0].y)]
	if obj[0].y >= BlockIndex {
		seed += obj[1].x
	}
	return colorize(p, -obj[0].x, obj[0].y, seed)
}

func sdScene(p vec3) mat3 {
	scene := sdPlane(p, normalize(vec3(0., -1., 0.)), 1., PlaneIndex) // default floor
	// Plane vertical displacement only off-road
	if p.x < -1. || p.x > 1. {
		dp := vec2(p.x, p.z)+vec2(0., -Distance)
		c := abs(1.-(abs(p.x)))*0.1
		scene[0].x -= noise(dp.xy*10., BlockSeeds[0])*c
	}

	roadl := 100.
	roadh := 0.35
	roadw := 1.0
	roadOffset := vec3(0., 0.999+roadh, -1.)
	road := sdBox(p, vec3(roadw, roadh, roadl), roadOffset, RoadIndex)
	// Road vertical displacement
	dp := vec2(p.x, p.z)+vec2(0., -Distance)
	road[0].x -= noise(dp.xy*10., BlockSeeds[0])*0.02

	sphereOffset := vec3(0., 0.999, 0.)
	sphereOffset = translate(sphereOffset, PlayerPosition)
	// TODO: * 2 radius is a hack to make sense with software value
	spherePlayer := sdSphere(p, PlayerRadius*2., sphereOffset, PlayerIndex)
	
	scene = minWithData(scene, minWithData(road, spherePlayer))
	for i := 0.; i < MaxBlocks; i++ {
		if i >= BlockCount {
			break
		}
		block := sdBlock(p, i)
		scene = minWithData(scene, block)
	}
	
	return scene
}

func rayMarch(ro, rd vec3, start, end float) mat3 {
	const (
		MaxSteps = 64. // TODO: Can lower this constant on-need for performance
		Precision = 0.005 // TODO: was 0.001
	)

	depth := start
	var obj mat3
	for i := 0; i < MaxSteps; i++ {
		p := ro + depth * rd
		obj = sdScene(p)
		depth += obj[0].x
    	if obj[0].x < Precision || depth > end {
			break
		}
	}

	obj[0].x = depth
	return obj
}

func calcNormal(p vec3) vec3 {
    e := vec2(1.0, -1.0) * 0.0005 // epsilon
    
	return normalize(
    	e.xyy * sdScene(p + e.xyy)[0].x +
    	e.yyx * sdScene(p + e.yyx)[0].x +
    	e.yxy * sdScene(p + e.yxy)[0].x +
    	e.xxx * sdScene(p + e.xxx)[0].x)
}

func phong(lightDir, normal, rd, clr vec3) vec3 {
	// ambient
	ambient := clr * 0.5
  
	// diffuse
	dotLN := clamp(dot(lightDir, normal), 0., 1.)
	diffuse := clr * dotLN
  
	// specular
	halfwayDir := normalize(lightDir + normal)
    specular := vec3(0.25)*pow(max(dot(normal, halfwayDir), 0.0), 16.0)
  
	return ambient + diffuse + specular
}
  
func softShadow(ro, rd vec3, mint, tmax float) float {
	const (
		MaxSteps = 8.
		Precision = 0.001
	)

	res := 1.0
	t := mint
	for i := 0.; i < MaxSteps; i++ {
		h := sdScene(ro + rd * t)[0].x
		res = min(res, 8.0*h/t)
		t += clamp(h, 0.02, 0.10)
		if h < Precision || t > tmax {
			break
		}
	}
  
	return clamp(res, 0.0, 1.0)
}

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	uv := (position.xy / ScreenSize) * 2. - 1.
	bgColor := background(vec3(uv.xy, Distance))

	// Early abort if at top part of screen
	if uv.y < 0. {
		return vec4(bgColor, 1.)
	}

  	ro := Camera
	rd := normalize(vec3(uv, -1.)) // ray direction

	obj := rayMarch(ro, rd, 0., MaxDepth)
	d := obj[0].x

	var clr vec3

	if (d > MaxDepth) {
		clr = bgColor // ray didn't hit anything
	} else {
		p := ro + rd * d

		clr = colorFromObj(p, obj)

		// Light stuff
		normal := calcNormal(p)
    	lightPosition := ro - vec3(0, 16., -32.) // let's say light is at camera position
    	lightDirection := normalize(lightPosition - p)
		lightIntensity := 1.0

		softShadows := clamp(softShadow(p, lightDirection, 0.02, 2.5), 0.1, 1.0)
	
		clr = lightIntensity * phong(lightDirection, normal, rd, clr)
		clr *= softShadows
	}

	clr = mix(clr, bgColor, 1.0-exp(-0.0002 * d * d * d)) // Fog
	
	return vec4(clr, 1.)
}
`)

	RaymarchShader *ebiten.Shader
)

func init() {
	var err error

	RaymarchShader, err = ebiten.NewShader(raymarchShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}

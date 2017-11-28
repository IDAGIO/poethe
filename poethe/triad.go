package poethe

import (
	"image/color"
	"math/rand"
	"time"
)

const (
	greyControl = 1
	alpha       = 255
)

type TriadGenerator struct {
}

func (TriadGenerator) Description() string {
	return "The awesome color generator (using triad mixing)"
}

// RandomMix we receive three RGBA colors and we return five RGBA colors (the triad)
func (TriadGenerator) Generate(colors []*color.RGBA) []*color.RGBA {
	rand.Seed(time.Now().UnixNano())

	pallete := make([]*color.RGBA, 5)

	for index := range pallete {
		pallete[index] = generateColor(colors)
	}

	return pallete
}

// generateColor generate one color at time
func generateColor(colors []*color.RGBA) *color.RGBA {
	randomIndex := rand.Int() % 3

	mixRatio1 := rand.Float32()
	mixRatio2 := rand.Float32()
	mixRatio3 := rand.Float32()

	if randomIndex == 0 {
		mixRatio1 = rand.Float32() * greyControl
	}

	if randomIndex == 1 {
		mixRatio2 = rand.Float32() * greyControl
	}

	if randomIndex == 2 {
		mixRatio3 = rand.Float32() * greyControl
	}

	sum := mixRatio1 + mixRatio2 + mixRatio3

	mixRatio1 /= sum
	mixRatio2 /= sum
	mixRatio3 /= sum

	return &color.RGBA{
		R: uint8(mixRatio1*float32(colors[0].R) + mixRatio2*float32(colors[1].R) + mixRatio3*float32(colors[2].R)),
		G: uint8(mixRatio1*float32(colors[0].G) + mixRatio2*float32(colors[1].G) + mixRatio3*float32(colors[2].G)),
		B: uint8(mixRatio1*float32(colors[0].B) + mixRatio2*float32(colors[1].B) + mixRatio3*float32(colors[2].B)),
		A: alpha,
	}
}

package poethe_test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/idagio/poethe/poethe"
)

func TestReturnColorListGivenFewColors(t *testing.T) {
	colors := []*color.RGBA{
		&color.RGBA{11, 12, 113, 255},
		&color.RGBA{21, 53, 67, 255},
		&color.RGBA{117, 228, 90, 255},
	}

	triad := poethe.TriadGenerator{}
	pallete := triad.Generate(colors)

	if pallete == nil {
		t.Fatalf("Returning nil when expecting %+v", pallete)
	}

	for _, value := range pallete {
		fmt.Printf("Pallete: %+v", value)
	}
}

func TestReturnFiveColorsList(t *testing.T) {
	colors := []*color.RGBA{
		&color.RGBA{11, 12, 113, 255},
		&color.RGBA{21, 53, 67, 255},
		&color.RGBA{117, 228, 90, 255},
	}

	triad := poethe.TriadGenerator{}
	pallete := triad.Generate(colors)

	if len(pallete) != 5 {
		t.Fatalf("Returning nil when expecting %+v", pallete)
	}

	for _, value := range pallete {
		fmt.Printf("Pallete: %+v", value)
	}
}

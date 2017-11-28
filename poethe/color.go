package poethe

import (
	"fmt"
	"image/color"
)

type col color.RGBA

type Color interface {
	Rgb255()(uint8, uint8, uint8)
}

func (c col) Rgb255() (uint8, uint8, uint8){
	return c.R, c.G, c.B
}


func Parse (s string) (Color, error){
	var r, g, b uint8
	_, err := fmt.Sscanf(s,"#%2x%2x%2x", &r, &g, &b)
	if err != nil {
		return nil, err
	}

	return col {
		r, g, b, 1,
	},nil
}
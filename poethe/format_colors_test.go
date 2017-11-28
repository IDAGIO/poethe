package poethe

import "testing"

type color struct {
	R, G, B int
}

func (ColorMock color) Color255() (r, g, b int) {

	return ColorMock.R, ColorMock.G, ColorMock.B
}

func TestFormatColors(t *testing.T) {

	var tests = []struct {
		Input  []Color
		Output string
	}{
		{
			Input:  []Color{color{255, 255, 0}, color{255, 0, 0}, color{255, 255, 255}, color{0, 255, 0}},
			Output: "ffff00,ff0000,ffffff,00ff00",
		},
	}

	for _, colorTest := range tests {
		output := FormatColors(colorTest.Input)
		if output != colorTest.Output {
			//t.Error("Output '" + output + "' does not match expected value of '" + colorTest.Output + "'")
			t.Errorf("Output '%s' does not match expected value of '%s'", output, colorTest.Output)
		}
	}
}

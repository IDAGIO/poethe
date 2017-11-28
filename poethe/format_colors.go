package poethe

import "fmt"

func FormatColors(colors []Color) string {
	// exampe output: ff0000,0000ff
	colorString := ""
	for _, color := range colors {
		var r, g, b = color.Color255()
		_, _, _ = r, g, b
		tmp := fmt.Sprintf("%.2x%.2x%.2x", r, g, b)
		if colorString != "" {
			colorString = colorString + "," + tmp
		} else {
			colorString = tmp
		}
	}

	return colorString
}

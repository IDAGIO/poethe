package poethe

import (
	"testing"
)

func TestParse(t *testing.T) {
	result, _:= Parse("#FFFFFF")
	r, g, b := result.Rgb255()
	if r != 255 || g != 255 || b != 255 {
		t.Fail()
	}
}

package bicycle

import (
	"testing"
)

func TestBaseGearRatio(t *testing.T) {
	g := BaseGear{2, 1}
	result := g.Ratio()
	if result != float64(2) {
		t.Error("Expected 2, got ", result)
	}
}

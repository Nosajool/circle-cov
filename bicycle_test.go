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

func TestBaseGearInches(t *testing.T) {
	g := BaseGear{2, 1}
	result := g.GearInches(3)
	if result != float64(6) {
		t.Error("Expected 6, got ", result)
	}
}

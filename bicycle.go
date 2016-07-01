package bicycle

type Gear interface {
	GearInches(diameter int) float64
	Ratio() float64
}

type BaseGear struct {
	ChainRing int
	Cog       int
}

type Wheel struct {
	Rim  int
	Tire float64
	Gear Gear
}

func (g *BaseGear) Ratio() float64 {
	return float64(g.ChainRing) / float64(g.Cog)
}

func (g *BaseGear) GearInches(diameter int) float64 {
	return g.Ratio() * float64(diameter)
}

func (w *Wheel) Diameter() float64 {
	return float64(w.Rim) + (w.Tire * float64(2))
}

func (w *Wheel) GearInches() float64 {
	return w.Gear.GearInches(int(w.Diameter()))
}

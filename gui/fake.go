package gui

import (
	"math"
	"math/rand"

	"github.com/gurupras/dota2bot"
	"github.com/sirupsen/logrus"
)

func FakeGameInfo() dota2bot.GameInfo {
	g := dota2bot.GameInfo{}
	g.TeamTypes.Dire = 1
	g.TeamTypes.Radiant = 2
	g.WorldBounds.MinX = -8288
	g.WorldBounds.MinY = -8288
	g.WorldBounds.MaxX = 8288
	g.WorldBounds.MaxY = 8288
	return g
}

func FakeUpdate() []dota2bot.Unit {
	gi := FakeGameInfo()

	numUnits := rand.Intn(100)
	units := make([]dota2bot.Unit, numUnits)
	for idx := 0; idx < numUnits; idx++ {
		u := dota2bot.Unit{}

		xAbs := math.Abs(gi.WorldBounds.MinX) + math.Abs(gi.WorldBounds.MaxX)
		x := float64(rand.Intn(int(xAbs))) - math.Abs(gi.WorldBounds.MinX)

		yAbs := math.Abs(gi.WorldBounds.MinY) + math.Abs(gi.WorldBounds.MaxY)
		y := float64(rand.Intn(int(yAbs))) - math.Abs(gi.WorldBounds.MinY)

		logrus.Infof("fake location: {%v,%v}", int(x), int(y))

		u.Location = dota2bot.Location{
			X: x,
			Y: y,
			Z: 0,
		}
		isHero := rand.Float32() > 0.5
		u.IsHero = isHero
		u.IsCreep = !isHero
		if rand.Float32() > 0.5 {
			u.Team = gi.TeamTypes.Dire
		} else {
			u.Team = gi.TeamTypes.Radiant
		}
		units[idx] = u
	}
	return units
}

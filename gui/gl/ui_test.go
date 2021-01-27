package dota2bot

import (
	"sync"
	"testing"

	"github.com/gurupras/dota2bot"
	"github.com/stretchr/testify/require"
)

func TestBasicUI(t *testing.T) {
	require := require.New(t)
	g := dota2bot.GameInfo{}
	g.TeamTypes.Dire = 1
	g.WorldBounds.MinX = -8288
	g.WorldBounds.MinY = -8288
	g.WorldBounds.MaxX = 8288
	g.WorldBounds.MaxY = 8288

	m, err := NewMiniMap(g)
	require.Nil(err)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Run()
	}()

	u := dota2bot.Unit{}
	u.Location = dota2bot.Location{
		4096,
		4096,
		0,
	}
	u.IsHero = true
	u.Team = g.TeamTypes.Dire

	units := make([]dota2bot.Unit, 0)
	units = append(units, u)

	m.Update(units)

	wg.Wait()
}

package web

import (
	"testing"

	"github.com/gurupras/dota2bot"
	"github.com/gurupras/dota2bot/gui"
	"github.com/stretchr/testify/require"
)

func TestCoordinateConversion(t *testing.T) {
	require := require.New(t)

	gi := gui.FakeGameInfo()

	w := NewWebUI(gi)

	l := dota2bot.Location{gi.WorldBounds.MinX, gi.WorldBounds.MinY, 0}
	p := w.GameCoordinatesToImageCoordinates(l)
	require.Equal(p, gui.Point{0, 0})
}

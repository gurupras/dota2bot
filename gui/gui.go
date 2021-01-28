package gui

import (
	"image"
	"math"
	"sync"

	"github.com/gurupras/dota2bot"
)

type UnitUpdate struct {
	Name   string `json:"name"`
	IsHero bool   `json:"isHero"`
	Point  `json:"point"`
	Radius float64 `json:"radius"`
	Red    int     `json:"red"`
	Green  int     `json:"green"`
	Blue   int     `json:"blue"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type BaseGUI struct {
	sync.Mutex
	*dota2bot.GameInfo
	Bounds     image.Rectangle
	LastUpdate []*UnitUpdate
}

type GUI interface {
	Update(units []dota2bot.Unit)
	Run() error
}

func NewBaseGUI(gameInfo *dota2bot.GameInfo, bounds image.Rectangle) *BaseGUI {
	b := &BaseGUI{
		sync.Mutex{},
		gameInfo,
		bounds,
		nil,
	}
	return b
}

func (b *BaseGUI) GameCoordinatesToImageCoordinates(location dota2bot.Location) Point {
	// We need to convert to image bounds
	pctX := (math.Abs(b.WorldBounds.MinX) + location.X) / (b.WorldBounds.MaxX + math.Abs(b.WorldBounds.MinX))
	mapX := float64(b.Bounds.Max.X) * pctX

	pctY := (math.Abs(b.WorldBounds.MinY) + location.Y) / (b.WorldBounds.MaxY + math.Abs(b.WorldBounds.MinY))
	pctY = 1 - pctY
	mapY := float64(b.Bounds.Max.Y) * pctY

	return Point{
		mapX,
		mapY,
	}
}

func (b *BaseGUI) Update(units []dota2bot.Unit) {
	unitUpdates := make([]*UnitUpdate, 0)
	for _, unit := range units {
		point := b.GameCoordinatesToImageCoordinates(unit.Location)
		var (
			red    int
			green  int
			blue   int
			radius float64
		)
		if unit.IsHero {
			radius = 0.0075 * float64(b.Bounds.Dx())
		} else if unit.IsCreep {
			radius = 0.005 * float64(b.Bounds.Dx())
		}
		if unit.Team == b.GameInfo.TeamTypes.Dire {
			red = 255
		} else {
			green = 255
		}
		if radius > 0 {
			unitUpdates = append(unitUpdates, &UnitUpdate{
				Name:   unit.Name,
				IsHero: unit.IsHero,
				Point:  point,
				Radius: radius,
				Red:    red,
				Green:  green,
				Blue:   blue,
			})
		}
	}
	b.Lock()
	defer b.Unlock()
	b.LastUpdate = unitUpdates
}

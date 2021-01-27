package dota2bot

import "sync"

type Team struct {
	Name     string
	TeamID   int
	DriverID int
	WorldInfo
	Bots map[int]Unit
	sync.Mutex
}

func NewTeam(name string) *Team {
	team := Team{}
	team.Name = name
	team.Bots = make(map[int]Unit)
	return &team
}

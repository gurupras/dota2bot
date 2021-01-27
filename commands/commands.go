package commands

import "github.com/gurupras/dota2bot"

type Command struct {
	Name string      `json:"name"`
	Args interface{} `json:"args"`
}

const moveToLocation = "moveToLocation"
const attackTarget = "attackTarget"
const sendWorldUpdate = "sendWorldUpdate"

func CommandMoveToLocation(l *dota2bot.Location) Command {
	return Command{
		moveToLocation,
		l,
	}
}

func CommandAttackTarget(id int) Command {
	return Command{
		attackTarget,
		id,
	}
}

func CommandDesignateToSendWorldInfo() Command {
	return Command{
		sendWorldUpdate,
		true,
	}
}

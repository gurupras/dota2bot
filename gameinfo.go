package dota2bot

type Player struct {
	TeamID           int    `json:"teamID"`
	PlayerID         int    `json:"playerID"`
	SelectedHeroName string `json:"selectedHeroName"`
	IsBot            bool   `json:"isPlayerBot"`
}

type WorldBounds struct {
	MinX float64 `json:"minX"`
	MinY float64 `json:"minY"`
	MaxX float64 `json:"maxX"`
	MaxY float64 `json:"maxY"`
}

type ShopLocation struct {
	ShopID   int `json:"shopID"`
	Location `json:"location"`
}

type TeamShopLocations struct {
	TeamID        int            `json:"teamID"`
	ShopLocations []ShopLocation `json:"shopLocations"`
}

type RuneSpawnLocation struct {
	RuneSpawnType int `json:"runeSpawnType"`
	Location      `json:"location"`
}

type AncientLocation struct {
	TeamID   int `json:"teamID"`
	Location `json:"location"`
}

type TeamPlayers []Player

type GameInfo struct {
	GameTime               float64                `json:"gameTime"`
	DotaTime               float64                `json:"dotaTime"`
	DamageTypes            DamageTypes            `json:"damageTypes"`
	TeamTypes              TeamTypes              `json:"teamTypes"`
	RuneTypes              RuneTypes              `json:"runeTypes"`
	RuneStatusTypes        RuneStatusTypes        `json:"runeStatusTypes"`
	RuneSpawnTypes         RuneSpawnTypes         `json:"runeSpawnTypes"`
	ItemSlotTypes          ItemSlotTypes          `json:"itemSlotTypes"`
	BotActionTypes         BotActionTypes         `json:"botActionTypes"`
	CourierActionTypes     CourierActionTypes     `json:"courierActionTypes"`
	CourierStateTypes      CourierStateTypes      `json:"courierStateTypes"`
	ShopTypes              ShopTypes              `json:"shopTypes"`
	AbilityTargetTeamTypes AbilityTargetTeamTypes `json:"abilityTargetTeamTypes"`
	AbilityTargetTypes     AbilityTargetTypes     `json:"abilityTargetTypes"`
	LaneTypes              LaneTypes              `json:"laneTypes"`
	TeamPlayers            [][]Player             `json:"teams"`
	WorldBounds            WorldBounds            `json:"worldBounds"`
	ShopLocations          []TeamShopLocations    `json:"shopLocations"`
	RuneSpawnLocations     []RuneSpawnLocation    `json:"runeSpawnLocations"`
	AncientLocations       []AncientLocation      `json:"ancientLocations"`
	ItemTypes              ItemsMap               `json:"itemTypes"`
	BotModeTypes           BotModeTypes           `json:"botModeTypes"`
	BotActionDesireTypes   BotActionDesireTypes   `json:"botActionDesireTypes"`
	BotModeDesireTypes     BotModeDesireTypes     `json:"botModeDesireTypes"`
	TowerTypes             TowerTypes             `json:"towerTypes"`
	BarrackTypes           BarrackTypes           `json:"barrackTypes"`
	ShrineTypes            ShrineTypes            `json:"shrineTypes"`
	AbilityTargetFlagTypes AbilityTargetFlagTypes `json:"abilityTargetFlagTypes"`
	AnimationActivityTypes AnimationActivityTypes `json:"animationActivityTypes"`
	PlayerID               int                    `json:"playerID"`
	TeamID                 int                    `json:"teamID"`
}

func (gi GameInfo) GetBotModeString(mode int) string {
	switch mode {
	case gi.BotModeTypes.Laning:
		return "laning"
	case gi.BotModeTypes.Attack:
		return "attack"
	case gi.BotModeTypes.Roam:
		return "roam"
	case gi.BotModeTypes.Retreat:
		return "retreat"
	case gi.BotModeTypes.SecretShop:
		return "secret_shop"
	case gi.BotModeTypes.SideShop:
		return "side_shop"
	case gi.BotModeTypes.PushTowerTop:
		return "push_tower_top"
	case gi.BotModeTypes.PushTowerMid:
		return "push_tower_mid"
	case gi.BotModeTypes.PushTowerBot:
		return "push_tower_bot"
	case gi.BotModeTypes.DefendTowerTop:
		return "defend_tower_top"
	case gi.BotModeTypes.DefendTowerMid:
		return "defend_tower_mid"
	case gi.BotModeTypes.DefendTowerBot:
		return "defend_tower_bot"
	case gi.BotModeTypes.Assemble:
		return "assemble"
	case gi.BotModeTypes.TeamRoam:
		return "team_roam"
	case gi.BotModeTypes.Farm:
		return "farm"
	case gi.BotModeTypes.DefendAlly:
		return "defend_ally"
	case gi.BotModeTypes.EvasiveManeuvers:
		return "evasive_maneuvers"
	case gi.BotModeTypes.Roshan:
		return "roshan"
	case gi.BotModeTypes.Item:
		return "item"
	case gi.BotModeTypes.Ward:
		return "ward"
	default:
		return "UNIMPLEMENTED"
	}
}

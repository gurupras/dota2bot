package dota2bot

type DamageTypes struct {
	Physical int `json:"DAMAGE_TYPE_PHYSICAL"`
	Magical  int `json:"DAMAGE_TYPE_MAGICAL"`
	Pure     int `json:"DAMAGE_TYPE_PURE"`
	All      int `json:"DAMAGE_TYPE_ALL"`
}

type TeamTypes struct {
	Radiant int `json:"TEAM_RADIANT"`
	Dire    int `json:"TEAM_DIRE"`
	Neutral int `json:"TEAM_NEUTRAL"`
	None    int `json:"TEAM_NONE"`
}

type RuneTypes struct {
	Invalid      int `json:"RUNE_INVALID"`
	DoubleDamage int `json:"RUNE_DOUBLEDAMAGE"`
	Haste        int `json:"RUNE_HASTE"`
	Illusion     int `json:"RUNE_ILLUSION"`
	Invisibility int `json:"RUNE_INVISIBILITY"`
	Regeneration int `json:"RUNE_REGENERATION"`
	Bounty       int `json:"RUNE_BOUNTY"`
	Arcane       int `json:"RUNE_ARCANE"`
}

type RuneStatusTypes struct {
	Unknown   int `json:"RUNE_STATUS_UNKNOWN"`
	Available int `json:"RUNE_STATUS_AVAILABLE"`
	Missing   int `json:"RUNE_STATUS_MISSING"`
}

type RuneSpawnTypes struct {
	PowerUp1 int `json:"RUNE_POWERUP_1"`
	PowerUp2 int `json:"RUNE_POWERUP_2"`
	Bounty1  int `json:"RUNE_BOUNTY_1"`
	Bounty2  int `json:"RUNE_BOUNTY_2"`
	Bounty3  int `json:"RUNE_BOUNTY_3"`
	Bounty4  int `json:"RUNE_BOUNTY_4"`
}

type ItemSlotTypes struct {
	Invalid  int `json:"ITEM_SLOT_TYPE_INVALID"`
	Main     int `json:"ITEM_SLOT_TYPE_MAIN"`
	Backpack int `json:"ITEM_SLOT_TYPE_BACKPACK"`
	Stash    int `json:"ITEM_SLOT_TYPE_STASH"`
}

type BotActionTypes struct {
	None           int `json:"BOT_ACTION_TYPE_NONE"`
	Idle           int `json:"BOT_ACTION_TYPE_IDLE"`
	MoveTo         int `json:"BOT_ACTION_TYPE_MOVE_TO"`
	MoveToDirectly int `json:"BOT_ACTION_TYPE_MOVE_TO_DIRECTLY"`
	Attack         int `json:"BOT_ACTION_TYPE_ATTACK"`
	AttackMove     int `json:"BOT_ACTION_TYPE_ATTACKMOVE"`
	UseAbility     int `json:"BOT_ACTION_TYPE_USE_ABILITY"`
	PickUpRune     int `json:"BOT_ACTION_TYPE_PICK_UP_RUNE"`
	PickUpItem     int `json:"BOT_ACTION_TYPE_PICK_UP_ITEM"`
	DropItem       int `json:"BOT_ACTION_TYPE_DROP_ITEM"`
	Shrine         int `json:"BOT_ACTION_TYPE_SHRINE"`
	Delay          int `json:"BOT_ACTION_TYPE_DELAY"`
}

type CourierActionTypes struct {
	Burst                int `json:"COURIER_ACTION_BURST"`
	EnemySecretShop      int `json:"COURIER_ACTION_ENEMY_SECRET_SHOP"`
	Return               int `json:"COURIER_ACTION_RETURN"`
	SecretShop           int `json:"COURIER_ACTION_SECRET_SHOP"`
	SideShop             int `json:"COURIER_ACTION_SIDE_SHOP"`
	SideShop2            int `json:"COURIER_ACTION_SIDE_SHOP2"`
	TakeStashItems       int `json:"COURIER_ACTION_TAKE_STASH_ITEMS"`
	TakeAndTransferItems int `json:"COURIER_ACTION_TAKE_AND_TRANSFER_ITEMS"`
	TransferItems        int `json:"COURIER_ACTION_TRANSFER_ITEMS"`
}

type CourierStateTypes struct {
	Idle            int `json:"COURIER_STATE_IDLE"`
	AtBase          int `json:"COURIER_STATE_AT_BASE"`
	Moving          int `json:"COURIER_STATE_MOVING"`
	DeliveringItems int `json:"COURIER_STATE_DELIVERING_ITEMS"`
	ReturningToBase int `json:"COURIER_STATE_RETURNING_TO_BASE"`
	Dead            int `json:"COURIER_STATE_DEAD"`
}

type ShopTypes struct {
	Home    int `json:"SHOP_HOME"`
	Side    int `json:"SHOP_SIDE"`
	Side2   int `json:"SHOP_SIDE2"`
	Secret  int `json:"SHOP_SECRET"`
	Secret2 int `json:"SHOP_SECRET2"`
}

type AbilityTargetTeamTypes struct {
	None     int `json:"ABILITY_TARGET_TEAM_NONE"`
	Friendly int `json:"ABILITY_TARGET_TEAM_FRIENDLY"`
	Enemy    int `json:"ABILITY_TARGET_TEAM_ENEMY"`
}

type AbilityTargetTypes struct {
	None     int `json:"ABILITY_TARGET_TYPE_NONE"`
	Hero     int `json:"ABILITY_TARGET_TYPE_HERO"`
	Creep    int `json:"ABILITY_TARGET_TYPE_CREEP"`
	Building int `json:"ABILITY_TARGET_TYPE_BUILDING"`
	Courier  int `json:"ABILITY_TARGET_TYPE_COURIER"`
	Other    int `json:"ABILITY_TARGET_TYPE_OTHER"`
	Tree     int `json:"ABILITY_TARGET_TYPE_TREE"`
	Basic    int `json:"ABILITY_TARGET_TYPE_BASIC"`
	All      int `json:"ABILITY_TARGET_TYPE_ALL"`
}

type LaneTypes struct {
	None int `json:"LANE_NONE"`
	Top  int `json:"LANE_TOP"`
	Mid  int `json:"LANE_MID"`
	Bot  int `json:"LANE_BOT"`
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type LocationMap = map[int]Location

type ShopLocations = LocationMap
type RuneSpawnLocations = LocationMap
type AncientLocations = LocationMap

type Item struct {
	Name                      string `json:"name"`
	Cost                      int    `json:"cost"`
	IsPurchasedFromSecretShop bool   `json:"isItemPurchasedFromSecretShop"`
	IsPurchasedFromSideShop   bool   `json:"isItemPurchasedFromSideShop"`
}

type ItemsMap = map[string]Item

type BotModeTypes struct {
	None             int `json:"BOT_MODE_NONE"`
	Laning           int `json:"BOT_MODE_LANING"`
	Attack           int `json:"BOT_MODE_ATTACK"`
	Roam             int `json:"BOT_MODE_ROAM"`
	Retreat          int `json:"BOT_MODE_RETREAT"`
	SecretShop       int `json:"BOT_MODE_SECRET_SHOP"`
	SideShop         int `json:"BOT_MODE_SIDE_SHOP"`
	PushTowerTop     int `json:"BOT_MODE_PUSH_TOWER_TOP"`
	PushTowerMid     int `json:"BOT_MODE_PUSH_TOWER_MID"`
	PushTowerBot     int `json:"BOT_MODE_PUSH_TOWER_BOT"`
	DefendTowerTop   int `json:"BOT_MODE_DEFEND_TOWER_TOP"`
	DefendTowerMid   int `json:"BOT_MODE_DEFEND_TOWER_MID"`
	DefendTowerBot   int `json:"BOT_MODE_DEFEND_TOWER_BOT"`
	Assemble         int `json:"BOT_MODE_ASSEMBLE"`
	TeamRoam         int `json:"BOT_MODE_TEAM_ROAM"`
	Farm             int `json:"BOT_MODE_FARM"`
	DefendAlly       int `json:"BOT_MODE_DEFEND_ALLY"`
	EvasiveManeuvers int `json:"BOT_MODE_EVASIVE_MANEUVERS"`
	Roshan           int `json:"BOT_MODE_ROSHAN"`
	Item             int `json:"BOT_MODE_ITEM"`
	Ward             int `json:"BOT_MODE_WARD"`
}

type BotActionDesireTypes struct {
	None     float64 `json:"BOT_ACTION_DESIRE_NONE"`
	Verylow  float64 `json:"BOT_ACTION_DESIRE_VERYLOW"`
	Low      float64 `json:"BOT_ACTION_DESIRE_LOW"`
	Moderate float64 `json:"BOT_ACTION_DESIRE_MODERATE"`
	High     float64 `json:"BOT_ACTION_DESIRE_HIGH"`
	Veryhigh float64 `json:"BOT_ACTION_DESIRE_VERYHIGH"`
	Absolute float64 `json:"BOT_ACTION_DESIRE_ABSOLUTE"`
}

type BotModeDesireTypes struct {
	None     float64 `json:"BOT_MODE_DESIRE_NONE"`
	Verylow  float64 `json:"BOT_MODE_DESIRE_VERYLOW"`
	Low      float64 `json:"BOT_MODE_DESIRE_LOW"`
	Moderate float64 `json:"BOT_MODE_DESIRE_MODERATE"`
	High     float64 `json:"BOT_MODE_DESIRE_HIGH"`
	Veryhigh float64 `json:"BOT_MODE_DESIRE_VERYHIGH"`
	Absolute float64 `json:"BOT_MODE_DESIRE_ABSOLUTE"`
}

type TowerTypes struct {
	Top1  int `json:"TOWER_TOP_1"`
	Top2  int `json:"TOWER_TOP_2"`
	Top3  int `json:"TOWER_TOP_3"`
	Mid1  int `json:"TOWER_MID_1"`
	Mid2  int `json:"TOWER_MID_2"`
	Mid3  int `json:"TOWER_MID_3"`
	Bot1  int `json:"TOWER_BOT_1"`
	Bot2  int `json:"TOWER_BOT_2"`
	Bot3  int `json:"TOWER_BOT_3"`
	Base1 int `json:"TOWER_BASE_1"`
	Base2 int `json:"TOWER_BASE_2"`
}

type BarrackTypes struct {
	TopMelee  int `json:"BARRACKS_TOP_MELEE"`
	TopRanged int `json:"BARRACKS_TOP_RANGED"`
	MidMelee  int `json:"BARRACKS_MID_MELEE"`
	MidRanged int `json:"BARRACKS_MID_RANGED"`
	BotMelee  int `json:"BARRACKS_BOT_MELEE"`
	BotRanged int `json:"BARRACKS_BOT_RANGED"`
}

type ShrineTypes struct {
	Base1   int `json:"SHRINE_BASE_1"`
	Base2   int `json:"SHRINE_BASE_2"`
	Base3   int `json:"SHRINE_BASE_3"`
	Base4   int `json:"SHRINE_BASE_4"`
	Base5   int `json:"SHRINE_BASE_5"`
	Jungle1 int `json:"SHRINE_JUNGLE_1"`
	Jungle2 int `json:"SHRINE_JUNGLE_2"`
}

type AbilityTargetFlagTypes struct {
	None                 int `json:"ABILITY_TARGET_FLAG_NONE"`
	RangedOnly           int `json:"ABILITY_TARGET_FLAG_RANGED_ONLY"`
	MeleeOnly            int `json:"ABILITY_TARGET_FLAG_MELEE_ONLY"`
	Dead                 int `json:"ABILITY_TARGET_FLAG_DEAD"`
	MagicImmuneEnemies   int `json:"ABILITY_TARGET_FLAG_MAGIC_IMMUNE_ENEMIES"`
	NotMagicImmuneAllies int `json:"ABILITY_TARGET_FLAG_NOT_MAGIC_IMMUNE_ALLIES"`
	Invulnerable         int `json:"ABILITY_TARGET_FLAG_INVULNERABLE"`
	FowVisible           int `json:"ABILITY_TARGET_FLAG_FOW_VISIBLE"`
	NoInvis              int `json:"ABILITY_TARGET_FLAG_NO_INVIS"`
	NotAncients          int `json:"ABILITY_TARGET_FLAG_NOT_ANCIENTS"`
	PlayerControlled     int `json:"ABILITY_TARGET_FLAG_PLAYER_CONTROLLED"`
	NotDominated         int `json:"ABILITY_TARGET_FLAG_NOT_DOMINATED"`
	NotSummoned          int `json:"ABILITY_TARGET_FLAG_NOT_SUMMONED"`
	NotIllusions         int `json:"ABILITY_TARGET_FLAG_NOT_ILLUSIONS"`
	NotAttackImmune      int `json:"ABILITY_TARGET_FLAG_NOT_ATTACK_IMMUNE"`
	ManaOnly             int `json:"ABILITY_TARGET_FLAG_MANA_ONLY"`
	CheckDisableHelp     int `json:"ABILITY_TARGET_FLAG_CHECK_DISABLE_HELP"`
	NotCreepHero         int `json:"ABILITY_TARGET_FLAG_NOT_CREEP_HERO"`
	OutOfWorld           int `json:"ABILITY_TARGET_FLAG_OUT_OF_WORLD"`
	NotNightmared        int `json:"ABILITY_TARGET_FLAG_NOT_NIGHTMARED"`
	PreferEnemies        int `json:"ABILITY_TARGET_FLAG_PREFER_ENEMIES"`
}

type AnimationActivityTypes struct {
	Idle               int `json:"ACTIVITY_IDLE"`
	IdleRare           int `json:"ACTIVITY_IDLE_RARE"`
	Run                int `json:"ACTIVITY_RUN"`
	Attack             int `json:"ACTIVITY_ATTACK"`
	Attack2            int `json:"ACTIVITY_ATTACK2"`
	AttackEvent        int `json:"ACTIVITY_ATTACK_EVENT"`
	Die                int `json:"ACTIVITY_DIE"`
	Flinch             int `json:"ACTIVITY_FLINCH"`
	Flail              int `json:"ACTIVITY_FLAIL"`
	Disabled           int `json:"ACTIVITY_DISABLED"`
	CastAbility1       int `json:"ACTIVITY_CAST_ABILITY_1"`
	CastAbility2       int `json:"ACTIVITY_CAST_ABILITY_2"`
	CastAbility3       int `json:"ACTIVITY_CAST_ABILITY_3"`
	CastAbility4       int `json:"ACTIVITY_CAST_ABILITY_4"`
	CastAbility5       int `json:"ACTIVITY_CAST_ABILITY_5"`
	CastAbility6       int `json:"ACTIVITY_CAST_ABILITY_6"`
	OverrideAbility1   int `json:"ACTIVITY_OVERRIDE_ABILITY_1"`
	OverrideAbility2   int `json:"ACTIVITY_OVERRIDE_ABILITY_2"`
	OverrideAbility3   int `json:"ACTIVITY_OVERRIDE_ABILITY_3"`
	OverrideAbility4   int `json:"ACTIVITY_OVERRIDE_ABILITY_4"`
	ChannelAbility1    int `json:"ACTIVITY_CHANNEL_ABILITY_1"`
	ChannelAbility2    int `json:"ACTIVITY_CHANNEL_ABILITY_2"`
	ChannelAbility3    int `json:"ACTIVITY_CHANNEL_ABILITY_3"`
	ChannelAbility4    int `json:"ACTIVITY_CHANNEL_ABILITY_4"`
	ChannelAbility5    int `json:"ACTIVITY_CHANNEL_ABILITY_5"`
	ChannelAbility6    int `json:"ACTIVITY_CHANNEL_ABILITY_6"`
	ChannelEndAbility1 int `json:"ACTIVITY_CHANNEL_END_ABILITY_1"`
	ChannelEndAbility2 int `json:"ACTIVITY_CHANNEL_END_ABILITY_2"`
	ChannelEndAbility3 int `json:"ACTIVITY_CHANNEL_END_ABILITY_3"`
	ChannelEndAbility4 int `json:"ACTIVITY_CHANNEL_END_ABILITY_4"`
	ChannelEndAbility5 int `json:"ACTIVITY_CHANNEL_END_ABILITY_5"`
	ChannelEndAbility6 int `json:"ACTIVITY_CHANNEL_END_ABILITY_6"`
	ConstantLayer      int `json:"ACTIVITY_CONSTANT_LAYER"`
	Capture            int `json:"ACTIVITY_CAPTURE"`
	Spawn              int `json:"ACTIVITY_SPAWN"`
	Killtaunt          int `json:"ACTIVITY_KILLTAUNT"`
	Taunt              int `json:"ACTIVITY_TAUNT"`
}

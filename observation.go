package dota2bot

type WorldInfo struct {
	GameTime       float64       `json:"gameTime"`
	DotaTime       float64       `json:"dotaTime"`
	TimeOfDay      float64       `json:"timeOfDay"`
	RoshanKillTime float64       `json:"roshanKillTime"`
	Units          []Unit        `json:"units"`
	Runes          []Rune        `json:"runes"`
	DroppedItems   []DroppedItem `json:"droppedItems"`
	LaneFronts     []LaneFront   `json:"laneFronts"`
	GlyphCooldown  float64       `json:"glyphCooldown"`
}

type Rune struct {
	RuneSpawn     int     `json:"runeSpawn"`
	Type          int     `json:"type"`
	Status        int     `json:"status"`
	TimeSinceSeen float64 `json:"timeSinceSeen"`
}

type DroppedItem struct {
	HandleID int    `json:"handleID"`
	Name     string `json:"name"`
	Location `json:"location"`
}

type LaneFront struct {
	Lane     int `json:"lane"`
	Location `json:"front"`
}

type Ability struct {
	Name                          string          `json:"name"`
	Slot                          int             `json:"slot"`
	CanAbilityBeUpgraded          bool            `json:"canAbilityBeUpgraded"`
	GetAutoCastState              bool            `json:"getAutoCastState"`
	GetToggleState                bool            `json:"getToggleState"`
	IsToggle                      bool            `json:"isToggle"`
	IsActivated                   bool            `json:"isActivated"`
	IsChanneling                  bool            `json:"isChanneling"`
	IsCooldownReady               bool            `json:"isCooldownReady"`
	IsFullyCastable               bool            `json:"isFullyCastable"`
	IsHidden                      bool            `json:"isHidden"`
	IsInAbilityPhase              bool            `json:"isInAbilityPhase"`
	IsOwnersManaEnough            bool            `json:"isOwnersManaEnough"`
	IsPassive                     bool            `json:"isPassive"`
	IsStealable                   bool            `json:"isStealable"`
	IsStolen                      bool            `json:"isStolen"`
	IsTrained                     bool            `json:"isTrained"`
	ProcsMagicStick               bool            `json:"procsMagicStick"`
	Behavior                      AbilityBehavior `json:"behavior"`
	GetCastPoint                  float64         `json:"getCastPoint"`
	GetCastRange                  float64         `json:"getCastRange"`
	GetChanneledManaCostPerSecond int             `json:"getChanneledManaCostPerSecond"`
	GetChannelTime                float64         `json:"getChannelTime"`
	GetDuration                   float64         `json:"getDuration"`
	GetCooldownTimeRemaining      float64         `json:"getCooldownTimeRemaining"`
	GetCurrentCharges             int             `json:"getCurrentCharges"`
	GetAbilityDamage              float64         `json:"getAbilityDamage"`
	GetDamageType                 int             `json:"getDamageType"`
	GetHeroLevelRequiredToUpgrade int             `json:"getHeroLevelRequiredToUpgrade"`
	GetInitialCharges             int             `json:"getInitialCharges"`
	GetLevel                      int             `json:"getLevel"`
	GetManaCost                   int             `json:"getManaCost"`
	GetMaxLevel                   int             `json:"getMaxLevel"`
	GetSecondaryCharges           int             `json:"getSecondaryCharges"`
	GetTargetTeam                 int             `json:"getTargetTeam"`
	GetTargetType                 int             `json:"getTargetType"`
	IsItem                        bool            `json:"isItem"`
	CanBeDisassembled             bool            `json:"canBeDisassembled"`
	IsCombineLocked               bool            `json:"isCombineLocked"`
}

const (
	AbilityBehaviorHidden       = 0x0000000001 // This ability can be owned by a unit but can't be casted and wont show up on the HUD.
	AbilityBehaviorPassive      = 0x0000000002 // Can't be casted like above but this one shows up on the ability HUD
	AbilityBehaviorNoTarget     = 0x0000000100 // Doesn't need a target to be cast, ability fires off as soon as the button is pressed
	AbilityBehaviorUnitTarget   = 0x0000001000 // Ability needs a target to be casted on.
	AbilityBehaviorPoint        = 0x0000010000 // Ability can be cast anywhere the mouse cursor is (If a unit is clicked it will just be cast where the unit was standing)
	AbilityBehaviorAoe          = 0x0000100000 // This ability draws a radius where the ability will have effect. YOU STILL NEED A TARGETTING BEHAVIOR LIKE ABILITY_BEHAVIOR_POINT FOR THIS TO WORK.
	AbilityBehaviorNotLearnable = 0x0001000000 // This ability probably can be casted or have a casting scheme but cannot be learned (these are usually abilities that are temporary like techie's bomb detonate)
	AbilityBehaviorChannelled   = 0x0010000000 // This abillity is channelled. If the user moves or is silenced the ability is interrupted.
	AbilityBehaviorItem         = 0x0100000000 // This ability is tied up to an item.
	AbilityBehaviorToggle       = 0x1000000000 // This ability can be toggled on/off
)

type AbilityBehavior int

// Passive ability
func (a AbilityBehavior) Passive() bool {
	return a&AbilityBehaviorPassive > 0
}

// NoTarget for ability
func (a AbilityBehavior) NoTarget() bool {
	return a&AbilityBehaviorNoTarget > 0
}

// UnitTarget ability
func (a AbilityBehavior) UnitTarget() bool {
	return a&AbilityBehaviorUnitTarget > 0
}

// Point target ability
func (a AbilityBehavior) Point() bool {
	return a&AbilityBehaviorPoint > 0
}

// AOE ability
func (a AbilityBehavior) AOE() bool {
	return a&AbilityBehaviorAoe > 0
}

// NotLearnable ability
func (a AbilityBehavior) NotLearnable() bool {
	return a&AbilityBehaviorNotLearnable > 0
}

// Channelled ability
func (a AbilityBehavior) Channelled() bool {
	return a&AbilityBehaviorChannelled > 0
}

type Unit struct {
	UUID                      string    `json:"handleID"`
	IsBot                     bool      `json:"isBot"`
	ActiveMode                int       `json:"activeMode"`
	ActiveModeDesire          float64   `json:"activeModeDesire"`
	Difficulty                int       `json:"difficulty"`
	Name                      string    `json:"name"`
	PlayerID                  int       `json:"playerID"`
	Team                      int       `json:"team"`
	IsHero                    bool      `json:"isHero"`
	IsIllusion                bool      `json:"isIllusion"`
	IsCreep                   bool      `json:"isCreep"`
	IsAncientCreep            bool      `json:"isAncientCreep"`
	IsBuilding                bool      `json:"isBuilding"`
	IsTower                   bool      `json:"isTower"`
	IsFort                    bool      `json:"isFort"`
	CanBeSeen                 bool      `json:"canBeSeen"`
	Health                    int       `json:"health"`
	MaxHealth                 int       `json:"maxHealth"`
	HealthRegen               float64   `json:"healthRegen"`
	Mana                      int       `json:"mana"`
	MaxMana                   int       `json:"maxMana"`
	ManaRegen                 float64   `json:"manaRegen"`
	BaseMovementSpeed         int       `json:"baseMovementSpeed"`
	CurrentMovementSpeed      int       `json:"currentMovementSpeed"`
	IsAlive                   bool      `json:"isAlive"`
	RespawnTime               float64   `json:"respawnTime"`
	HasBuyback                bool      `json:"hasBuyback"`
	BuybackCost               int       `json:"buybackCost"`
	BuybackCooldown           float64   `json:"buybackCooldown"`
	RemainingLifespan         float64   `json:"remainingLifespan"`
	BaseDamage                float64   `json:"baseDamage"`
	BaseDamageVariance        float64   `json:"baseDamageVariance"`
	AttackDamage              float64   `json:"attackDamage"`
	AttackRange               int       `json:"attackRange"`
	AttackSpeed               float64   `json:"attackSpeed"`
	SecondsPerAttack          float64   `json:"secondsPerAttack"`
	AttackPoint               float64   `json:"attackPoint"`
	LastAttackTime            float64   `json:"lastAttackTime"`
	AttackTarget              string    `json:"attackTarget"` // UUID of target being attacked
	AcquisitionRange          int       `json:"acquisitionRange"`
	AttackProjectileSpeed     int       `json:"attackProjectileSpeed"`
	SpellAmp                  float64   `json:"spellAmp"`
	Armor                     float64   `json:"armor"`
	MagicResist               float64   `json:"magicResist"`
	Evasion                   float64   `json:"evasion"`
	PrimaryAttribute          int       `json:"primaryAttribute"`
	BountyXP                  int       `json:"bountyXP"`
	BountyGoldMin             int       `json:"bountyGoldMin"`
	BountyGoldMax             int       `json:"bountyGoldMax"`
	XPNeededToLevel           int       `json:"XPNeededToLevel"`
	AbilityPoints             int       `json:"abilityPoints"`
	Level                     int       `json:"level"`
	Gold                      int       `json:"gold"`
	NetWorth                  int       `json:"netWorth"`
	StashValue                int       `json:"stashValue"`
	CourierValue              int       `json:"courierValue"`
	LastHits                  int       `json:"lastHits"`
	Denies                    int       `json:"denies"`
	BoundingRadius            float64   `json:"boundingRadius"`
	Location                  Location  `json:"location"`
	Facing                    int       `json:"facing"`
	Velocity                  Location  `json:"velocity"`
	DayTimeVisionRange        int       `json:"dayTimeVisionRange"`
	NightTimeVisionRange      int       `json:"nightTimeVisionRange"`
	CurrentVisionRange        int       `json:"currentVisionRange"`
	HealthRegenPerStr         float64   `json:"healthRegenPerStr"`
	ManaRegenPerInt           float64   `json:"manaRegenPerInt"`
	AnimationActivity         int       `json:"animationActivity"`
	AnimationCycle            float64   `json:"animationCycle"`
	IsChanneling              bool      `json:"isChanneling"`
	IsUsingAbility            bool      `json:"isUsingAbility"`
	IsCastingAbility          bool      `json:"isCastingAbility"`
	IsAttackImmune            bool      `json:"isAttackImmune"`
	IsBlind                   bool      `json:"isBlind"`
	IsBlockDisabled           bool      `json:"isBlockDisabled"`
	IsDisarmed                bool      `json:"isDisarmed"`
	IsDominated               bool      `json:"isDominated"`
	IsEvadeDisabled           bool      `json:"isEvadeDisabled"`
	IsHexed                   bool      `json:"isHexed"`
	IsInvisible               bool      `json:"isInvisible"`
	IsInvulnerable            bool      `json:"isInvulnerable"`
	IsMagicImmune             bool      `json:"isMagicImmune"`
	IsMuted                   bool      `json:"isMuted"`
	IsNightmared              bool      `json:"isNightmared"`
	IsRooted                  bool      `json:"isRooted"`
	IsSilenced                bool      `json:"isSilenced"`
	IsSpeciallyDeniable       bool      `json:"isSpeciallyDeniable"`
	IsStunned                 bool      `json:"isStunned"`
	IsUnableToMiss            bool      `json:"isUnableToMiss"`
	HasScepter                bool      `json:"hasScepter"`
	TimeSinceDamagedByAnyHero float64   `json:"timeSinceDamagedByAnyHero"`
	TimeSinceDamagedByCreep   float64   `json:"timeSinceDamagedByCreep"`
	TimeSinceDamagedByTower   float64   `json:"timeSinceDamagedByTower"`
	CurrentActionType         int       `json:"currentActionType"`
	CourierState              int       `json:"courierState"`
	IsFlyingCourier           bool      `json:"isFlyingCourier"`
	Abilities                 []Ability `json:"abilities"`
	Items                     []Ability `json:"item"`
}

package augments

const (
	LegendaryRarityPercent = 0.10
	EpicRarityPercent      = 0.30
	CommonRarityPercent    = 1.
	NegativeRarityPercent  = 1. / 5.
)

type ID byte

const (
	// Common
	IDDebugLines ID = iota
	IDFreeRoll
	// Epic
	IDNegateNextCost
	IDHighSpawn
	IDHeartSpawn
	IDGoldHeartSpawn
	IDSlowMotion
	IDHeartContainer
	// Legendary
	IDNegativeHearts
	IDCircular
	IDPerfectStep
	IDRemoveLastNegative
	// Negative
	IDMoreBlocks
	IDTallerBlocks
	IDMoreSpawns
	IDEvenMoreSpawns
	IDCloserSpawns
	IDCloserSpawns2
	IDNothing
	IDNothing2
	IDNothing3
	IDNothing4
	IDRemoveLastPositive
	IDHarderBlocks
	IDHarderBlocks2
	IDNoRegularBlocks
	// TODO: JumpLocked
	// TODO: MovementLocked

	IDMax
)

type CostKind byte

const (
	CostNone CostKind = iota
	CostHP
	CostHalfScore
)

type Cost struct {
	Kind  CostKind
	Value int
}

type Rarity byte

const (
	RarityLegendary Rarity = iota
	RarityEpic
	RarityCommon
	RarityNegative
)

type Augment struct {
	cost Cost

	ID          ID
	Symbol      string
	Name        string
	Description string
	Stackable   bool
	Rarity      Rarity
	Constraints []ID
}

func (a *Augment) GetCost() Cost {
	return a.cost
}

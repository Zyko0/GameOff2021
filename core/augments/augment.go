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
	IDIncreaseSpeed ID = iota
	IDDecreaseSpeed
	IDDebugLines
	IDActionJump
	IDActionDash
	IDFreeRoll
	// Epic
	IDNegateNextCost
	IDHighSpawn
	IDHeartSpawn
	IDSlowMotion
	IDHeartContainer
	// Legendary
	IDNegativeHearts
	IDCircular
	IDPerfectStep
	IDRemoveLastNegative
	// Negative
	IDOneMoreBlock
	IDTallerBlocks
	IDTopView
	IDMoreSpawns
	IDEvenMoreSpawns
	IDCloserSpawns
	IDCloserSpawns2
	IDNothing
	IDNothing2
	IDRemoveLastPositive
	IDLessAugments
	IDHarderBlocks
	IDHarderBlocks2
	IDNoRegularBlocks
	IDFourTimesFaster
	// TODO: Drunk ? Offseted block positions ?

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

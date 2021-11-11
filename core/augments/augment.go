package augments

const (
	LegendaryRarityPercent = 0.10
	EpicRarityPercent      = 0.20
	RareRarityPercent      = 0.30
	CommonRarityPercent    = 0.40
	NegativeRarityPercent  = 1. / 5.
)

type ID byte

const (
	// Common
	IDIncreaseSpeed ID = iota
	IDDebugLines
	// Rare
	IDActionJump
	// Epic
	IDHighSpawn
	IDHeartSpawn
	IDSlowMotion
	IDHeartContainer
	// Legendary
	IDNegativeHearts
	IDCircular
	IDCollisionCheck
	IDRemoveLastNegative
	// Negative
	IDOneMoreBlock
	IDTallerBlocks
	IDTopView
	IDMoreSpawns
	IDEvenMoreSpawns
	IDCloserSpawns
	IDEvenCloserSpawns
	IDNothing
	IDNothing2
	IDRemoveLastPositive
	IDLessAugments
	IDStrongerBlocks
	IDEvenStrongerBlocks
	// TODO: Drunk ? Offseted block positions ?

	IDMax
)

type CostKind byte

const (
	CostNone CostKind = iota
	CostKindHP
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
	RarityRare
	RarityCommon
	RarityNegative
)

type Augment struct {
	ID          ID
	Name        string
	Description string
	Stackable   bool
	Rarity      Rarity
	Cost        Cost
	Constraints []ID
}

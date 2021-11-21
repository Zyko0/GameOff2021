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
	IDFreeRoll // TODO: Useless
	// Epic
	IDNegateNextCost // TODO: Useless, almost no cost on augments anymore
	IDHighSpawn // TODO: Not sure this is even a bonus or not
	IDHeartSpawn
	IDGoldHeartSpawn
	IDSlowMotion // ?
	IDHeartContainer
	// Legendary
	IDNegativeHearts
	IDCircular // TODO: Not sure it's even a bonus
	IDPerfectStep
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
	IDHarderBlocks
	IDHarderBlocks2
	IDNoRegularBlocks
	// TODO: JumpFix "blabla not the kind of fix you expect" => jump speed depends on game's speed
	// TODO: charging beam on Z axis ? One shots ? 1hp ?
	// TODO: Hole rectangle block on Y=0 - width=1.

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

package augments

const (
	LegendaryRarityPercent = 0.05
	EpicRarityPercent      = 0.10
	RareRarityPercent      = 0.30
	UncommonRarityPercent  = 0.50
	NegativeRarityPercent  = 1. / 5.
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
	Name        string
	Description string
	Stackable   bool
	Rarity      Rarity
	Cost        Cost
}

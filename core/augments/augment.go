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
	// Epic
	IDHighSpawn      // TODO: Not sure this is even a bonus or not
	IDHeartSpawn
	IDGoldHeartSpawn
	IDHeartContainer
	// Legendary
	IDPerfectStep
	// Negative
	IDMoreBlocks
	IDTallerBlocks
	IDMoreSpawns
	IDEvenMoreSpawns
	IDCloserSpawns
	IDCloserSpawns2
	IDHarderBlocks
	IDHarderBlocks2
	IDLateralHoles
	IDLongHoles
	IDChargingBeam
	IDNoRegularBlocks
	IDCircular // TODO: I don't think this deserves to be a bonus anymore
	// TODO: JumpFix "blabla not the kind of fix you expect" => jump speed depends on game's speed
	// TODO: Charging beam on Z axis ? One shots ? 1hp ?
	// TODO: Hole rectangle block on Y=0 - width=1 => Substracts matter to the road
	// TODO: Hole rectangle block (same) but not lateral, deep instead ?

	IDMax
)

type Rarity byte

const (
	RarityLegendary Rarity = iota
	RarityEpic
	RarityCommon
	RarityNegative
)

type Augment struct {
	ID          ID
	Name        string
	Description string
	Stackable   bool
	Rarity      Rarity
	Constraints []ID
}

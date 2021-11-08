package augments

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
	IDCancelLastNegative
	IDSlowMotion
	// Legendary
	IDNegativeHearts
	IDCircular
	// Negative
	IDOneMoreBlocks
	IDTallerBlocks
	IDTopView
	IDMoreSpawns
	IDHalfwaySpawns
	IDNothing
	IDCancelLastPositive
	IDAugmentLessAugments
	// TODO: Drunk ? Offseted block positions ?
)

var (
	// Common
	AugmentIncreaseSpeed = &Augment{
		Name:        "Speed",
		Description: "Increases the speed of the unit by 10%.",
		Stackable:   true,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentDebugLines = &Augment{
		Name:        "Developer Mode",
		Description: "Traces lines between different blocks, disabled in production of course.",
		Stackable:   false,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

var (
	// Rare
	AugmentActionJump = &Augment{
		Name:        "Jump",
		Description: "It seems your space button now lets you jump. Sorry for having missed this core feature from the release.",
		Stackable:   false,
		Rarity:      RarityRare,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

var (
	// Epic
	AugmentHighSpawn = &Augment{
		Name:        "Weird gravity",
		Description: "How do these blocks not fall btw ?",
		Stackable:   false,
		Rarity:      RarityEpic,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHeartSpawn = &Augment{
		Name:        "Heart containers",
		Description: "This doesn't really look like a bug, but more like an omitted feature.",
		Stackable:   false,
		Rarity:      RarityEpic,
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 1,
		},
	}
	AugmentCancelLastNegative = &Augment{
		Name:        "Bug fix",
		Description: "Okay the last negative bug you encountered is now fixed.",
		Stackable:   true,
		Rarity:      RarityEpic,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentSlowMotion = &Augment{
		Name:        "Lag",
		Description: "Every N seconds, you will experience a 2 second lag.",
		Stackable:   false,
		Rarity:      RarityEpic,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

var (
	// Legendary
	AugmentNegativeHearts = &Augment{
		Name:        "Negative Hearts",
		Description: "Game is over at -3 hp, why though...",
		Stackable:   false,
		Rarity:      RarityLegendary,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentCircular = &Augment{
		Name:        "Circular",
		Description: "Can someone explain why there is no horizontal boundary anymore ?",
		Stackable:   false,
		Rarity:      RarityLegendary,
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 1,
		},
	}
)

var (
	// Negative
	AugmentOneMoreBlocks = &Augment{
		Name:        "More blocks",
		Description: "Wait, this game was designed with 3 blocks per spawn at maximum...",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentTallerBlocks = &Augment{
		Name:        "Taller blocks",
		Description: "Some blocks are taller than the other, how is this supposed to make it harder without a jump ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentTopView = &Augment{
		Name:        "Top view",
		Description: "The camera is now positionned on top, this is usefull for debugging purposes.",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentMoreSpawns = &Augment{
		Name:        "More spawns",
		Description: "Twice the amount of rows spawning... Who let that happen ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHalfwaySpawns = &Augment{
		Name:        "Halfway spawns",
		Description: "So blocks spawn closer now, how is the player supposed to react properly ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentNothing = &Augment{
		Name:        "Nothing",
		Description: "",
		Stackable:   true,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 1,
		},
	}
	AugmentCancelLastPositive = &Augment{
		Name:        "Broken feature",
		Description: "Sorry about that, it might break the last abusive bug, but hey it's a fix !",
		Stackable:   true,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentLessAugments = &Augment{
		Name:        "Less bugs",
		Description: "We are getting closer to a clean build, bugs will show less often, people don't like bugs, right ?",
		Stackable:   true,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

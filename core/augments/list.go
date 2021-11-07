package augments

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
	AugmentOneMoreBlock = &Augment{
		Name:        "More blocks",
		Description: "Wait, this game was designed with 3 blocks per spawn at maximum...",
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
)

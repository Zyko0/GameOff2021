package augments

var (
	List = make([]*Augment, IDMax)
)

func init() {
	List[IDDebugLines] = AugmentDebugLines
	List[IDFreeRoll] = AugmentFreeRoll

	List[IDNegateNextCost] = AugmentNegateNextCost
	List[IDHighSpawn] = AugmentHighSpawn
	List[IDHeartSpawn] = AugmentHeartSpawn
	List[IDGoldHeartSpawn] = AugmentGoldHeartSpawn
	List[IDSlowMotion] = AugmentSlowMotion
	List[IDHeartContainer] = AugmentHeartContainer

	List[IDNegativeHearts] = AugmentNegativeHearts
	List[IDCircular] = AugmentCircular
	List[IDPerfectStep] = AugmentPerfectStep

	List[IDMoreBlocks] = AugmentMoreBlocks
	List[IDTallerBlocks] = AugmentTallerBlocks
	List[IDMoreSpawns] = AugmentMoreSpawns
	List[IDEvenMoreSpawns] = AugmentEvenMoreSpawns
	List[IDCloserSpawns] = AugmentCloserSpawns
	List[IDCloserSpawns2] = AugmentCloserSpawns2
	List[IDNothing] = AugmentNothing
	List[IDNothing2] = AugmentNothing2
	List[IDNothing3] = AugmentNothing3
	List[IDNothing4] = AugmentNothing4
	List[IDHarderBlocks] = AugmentHarderBlocks
	List[IDHarderBlocks2] = AugmentHarderBlocks2
	List[IDNoRegularBlocks] = AugmentNoRegularBlocks

	// Reprocess description texts for them to fit in card caption
	for _, a := range List {
		var desc string

		for i, r := range a.Description {
			if i > 0 && i%18 == 0 {
				if r == ' ' {
					desc += "\n"
				} else if a.Description[i-1] == ' ' {
					desc += "\n"
					desc += string(r)
				} else {
					desc += "-"
					desc += "\n"
					desc += string(r)
				}
			} else {
				desc += string(r)
			}
		}
		a.Description = desc
	}
}

var (
	// Common
	AugmentDebugLines = &Augment{
		ID:          IDDebugLines,
		Name:        "Dev Mode",
		Description: "Traces lines between different blocks, disabled in production of course.",
		Stackable:   false,
		Rarity:      RarityCommon,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentFreeRoll = &Augment{
		ID:          IDFreeRoll,
		Name:        "Infinite Loop",
		Description: "Triggers a new roll.",
		Stackable:   true,
		Rarity:      RarityCommon,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

var (
	// Epic
	AugmentNegateNextCost = &Augment{
		ID:          IDNegateNextCost,
		Name:        "Free cost",
		Description: "The next exploited bug costs you nothing.",
		Stackable:   true,
		Rarity:      RarityEpic,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHighSpawn = &Augment{
		ID:          IDHighSpawn,
		Name:        "Weird gravity",
		Description: "How do these blocks not fall btw ?",
		Stackable:   false,
		Rarity:      RarityEpic,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHeartSpawn = &Augment{
		ID:          IDHeartSpawn,
		Name:        "Heart blocks",
		Description: "This doesn't really look like a bug, but more like an omitted feature.",
		Stackable:   false,
		Rarity:      RarityEpic,
		cost: Cost{
			Kind:  CostHP,
			Value: 1,
		},
	}
	AugmentGoldHeartSpawn = &Augment{
		ID:          IDGoldHeartSpawn,
		Name:        "Golden hearts",
		Description: "These hearts give you 2HP back.",
		Stackable:   false,
		Rarity:      RarityEpic,
		cost: Cost{
			Kind:  CostHP,
			Value: 0,
		},
		Constraints: []ID{
			IDHeartSpawn,
		},
	}
	AugmentSlowMotion = &Augment{
		ID:          IDSlowMotion,
		Name:        "Lag",
		Description: "Every N seconds, you will experience a 2 second lag.",
		Stackable:   false,
		Rarity:      RarityEpic,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHeartContainer = &Augment{
		ID:          IDHeartContainer,
		Name:        "Heart container",
		Description: "This is an additional heart container, devs are bad at their own game so we need this option for a moment.",
		Stackable:   true,
		Rarity:      RarityEpic,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDHeartSpawn,
		},
	}
)

var (
	// Legendary
	AugmentNegativeHearts = &Augment{
		ID:          IDNegativeHearts,
		Name:        "Negative Hearts",
		Description: "Game is over at -3 hp, why though...",
		Stackable:   false,
		Rarity:      RarityLegendary,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentCircular = &Augment{
		ID:          IDCircular,
		Name:        "Circular",
		Description: "Can someone explain why there is no horizontal boundary anymore ?",
		Stackable:   false,
		Rarity:      RarityLegendary,
		cost: Cost{
			Kind:  CostHP,
			Value: 0,
		},
	}
	AugmentPerfectStep = &Augment{
		ID:          IDPerfectStep,
		Name:        "Perfect Step",
		Description: "The sphere now steps exactly from a row to another, we needed an easy mode for devs.",
		Stackable:   false,
		Rarity:      RarityLegendary,
		cost: Cost{
			Kind:  CostHP,
			Value: 0, // TODO: Maybe I add this but make it cost a lot of hp ?
		},
		Constraints: []ID{
			IDDebugLines,
		},
	}
	// TODO: PerfectStepY ? For jump and Y axis
)

var (
	// Negative
	AugmentMoreBlocks = &Augment{
		ID:          IDMoreBlocks,
		Name:        "More blocks",
		Description: "Wait, this game was designed with 3 blocks per spawn at maximum...",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentTallerBlocks = &Augment{
		ID:          IDTallerBlocks,
		Name:        "Taller blocks",
		Description: "Some blocks are taller than the other, how is this supposed to make it harder without a jump ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentMoreSpawns = &Augment{
		ID:          IDMoreSpawns,
		Name:        "More spawns",
		Description: "Twice the amount of rows spawning... Who let that happen ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentEvenMoreSpawns = &Augment{
		ID:          IDEvenMoreSpawns,
		Name:        "Even more spawns",
		Description: "Three times the amount of rows spawning, this is for testing.",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDEvenMoreSpawns,
		},
	}
	AugmentCloserSpawns = &Augment{
		ID:          IDCloserSpawns,
		Name:        "Closer spawns",
		Description: "So blocks spawn closer now, how is the player supposed to react properly ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentCloserSpawns2 = &Augment{
		ID:          IDCloserSpawns2,
		Name:        "Closer Spawns II",
		Description: "So blocks spawn closer now, how is the player supposed to react properly ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDCloserSpawns,
		},
	}
	AugmentNothing = &Augment{
		ID:          IDNothing,
		Name:        "Nothing",
		Description: "",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostHP,
			Value: 1,
		},
	}
	AugmentNothing2 = &Augment{
		ID:          IDNothing2,
		Name:        "Nothing II",
		Description: "",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostHP,
			Value: 2,
		},
		Constraints: []ID{
			IDNothing,
		},
	}
	AugmentNothing3 = &Augment{
		ID:          IDNothing3,
		Name:        "Nothing III",
		Description: "",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostHP,
			Value: 3,
		},
		Constraints: []ID{
			IDNothing2,
		},
	}
	AugmentNothing4 = &Augment{
		ID:          IDNothing4,
		Name:        "Nothing IV",
		Description: "",
		Stackable:   true,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostHP,
			Value: 4,
		},
		Constraints: []ID{
			IDNothing3,
		},
	}
	AugmentHarderBlocks = &Augment{
		ID:          IDHarderBlocks,
		Name:        "Harder Blocks",
		Description: "Some blocks deal more damage, you should recognize them.",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHarderBlocks2 = &Augment{
		ID:          IDHarderBlocks2,
		Name:        "Harder Blocks II",
		Description: "Some blocks deal even more damage, you should also recognize them.",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
	AugmentNoRegularBlocks = &Augment{
		ID:          IDHarderBlocks2,
		Name:        "No Regular Blocks",
		Description: "The block you used to know doesn't exist anymore, but then it must be replaced by something, hmm...",
		Stackable:   false,
		Rarity:      RarityNegative,
		cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
)

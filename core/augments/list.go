package augments

var (
	List = make([]*Augment, IDMax)
)

func init() {
	List[IDIncreaseSpeed] = AugmentIncreaseSpeed
	List[IDDecreaseSpeed] = AugmentDecreaseSpeed
	List[IDDebugLines] = AugmentDebugLines
	List[IDActionJump] = AugmentActionJump
	List[IDActionDash] = AugmentActionDash
	List[IDHighSpawn] = AugmentHighSpawn
	List[IDHeartSpawn] = AugmentHeartSpawn
	List[IDSlowMotion] = AugmentSlowMotion
	List[IDHeartContainer] = AugmentHeartContainer
	List[IDNegativeHearts] = AugmentNegativeHearts
	List[IDCircular] = AugmentCircular
	List[IDPerfectStep] = AugmentPerfectStep
	List[IDRemoveLastNegative] = AugmentRemoveLastNegative
	List[IDOneMoreBlock] = AugmentOneMoreBlock
	List[IDTallerBlocks] = AugmentTallerBlocks
	List[IDTopView] = AugmentTopView
	List[IDMoreSpawns] = AugmentMoreSpawns
	List[IDEvenMoreSpawns] = AugmentEvenMoreSpawns
	List[IDCloserSpawns] = AugmentCloserSpawns
	List[IDCloserSpawns2] = AugmentCloserSpawns2
	List[IDNothing] = AugmentNothing
	List[IDNothing2] = AugmentNothing2
	List[IDRemoveLastPositive] = AugmentRemoveLastPositive
	List[IDLessAugments] = AugmentLessAugments
	List[IDHarderBlocks] = AugmentHarderBlocks
	List[IDHarderBlocks2] = AugmentHarderBlocks2

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
	AugmentIncreaseSpeed = &Augment{
		ID:          IDIncreaseSpeed,
		Name:        "Speed hack",
		Description: "Increases the lateral speed of the sphere by 10%.",
		Stackable:   true,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentDecreaseSpeed = &Augment{
		ID:          IDDecreaseSpeed,
		Name:        "Speed fix",
		Description: "Decreases the lateral speed of the sphere by 10%.",
		Stackable:   true,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentDebugLines = &Augment{
		ID:          IDDebugLines,
		Name:        "Dev Mode",
		Description: "Traces lines between different blocks, disabled in production of course.",
		Stackable:   false,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentActionJump = &Augment{
		ID:          IDActionJump,
		Name:        "Jump",
		Description: "It seems your space button now lets you jump. Sorry for having missed this core feature from the release.",
		Stackable:   false,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentActionDash = &Augment{
		ID:          IDActionDash,
		Name:        "Dash",
		Description: "Your space button now lets you dash in your last inputted direction.",
		Stackable:   false,
		Rarity:      RarityCommon,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

var (
	// Epic
	AugmentHighSpawn = &Augment{
		ID:          IDHighSpawn,
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
		ID:          IDHeartSpawn,
		Name:        "Heart blocks",
		Description: "This doesn't really look like a bug, but more like an omitted feature.",
		Stackable:   false,
		Rarity:      RarityEpic,
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 1,
		},
	}
	AugmentSlowMotion = &Augment{
		ID:          IDSlowMotion,
		Name:        "Lag",
		Description: "Every N seconds, you will experience a 2 second lag.",
		Stackable:   false,
		Rarity:      RarityEpic,
		Cost: Cost{
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
		Cost: Cost{
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
		Cost: Cost{
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
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 1,
		},
	}
	AugmentPerfectStep = &Augment{
		ID:          IDPerfectStep,
		Name:        "Perfect Step",
		Description: "The sphere now steps exactly from a row to another, broken TODO: really broken",
		Stackable:   false,
		Rarity:      RarityLegendary,
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 3, // TODO: Maybe I add this but make it cost a lot of hp ?
		},
	}
	AugmentRemoveLastNegative = &Augment{
		ID:          IDRemoveLastNegative,
		Name:        "Bug fix",
		Description: "Okay the last negative bug you encountered is now fixed.",
		Stackable:   true,
		Rarity:      RarityEpic,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
)

var (
	// Negative
	AugmentOneMoreBlock = &Augment{
		ID:          IDOneMoreBlock,
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
		ID:          IDTallerBlocks,
		Name:        "Taller blocks",
		Description: "Some blocks are taller than the other, how is this supposed to make it harder without a jump ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDActionJump,
		},
	}
	AugmentTopView = &Augment{
		ID:          IDTopView,
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
		ID:          IDMoreSpawns,
		Name:        "More spawns",
		Description: "Twice the amount of rows spawning... Who let that happen ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
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
		Cost: Cost{
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
		Cost: Cost{
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
		Cost: Cost{
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
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 1,
		},
	}
	AugmentNothing2 = &Augment{
		ID:          IDNothing2,
		Name:        "Nothing II",
		Description: "",
		Stackable:   true,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostKindHP,
			Value: 2,
		},
		Constraints: []ID{
			IDNothing,
		},
	}
	AugmentRemoveLastPositive = &Augment{
		ID:          IDRemoveLastPositive,
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
		ID:          IDLessAugments,
		Name:        "Less bugs",
		Description: "We are getting closer to a clean build, bugs will show less often, people don't like bugs, right ?",
		Stackable:   true,
		Rarity:      RarityNegative,
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
	}
	AugmentHarderBlocks = &Augment{
		ID:          IDHarderBlocks,
		Name:        "Harder Blocks II",
		Description: "Some blocks deal more damage, you should recognize them.",
		Stackable:   false,
		Rarity:      RarityNegative,
		Cost: Cost{
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
		Cost: Cost{
			Kind:  CostNone,
			Value: 0,
		},
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
)

package augments

var (
	List = make([]*Augment, IDMax)
)

func init() {
	List[IDDebugLines] = AugmentDebugLines

	List[IDHighSpawn] = AugmentHighSpawn
	List[IDHeartSpawn] = AugmentHeartSpawn
	List[IDGoldHeartSpawn] = AugmentGoldHeartSpawn
	List[IDHeartContainer] = AugmentHeartContainer

	List[IDPerfectStep] = AugmentPerfectStep

	List[IDMoreBlocks] = AugmentMoreBlocks
	List[IDTallerBlocks] = AugmentTallerBlocks
	List[IDMoreSpawns] = AugmentMoreSpawns
	List[IDEvenMoreSpawns] = AugmentEvenMoreSpawns
	List[IDCloserSpawns] = AugmentCloserSpawns
	List[IDCloserSpawns2] = AugmentCloserSpawns2
	List[IDHarderBlocks] = AugmentHarderBlocks
	List[IDHarderBlocks2] = AugmentHarderBlocks2
	List[IDLateralHoles] = AugmentLateralHoles
	List[IDLongHoles] = AugmentLongHoles
	List[IDChargingBeam] = AugmentChargingBeam
	List[IDNoRegularBlocks] = AugmentNoRegularBlocks
	List[IDCircular] = AugmentCircular

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
	}
)

var (
	AugmentHighSpawn = &Augment{
		ID:          IDHighSpawn,
		Name:        "Weird gravity",
		Description: "How do these blocks not fall btw ?",
		Stackable:   false,
		Rarity:      RarityEpic,
	}
	AugmentHeartSpawn = &Augment{
		ID:          IDHeartSpawn,
		Name:        "Heart blocks",
		Description: "This doesn't really look like a bug, but more like an omitted feature.",
		Stackable:   false,
		Rarity:      RarityEpic,
	}
	AugmentGoldHeartSpawn = &Augment{
		ID:          IDGoldHeartSpawn,
		Name:        "Golden hearts",
		Description: "These hearts give you 2HP back.",
		Stackable:   false,
		Rarity:      RarityEpic,
		Constraints: []ID{
			IDHeartSpawn,
		},
	}
	AugmentHeartContainer = &Augment{
		ID:          IDHeartContainer,
		Name:        "Heart container",
		Description: "This is an additional heart container, devs are bad at their own game so we need this option for a moment.",
		Stackable:   true,
		Rarity:      RarityEpic,
		Constraints: []ID{
			IDHeartSpawn,
		},
	}
)

var (
	// Legendary
	AugmentPerfectStep = &Augment{
		ID:          IDPerfectStep,
		Name:        "Perfect Step",
		Description: "The sphere now steps exactly from a row to another, we needed an easy mode for devs.",
		Stackable:   false,
		Rarity:      RarityLegendary,
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
	}
	AugmentTallerBlocks = &Augment{
		ID:          IDTallerBlocks,
		Name:        "Taller blocks",
		Description: "Some blocks are taller than the other, how is this supposed to make it harder without a jump ?",
		Stackable:   false,
		Rarity:      RarityNegative,
	}
	AugmentMoreSpawns = &Augment{
		ID:          IDMoreSpawns,
		Name:        "More spawns",
		Description: "Twice the amount of rows spawning... Who let that happen ?",
		Stackable:   false,
		Rarity:      RarityNegative,
	}
	AugmentEvenMoreSpawns = &Augment{
		ID:          IDEvenMoreSpawns,
		Name:        "Even more spawns",
		Description: "Three times the amount of rows spawning, this is for testing.",
		Stackable:   false,
		Rarity:      RarityNegative,
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
	}
	AugmentCloserSpawns2 = &Augment{
		ID:          IDCloserSpawns2,
		Name:        "Closer Spawns II",
		Description: "So blocks spawn closer now, how is the player supposed to react properly ?",
		Stackable:   false,
		Rarity:      RarityNegative,
		Constraints: []ID{
			IDCloserSpawns,
		},
	}
	AugmentHarderBlocks = &Augment{
		ID:          IDHarderBlocks,
		Name:        "Harder Blocks",
		Description: "Some blocks deal more damage, you should recognize them.",
		Stackable:   false,
		Rarity:      RarityNegative,
	}
	AugmentHarderBlocks2 = &Augment{
		ID:          IDHarderBlocks2,
		Name:        "Harder Blocks II",
		Description: "Some blocks deal even more damage, you should also recognize them.",
		Stackable:   false,
		Rarity:      RarityNegative,
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
	AugmentLateralHoles = &Augment{
		ID:          IDLateralHoles,
		Name:        "Lateral Holes",
		Description: "Some holes can spawn on the road now, please jump.",
		Stackable:   false,
		Rarity:      RarityNegative,
	}
	AugmentLongHoles = &Augment{
		ID:          IDLongHoles,
		Name:        "Long Holes",
		Description: "Some holes can spawn on the road now, please dodge.",
		Stackable:   false,
		Rarity:      RarityNegative,
	}
	AugmentChargingBeam = &Augment{
		ID:          IDChargingBeam,
		Name:        "Laser Beam",
		Description: "You might see a long laser charging, be careful about the detonation.",
		Stackable:   false,
		Rarity:      RarityNegative,
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
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
	AugmentCircular = &Augment{
		ID:          IDCircular,
		Name:        "Circular",
		Description: "Can someone explain why there is no horizontal boundary anymore ?",
		Stackable:   false,
		Rarity:      RarityNegative,
	}
)

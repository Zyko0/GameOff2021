package augments

var (
	List = make([]*Augment, IDMax)
)

func init() {
	List[IDHighSpawn] = AugmentHighSpawn
	List[IDMoreBlocks] = AugmentMoreBlocks
	List[IDTallerBlocks] = AugmentTallerBlocks
	List[IDMoreSpawns] = AugmentMoreSpawns
	List[IDEvenMoreSpawns] = AugmentEvenMoreSpawns
	List[IDCloserSpawns] = AugmentCloserSpawns
	List[IDCloserSpawns2] = AugmentCloserSpawns2
	List[IDHarderBlocks] = AugmentHarderBlocks
	List[IDHarderBlocks2] = AugmentHarderBlocks2
	List[IDNoRegularBlocks] = AugmentNoRegularBlocks
	List[IDCircular] = AugmentCircular
	List[IDFallenCamera] = AugmentFallenCamera

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
	AugmentHighSpawn = &Augment{
		ID:          IDHighSpawn,
		Name:        "Weird gravity",
		Description: "How do these blocks not fall btw ?",
		Stackable:   false,
	}
	AugmentMoreBlocks = &Augment{
		ID:          IDMoreBlocks,
		Name:        "More blocks",
		Description: "This game was balanced around 3 blocks per spawn at maximum...",
		Stackable:   false,
	}
	AugmentTallerBlocks = &Augment{
		ID:          IDTallerBlocks,
		Name:        "Taller blocks",
		Description: "Some blocks are taller than others, don't drop the camera !",
		Stackable:   false,
	}
	AugmentMoreSpawns = &Augment{
		ID:          IDMoreSpawns,
		Name:        "More spawns",
		Description: "More rows spawning... Who let that happen ?",
		Stackable:   false,
	}
	AugmentEvenMoreSpawns = &Augment{
		ID:          IDEvenMoreSpawns,
		Name:        "More Spawns II",
		Description: "Way more rows spawning ! TODO(me): remove, this was for testing.",
		Stackable:   false,
		Constraints: []ID{
			IDMoreSpawns,
		},
	}
	AugmentCloserSpawns = &Augment{
		ID:          IDCloserSpawns,
		Name:        "Closer Spawns",
		Description: "So blocks spawn closer now, how is the player supposed to react properly ?",
		Stackable:   false,
	}
	AugmentCloserSpawns2 = &Augment{
		ID:          IDCloserSpawns2,
		Name:        "Closer Spawns II",
		Description: "This looks like an unfair setting, unless you're a robot.",
		Stackable:   false,
		Constraints: []ID{
			IDCloserSpawns,
		},
	}
	AugmentHarderBlocks = &Augment{
		ID:          IDHarderBlocks,
		Name:        "Harder Blocks",
		Description: "Some blocks deal more damage, you should recognize them.",
		Stackable:   false,
	}
	AugmentHarderBlocks2 = &Augment{
		ID:          IDHarderBlocks2,
		Name:        "Harder Blocks II",
		Description: "Some blocks deal even more damage, you should also recognize them.",
		Stackable:   false,
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
	AugmentNoRegularBlocks = &Augment{
		ID:          IDNoRegularBlocks,
		Name:        "No Regular Blocks",
		Description: "The block you used to know doesn't exist anymore, we need a default block now.",
		Stackable:   false,
		Constraints: []ID{
			IDHarderBlocks,
		},
	}
	AugmentCircular = &Augment{
		ID:          IDCircular,
		Name:        "Circular",
		Description: "Can someone explain why there is no horizontal boundary anymore ?",
		Stackable:   false,
	}
	AugmentFallenCamera = &Augment{
		ID:          IDFallenCamera,
		Name:        "Camera fall",
		Description: "Someone dropped the camera, you might have to deal with it...",
		Stackable:   false,
	}
)

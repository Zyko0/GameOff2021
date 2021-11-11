package core

import (
	"math/rand"

	"github.com/Zyko0/GameOff2021/core/augments"
)

type AugmentManager struct {
	currentAugmentsByID []*augments.Augment

	possibleAugmentsCommon    []*augments.Augment
	possibleAugmentsRare      []*augments.Augment
	possibleAugmentsEpic      []*augments.Augment
	possibleAugmentsLegendary []*augments.Augment
	possibleAugmentsNegative  []*augments.Augment

	CurrentAugments []*augments.Augment
}

func NewAugmentManager() *AugmentManager {
	return &AugmentManager{
		currentAugmentsByID: make([]*augments.Augment, augments.IDMax),

		possibleAugmentsCommon:    make([]*augments.Augment, 0, augments.IDMax),
		possibleAugmentsRare:      make([]*augments.Augment, 0, augments.IDMax),
		possibleAugmentsEpic:      make([]*augments.Augment, 0, augments.IDMax),
		possibleAugmentsLegendary: make([]*augments.Augment, 0, augments.IDMax),
		possibleAugmentsNegative:  make([]*augments.Augment, 0, augments.IDMax),

		CurrentAugments: []*augments.Augment{},
	}
}

func (am *AugmentManager) generatePossibleAugments() {
	am.possibleAugmentsCommon = am.possibleAugmentsCommon[:0]
	am.possibleAugmentsRare = am.possibleAugmentsRare[:0]
	am.possibleAugmentsEpic = am.possibleAugmentsEpic[:0]
	am.possibleAugmentsLegendary = am.possibleAugmentsLegendary[:0]
	am.possibleAugmentsNegative = am.possibleAugmentsNegative[:0]

	for _, a := range augments.List {
		// Check constraints satisfaction
		var possible = true
		for _, id := range a.Constraints {
			if am.currentAugmentsByID[id] == nil {
				possible = false
				break
			}
		}
		if !possible {
			continue
		}
		// If already equipped and not stackable, skip it
		if am.currentAugmentsByID[a.ID] != nil && !a.Stackable {
			continue
		}
		// Add and order by rarity
		switch a.Rarity {
		case augments.RarityCommon:
			am.possibleAugmentsCommon = append(am.possibleAugmentsCommon, a)
		case augments.RarityRare:
			am.possibleAugmentsRare = append(am.possibleAugmentsRare, a)
		case augments.RarityEpic:
			am.possibleAugmentsEpic = append(am.possibleAugmentsEpic, a)
		case augments.RarityLegendary:
			am.possibleAugmentsLegendary = append(am.possibleAugmentsLegendary, a)
		case augments.RarityNegative:
			am.possibleAugmentsNegative = append(am.possibleAugmentsNegative, a)
		}
	}
}

func (am *AugmentManager) rollAugments(possible []*augments.Augment) []*augments.Augment {
	rolls := make([]*augments.Augment, 0, 3)
	indices := make([]int, len(possible))
	for i := range indices {
		indices[i] = i
	}
	for i := 0; i < 3 && i < len(indices); i++ {
		idx := rand.Intn(len(indices))
		rolls = append(rolls, possible[indices[idx]])
		// Remove from local indices
		indices[idx] = indices[len(indices)-1]
		indices = indices[:len(indices)-1]
	}

	return rolls
}

func (am *AugmentManager) RollAugments() []*augments.Augment {
	am.generatePossibleAugments()

	rarity := rand.Float64()
	// Return negative rolls
	if rarity < augments.NegativeRarityPercent {
		return am.rollAugments(am.possibleAugmentsNegative)
	}
	// Return positive rolls
	rarity = rand.Float64()
	switch {
	case rarity < augments.LegendaryRarityPercent:
		return am.rollAugments(am.possibleAugmentsLegendary)
	case rarity < augments.EpicRarityPercent:
		return am.rollAugments(am.possibleAugmentsEpic)
	case rarity < augments.RareRarityPercent:
		return am.rollAugments(am.possibleAugmentsRare)
	default:
		return am.rollAugments(am.possibleAugmentsCommon)
	}
}

func (am *AugmentManager) AddAugment(augment *augments.Augment) {
	am.currentAugmentsByID[augment.ID] = augment
	am.CurrentAugments = append(am.CurrentAugments, augment)
}

func (am *AugmentManager) RemoveAugment(augment *augments.Augment) {
	am.currentAugmentsByID[augment.ID] = nil
	for i, a := range am.CurrentAugments {
		if a.ID == augment.ID {
			am.CurrentAugments[i] = am.CurrentAugments[len(am.CurrentAugments)-1]
			am.CurrentAugments = am.CurrentAugments[:len(am.CurrentAugments)-1]
			break
		}
	}
}

package augments

import "math/rand"

type Manager struct {
	currentAugmentsByID []*Augment

	possibleAugmentsCommon    []*Augment
	possibleAugmentsEpic      []*Augment
	possibleAugmentsLegendary []*Augment
	possibleAugmentsNegative  []*Augment

	CurrentAugments []*Augment
}

func NewManager() *Manager {
	return &Manager{
		currentAugmentsByID: make([]*Augment, IDMax),

		possibleAugmentsCommon:    make([]*Augment, 0, IDMax),
		possibleAugmentsEpic:      make([]*Augment, 0, IDMax),
		possibleAugmentsLegendary: make([]*Augment, 0, IDMax),
		possibleAugmentsNegative:  make([]*Augment, 0, IDMax),

		CurrentAugments: []*Augment{},
	}
}

func (m *Manager) Reset() {
	for i := range m.currentAugmentsByID {
		m.currentAugmentsByID[i] = nil
	}
	m.CurrentAugments = m.CurrentAugments[:0]
}

func (am *Manager) generatePossibleAugments() {
	am.possibleAugmentsCommon = am.possibleAugmentsCommon[:0]
	am.possibleAugmentsEpic = am.possibleAugmentsEpic[:0]
	am.possibleAugmentsLegendary = am.possibleAugmentsLegendary[:0]
	am.possibleAugmentsNegative = am.possibleAugmentsNegative[:0]

	for _, a := range List {
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
		// Hard fixes
		switch a.ID {

		}
		// Add and order by rarity
		switch a.Rarity {
		case RarityCommon:
			am.possibleAugmentsCommon = append(am.possibleAugmentsCommon, a)
		case RarityEpic:
			am.possibleAugmentsEpic = append(am.possibleAugmentsEpic, a)
		case RarityLegendary:
			am.possibleAugmentsLegendary = append(am.possibleAugmentsLegendary, a)
		case RarityNegative:
			am.possibleAugmentsNegative = append(am.possibleAugmentsNegative, a)
		}
	}
}

func (am *Manager) rollAugments(possible []*Augment) []*Augment {
	rolls := make([]*Augment, 0, 3)
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

func (am *Manager) RollAugments() []*Augment {
	am.generatePossibleAugments()

	rarity := rand.Float64()
	// Return negative rolls
	if rarity < NegativeRarityPercent {
		return am.rollAugments(am.possibleAugmentsNegative)
	}
	// Return positive rolls
	rarity = rand.Float64()
	switch {
	case rarity < LegendaryRarityPercent:
		return am.rollAugments(am.possibleAugmentsLegendary)
	case rarity < EpicRarityPercent:
		return am.rollAugments(am.possibleAugmentsEpic)
	default:
		return am.rollAugments(am.possibleAugmentsCommon)
	}
}

func (am *Manager) AddAugment(augment *Augment) Cost {
	cost := augment.cost
	if negCostAugment := am.currentAugmentsByID[IDNegateNextCost]; negCostAugment != nil {
		cost = Cost{
			Kind: CostNone,
		}
		am.RemoveAugment(negCostAugment)
	}
	// Hard check for the removal of last negative augment
	if augment.ID == IDRemoveLastNegative {
		var last *Augment
		for i := len(am.CurrentAugments) - 1; i > 0; i-- {
			if am.CurrentAugments[i].Rarity == RarityNegative {
				last = am.CurrentAugments[i]
				break
			}
		}
		if last != nil {
			am.RemoveAugment(last)
		}
		return cost
	}
	// Hard check for the removal of last positive augment
	if augment.ID == IDRemoveLastNegative {
		var last *Augment
		for i := len(am.CurrentAugments) - 1; i > 0; i-- {
			if am.CurrentAugments[i].Rarity == RarityNegative {
				last = am.CurrentAugments[i]
				break
			}
		}
		if last != nil {
			am.RemoveAugment(last)
		}
		return cost
	}
	// Add augment
	am.currentAugmentsByID[augment.ID] = augment
	am.CurrentAugments = append(am.CurrentAugments, augment)

	return cost
}

func (am *Manager) RemoveAugment(augment *Augment) {
	am.currentAugmentsByID[augment.ID] = nil
	for i, a := range am.CurrentAugments {
		if a.ID == augment.ID {
			am.CurrentAugments[i] = am.CurrentAugments[len(am.CurrentAugments)-1]
			am.CurrentAugments = am.CurrentAugments[:len(am.CurrentAugments)-1]
			break
		}
	}
}

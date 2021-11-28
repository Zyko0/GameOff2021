package augments

import "math/rand"

type Manager struct {
	currentAugmentsByID []*Augment

	possibleAugments    []*Augment

	CurrentAugments []*Augment
}

func NewManager() *Manager {
	return &Manager{
		currentAugmentsByID: make([]*Augment, IDMax),

		possibleAugments:    make([]*Augment, 0, IDMax),

		CurrentAugments: []*Augment{},
	}
}

func (m *Manager) Reset() {
	for i := range m.currentAugmentsByID {
		m.currentAugmentsByID[i] = nil
	}
	m.CurrentAugments = m.CurrentAugments[:0]
}

func (am *Manager) generatePossibleAugments(waveNumber int) {
	am.possibleAugments = am.possibleAugments[:0]

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
		// Add and order by rarity
		am.possibleAugments = append(am.possibleAugments, a)
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

func (am *Manager) RollAugments(waveNumber int) []*Augment {
	am.generatePossibleAugments(waveNumber)
	
	return am.rollAugments(am.possibleAugments)
}

func (am *Manager) AddAugment(augment *Augment) {
	am.currentAugmentsByID[augment.ID] = augment
	am.CurrentAugments = append(am.CurrentAugments, augment)
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

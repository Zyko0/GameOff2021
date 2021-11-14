package core

type ActionManager struct {
	icd      uint
	duration uint

	cooldown        uint
	currentDuration uint
}

func NewActionManager() *ActionManager {
	return &ActionManager{}
}

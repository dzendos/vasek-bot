package state

type State int

type CurrentState struct {
	UserID  int64
	Name    string
	Score   int
	WasSent bool
}

var MomState = CurrentState{
	UserID:  977443987,
	Name:    "Мама",
	Score:   1,
	WasSent: false,
}

var RomkaState = CurrentState{
	UserID:  1949227623,
	Name:    "Ромка",
	Score:   2,
	WasSent: false,
}

var EvgState = CurrentState{
	UserID:  823601282,
	Name:    "Женя",
	Score:   1,
	WasSent: false,
}

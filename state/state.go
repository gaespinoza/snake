package state

type State string

const (
	HomeState State = "home"
	GameState State = "game"
)

type MainUi struct {
	CurrentState State

	Home *HomeUi
	Game *GameUi
}

func NewMainState() *MainUi {
	homeState := NewHomeState()
	return &MainUi{
		CurrentState: HomeState,
		Home:         homeState,
		Game:         nil,
	}
}

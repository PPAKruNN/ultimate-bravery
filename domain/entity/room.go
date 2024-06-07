package entity

import "fmt"

type RoomState int8

const (
	LobbyState RoomState = iota
	InGameState
	FinishedGameState
)

type Room struct {
	ID           int
	game         string
	state        RoomState
	Participants []User
}

func (r Room) IsValid() error {
	switch {
	case validateID(r.ID) != nil:
		return fmt.Errorf("Room ID is required.")
	case r.game == "":
		return fmt.Errorf("Room game is required.")
	case r.state < 0 || r.state > FinishedGameState:
		return fmt.Errorf("Room state is not valid.")
	}
	return nil
}

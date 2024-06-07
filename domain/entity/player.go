package entity

import "fmt"

type PlayerStatus int8

const (
	NotReadyPlayerStatus PlayerStatus = iota
	ReadyPlayerStatus
	InGamePlayerStatus
	DisconnectedPlayerStatus
)

type Player struct {
	ID     int
	Room   Room
	User   User
	status PlayerStatus
}

// Check if player data is valid.
func (p Player) IsValid() error {

	switch {

	case validateID(p.ID) != nil:
		return fmt.Errorf("Player ID is required.")

	case p.Room.IsValid() != nil:
		return fmt.Errorf("Player RoomID is required.")

	case p.User.IsValid() != nil:
		return fmt.Errorf("Player User is not valid.")

	case p.status < 0 || p.status > DisconnectedPlayerStatus:
		return fmt.Errorf("Player status is not valid.")
	}

	return nil
}

func (p *Player) SetReady() error {
	if p.status != ReadyPlayerStatus {
		return p.toggleReady()
	}
	return nil
}

func (p *Player) SetNotReady() error {
	if p.status != NotReadyPlayerStatus {
		return p.toggleReady()
	}
	return nil
}

func (p *Player) SetInGame() error {

	if p.Room.state != InGameState {
		return fmt.Errorf("Cannot change player status. The Player's (ID: %d) room (ID: %d) is not in game.", p.ID, p.Room.ID)
	}

	if p.status != ReadyPlayerStatus {
		return fmt.Errorf("Cannot change player status. Player (ID: %d) is not ready.", p.ID)
	}
	return nil
}

// Set player status to be ready.
// If player is not in a room or not able to be ready, return an error.
func (p *Player) toggleReady() error {

	// Verify
	if p.status != NotReadyPlayerStatus && p.status != ReadyPlayerStatus {
		return fmt.Errorf("Cannot change player ready status. Player (ID: %d) is not at a valid status.", p.ID)
	}

	// Set
	if p.status == NotReadyPlayerStatus {
		p.status = ReadyPlayerStatus
	} else {
		p.status = NotReadyPlayerStatus
	}

	return nil
}

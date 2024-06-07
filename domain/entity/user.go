package entity

import "fmt"

type User struct {
	ID          int
	DisplayName string
}

func (u *User) IsValid() error {
	if validateID(u.ID) != nil {
		return fmt.Errorf("User ID is required.")
	}
	if u.DisplayName == "" {
		return fmt.Errorf("User DisplayName is required.")
	}
	return nil
}

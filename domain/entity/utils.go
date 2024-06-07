package entity

import "fmt"

func validateID(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID is required.")
	}

	return nil
}

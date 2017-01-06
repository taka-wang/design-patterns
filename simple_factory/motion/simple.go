package motion

import (
	"fmt"
)

// Create package level factory method to simulate 'static' factory method.
//  static factory pattern violate the open-close principle (OCP).
func Create(s string) ICard {
	switch s {
	case "8338":
		return &ethercatCard{"8338"}
	case "7856":
		return &ethercatCard{"7856"}
	default:
		return nil
	}
}

// ICard motion card interface.
//  we prefer composition than inheritance.
type ICard interface {
	// GetName get card number
	GetName() string
	// Move execute move plan
	Move()
}

// Implementation
type ethercatCard struct {
	name string
}

func (e *ethercatCard) GetName() string {
	return e.name
}

func (e *ethercatCard) Move() {
	fmt.Println(e.name + ": Ethercat Move..")
}

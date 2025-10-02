package presenter

import "fiberent/entity"

// Pet data
type Pet struct {
	ID   entity.ID `json:"id,omitempty"`
	Name string    `json:"first_name"`
	Age  int       `json:"age"`
}

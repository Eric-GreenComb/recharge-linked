package bean

import (
	"time"
)

// School School
type School struct {
	Name string `json:"name,omitempty"`
}

// Loc Loc
type Loc struct {
	Type   string    `json:"type,omitempty"`
	Coords []float64 `json:"coordinates,omitempty"`
}

// If omitempty is not set, then edges with empty values (0 for int/float, "" for string, false
// for bool) would be created for values not specified explicitly.

// Person Person
type Person struct {
	UID      string     `json:"uid,omitempty"`
	Name     string     `json:"name,omitempty"`
	Age      int        `json:"age,omitempty"`
	Dob      *time.Time `json:"dob,omitempty"`
	Married  bool       `json:"married,omitempty"`
	Raw      []byte     `json:"raw_bytes,omitempty"`
	Friends  []Person   `json:"friend,omitempty"`
	Location Loc        `json:"loc,omitempty"`
	School   []School   `json:"school,omitempty"`
}

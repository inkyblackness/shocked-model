package model

// Levels is the collection of currently available levels.
type Levels struct {
	Referable

	// Items refers to the available levels.
	Items []Identifiable `json:"items"`
}

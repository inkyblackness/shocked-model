package model

// Color describes a single color value. The channels have a range of 0..255 .
type Color struct {
	Red   int
	Green int
	Blue  int
}

// Palette is an identifiable list of colors.
type Palette struct {
	// Colors contains the color values of the palette.
	Colors [256]Color
}

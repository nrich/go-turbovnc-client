package vncclient

import "unsafe"

const colorSize = int(unsafe.Sizeof(Color{}))
const colorAlphaSize = int(unsafe.Sizeof(ColorAlpha{}))

// Color represents a single color in a color map.
type Color struct {
	R, G, B uint8
}

type ColorAlpha struct {
	R, G, B, A uint8
}

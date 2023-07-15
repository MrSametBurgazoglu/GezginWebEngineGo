package drawerBackend

import (
	"github.com/golang/freetype"
)

type GezginFontInterface interface {
	GetSize() float64
	GetSpacing() float64
	GetContext() *freetype.Context
}

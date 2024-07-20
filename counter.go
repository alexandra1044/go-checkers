//go:generate fyne bundle -o assets.go assets

package checkers

import "fyne.io/fyne/v2"

type CounterType int8

const (
	Empty = iota
	Standard
	King
)

type CounterColour int8

const (
	Colourless = iota
	Black
	Yellow
)

func GetCounterColour() CounterColour {

	return Black
}

func GetCounterSVG(c CounterType, cc CounterColour) fyne.Resource {
	switch c {
	case Standard:
		switch cc {
		case Black:
			return resourceBlueGreycounterSvg

		case Yellow:
			return resourceAmbercounterSvg
		}
	case King:
		switch cc {
		case Black:
			return resourceBlueGreycounterSvg

		case Yellow:
			return resourceKingyellowSvg
		}
	}

	return nil

}

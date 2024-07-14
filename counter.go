//go:generate fyne bundle -o assets.go assets

package main

import "fyne.io/fyne/v2"

func GetCounterSVG() fyne.Resource {
	return resourceAmbercounterSvg

}

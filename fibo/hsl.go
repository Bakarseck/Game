package main

import (
	"fmt"
	"math"
)

func rgbToHSL(r, g, b float64) (float64, float64, float64) {
	// Convert RGB values to the range 0 to 1
	rPrime := r / 255.0
	gPrime := g / 255.0
	bPrime := b / 255.0

	colo := []float64{rPrime, gPrime, bPrime}
	max, min := rPrime, rPrime

	for _, v := range colo {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	cMax, cMin := max, min

	// Calculate the difference between the maximum and minimum values
	delta := cMax - cMin

	// Calculate the Hue (H) value
	var h float64
	switch cMax {
	case rPrime:
		h = ((gPrime - bPrime) / delta)
		if delta != 0 {
			if gPrime < bPrime {
				h += 6.0
			}
		}
	case gPrime:
		h = ((bPrime-rPrime)/delta + 2.0)
	case bPrime:
		h = ((rPrime-gPrime)/delta + 4.0)
	}
	h *= 60.0 // Convert H to degrees

	// Calculate the Lightness (L) value
	l := (cMax + cMin) / 2.0

	// Calculate the Saturation (S) value
	var s float64
	if delta == 0 {
		s = 0
	} else {
		s = delta / (1.0 - math.Abs(2.0*l-1.0))
	}

	return h, s, l
}


func main() {
	r, g, b := 255.0, 0.0, 0.0
	h, s, l := rgbToHSL(r, g, b)
	fmt.Printf("RGB(%v, %v, %v) to HSL(%v, %v%%, %v%%)\n", r, g, b, h, s*100, l*100)
}

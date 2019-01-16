// Copyright © 2018 Adrian Simmons <adrian@perlucida.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	colorful "github.com/lucasb-eyer/go-colorful"
)

var (
	h    float64
	c    float64
	l    float64
	hues = map[string]float64{
		"red":     0, // and 360
		"orange":  30,
		"yellow":  60,
		"lime":    90,
		"green":   120,
		"teal":    150,
		"cyan":    180,
		"blue":    210,
		"indigo":  240,
		"violet":  270,
		"fuschia": 300,
		"pink":    330,
		"grey":    360, // desaturate input to create a grey
	}
)

func main() {
	// ic = input color as CSS style hexadecimal triplet
	ic := "#" + os.Args[1]

	// hc = convert input to colorful.color
	hc, _ := colorful.Hex(ic)

	// colorful.Color converted to hcl space (hue, chroma, lightness)
	h, c, l := hc.Hcl()

	fmt.Println("Input color:", ic)
	fmt.Println("Converted to HCL space:", h, c, l)

	genHues(h, c, l)

}

func genHues(h, c, l float64) {
	// generate hues
	// rotate the hue value around the full 360 degress with 12 steps
	steps := 12
	size := float64(360 / steps)

	for i := 0; i < steps; i++ {
		if (h + size) <= 360 {
			h = h + size
		} else if (h + size) > 360 {
			h = (h + size) - 360
		}
		fmt.Printf("hcl is: %v %v %v\n", h, c, l)

		switch {
		case h >= 0 && h <= 30:
			// add to the hues map for the red key
			fmt.Println("Assigned to: red")
		case h > 30 && h <= 60:
			// add to the hues map for the red key
			fmt.Println("Assigned to: orange")
		}

	}

	// need these in a map
}

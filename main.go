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
	"math"
	"os"
	"strconv"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// named color in hsl space
type nc struct {
	name string
	hu   float64
	sa   float64
	li   float64
}

var (
	h float64
	s float64
	l float64

	// hueNames = map[string]float64{
	// 	"red":     0, // and 360
	// 	"orange":  30,
	// 	"yellow":  60,
	// 	"lime":    90,
	// 	"green":   120,
	// 	"teal":    150,
	// 	"cyan":    180,
	// 	"blue":    210,
	// 	"indigo":  240,
	// 	"violet":  270,
	// 	"fuschia": 300,
	// 	"pink":    330,
	// 	//"red":     360,
	// }

	// define a slice of structs containig main colors
	clrs = []nc{
		{"red", 0.0, 1.0, 0.5},
		{"orange", 0.0, 1.0, 0.5},
		{"yellow", 0.0, 1.0, 0.5},
		{"lime", 0.0, 1.0, 0.5},
		{"green", 0.0, 1.0, 0.5},
		{"teal", 0.0, 1.0, 0.5},
		{"cyan", 0.0, 1.0, 0.5},
		{"blue", 0.0, 1.0, 0.5},
		{"indigo", 0.0, 1.0, 0.5},
		{"violet", 0.0, 1.0, 0.5},
		{"fuschia", 0.0, 1.0, 0.5},
		{"pink", 0.0, 1.0, 0.5},
		{"grey", 0.0, 1.0, 0.5},
	}

	shades = []nc{}
)

// TODO generate shades and tints from hues
// TODO fix shades/tints of reds

func main() {
	// ic = input color as CSS style hex
	ic := "#" + os.Args[1]

	// hc = convert input to colorful.color as hex
	hc, _ := colorful.Hex(ic)

	// colorful.Color converted to hsl space (hue, saturation, lightness)
	h, s, l := hc.Hsl()

	//fmt.Println("Input color:", ic)
	//fmt.Println("Converted to HSL space:", h, s, l)

	rotateHue(h, s, l)
	grey(h, s, l)
	genShades(h, s, l)

	// Output file with clrs as css vars
	// Open file for writing
	f, err := os.Create("colors.css")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// defers close file to when function finishes running
	defer f.Close()

	_, err = f.WriteString(`:root {
/* Generated by palegen */
`)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Write the base input color directly
	_, err = fmt.Fprintf(f, "  --base: hsl(%1.f, %d%%, %d%%);\n", math.Floor(h), int(s*100), int(l*100))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// convert clrs to CSS output
	for i := range clrs {

		hs := int(clrs[i].sa * 100)
		hl := int(clrs[i].li * 100)
		fmt.Fprintf(f, "  --%s: hsl(%1.f, %d%%, %d%%);\n", clrs[i].name, clrs[i].hu, hs, hl)

		//cv := colorful.Hsl(clrs[i].hu, clrs[i].sa, clrs[i].li)

		// TODO logic switch here based on command line flag for -rgb vs --hex
		//r, g, b := cv.RGB255()
		//fmt.Fprintf(f, "  --%s: rgb(%d, %d, %d);\n", clrs[i].name, r, g, b)
		//hex := cv.Hex()
		//fmt.Fprintf(f, "  --%s: %v;\n", clrs[i].name, hex)
	}
	_, err = fmt.Fprintf(f, "\n")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// Add a black tinted to background via alpha
	_, err = f.WriteString("  --black: hsla(0, 100%, 0%, 0.9);\n")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// empty line
	_, err = f.WriteString("\n")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// Print shades as CSS
	for i := range shades {
		s := int(shades[i].sa * 100)
		l := int(shades[i].li * 10)
		fmt.Fprintf(f, "  --%s: hsl(%1.f, %d%%, %d%%);\n", shades[i].name, shades[i].hu, s, l)
		//hex := cv.Hex()
		//fmt.Fprintf(f, "  --%s: %v;\n", shades[i].name, hex)
	}
	_, err = fmt.Fprintf(f, "\n")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	_, err = f.WriteString("\n}")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("Color generation complete.")

}

func rotateHue(h float64, s float64, l float64) []nc {

	// rotate the hue value around the full 360 degress with 12 steps
	steps := 12
	hueStep := float64(360 / steps)

	for i := 0; i < steps; i++ {

		h := math.Mod(((float64(i) * hueStep) + h), 360)

		//fmt.Printf("h is: %v\n", h)

		switch {
		case h >= 350 && h <= 360:
			for i := range clrs {
				if clrs[i].name == "red" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h >= 0 && h <= 20:
			for i := range clrs {
				if clrs[i].name == "red" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 20 && h <= 50:
			for i := range clrs {
				if clrs[i].name == "orange" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 50 && h <= 80:
			for i := range clrs {
				if clrs[i].name == "yellow" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 80 && h <= 110:
			for i := range clrs {
				if clrs[i].name == "lime" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 110 && h <= 140:
			for i := range clrs {
				if clrs[i].name == "green" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 140 && h <= 170:
			for i := range clrs {
				if clrs[i].name == "teal" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 170 && h <= 200:
			for i := range clrs {
				if clrs[i].name == "cyan" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 200 && h <= 230:
			for i := range clrs {
				if clrs[i].name == "blue" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 230 && h <= 260:
			for i := range clrs {
				if clrs[i].name == "indigo" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 260 && h <= 290:
			for i := range clrs {
				if clrs[i].name == "violet" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 290 && h <= 320:
			for i := range clrs {
				if clrs[i].name == "fuschia" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		case h > 320 && h <= 350:
			for i := range clrs {
				if clrs[i].name == "pink" {
					clrs[i].hu = math.Floor(h)
					clrs[i].sa = s
					clrs[i].li = l
				}
			}
		}
	}
	return clrs
}

func grey(h float64, s float64, l float64) []nc {

	ns := (s + 0.4) / 10

	for i := range clrs {
		if clrs[i].name == "grey" {
			clrs[i].hu = math.Floor(h)
			clrs[i].sa = ns
			clrs[i].li = 0.7
		}
	}
	return clrs
}

func genShades(h float64, s float64, l float64) []nc {

	// TODO Work in actual color from clrs
	// TODO Spread of shades is not as good as palx

	// step through lightness values to give tints and shades of main colors
	lightness := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for x := range clrs {
		for i := 0; i < len(lightness); i++ {
			nnc := new(nc)
			nnc.name = clrs[x].name + strconv.Itoa(i)
			nnc.hu = clrs[x].hu
			nnc.sa = clrs[x].sa
			nnc.li = float64(lightness[i])
			shades = append(shades, *nnc)
		}
	}
	return shades

}

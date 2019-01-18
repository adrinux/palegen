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
	}
	clrs []string, h float64, c float64, l float64
	// {
	// 	"red":     {0.0 1.0 0.5},
	// 	"orange":  {0.0 1.0 0.5},
	// 	"yellow":  {0.0 1.0 0.5},
	// 	"lime":    {0.0 1.2 0.5},
	// 	"green":   {0.0 1.2 0.5},
	// 	"teal":    {0.0 1.2 0.5},
	// 	"cyan":    {0.0 1.2 0.5},
	// 	"blue":    {0.0 1.2 0.5},
	// 	"indigo":  {0.0 1.2 0.5},
	// 	"violet":  {0.0 1.2 0.5},
	// 	"fuschia": {0.0 1.2 0.5},
	// 	"pink":    {0.0 1.2 0.5},
	// 	"grey":    {0.0 1.2 0.5}, // desaturate input to create a grey
	// }
)

// TODO ++++++++
// Create a declaration struct to hold colour name + h,c,l values
// convert genClrs to use a slice of structs
// output colours from the slice of structs

func main() {
	// ic = input color as CSS style hexadecimal triplet
	ic := "#" + os.Args[1]

	// hc = convert input to colorful.color
	hc, _ := colorful.Hex(ic)

	// colorful.Color converted to hcl space (hue, chroma, lightness)
	h, c, l := hc.Hcl()

	fmt.Println("Input color:", ic)
	//fmt.Println("Converted to HCL space:", h, c, l)

	genHues(h)
	//fmt.Println("hues map:", hues)

	genClrs(hues, h, c, l)
	//fmt.Println("Colours map:", clrs)

	// TODO generate shades and tints from hues

	// output file with clrs as css vars
	// Open file for writing
	f, err := os.Create("colors.css")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// defer close file when function finishes running
	defer f.Close()

	n1, err := f.WriteString(`:root {
/* Generated by palegen */
`)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Print("Wrote bytes:", n1)
	}

	// convert clrs to output
	for k, v := range clrs {
		c := colorful.Hcl(v)
		fmt.Fprintf(f, "  --%v: %v;\n", k, c.Hex())
	}
	// n3, err := fmt.Fprintf(f, "\n/* Color map here */\n")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// } else {
	// 	fmt.Print("Wrote bytes:", n3)
	// }

	n6, err := f.WriteString("\n}")
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Print("Wrote bytes:", n6)
	}

}

func genHues(h float64) {
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
		fmt.Printf("h is: %v \n", h)

		switch {
		case h >= 0 && h <= 30:
			// add h to the hues map for the red key
			hues["red"] = h
			//fmt.Println("Assigned to: red")
		case h > 30 && h <= 60:
			hues["orange"] = h
			//fmt.Println("Assigned to: orange")
		case h > 60 && h <= 90:
			hues["yellow"] = h
			//fmt.Println("Assigned to: yellow")
		case h > 90 && h <= 120:
			hues["lime"] = h
			//fmt.Println("Assigned to: lime")
		case h > 120 && h <= 150:
			hues["green"] = h
			//fmt.Println("Assigned to: green")
		case h > 150 && h <= 180:
			hues["teal"] = h
			//fmt.Println("Assigned to: teal")
		case h > 180 && h <= 210:
			hues["cyan"] = h
			//fmt.Println("Assigned to: cyan")
		case h > 210 && h <= 240:
			hues["blue"] = h
			//fmt.Println("Assigned to: blue")
		case h > 240 && h <= 270:
			hues["indigo"] = h
			//fmt.Println("Assigned to: indigo")
		case h > 270 && h <= 300:
			hues["violet"] = h
			//fmt.Println("Assigned to: violet")
		case h > 300 && h <= 330:
			hues["fuschia"] = h
			//fmt.Println("Assigned to: fuschia")
		case h > 330 && h <= 360:
			hues["pink"] = h
			//fmt.Println("Assigned to: pink")
		}
	}
}

func genClrs(hues map[string]float64, h float64, c float64, l float64) map[string]colorful.Color {

	clrs = make(map[string]colorful.Color)

	// range of hues map converting h back to color in hcl space
	for k, v := range hues {
		// get value h from hues create a colorful.Colour with it
		colr := colorful.Hcl(v, c, l)
		// add to clrs with key
		clrs[k] = colr
		fmt.Println("key:", k, "/ value:", colr)
	}

	n := 0.0
	clrs["grey"] = colorful.Hcl(h, n, l)

	return clrs
}

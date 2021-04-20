package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	ini "github.com/gookit/ini"
	colorful "github.com/lucasb-eyer/go-colorful"
)

// named color in hsl space
type nc struct {
	name string
	hu   float64
	ch   float64
	li   float64
}

var (
	h float64
	c float64
	l float64

	// define a slice of structs containig main hsl hues
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

	variants = []nc{}
)

func main() {

	// Load configuration from file
	config, err := ini.LoadExists("palegen.ini")
	if err != nil {
		panic(err)
	}

	// Read base colour from config
	base, ok := config.String("hex")
	if ok {
		fmt.Println("Base color found: hex", base)
	}
	// input color as CSS style hex
	ic := base

	// convert input to colorful.color as hex
	hc, err := colorful.Hex(ic)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// colorful.Color converted to hcl space (hue, chroma, lightness)
	h, c, l := hc.Hcl()

	fmt.Println("Input color:", ic)
	fmt.Println("Converted to HCL space:", h, c, l)

	// generate colors through 12 hue steps
	rotateHue(h, c, l)

	//generate a grey tinted with input (base) color
	grey(h, c, l)

	// generate 10 lightness level variants of each hue
	genVariants(h, c, l)

	// Output file with clrs as css vars
	// Get output destination
	outputFile, ok := config.String("outputFile")
	if !ok {
		outputFile = "colors.css"
	}
	// Open outputFile for writing
	if f, err := os.Create(outputFile); err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	} else {
		// defers close file to when function finishes running
		defer f.Close()

		// start root section of css
		if _, err = f.WriteString(`:root {
/* Generated by palegen */
`); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// Write the input base color directly as css vars
		_, err = fmt.Fprintf(f, "  --base: hsla(%1.f, %d%%, %d%%, 1);\n", math.Floor(h), int(c*100), int(l*100))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// convert clrs slice to CSS output as css vars
		for i := range clrs {
			hs := int(clrs[i].ch * 100)
			hl := int(clrs[i].li * 100)
			ha := 1
			fmt.Fprintf(f, "  --%s: hsla(%1.f, %d%%, %d%%, %d);\n", clrs[i].name, clrs[i].hu, hs, hl, ha)
		}

		// black line in css
		_, err = fmt.Fprintf(f, "\n")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// Directly Add a black tinted to background via alpha
		_, err = f.WriteString("  --textblack: hsla(0, 100%, 0%, 0.9);\n")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// Directly Add a wite tinted to background via alpha
		_, err = f.WriteString("  --textwhite: hsla(360, 100%, 100%, 0.9);\n")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// empty line
		_, err = f.WriteString("\n")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// Print lightness variants as CSS
		for i := range variants {
			s := int(variants[i].ch * 100)
			l := int(variants[i].li)
			a := 1
			fmt.Fprintf(f, "  --%s: hsla(%1.f, %d%%, %d%%, %d);\n", variants[i].name, variants[i].hu, s, l, a)
		}

		// end root section of css
		_, err = f.WriteString("\n}")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Println("Color generation complete.")
	}
}

func rotateHue(h float64, c float64, l float64) []nc {

	// rotate the hue value around the full 360 degress with 12 steps
	steps := 12
	hueStep := float64(360 / steps)

	for i := 0; i < steps; i++ {

		h := math.Mod(((float64(i) * hueStep) + h), 360)

		switch {
		case h >= 350 && h <= 360:
			for i := range clrs {
				if clrs[i].name == "red" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h >= 0 && h <= 20:
			for i := range clrs {
				if clrs[i].name == "red" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 20 && h <= 50:
			for i := range clrs {
				if clrs[i].name == "orange" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 50 && h <= 80:
			for i := range clrs {
				if clrs[i].name == "yellow" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 80 && h <= 110:
			for i := range clrs {
				if clrs[i].name == "lime" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 110 && h <= 140:
			for i := range clrs {
				if clrs[i].name == "green" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 140 && h <= 170:
			for i := range clrs {
				if clrs[i].name == "teal" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 170 && h <= 200:
			for i := range clrs {
				if clrs[i].name == "cyan" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 200 && h <= 230:
			for i := range clrs {
				if clrs[i].name == "blue" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 230 && h <= 260:
			for i := range clrs {
				if clrs[i].name == "indigo" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 260 && h <= 290:
			for i := range clrs {
				if clrs[i].name == "violet" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 290 && h <= 320:
			for i := range clrs {
				if clrs[i].name == "fuschia" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		case h > 320 && h <= 350:
			for i := range clrs {
				if clrs[i].name == "pink" {
					clrs[i].hu = math.Floor(h)
					clrs[i].ch = c
					clrs[i].li = l
				}
			}
		}
	}
	return clrs
}

func grey(h float64, c float64, l float64) []nc {

	nc := c / 10

	for i := range clrs {
		if clrs[i].name == "grey" {
			clrs[i].hu = math.Floor(h)
			clrs[i].ch = nc
			clrs[i].li = 0.5
		}
	}
	return clrs
}

func genVariants(h float64, c float64, l float64) []nc {

	// step through fixed lightness values
	lightness := [10]float64{98, 88, 78, 67, 57, 47, 37, 26, 16, 6}

	for x := range clrs {
		for i := 0; i < len(lightness); i++ {
			nnc := new(nc)
			nnc.name = clrs[x].name + strconv.Itoa(i)
			nnc.hu = clrs[x].hu
			nnc.ch = clrs[x].ch
			nnc.li = lightness[i]
			variants = append(variants, *nnc)
		}
	}
	return variants

}

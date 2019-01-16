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

func main() {
	// ic = input color as CSS style hexadecimal triplet
	ic := "#" + os.Args[1]

	// hc = convert input to colorful.color
	hc, _ := colorful.Hex(ic)

	// colorful.color converted to hcl space (hue, chroma, lightness)
	h, c, l := hc.Hcl()

	fmt.Println("Input color:", ic)
	fmt.Println("Converted to HCL space:", h, c, l)

	// generate hues
	// rotate the hue value around the full 360 degress with 12 steps
	steps := 12
	size := float64(360 / steps)

	for i := 0; i < steps; i++ {
		h = h + size
		fmt.Printf("hcl is: %v %v %v\n", h, c, l)
	}

	// need these in a map
}

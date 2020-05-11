package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"

	"github.com/mattn/go-forlines"
	hsbcolor "github.com/xyproto/color"
)

func main() {
	err := forlines.Do(os.Stdin, func(line string) error {
		p, err := strconv.ParseFloat(line, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		rgb := percent2rgb(p)
		fmt.Printf("%07f: %02X%02X%02X\n", p, rgb.R, rgb.G, rgb.B)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func percent2rgb(p float64) color.RGBA {
	if p < 0 {
		p = 0
	}
	if p > 1 {
		p = 1
	}

	hsba := &hsbcolor.HSBA{
		H: int(math.Round(120 * p)), // 120 means green
		S: 1.0,
		B: 1.0,
		A: 1.0,
	}
	rgb := hsba.RGB()
	return rgb
}

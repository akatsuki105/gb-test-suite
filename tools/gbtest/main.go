package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
)

const (
	width  = 160
	height = 144
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Command: gbtest <FILE1> <FILE2>")
		os.Exit(1)
	}

	file1, err := os.Open(args[0])
	if err != nil {
		panic("FILE1 is not found")
	}

	file2, err := os.Open(args[1])
	if err != nil {
		panic("FILE2 is not found")
	}

	image1, _, err := image.Decode(file1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to decode FILE1 image: %s\n", err.Error())
		os.Exit(1)
	}
	width1, height1 := image1.Bounds().Dx(), image1.Bounds().Dy()
	if width1 != width || height1 != height {
		fmt.Fprintln(os.Stderr, "FILE1 is not GameBoy screenshot")
		os.Exit(1)
	}

	image2, _, err := image.Decode(file2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to decode FILE2 image: %s\n", err.Error())
		os.Exit(1)
	}
	width2, height2 := image2.Bounds().Dx(), image2.Bounds().Dy()
	if width2 != width || height2 != height {
		fmt.Fprintln(os.Stderr, "FILE2 is not GameBoy screenshot")
		os.Exit(1)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r1, g1, b1, _ := image1.At(x, y).RGBA()
			black1 := r1 == 0 && g1 == 0 && b1 == 0

			r2, g2, b2, _ := image2.At(x, y).RGBA()
			black2 := r2 == 0 && g2 == 0 && b2 == 0

			if black1 || black2 {
				if !(black1 && black2) {
					fmt.Fprintln(os.Stderr, "Two images are not the same")
					os.Exit(1)
				}
			}
		}
	}
}

package main

import (
	"flag"
	"github.com/disintegration/imaging"
	"image"
	"strings"
)

func main() {

	filesArg := flag.String("files", "", "list of images separated by commas")
	outputArg := flag.String("output", "output.png", "output image")
	flag.Parse()

	files := strings.Split(*filesArg, ",")
	output := *outputArg

	if files[0] == "" {
		flag.Usage()
		return
	}

	newImg, err := imaging.Open(files[0])

	if err != nil {
		panic(err)
	}

	for i, size := 1, len(files); size > 1 && i < size; i++ {

		overImg, err := imaging.Open(files[i])

		if err != nil {
			panic(err)
		}

		newImg = imaging.Overlay(newImg, overImg, image.Pt(0, 0), 1)
	}

	err = imaging.Save(newImg, output)

	if err != nil {
		panic(err)
	}

}

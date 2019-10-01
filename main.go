package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type dot struct {
	X int
	Y int
}

const width = 100
const height = 100

var in, out int = 0, 0

func setup(imd *imdraw.IMDraw) {
	imd.Color = colornames.Gray
	imd.Push(pixel.V(0, 0))
	imd.Circle(width, 5)
}

func render(imd *imdraw.IMDraw) {

	newDot := dot{
		X: rand.Intn(width + 1),
		Y: rand.Intn(height + 1),
	}

	if math.Sqrt(math.Pow(float64(newDot.X), 2)+math.Pow(float64(newDot.Y), 2)) < width {
		imd.Color = colornames.Green
		in++
	} else {
		imd.Color = colornames.Red
		out++
	}

	imd.Push(pixel.V(float64(newDot.X), float64(newDot.Y)))
	imd.Circle(1, 5)

	total := in + out

	pi := (float64(in) / float64(total)) * 4

	fmt.Printf("pi = %v\n", pi)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "PiCulate",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	win.Clear(colornames.Skyblue)
	setup(imd)

	for !win.Closed() {
		render(imd)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

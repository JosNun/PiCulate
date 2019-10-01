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

var dots = make([]dot, 0)

func update() {
	newDot := dot{
		X: rand.Intn(width + 1),
		Y: rand.Intn(height + 1),
	}

	dots = append(dots, newDot)
}

func render(imd *imdraw.IMDraw) {
	var in, out int

	imd.Color = colornames.Gray
	imd.Push(pixel.V(0, 0))
	imd.Circle(width, 5)

	for _, dot := range dots {
		if math.Sqrt(math.Pow(float64(dot.X), 2)+math.Pow(float64(dot.Y), 2)) < width {
			imd.Color = colornames.Green
			in++
		} else {
			imd.Color = colornames.Red
			out++
		}

		imd.Push(pixel.V(float64(dot.X), float64(dot.Y)))
		imd.Circle(1, 5)
	}

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

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		update()
		render(imd)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

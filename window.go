package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var win *pixelgl.Window

// Create launches a window designed for the given game of life parameters
// (number of blocks in width and height, and the size of one block)
func run(g *game, blockSize int) {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title: "Game Of Life",
			Bounds: pixel.R(0.0, 0.0,
				float64(g.size*blockSize), float64(g.size*blockSize)),
			VSync: false,
		}

		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		imd := imdraw.New(nil)

		for !win.Closed() {
			win.Clear(colornames.White)

			// Block rendering
			imd.Clear()
			for i := 0; i < g.size; i++ {
				for j := 0; j < g.size; j++ {
					imd.Color = colornames.White
					if g.get(i, j) == 1 {
						imd.Color = colornames.Black
					}
					imd.Push(pixel.V(float64(i*blockSize), float64(j*blockSize)))
					imd.Push(pixel.V(float64((i+1)*blockSize), float64((j+1)*blockSize)))
					imd.Rectangle(0)
				}
			}

			// Game logic
			g.run()

			imd.Draw(win)
			win.Update()

			time.Sleep(time.Millisecond * time.Duration(80))
		}
	})
}

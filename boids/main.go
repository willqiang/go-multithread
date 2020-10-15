package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
)

var (
	grenn = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), grenn)
			screen.Set(int(boid.position.x-1), int(boid.position.y), grenn)
			screen.Set(int(boid.position.x), int(boid.position.y-1), grenn)
			screen.Set(int(boid.position.x), int(boid.position.y+1), grenn)
		}
	}
	return nil
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Boids in a box"); err != nil {
		log.Fatalln(err)
	}
}

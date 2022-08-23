package main

import "github.com/gen2brain/raylib-go/raylib"

const (
  screenWidth = 1000
  screenHeight = 480
)

var (
  running = true
)

func drawScene() {}

func input() {}
func update() {
  running = !rl.WindowShouldClose()
}

func render() {
  rl.BeginDrawing()
  rl.ClearBackground(rl.RayWhite)

  drawScene()

  rl.EndDrawing()
}

func init() {
  rl.InitWindow(screenWidth, screenHeight, "seeder")
	rl.SetTargetFPS(60)
}

func quit() {
  rl.CloseWindow()
}


func main() {
	for running {
    input()
    update()
    render()
	}

  quit()
}

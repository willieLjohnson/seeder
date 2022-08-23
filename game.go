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

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	for running {
    input()
    update()
    render()
	}

	rl.CloseWindow()
}

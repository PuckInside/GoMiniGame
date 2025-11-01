package main

import (
	"GoEngine/asset"
	"GoEngine/engine"
	"GoEngine/resource"
	"fmt"
	"time"
)

const FPS = 120
const DeltaTime = 1.0 / FPS

var AxisMapping = map[string]engine.Vector2i{
	"w": {X: 0, Y: -1},
	"s": {X: 0, Y: 1},
	"a": {X: -1, Y: 0},
	"d": {X: 1, Y: 0},
}

var turn_timer *engine.Timer
var camera *engine.Camera
var collision_layer *engine.CollisionLayer
var direction *engine.Vector2i
var player *asset.Player

func init_main() {
	turn_timer = engine.Timer{}.New(1)
	turn_timer.Start()

	var window engine.Vector2i = engine.Vector2i{X: 32, Y: 12}
	camera = engine.Camera{}.New(window, engine.Vector2i{X: 0, Y: 0})

	colliders := resource.GenerateCollisionMap(resource.WorldMap)
	collision_layer = &engine.CollisionLayer{
		Position:  engine.Vector2i{X: 0, Y: 0},
		Colliders: colliders,
	}

	direction = &engine.Vector2i{X: 0, Y: 0}
	player = asset.Player{}.New(engine.Vector2f{X: 24.0, Y: 6.0}, asset.PlayerSprite, 2)
}

func loop() {
	for {
		fmt.Print("\033c")
		fmt.Println("FPS -", FPS)
		turn_timer.Update(DeltaTime)

		camera.InitBuffer()
		camera.Draw(engine.Vector2i{X: 0, Y: 0}, resource.WorldMap)

		player.MoveAndCollide(*direction, *collision_layer, DeltaTime)
		player_layer := [][]rune{{player.Sprite}}
		fmt.Printf("🐸 Позиция игрока: X=%.3f, Y=%.3f\n", player.Position.X, player.Position.Y)

		camera.Draw(player.Position.GetVector2i(), player_layer)
		camera.Render()

		if turn_timer.Is_stopped {
			var input string = "w"
			input = turn_input()
			*direction = AxisMapping[input]
			turn_timer.Start()
		}

		time.Sleep(time.Second / FPS)
	}
}

func main() {
	init_main()
	loop()
}

func turn_input() string {
	var input string

	fmt.Print("Введите направление(w,a,s,d):")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("\nПроизошла ошибка при вводе:", err)
		return "!Error!"
	}

	return input
}

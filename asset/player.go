package asset

import "GoEngine/engine"

const PlayerSprite = '🐸'

type Player struct {
	Position engine.Vector2f
	Sprite   rune
	Speed    float32
}

func (player *Player) MoveAndCollide(input engine.Vector2i, collision engine.CollisionLayer, delta float32) {
	var velocity engine.Vector2f

	if input.X != 0.0 {
		velocity.X += input.GetVector2i().X * player.Speed * delta
	}
	if input.Y != 0.0 {
		velocity.Y += input.GetVector2i().Y * player.Speed * delta
	}

	if player.isCollide(collision, velocity.X, velocity.Y) {
		return
	}

	player.Position.X += velocity.X
	player.Position.Y += velocity.Y
}

func (player *Player) isCollide(collision engine.CollisionLayer, d_x float32, d_y float32) bool {
	new_x := player.Position.X + d_x
	new_y := player.Position.Y + d_y

	target_tile_x := int(new_x)
	target_tile_y := int(new_y)

	map_height := len(collision.Colliders)
	if map_height == 0 {
		return false
	}

	map_width := len(collision.Colliders[0])
	if target_tile_x < 0 || target_tile_x >= map_width ||
		target_tile_y < 0 || target_tile_y >= map_height {
		return false
	}

	return collision.Colliders[target_tile_y][target_tile_x]
}

func (p Player) New(position engine.Vector2f, sprite rune, speed float32) *Player {
	player := &Player{
		Position: position,
		Sprite:   sprite,
		Speed:    speed,
	}

	return player
}

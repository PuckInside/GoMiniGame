package engine

type Vector2f struct {
	X float32
	Y float32
}

func (v Vector2f) GetVector2i() Vector2i {
	return Vector2i{X: int(v.X), Y: int(v.Y)}
}

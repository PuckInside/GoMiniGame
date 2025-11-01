package engine

type Vector2i struct {
	X int
	Y int
}

func (v Vector2i) GetVector2i() Vector2f {
	return Vector2f{X: float32(v.X), Y: float32(v.Y)}
}

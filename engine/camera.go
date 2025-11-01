package engine

import "fmt"

type Camera struct {
	Position   Vector2i
	Resolution Vector2i
	buffer     [][]rune
}

func (c *Camera) Draw(layer_position Vector2i, layer [][]rune) {
	cam_width := c.Resolution.X
	cam_height := c.Resolution.Y

	for y := 0; y < len(layer); y++ {
		layer_row := layer[y]
		target_y := layer_position.Y + y - c.Position.Y

		if target_y < 0 || target_y >= cam_height {
			continue
		}

		for x := 0; x < len(layer_row); x++ {
			char := layer_row[x]
			target_x := layer_position.X + x - c.Position.X

			if target_x < 0 || target_x >= cam_width {
				continue
			}

			c.buffer[target_y][target_x] = char
		}
	}
}

func (c *Camera) Render() {
	for _, row := range c.buffer {
		fmt.Println(string(row))
	}
}

func (c *Camera) InitBuffer() {
	height := c.Resolution.Y
	width := c.Resolution.X

	c.buffer = make([][]rune, height)

	for i := 0; i < height; i++ {
		c.buffer[i] = make([]rune, width)

		for j := 0; j < width; j++ {
			c.buffer[i][j] = ' '
		}
	}
}

func (c Camera) New(resolution Vector2i, position Vector2i) *Camera {
	camera := &Camera{
		Position:   position,
		Resolution: resolution,
	}

	camera.InitBuffer()
	return camera
}

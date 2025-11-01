package engine

type Timer struct {
	timer      float32
	duration   float32
	Is_stopped bool
}

func (t *Timer) Update(delta float32) {
	t.timer -= delta
	if t.timer <= 0.0 {
		t.Is_stopped = true
	}
}

func (t *Timer) Start() {
	t.timer = t.duration
	t.Is_stopped = false
}

func (t *Timer) Stop() {
	t.timer = 0.0
	t.Is_stopped = true
}

func (t Timer) New(duration float32) *Timer {
	timer := &Timer{
		timer:      0.0,
		duration:   duration,
		Is_stopped: true,
	}

	return timer
}

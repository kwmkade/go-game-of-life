package main

import (
	"go-game-of-life/life"
	"go-game-of-life/ui"
	"time"
)

type context struct {
	container *life.Container
	ui        *ui.UI
}

func update(c *context) {
	for c.container.Begin(); ; {
		if cell := c.container.Next(); cell != nil {
			cell.Judge()
			continue
		}
		break
	}

	for c.container.Begin(); ; {
		if cell := c.container.Next(); cell != nil {
			cell.Fix()
			c.ui.UpdateCell(cell)
			continue
		}
		break
	}

	c.ui.Update()
}

func main() {

	context := context{
		new(life.Container),
		new(ui.UI)}

	seed := time.Now().UnixNano()
	context.container.Initialize(100, 100, seed)
	context.ui.Initialize()

	defer context.ui.Finalize()

	ticker := time.NewTicker(20 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				update(&context)
			}
		}
	}()

	running := true
	for running {
		if context.ui.HasQuit() {
			running = false
			break
		}
	}

	ticker.Stop()
}

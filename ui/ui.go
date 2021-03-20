package ui

import (
	"go-game-of-life/life"

	"github.com/veandco/go-sdl2/sdl"
)

type rectkey struct {
	x int
	y int
}

// UI は画面表示を管理するため構造体です
type UI struct {
	window  *sdl.Window
	surface *sdl.Surface
	rects   map[rectkey]*sdl.Rect
}

// Initialize はUIの初期化を行います
func (ui *UI) Initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Game of Life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		500, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	window.UpdateSurface()

	ui.window = window
	ui.surface = surface
	ui.rects = make(map[rectkey]*sdl.Rect)
}

// Finalize はUIの終了処理を行います
func (ui *UI) Finalize() {
	sdl.Quit()
	ui.window.Destroy()
}

// HasQuit は停止要求があった場合にtrueを返します
func (ui UI) HasQuit() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

// UpdateCell はセルの色を更新します
func (ui *UI) UpdateCell(cell *life.Cell) {

	if _, exists := ui.rects[rectkey{cell.X, cell.Y}]; !exists {
		rect := new(sdl.Rect)
		rect.X = int32(cell.X) * 5
		rect.Y = int32(cell.Y) * 5
		rect.W = 5
		rect.H = 5
		ui.rects[rectkey{cell.X, cell.Y}] = rect
		ui.surface.FillRect(rect, cellColor(cell.Status == life.Alive))
		return
	}

	if cell.HasChanged {
		rect := ui.rects[rectkey{cell.X, cell.Y}]
		ui.surface.FillRect(rect, cellColor(cell.Status == life.Alive))
	}
}

func cellColor(isAlive bool) uint32 {
	if isAlive {
		return 0xffffd7ff
	}
	return 0xffffffff
}

// Update は表示を更新します
func (ui *UI) Update() {
	ui.window.UpdateSurface()
}

package life

import (
	"math/rand"
)

// Status はセルの生死状態
type Status int

const (
	// Alive は生きている状態
	Alive Status = 1
	// Dead は死んでいる状態
	Dead Status = -1
)

// Cell はセルが持つ属性を管理する構造体
type Cell struct {
	Status     Status
	willChange bool
	HasChanged bool
	X          int
	Y          int
	container  *Container
}

// Judge では次フレームの生成状態を判定します
func (s *Cell) Judge() {

	aliveNum := 0

	// 北
	if s.Y < int(s.container.height()-1) {
		o := s.container.cells[s.X][s.Y+1]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 北東
	if s.X < int(s.container.width()-1) && s.Y < int(s.container.height()-1) {
		o := s.container.cells[s.X+1][s.Y+1]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 東
	if s.X < int(s.container.width()-1) {
		o := s.container.cells[s.X+1][s.Y]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 南東
	if s.X < int(s.container.width()-1) && 0 < s.Y {
		o := s.container.cells[s.X+1][s.Y-1]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 南
	if 0 < s.Y {
		o := s.container.cells[s.X][s.Y-1]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 南西
	if 0 < s.X && 0 < s.Y {
		o := s.container.cells[s.X-1][s.Y-1]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 西
	if 0 < s.X {
		o := s.container.cells[s.X-1][s.Y]
		if o.Status == Alive {
			aliveNum++
		}
	}
	// 北西
	if 0 < s.X && s.Y < int(s.container.height()-1) {
		o := s.container.cells[s.X-1][s.Y+1]
		if o.Status == Alive {
			aliveNum++
		}
	}

	s.willChange = false
	s.HasChanged = false

	switch {
	// 【誕生】死んでいるセルに隣接する生きたセルがちょうど3つあれば、次の世代が誕生する。
	case s.Status == Dead && aliveNum == 3:
		s.willChange = true
	// 【生存】 生きているセルに隣接する生きたセルが2つか3つならば、次の世代でも生存する。
	case s.Status == Alive && (aliveNum == 2 || aliveNum == 3):
	// 【過疎】 生きているセルに隣接する生きたセルが1つ以下ならば、過疎により死滅する。
	case s.Status == Alive && aliveNum <= 1:
		s.willChange = true
	// 【過密】 生きているセルに隣接する生きたセルが4つ以上ならば、過密により死滅する。
	case s.Status == Alive && aliveNum >= 4:
		s.willChange = true
	}
}

// Fix では次フレームの生死状態を確定させます
func (s *Cell) Fix() {
	if s.willChange {
		s.Status *= -1 // NOTE: Alive->Dead or Dead->Alive
		s.HasChanged = true
	}
}

// Container はセルを保持する役割を持ち、全てのセルはContainerに所属します
type Container struct {
	cells [][]*Cell
	iterX uint
	iterY uint
}

// Initialize ではContainerを初期化します
func (c *Container) Initialize(width uint, height uint, seed int64) bool {

	c.cells = make([][]*Cell, int(width))
	for i := range c.cells {
		c.cells[i] = make([]*Cell, int(height))
	}

	rand.Seed(seed)

	for i := 0; i < int(width); i++ {
		for j := 0; j < int(height); j++ {

			cell := new(Cell)

			if rand.Intn(100) > 60 {
				cell.Status = Alive
			} else {
				cell.Status = Dead
			}

			cell.X = i
			cell.Y = j
			cell.container = c

			c.cells[i][j] = cell
		}
	}

	return true
}

func (c Container) width() uint {
	return uint(len(c.cells))
}

func (c Container) height() uint {
	if c.width() == 0 {
		return 0
	}
	return uint(len(c.cells[0]))
}

// Next はすべてのセルに対する反復処理において次のセルを返します
// 最後の要素を返した後は nil を返します
func (c *Container) Next() *Cell {

	if c.iterY >= c.height() {
		return nil
	}

	cell := c.cells[c.iterX][c.iterY]

	c.iterX++
	if c.iterX >= c.width() {
		c.iterX = 0
		c.iterY++
	}

	return cell
}

// Begin はすべてのセルに対する反復処理を開始するときに一度呼ぶ必要があります
func (c *Container) Begin() {
	c.iterX = 0
	c.iterY = 0
}

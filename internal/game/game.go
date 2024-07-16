package game

import (
	"math/rand"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	X, Y int
}

type Game struct {
	Width, Height  int
	Snake          *Snake
	Food           Position
	lastUpdateTime time.Time
	speed          time.Duration
	gameOver       bool // 游戏结束标志
}

func NewGame(width, height int) *Game {
	g := &Game{
		Width:  width,
		Height: height,
		Snake:  NewSnake(Position{X: width / 2, Y: height / 2}),
		speed:  100 * time.Millisecond,
	}
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	g.spawnFood()
	g.lastUpdateTime = time.Now()
	return g
}

func (g *Game) Update() {
	if g.gameOver {
		return
	}

	if time.Since(g.lastUpdateTime) < g.speed {
		return
	}

	g.Snake.Move()

	if g.isCollision() {
		g.gameOver = true
	}

	if g.Snake.Head() == g.Food {
		g.Snake.Grow()
		g.spawnFood()
	}

	g.lastUpdateTime = time.Now()
}

func (g *Game) ChangeDirection(dir Direction) {
	g.Snake.ChangeDirection(dir)
}

func (g *Game) spawnFood() {
	g.Food = Position{X: rand.Intn(g.Width), Y: rand.Intn(g.Height)}
}

func (g *Game) IsGameOver() bool {
	return g.gameOver
}

func (g *Game) isCollision() bool {
	head := g.Snake.Head()

	// 检查是否撞到墙壁
	if head.X < 0 || head.X >= g.Width || head.Y < 0 || head.Y >= g.Height {
		return true
	}

	// 检查是否撞到自己的身体
	for _, body := range g.Snake.Body[1:] { // 蛇身体从索引1开始，因为索引0是头部
		if body == head {
			return true
		}
	}

	return false
}

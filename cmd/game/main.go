package main

import (
	"image/color"
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 1440
	screenHeight = 720
	tileSize     = 15
	gridCols     = screenWidth / tileSize
	gridRows     = screenHeight / tileSize
	tickRate     = 10
)

type Game struct {
	grid          [gridCols][gridRows]int
	pressedMemory [][2]int
	epochs        int
	lastUpdate    time.Time
	clearSet      bool
	gameOn        bool
	tickRate      time.Duration
}

func (g *Game) Update(game *ebiten.Image) error {
	now := time.Now()

	if now.Sub(g.lastUpdate) < time.Second/g.tickRate {
		return nil
	}
	g.lastUpdate = now

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Get Cursor position
		cursorX, cursorY := ebiten.CursorPosition()
		x := cursorX / tileSize
		y := cursorY / tileSize

		if y < gridRows { // Check if within the grid
			if g.grid[x][y] == 0 {
				// Remember action
				var action [2]int = [2]int{x, y}
				g.pressedMemory = append(g.pressedMemory, action)
			}
			// Set grid
			g.grid[x][y] = 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		lenght := len(g.pressedMemory)
		if lenght > 0 {
			last_actions := g.pressedMemory[len(g.pressedMemory)-1]
			g.grid[last_actions[0]][last_actions[1]] = 0
			g.pressedMemory = g.pressedMemory[:lenght-1]
		}

	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.clearSet = !g.clearSet
		g.gameOn = !g.gameOn
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.tickRate++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.tickRate--
	}

	if g.gameOn {
		g.applyConwaysGameOfLifeRules()
		g.epochs++
	}

	return nil
}

func (g *Game) applyConwaysGameOfLifeRules() {
	newGrid := [gridCols][gridRows]int{}
	for y := 0; y < gridRows; y++ {
		for x := 0; x < gridCols; x++ {
			liveNeighbors := g.countLiveNeighbors(x, y)
			if g.grid[x][y] == 1 {
				if liveNeighbors == 2 || liveNeighbors == 3 {
					newGrid[x][y] = 1
				} else {
					newGrid[x][y] = 0
				}
			} else {
				if liveNeighbors == 3 {
					newGrid[x][y] = 1
				}
			}
		}
	}
	g.grid = newGrid
}

func (g *Game) countLiveNeighbors(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < gridCols && ny >= 0 && ny < gridRows {
				if g.grid[nx][ny] == 1 {
					count++
				}
			}
		}
	}
	return count
}

func (g *Game) Draw(screen *ebiten.Image) {
	cursorX, cursorY := ebiten.CursorPosition()

	for y := 0; y < gridRows; y++ {
		for x := 0; x < gridCols; x++ {
			rectX := x * tileSize
			rectY := y * tileSize
			borderColor := color.Black
			tileColor := color.Black
			if cursorX >= rectX && cursorX < rectX+tileSize && cursorY >= rectY && cursorY < rectY+tileSize {
				borderColor = color.White
			}
			if g.grid[x][y] == 1 {
				tileColor = color.White
			}

			ebitenutil.DrawRect(screen, float64(rectX), float64(rectY), tileSize, tileSize, tileColor)
			ebitenutil.DrawRect(screen, float64(rectX), float64(rectY), 1, tileSize, borderColor)
			ebitenutil.DrawRect(screen, float64(rectX), float64(rectY), tileSize, 1, borderColor)
			ebitenutil.DrawRect(screen, float64(rectX+tileSize-1), float64(rectY), 1, tileSize, borderColor)
			ebitenutil.DrawRect(screen, float64(rectX), float64(rectY+tileSize-1), tileSize, 1, borderColor)
		}
	}

	if g.gameOn {
		ebitenutil.DebugPrint(
			screen, "On\nEpochs "+
				strconv.Itoa(g.epochs)+
				"\nUP/DOWN to adjust tick-rate"+
				"\nTick Rate:"+g.tickRate.String())
	} else {
		ebitenutil.DebugPrint(
			screen, "Off\nEpochs "+
				strconv.Itoa(g.epochs)+
				"\nUP/DOWN to adjust tick-rate"+
				"\nTick Rate:"+g.tickRate.String())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1400, 700
}

func main() {
	ebiten.SetWindowSize(1440, 720)
	ebiten.SetWindowTitle("Gonway's Game of Life")
	if err := ebiten.RunGame(&Game{tickRate: tickRate}); err != nil {
		log.Fatal(err)
	}
}

# Conway's Game of Life in Go

## Overview

This is a simple implementation of Conway's Game of Life using the Go programming language and the Ebiten game library.
The Game of Life is a cellular automaton devised by the British mathematician John Horton Conway in 1970.
It is a zero-player game, meaning its evolution is determined by its initial state, with no further input from the user.
This implementation allows users to interact with the grid, toggle cells on and off, start and stop the simulation, and adjust the tick rate (speed) of the simulation.

## Features

- Interactive Grid: Users can click on the grid to toggle cells between alive (white) and dead (black).
- Undo Feature: Users can undo the last placed tile by pressing Backspace.
- Simulation Control: Users can start and stop the simulation by pressing Enter.
- Adjustable Speed: The simulation speed (tick rate) can be adjusted using the Up and Down arrow keys.

## Controls

- Left Mouse Button: Toggle the state of a cell (alive/dead).
- Backspace: Undo the last tile placement.
- Enter: Start/stop the simulation.
- Up Arrow: Increase the tick rate (speed up the simulation).
- Down Arrow: Decrease the tick rate (slow down the simulation).

## Installation

Prerequisites:

- Go 1.16 or higher
- Ebiten game library

### Installation Steps

Clone the repository:

```bash
    git clone <https://github.com/yourusername/conway-go-game.git>
    cd conway-go-game
```

Install Ebiten:

```bash
    go get -u github.com/hajimehoshi/ebiten/v2
```

Run the game:

```bash
    go run cmd/game/main.go 
```

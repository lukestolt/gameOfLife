package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe [][]bool

const (
	width  = 80
	height = 15
)

func main() {
	universe := NewUniverse()
	fmt.Println(len(universe))
	fmt.Println(len(universe[0]))
	Seed(universe)
	Show(universe)
}

func NewUniverse() Universe {
	u := make(Universe, width)
	for x := range u {
		u[x] = make([]bool, height)
	}
	return u
}

const (
	dead  = " "
	alive = "*"
)

func Show(u Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cellVal := u[j][i]
			if cellVal {
				fmt.Print(alive)
			} else {
				fmt.Print(dead)
			}
		}
		fmt.Print("\n")
	}
}

func Seed(u Universe) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			randNum := rand.Intn(4)
			// 25% chance of being alive
			if randNum == 0 {
				u[i][j] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	if x > width-1 {
		x = 0
	} else if x < 0 {
		x = width - 1
	}
	if y > height-1 {
		y = 0
	} else if y < 0 {
		y = height - 1
	}
	return u[x][y]
}

func (u Universe) GetNumAliveNeighbors(x, y int) int {
	var numNeighbors = 0
	if u.Alive(x+1, y) {
		numNeighbors++
	}
	if u.Alive(x+1, y+1) {
		numNeighbors++
	}
	if u.Alive(x+1, y-1) {
		numNeighbors++
	}
	if u.Alive(x-1, y) {
		numNeighbors++
	}
	if u.Alive(x-1, y+1) {
		numNeighbors++
	}
	if u.Alive(x-1, y-1) {
		numNeighbors++
	}
	if u.Alive(x, y+1) {
		numNeighbors++
	}
	if u.Alive(x, y-1) {
		numNeighbors++
	}

	return numNeighbors
}

// should be called next
func Next(x int, y int, universe Universe) {
	numNeighbors := universe.GetNumAliveNeighbors(x, y)
	if numNeighbors < 2 || numNeighbors > 3 {
		universe[x][y] = false
	} else if numNeighbors == 3 && !universe[x][y] {
		universe[x][y] = true
	}
	// < 2 alive neighbors - dies
	// > 3 alive neighbors - dies
	// alive cell with [2, 3] alive neighbors - lives
	// dead cell [3] alive neighbors - turns alive
}

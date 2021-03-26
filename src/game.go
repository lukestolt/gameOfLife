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
	RunGame()
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

func Seed(u Universe, percentAlive int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			randNum := rand.Intn(101)
			if randNum <= percentAlive {
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
func (u Universe) Next(x int, y int) bool {
	numNeighbors := u.GetNumAliveNeighbors(x, y)
	if numNeighbors < 2 || numNeighbors > 3 {
		return false
	} else if numNeighbors == 3 && !u[x][y] {
		return true
	} else {
		return u[x][y]
	}
	// < 2 alive neighbors - dies
	// > 3 alive neighbors - dies
	// alive cell with [2, 3] alive neighbors - lives
	// dead cell [3] alive neighbors - turns alive
}

func Step(a, b Universe) {
	for x := range a {
		for y := range a[x] {
			b[x][y] = a.Next(x, y)
		}
	}
}

func ClearScreen() {
	// should check for the os here before doing this
	fmt.Print("\033[H\033[2J") // clear screen for macos
}

func RunGame() {
	generation := 0
	universe := NewUniverse()
	b := NewUniverse()
	Seed(universe, 50)
	for {
		fmt.Println("Generation: ", generation)
		Show(universe)
		time.Sleep(3 * time.Second)
		// ClearScreen()
		fmt.Println()
		Step(universe, b)
		universe = b
		generation++
	}

}

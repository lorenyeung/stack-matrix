package main

import (
	"time"

	"github.com/lorenyeung/stack-matrix/stack"
	log "github.com/sirupsen/logrus"
)

type Matrix [][]int

func main() {
	start := time.Now()
	//take in input or read in default
	size := 5
	m := Matrix{{1, 1, 1, -1, 1},
		{-1, 1, 1, -1, -1},
		{1, 1, 1, -1, 1},
		{-1, 1, 1, 1, 1},
		{1, 1, -1, 1, 1}}
	findPath(size, m)
	duration := time.Since(start)
	log.Info("runtime:", duration)
}

func findPath(size int, m Matrix) {
	//right down left up
	heading := Matrix{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	stack := stack.New()
	stack.Push([]int{0, 0})

	for {
		if stack.Len() == 0 {
			break
		}
		current := stack.Peek()
		stack.Pop()
		currentAssert := current.([]int)

		//mark current position as blocked to prevent backtracking
		m[currentAssert[0]][currentAssert[1]] = -1

		// base case: reached bottom right, can exit
		if currentAssert[0] == size-1 && currentAssert[1] == size-1 {
			log.Info("path found")
			break
		}

		for i := 0; i < 4; i++ {

			// Using the direction array
			x := currentAssert[0] + heading[i][0]
			y := currentAssert[1] + heading[i][1]

			// check next valid move:
			if x > -1 && y > -1 && x < size && y < size && m[x][y] != -1 {
				log.Info("path[", printDirection(i), "]:", x, y, " pos:", currentAssert[0], currentAssert[1])
				if x >= size-1 && y >= size-1 {
					log.Info("path found")
					return
				}
				stack.Push([]int{x, y})
			} else {
				log.Debug("path[", printDirection(i), "] invalid:", x, y, " pos:", currentAssert[0], currentAssert[1])
			}
		}
	}
	log.Info("no path found")
}

func printDirection(dir int) string {
	switch dir {
	case 0:
		return ">"
	case 1:
		return "v"
	case 2:
		return "<"
	case 3:
		return "^"
	}
	return ""
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/julienroland/usg"

	"github.com/lorenyeung/stack-matrix/stack"
	log "github.com/sirupsen/logrus"
)

type Matrix [][]int

func main() {
	start := time.Now()
	//take in input or read in default
	size := 5
	m := Matrix{
		{2, -10, 1, -1, 1},
		{-1, 1, 1, -1, -1},
		{1, 1, 1, -1, 1},
		{-1, -100, -101, -110, -101},
		{1, 1, -1, 2, 1}}
	printProblem(size, m)
	findPath(size, m)
	duration := time.Since(start)
	log.Info("runtime:", duration)
}

func findPath(size int, m Matrix) {

	//right down left up
	var sum = m[0][0]
	heading := Matrix{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	stack := stack.New()
	//x,y,sum at point in time,value, direction
	stack.Push([]int{0, 0, m[0][0], m[0][0], 0})
	log.Info("path[x]:", 0, 0, " pos:", 0, 0, " sum:", sum)

	for {
		if stack.Len() == 0 {
			stack.Push([]int{0, 0, m[0][0], m[0][0], 0})
		}
		current := stack.Peek()
		currentAssert := current.([]int)

		log.Debug("attempt new set")
		for i := 0; i < 4; i++ {

			// Using the direction array
			x := currentAssert[0] + heading[i][0]
			y := currentAssert[1] + heading[i][1]

			// check next valid move:
			if x > -1 && y > -1 && x < size && y < size {
				//potential answer
				if x >= size-1 && y >= size-1 {
					currTop := stack.Peek()
					currSum := currTop.([]int)
					if currSum[2] < 0 {
						log.Warn("sum is less than zero, find another path")

						//pop random lol
						rand.Seed(time.Now().UnixNano())
						min := 1
						max := stack.Len() - 1
						k := rand.Intn(max-min+1) + min

						for j := 0; j < k; j++ {
							stack.Pop()
						}
						//reset track
						top := stack.Peek()
						current := top.([]int)

						//pick new direction
						rand.Seed(time.Now().UnixNano())
						min = 0
						max = 3
						i = rand.Intn(max-min+1) + min
						sum = current[2]
						currentAssert[0] = current[0]
						currentAssert[1] = current[1]
						continue
					} else {
						log.Info("path found:", sum)
						for {
							if stack.Len() == 0 {
								return
							}
							top := stack.Peek()
							topData := top.([]int)
							log.Info(topData)
							stack.Pop()
						}
					}
				} else {
					sum = sum + m[x][y]
					log.Info("path[", printDirection(i), "]:", x, y, " pos:", currentAssert[0], currentAssert[1], " sum:", sum, " size:", stack.Len())
					stack.Push([]int{x, y, sum, m[x][y], i})
					break
				}
			} else {
				//stack.Pop()
				log.Debug("pathInv[", printDirection(i), "]:", x, y, " pos:", currentAssert[0], currentAssert[1], " sum:", sum)
			}
		}
	}
	log.Info("no path found")
}

func printDirection(dir int) string {
	switch dir {
	case 0:
		return usg.Get.ArrowRight
	case 1:
		return usg.Get.ArrowDown
	case 2:
		return usg.Get.ArrowLeft
	case 3:
		return usg.Get.ArrowUp
	}
	return ""
}

func printProblem(size int, m Matrix) {

	for _, v := range m {
		for _, w := range v {
			if w == -1 {
				fmt.Print(usg.Get.SquareSmallFilled)
			} else {
				fmt.Print(usg.Get.SquareSmall)
			}
		}
		fmt.Println()
	}
}

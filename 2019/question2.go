package main

import (
	"fmt"
	"math"
)

func question2() {

	order := 2
	numSquares := int(math.Pow(3, float64(order)))
	midPoint := int(math.Ceil(float64(numSquares / 2)))
	fmt.Printf("Mid cell %v %v\n\n", math.Ceil(float64(numSquares/2)), math.Ceil(float64(numSquares/2)))

	carpet := make([][]int, numSquares)
	for l := 0; l < numSquares; l++ {
		carpet[l] = make([]int, numSquares)
	}

	var points []int

	points = append(points, 1)
	points = append(points, 2)

	points = append(points, 1)
	points = append(points, 0)

	points = append(points, 1)
	points = append(points, 1)

	for i := 0; i < numSquares; i++ {
		for j := 0; j < numSquares; j++ {
			carpet[i][j] = 0
		}
	}

	carpet[midPoint][midPoint] = 1

	switch order {
	case 1:
		carpet[midPoint][midPoint] = 1

	case 2:

		for i := -1; i < 2; i++ {
			carpet[midPoint+i][midPoint-1] = 1
			carpet[midPoint+i][midPoint] = 1
			carpet[midPoint+i][midPoint+1] = 1
		}

	}

	for i := 0; i < numSquares; i++ {
		for j := 0; j < numSquares; j++ {
			fmt.Printf("%d ", carpet[i][j])
			if j == numSquares-1 {
				fmt.Printf("\n")
			}
		}
	}
}

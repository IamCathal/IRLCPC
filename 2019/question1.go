package main

import (
	"fmt"
)

func question1() {
	var pointArr []int
	pointArr = append(pointArr, 0)
	pointArr = append(pointArr, 0)

	pointArr = append(pointArr, 1)
	pointArr = append(pointArr, 0)

	pointArr = append(pointArr, 1)
	pointArr = append(pointArr, 1)

	pointArr = append(pointArr, 0)
	pointArr = append(pointArr, 1)

	for i := 0; i < len(pointArr); i += 2 {
		fmt.Printf("%d %d\n", pointArr[i], pointArr[i+1])
	}

	for k := 0; k < len(pointArr); k += 4 {
		// if x1 - x2 is 0, it's level
		// if y1 - y2 is 0
		if pointArr[k]-pointArr[k+1] == 0 {

		}
	}

	// fmt.Printf("%v", pointArr)
}

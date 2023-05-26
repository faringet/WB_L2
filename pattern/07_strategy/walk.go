package main

import "fmt"

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] Total[%d] time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

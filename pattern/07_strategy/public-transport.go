package main

import "fmt"

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] Total[%d] time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

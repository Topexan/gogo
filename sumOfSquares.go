package main

import (
	"math"
	"time"
)

func sumOfSquares(n int) int {
	res := 0
	sq := make(chan float64)

	for i := 1; i <= n; i++ {
		go func() { sq <- math.Pow(float64(i), 2) }()
		res += int(<-sq)
	}
	time.Sleep(time.Millisecond)
	return res
}

package main

import "fmt"

type MovingAverage struct {
	size int
	elts []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.elts) == this.size {
		this.elts = this.elts[1:]
	}
	this.elts = append(this.elts, val)
	var sum int
	for _, v := range this.elts {
		sum += v
	}
	return float64(sum) / float64(len(this.elts))
}

func main() {
	ma := Constructor(3)
	fmt.Printf("= %f\n", ma.Next(1))
	fmt.Printf("= %f\n", ma.Next(10))
}

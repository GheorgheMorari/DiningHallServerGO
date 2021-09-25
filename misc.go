package main

import "math/rand"

func randU32( max uint32 ) uint32 {
	var a = rand.Uint32()
	a %= max
	return a
}
func randRangeU32( min,max uint32 ) uint32 {
	var a = rand.Uint32()
	a %= (max - min)
	a += min
	return a
}
func MinMax(array []uint32) (uint32, uint32) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

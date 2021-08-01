package main

import (
	"fmt"
)

func removeIndexFromSlice(src []float64, i int) []float64 {
	return append(src[:i], src[i+1:]...)
}

func temperatureGrouping(temps []float64) (res map[int][]float64) {
	res = make(map[int][]float64)
	tempsCopy := make([]float64, len(temps))
	copy(tempsCopy, temps)

	for i := 0; len(tempsCopy) != 0; {
		tempDiv10 := int(tempsCopy[i]) / 10
		if _, ok := res[tempDiv10*10]; !ok {
			for j := i + 1; j < len(tempsCopy); j++ {
				if tempDiv10 == int(tempsCopy[j])/10 {
					res[tempDiv10*10] = append(res[tempDiv10*10], tempsCopy[j])
					tempsCopy = removeIndexFromSlice(tempsCopy, j)
					j--
				}
			}
		}
		res[tempDiv10*10] = append(res[tempDiv10*10], tempsCopy[i])
		tempsCopy = removeIndexFromSlice(tempsCopy, i)
	}
	return
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.5, 3.2, 7.1, -1.4, -3.6}
	tempGroup := temperatureGrouping(temps)
	fmt.Println(tempGroup)
}

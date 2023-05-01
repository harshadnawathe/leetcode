package averagesalaryexcludingtheminimumandmaximumsalary

import "math"

func average(salary []int) float64 {
	min, max := math.MaxInt32, math.MinInt32
	for _, sal := range salary {
		if sal < min {
			min = sal
		}
		if sal > max {
			max = sal
		}
	}

	sum := -(min + max)
	for _, sal := range salary {
		sum += sal
	}

	return float64(sum) / float64(len(salary)-2)
}

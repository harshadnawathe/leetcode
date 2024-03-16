package mycalendariii

import (
	"github.com/emirpasic/gods/maps"
	"github.com/emirpasic/gods/maps/treemap"
)

// Implementation using TreeMap
// This implementation ran slower on leetcode than slice and binary search
type MyCalendarThree struct {
	timestampFreq maps.Map
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		timestampFreq: treemap.NewWithIntComparator(),
	}
}

func (cal *MyCalendarThree) Book(startTime int, endTime int) int {
	insertStartTime(cal.timestampFreq, startTime)
	insertEndTime(cal.timestampFreq, endTime)
	return maxOverlaps(cal.timestampFreq)
}

func insertStartTime(m maps.Map, startTime int) {
	if x, ok := m.Get(startTime); ok {
		m.Put(startTime, x.(int)+1)
	} else {
		m.Put(startTime, 1)
	}
}

func insertEndTime(m maps.Map, endTime int) {
	if x, ok := m.Get(endTime); ok {
		m.Put(endTime, x.(int)-1)
	} else {
		m.Put(endTime, -1)
	}
}

func maxOverlaps(m maps.Map) int {
	max := 0
	curr := 0
	for _, key := range m.Keys() {
		x, _ := m.Get(key)
		curr += x.(int)

		if max < curr {
			max = curr
		}
	}

	return max
}

// // Implmentation using slice and binary search
// type MyCalendarThree struct {
// 	timestampFreq [][2]int
// }
//
// func Constructor() MyCalendarThree {
// 	return MyCalendarThree{}
// }
//
// func (cal *MyCalendarThree) Book(startTime int, endTime int) int {
// 	cal.timestampFreq = insertStartTime(cal.timestampFreq, startTime)
// 	cal.timestampFreq = insertEndTime(cal.timestampFreq, endTime)
//
// 	return maxOverlaps(cal.timestampFreq)
// }
//
// func insertStartTime(timestampFreq [][2]int, startTime int) [][2]int {
// 	x := sort.Search(len(timestampFreq), func(i int) bool { return startTime <= timestampFreq[i][0] })
//
// 	if x == len(timestampFreq) {
// 		return append(timestampFreq, [2]int{startTime, 1})
// 	}
//
// 	if timestampFreq[x][0] == startTime {
// 		timestampFreq[x][1] += 1
// 	} else {
// 		timestampFreq = append(timestampFreq[:x+1], timestampFreq[x:]...)
// 		timestampFreq[x] = [2]int{startTime, 1}
// 	}
//
// 	return timestampFreq
// }
//
// func insertEndTime(timestampFreq [][2]int, endTime int) [][2]int {
// 	x := sort.Search(len(timestampFreq), func(i int) bool { return endTime <= timestampFreq[i][0] })
//
// 	if x == len(timestampFreq) {
// 		return append(timestampFreq, [2]int{endTime, -1})
// 	}
//
// 	if timestampFreq[x][0] == endTime {
// 		timestampFreq[x][1] -= 1
// 	} else {
// 		timestampFreq = append(timestampFreq[:x+1], timestampFreq[x:]...)
// 		timestampFreq[x] = [2]int{endTime, -1}
// 	}
// 	return timestampFreq
// }
//
// func maxOverlaps(timestampFreq [][2]int) int {
// 	max, curr := 0, 0
// 	for _, checkpoint := range timestampFreq {
// 		curr += checkpoint[1]
// 		if curr > max {
// 			max = curr
// 		}
// 	}
// 	return max
// }

package meetingroomsiii

import (
	"container/heap"
	"math"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	var bookedRooms busyRooms

	var availableRooms rooms
	for room := range n {
		heap.Push(&availableRooms, room)
	}

	roomMeetingCount := make([]int, n)

	maxMeetingCount := math.MinInt32

	for _, meeting := range meetings {
		currTime := int64(meeting[0])
		duration := int64(meeting[1] - meeting[0])

		for len(bookedRooms) > 0 && bookedRooms[0].BusyTill <= currTime {
			heap.Push(&availableRooms, bookedRooms[0].Index)
			_ = heap.Pop(&bookedRooms)
		}

		var room int
		if len(availableRooms) == 0 {
			room = bookedRooms[0].Index
			bookedRooms[0].BusyTill += int64(duration)
			heap.Fix(&bookedRooms, 0)
		} else {
			room = availableRooms[0]
			heap.Push(&bookedRooms, busyRoom{room, currTime + duration})
			_ = heap.Pop(&availableRooms)
		}

		roomMeetingCount[room] += 1
		if maxMeetingCount < roomMeetingCount[room] {
			maxMeetingCount = roomMeetingCount[room]
		}
	}

	var result int
	for room := range n {
		if maxMeetingCount == roomMeetingCount[room] {
			result = room
			break
		}
	}

	return result
}

type busyRoom struct {
	Index    int
	BusyTill int64
}

type busyRooms []busyRoom

func (b busyRooms) Len() int {
	return len(b)
}

func (b busyRooms) Less(i, j int) bool {
	if b[i].BusyTill == b[j].BusyTill {
		return b[i].Index < b[j].Index
	}

	return b[i].BusyTill < b[j].BusyTill
}

func (b busyRooms) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *busyRooms) Push(x any) {
	*b = append(*b, x.(busyRoom))
}

func (b *busyRooms) Pop() any {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}

type rooms []int

func (r rooms) Len() int {
	return len(r)
}

func (r rooms) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r rooms) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *rooms) Push(x any) {
	*r = append(*r, x.(int))
}

func (r *rooms) Pop() any {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

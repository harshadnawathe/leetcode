package dividearrayinsetsofkconsecutivenumbers

import "github.com/emirpasic/gods/maps/treemap"

func isPossibleDivide1(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	m := treemap.NewWithIntComparator()

	for _, num := range nums {
		addNum(m, num)
	}

	for m.Size() > 0 {
		it := m.Iterator()
		it.Next()
		num := it.Key().(int)

		removeNum(m, num)

		for i := 1; i < k; i++ {
			if _, ok := m.Get(num + i); !ok {
				return false
			}

			removeNum(m, num+i)
		}
	}

	return true
}

func addNum(m *treemap.Map, num int) {
	if cnt, ok := m.Get(num); ok {
		m.Put(num, cnt.(int)+1)
	} else {
		m.Put(num, 1)
	}
}

func removeNum(m *treemap.Map, num int) {
	if val, ok := m.Get(num); ok {
		cnt := val.(int)
		if cnt > 1 {
			m.Put(num, cnt-1)
		} else {
			m.Remove(num)
		}
	}
}

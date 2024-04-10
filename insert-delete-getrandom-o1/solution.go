package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	numToIndex map[int]int
	nums       []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		numToIndex: make(map[int]int),
		nums:       make([]int, 0),
	}
}

func (set *RandomizedSet) Insert(val int) bool {
	if _, ok := set.numToIndex[val]; ok {
		return false
	}
	set.nums = append(set.nums, val)
	set.numToIndex[val] = len(set.nums) - 1
	return true
}

func (set *RandomizedSet) Remove(val int) bool {
	if _, ok := set.numToIndex[val]; !ok {
		return false
	}
	set.nums[set.numToIndex[val]] = set.nums[len(set.nums)-1]
	set.numToIndex[set.nums[len(set.nums)-1]] = set.numToIndex[val]
	set.nums = set.nums[:len(set.nums)-1]
	delete(set.numToIndex, val)
	return true
}

func (set *RandomizedSet) GetRandom() int {
	return set.nums[rand.Intn(len(set.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

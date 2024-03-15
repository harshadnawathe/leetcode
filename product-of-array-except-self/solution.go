package productofarrayexceptself

// //time: O(n), space: O(1)
// //memory for result is not taken into account
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	leftProduct := 1
	for i := range nums {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// // time: O(n), space: O(n)
// func productExceptSelf(nums []int) []int {
// 	leftProduct := make([]int, len(nums))
// 	leftProduct[0] = 1
// 	for i:=1; i < len(nums); i++ {
//     leftProduct[i] = leftProduct[i-1] * nums[i-1]
//   }
//
// 	rightProduct := make([]int, len(nums))
// 	rightProduct[len(nums)-1] = 1
// 	for i := len(nums)-2; i >= 0; i-- {
// 		rightProduct[i] = rightProduct[i+1] * nums[i+1]
// 	}
//
// 	result := make([]int, len(nums))
// 	for i := range result {
//     result[i] = leftProduct[i] * rightProduct[i]
//   }
//
// 	return result
// }
//

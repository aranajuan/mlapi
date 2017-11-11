package math

// GetStats return statics of array in argumment
// max, min, avg , error(if array if nil or empty)
func GetStats(nums []float32) (float32, float32, float32, int) {
	if nums == nil {
		return 0, 0, 0, 1
	}
	lenN := len(nums)
	if lenN == 0 {
		return 0, 0, 0, 1
	}
	var max, min, sum, c float32
	max = nums[0]
	min = nums[0]

	for i := 0; i < lenN; i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
		c++
	}

	return max, min, sum / c, 0
}

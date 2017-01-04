package pearls

// MaxSubArrayCubed will return the sum of the maximum subarray (subslice)
// or arr. It does so in the brute-force, ϴ(n³) manner of summing up every
// possible sub-array. In *Programming Pearls* this would be an implementation
// of algorithm 1 in Column 8.
func MaxSubArrayCubed(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j <= len(arr); j++ {
			sum := 0
			for _, value := range arr[i:j] {
				sum += value
				if sum > max {
					max = sum
				}
			}
		}

	}
	return max
}

// MaxSubArraySquared will return the maximum sum that can be obtained by
// adding together consecutive entries in arr. It is similar to
// MaxSubArrayCubed but achieves a better run time complexity (ϴ(n²)) by
// hanging onto already completed work in the creation of the sums.
func MaxSubArraySquared(arr []int) int {
	max := 0

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// Max returns the larger of two ints. There is no function in the Go standard
// library that does this for ints.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxSubArrayLinear will return the maximum sum that can be obtained by
// adding together consecutive entries in the arr. It runs in linear (O(n))
// time.
func MaxSubArrayLinear(arr []int) int {
	max := arr[0]
	maxAtCurrent := 0

	for i := 0; i < len(arr); i++ {
		maxAtCurrent = Max(maxAtCurrent+arr[i], 0)
		max = Max(max, maxAtCurrent)
	}

	return max
}

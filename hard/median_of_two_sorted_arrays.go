func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	left, right := -1, -1
	x, y := 0, 0
	for i := 0; i <= length/2; i++ {
		left = right
		if x < m && (y >= n || nums1[x] < nums2[y]) {
			right = nums1[x]
			x++
		} else {
			right = nums2[y]
			y++
		}
	}
	if (length & 1) == 0 {
		return float64(left+right) / 2.0
	} else {
		return float64(right)
	}
}
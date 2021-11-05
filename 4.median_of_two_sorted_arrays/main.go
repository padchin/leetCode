package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	totalLength := length1 + length2
	resArray := make([]int, totalLength)

	j := 0
	k := 0

	for i := 0; i < totalLength; i++ {
		switch {
		case j < length1 && k < length2 && nums1[j] <= nums2[k]:
			resArray[i] = nums1[j]
			j++
		case j >= length1 && k < length2:
			resArray[i] = nums2[k]
			k++
		case j < length1 && k >= length2:
			resArray[i] = nums1[j]
			j++
		case j < length1 && k < length2 && nums1[j] > nums2[k]:
			resArray[i] = nums2[k]
			k++
		}
	}

	if totalLength%2 == 0 {
		return (float64(resArray[totalLength/2-1]) + float64(resArray[totalLength/2])) / 2.0
	}
	return float64(resArray[totalLength/2])
}

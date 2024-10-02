package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create a map to store numbers and their corresponding indices
	numToIndexMap := make(map[int]int)

	// Loop through the array
	for i, num := range nums {
		// Calculate the difference between the target and the current number
		diff := target - num

		// Check if the difference already exists in the map
		if idx, found := numToIndexMap[diff]; found {
			// If it exists, return the indices of the current number and the number that adds up to the target
			return []int{idx, i}
		}

		// If it doesn't exist, add the current number and its index to the map
		numToIndexMap[num] = i
	}

	// If no two numbers add up to the target, return nil
	return nil
}

func main() {
	case1 := []int{2, 7, 11, 15}
	case2 := []int{3, 2, 4}
	case3 := []int{3, 3}
	fmt.Println(twoSum(case1, 9))
	fmt.Println(twoSum(case2, 6))
	fmt.Println(twoSum(case3, 6))
}

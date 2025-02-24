package main

// Checks slice of int if it contains int i, return True of False
func ContainsInt(slice []int, i int) bool {
	res := false
	for _, num := range slice {
		if num == i {
			res = true
			break
		}
	}
	return res
}

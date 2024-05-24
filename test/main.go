package main

func main() {
	_ = isSubsequence2("axc", "ahbgdc")
}
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	t := 1
	b := 1
	result[0] = 1
	for i, v := range nums {
		if i != 0 {
			b *= v
		}
		t *= v
		if i == len(nums)-1 {
			continue
		}
		result[i+1] = t
	}
	result[0] = b
	c := 1
	for i := len(nums) - 2; i > 0; i-- {
		c *= nums[i+1]
		result[i] *= c
	}

	return result
}
func isSubsequence(s string, t string) bool {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		_, ok := m[v]
		if !ok {
			continue
		}
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}

	}
	return true
}
func isSubsequence2(s string, t string) bool {
	result := false
	index := 0
	i := 0
	for j, _ := range s {
		for i = index; i < len(t); i++ {
			if t[i] == s[j] {
				index = i + 1
				if j == len(s)-1 {
					result = true
				}
				break
			}
		}
		if i == len(t)-1 && j < len(s)-1 {
			return false
		}
	}
	return result
}

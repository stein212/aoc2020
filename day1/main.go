package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// twoSum()
	threeSum()
}

func twoSum() {
	reader := bufio.NewReader(os.Stdin)

	m := make(map[int]int)

	var num1 int

	for {
		_, err := fmt.Fscanf(reader, "%d\n", &num1)

		if err != nil {
			fmt.Println(err)
			break
		}

		num2 := 2020 - num1
		if _, ok := m[num2]; ok {

			fmt.Println(num1 * num2)
			break
		}

		m[num1] = 1
	}
}

func threeSum() {
	reader := bufio.NewReader(os.Stdin)

	nums := make([]int, 0, 10)

	for {
		var num int
		_, err := fmt.Fscanf(reader, "%d\n", &num)

		if err != nil {
			break
		}

		nums = append(nums, num)
	}

	sort.Ints(nums)

	for i, num := range nums {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			if num+nums[left]+nums[right] == 2020 {
				fmt.Println(num * nums[left] * nums[right])
				break
			} else if num+nums[left]+nums[right] < 2020 {
				left++
			} else {
				right--
			}
		}
	}
}

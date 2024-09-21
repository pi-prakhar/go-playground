package main

/**
PROBLEM :
leetcode -> medium -> sort colors

TAGS:
arrays, map
**/
func sortColors(nums []int) {
	freq := make(map[int]int)
	color, index := 0, 0
	for _, num := range nums {
		freq[num]++
	}

	for color <= 2 {
		if freq[color] > 0 {
			nums[index] = color
			freq[color]--
			index++
			continue
		}
		color++
	}
	clear(freq)

}

// func sortColors(nums []int)  {
//     count0 := 0
//     count1 := 0
//     count2 := 0
//     index := 0
//     for _, num := range(nums) {
//         switch num {
//             case 0:
//                 count0++;
//             case 1:
//                 count1++;
//             case 2:
//                 count2++;
//             default:
//                 break
//         }
//     }
//     for count0 > 0 {
//         nums[index] = 0
//         count0--
//         index++
//     }
//     for count1 > 0 {
//         nums[index] = 1
//         count1--
//         index++
//     }
//     for count2 > 0 {
//         nums[index] = 2
//         count2--
//         index++
//     }

// }

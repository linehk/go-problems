package design

import (
	"math/rand"
)

type Solution struct {
	resetNums   []int
	shuffleNums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		// resetNums 数组用于保存原数组
		resetNums: nums,
		// 这样可以重新生成一个数组
		shuffleNums: append([]int{}, nums...),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.resetNums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	// 对数组里的每个元素，随机生成一个索引与其交换
	for i := 0; i < len(this.shuffleNums); i++ {
		// 随机生成一个索引
		r := rand.Intn(len(this.shuffleNums))
		// 交换 r 和 i
		this.shuffleNums[r], this.shuffleNums[i] = this.shuffleNums[i], this.shuffleNums[r]
	}
	return this.shuffleNums
}

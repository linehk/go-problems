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
		resetNums:   nums,
		shuffleNums: append([]int{}, nums...),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.resetNums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.shuffleNums); i++ {
		r := rand.Intn(len(this.shuffleNums))
		this.shuffleNums[r], this.shuffleNums[i] =
			this.shuffleNums[i], this.shuffleNums[r]
	}
	return this.shuffleNums
}

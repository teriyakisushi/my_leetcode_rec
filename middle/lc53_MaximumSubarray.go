/*
@ Prolblem: LeeCode-53 Maximum Subarray
@ Solution: 贪心模拟
@ Time Complexity: O(n)
@ Detail:
	找到一个数组中连续的子数组，使得子数组的和最大。
	这题很显然和股票买卖是一样的，都很基础的贪心
*/


func maxSubArray(nums []int) int {
    maxNum := nums[0];
    for i:=1;i<len(nums);i++{
        if nums[i]+nums[i-1]>nums[i]{
            nums[i]+=nums[i-1]
        }
        if nums[i]>maxNum{
        maxNum=nums[i]
    }
    }

    return maxNum
}
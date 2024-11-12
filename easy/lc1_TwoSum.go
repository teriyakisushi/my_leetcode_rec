package main

import "fmt"

/*
@ Problem: Leetcode 1. Two Sum
@ Solution: 哈希表
@ Time Complexity: O(n)
@ Detail:
    考虑到题目要求返回的是下标，很容易想到用哈希表来存储遍历过的元素。
    遍历数组，对于每个元素num，判断target-num是否在哈希表中，如果在，返回两个元素的下标。
    如果不在，将num存入哈希表中。
*/


func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num := range nums {
        if j, ok := hashMap[target-num]; ok {
            return []int{j, i}
        }
        hashMap[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result)
}
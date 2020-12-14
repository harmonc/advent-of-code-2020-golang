package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "sort"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    lines := strings.Split(string(data),"\n")
    nums := make([]int,0)
    for _,value := range(lines){
        num,_ := strconv.Atoi(value)
        nums = append(nums,num) 
    }
    part1ans := part1(nums)
    fmt.Println(part1ans)
    memo := make([]int,nums[len(nums)-1]+1)
    part2Answer := part2(nums, nums[len(nums)-1],memo)
    fmt.Println(len(memo))
    fmt.Println(part2Answer)
}

func part1(nums []int) int {
    ones := 0
    threes := 1
    prev:=0
    sort.Ints(nums)
    for _,value := range(nums){
        if value - prev == 1 {
            ones = ones + 1
        }else if value - prev == 3{
            threes = threes + 1
        }
        prev = value
    }
    fmt.Println(nums)
    return ones * threes
}

func part2(nums []int, value int, memo []int) int{
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    fmt.Println(nums)
    memo[nums[0]] = 1
    for i := 1; i < len(nums); i++{
        fmt.Println(nums[i])
        sum := 0
        n := nums[i]
        if n + 3 < len(memo) {
            sum = sum + memo[n+3]
        }
        if n + 2 < len(memo) {
            sum = sum + memo[n+2]
        }
        if n + 1 < len(memo) {
            sum = sum + memo[n+1]
        }
        memo[n] = sum
    }
    memo[0] = memo[1] + memo[2] + memo[3]
    fmt.Println(memo)
    return memo[0]
}
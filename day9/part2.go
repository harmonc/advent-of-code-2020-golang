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
    fmt.Println(lines)
    fmt.Println(nums)
    
    invalid := firstNonsum(nums,25)
    fmt.Println(invalid)
    ans := findContigousSum(nums,invalid)
    fmt.Println(ans)
    sum := 0
    for _,value := range(ans){
        sum = sum + value
    }
    fmt.Println(sum)
    sort.Ints(ans)
    fmt.Println(ans[0],ans[len(ans)-1],ans[0]+ans[len(ans)-1])
}

func findContigousSum(nums []int, n int) []int {
    terminate := false
    i:=0
    start := 0
    end := 0
    for !terminate{
        curr := i
        sum := nums[i]
        for curr < len(nums) && sum < n {
            curr = curr + 1
            sum = sum + nums[curr]
        }
        if sum == n {
            start = i
            end = curr
            terminate = true
        }
        i = i+1
        if i>len(nums){
            terminate = true
        }
    }
    result := make([]int, 0)
    for j := start; j <= end; j++{
        result = append(result, nums[j])
    }
    return result
}

func firstNonsum(nums []int, n int) int {
    curr := n
    terminate := false
    for !terminate {
        sum := false
        for i := curr-n; i < curr; i++ {
            for j := curr-n; j<curr; j++{
                if i != j{
                    if nums[i] + nums[j] == nums[curr] {
                        sum = true
                    }
                }
            }
        }
        if !sum {
            terminate = true
        }else {
            curr = curr + 1
        }
    }
    return nums[curr]
}
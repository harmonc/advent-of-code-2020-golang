package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

type bag struct {
    color string
    count int
}


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
    fmt.Println(firstNonsum(nums,25))
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
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "sort"
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
    fmt.Println(part1(nums))
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
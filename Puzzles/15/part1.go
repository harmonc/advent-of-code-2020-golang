package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    data, err := ioutil.ReadFile("data.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    strs := strings.Split(string(data),",")
    nums := make([]int64,0)
    for _,value := range(strs){
        n,_ := strconv.Atoi(value)
        nums = append(nums, int64(n))
    }
    result := seq(nums, 2020)
    fmt.Println(result)
}

func seq(nums []int64, turns int64) int64 {
    set := make(map[int64][]int64)
    for i,value := range(nums) {
        ns := make([]int64,0)
        ns = append(ns,int64(i+1))
        set[value] = ns
    }
    turn := int64(len(nums) + 1)
    prev := nums[len(nums)-1]

    for turn <= turns {
        if len(set[prev]) <= 1 {
            prev = 0
        } else {
            prevTurns := set[prev]
            prev = prevTurns[len(prevTurns)-1] - prevTurns[len(prevTurns)-2]
        }
        _, in := set[prev]
        if !in{
            set[prev] = make([]int64,0)
        }
        set[prev] = append(set[prev],int64(turn))
        turn = turn + int64(1)
    }
    return int64(prev)
}
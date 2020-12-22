package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    lines := strings.Split(string(data),"\n")
    fmt.Println(Part1(lines))
}

func Part1(input []string) int {
    t,_ := strconv.Atoi(input[0])
    fmt.Println(t,t+1)
    min := math.MaxInt32
    result := 0
    for _,value := range(strings.Split(input[1],",")){
        fmt.Println(value)
        if strings.Compare(value,"x") == 0{
        }else{
            num,_ := strconv.Atoi(value)
            dist := num - t % num
            if dist < min {
                min = dist
                result = num
            }
        }
    }
    return result * min
}
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
    
    lines := strings.Split(string(data),"\n")
    result := runProgram(lines)
    fmt.Println(result)
}

func runProgram(lines []string) uint64{
    set := make(map[string]uint64)
    mask0 := uint64(0)
    mask1 := uint64(0)
    for _,value := range(lines){
        l := strings.Split(value," = ")
        if strings.Compare(l[0], "mask") == 0 {
            mask0,mask1 = setMask(l[1])
        }else{
            num,_ := strconv.Atoi(l[1])
            set[l[0]] = uint64(num) & ^mask0 | mask1
        }
    }
    fmt.Println(set)
    sum := uint64(0)
    for _,value := range(set){
        sum = sum + value
    }
    return sum
}

func setMask(mask string) (uint64,uint64) {
    c := strings.Split(mask,"")
    exp := uint64(1)
    result0 := uint64(0)
    result1 := uint64(0)
    for i := len(c)-1; i >=0; i--{
        if strings.Compare(c[i],"0") == 0 {
            result0 = result0 + exp
        }else if strings.Compare(c[i],"1") == 0 {
            result1 = result1 + exp
        }
        exp = uint64(2) * exp
    }
    return result0,result1
}
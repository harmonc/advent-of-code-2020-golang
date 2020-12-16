package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    strs := strings.Split(string(data),"\n\n")
    fields := strings.Split(strs[0],"\n")
    ranges := make([][]int,0)
    for _,value := range(fields){
        r := strings.Split(value,": ")[1]
        for _,x := range(strings.Split(r," or ")) {
            nums := strings.Split(x,"-")
            num1,_ := strconv.Atoi(nums[0])
            num2,_ := strconv.Atoi(nums[1])
            range1 := []int{num1,num2}
            ranges = append(ranges,range1)
        }
    }
    fmt.Println(ranges)
    
    tickets := strings.Split(strs[2],"\n")[1:]
    invalid := make([]int,0)
    for _,value := range(tickets){
        for _,num := range(strings.Split(value,",")){
            valid := false
            n,_ := strconv.Atoi(num)
            for _,r := range(ranges){
                if n >= r[0] && n <= r[1] {
                    valid = true
                } 
            }
            if !valid {
                invalid = append(invalid,n)
            }
        }
    }
    fmt.Println(invalid)
    sum := 0
    for _,value := range(invalid){
        sum = sum + value
    }
    fmt.Println(sum)
}
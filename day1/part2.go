package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "sort"
)

func main() {
    fmt.Println("Hello, World!")
    data, err := ioutil.ReadFile("data.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    //fmt.Println("Contents of file:", string(data))
    arr := strings.Split(string(data),"\n")
    //left := make([]int, leftLength)
    numArr := make([]int, len(arr))
    for i, value := range arr {
        num,_ := strconv.Atoi(value)
        numArr[i] = num
    }
    sort.Ints(numArr)
    fmt.Println(numArr)
    num1:=0
    num2:=0
    num3:=0
    for i,_ := range numArr{
        for j,_ := range numArr{
            if(i!=j && numArr[i]+numArr[j]<2020){
                lookFor := 2020-(numArr[i]+numArr[j])
                a := sort.SearchInts(numArr,lookFor)
                if(a!=len(numArr) && numArr[a]==lookFor){
                    num1 = numArr[i]
                    num2 = numArr[j]
                    num3 = lookFor
                }
            }
        }
    }
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(num3)
    fmt.Println(num1+num2+num3)
    fmt.Println(num1*num2*num3)
}
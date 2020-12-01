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
    for _, value := range numArr {
        a:=sort.SearchInts(numArr,2020-value)
        if(numArr[a]==2020-value){
            num1 = value
            num2 = numArr[a]
        }
    }
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(num1*num2)
    fmt.Println(num1+num2)
}
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    arr := strings.Split(string(data),"\n\n")
    result := 0
    for _,value := range(arr){
        result = result + getTotal(value)
    }
    fmt.Println(result)
}

func getTotal(input string) int {
    set := make(map[string]bool)
    for _,value := range(strings.Split(input,"\n")) {
        for _,c := range(strings.Split(value,"")){
            set[c] = true    
        }
    }
    return len(set)
}
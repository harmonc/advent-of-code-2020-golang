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
    players := strings.Split(string(data),"\n\n")

}
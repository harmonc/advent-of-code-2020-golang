package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    //"strconv"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    arr := strings.Split(string(data),"\n")
    result := countTrees(arr, 0, 0, 3, 1)
    fmt.Println(result)
}

func countTrees(grid []string, xPos int, yPos int, dx int, dy int) int {
    result := 0
    x := xPos
    y := yPos
    for y<len(grid){
        if(strings.Compare(string(grid[y][x]),"#")==0){
            result++;
        }
        x += dx
        x = x % len(grid[y])
        y += dy
    }
    return result
}
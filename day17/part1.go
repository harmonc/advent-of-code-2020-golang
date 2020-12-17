package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    data, err := ioutil.ReadFile("data1.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    cubes := make([][]int,0)
    lines := strings.Split(string(data),"\n")
    for i,value := range(lines){
        for j,cell := range(strings.Split(value,"")){
            if strings.Compare(cell,"#") == 0{
                newCube := []int{j,i,0}
                cubes = append(cubes, newCube)
            }
        }
    }
    fmt.Println(cubes)
    fmt.Println(minValues(cubes))
    fmt.Println(maxValues(cubes))
}

func minValues(cubes [][]int) []int{
    result := cubes[0]
    for _,cube := range(cubes){
        for i,value := range(result){
            if cube[i] < value {
                result[i] = cube[i]
            }
        }
    }
    return result
}

func maxValues(cubes [][]int) []int{
    result := cubes[0]
    for _,cube := range(cubes){
        for i,value := range(result){
            if cube[i] > value {
                result[i] = cube[i]
            }
        }
    }
    return result
}
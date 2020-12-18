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
    cubes := make([][]int,0)
    lines := strings.Split(string(data),"\n")
    for i,value := range(lines){
        for j,cell := range(strings.Split(value,"")){
            if strings.Compare(cell,"#") == 0{
                newCube := []int{j,i,0,0}
                cubes = append(cubes, newCube)
            }
        }
    }
    printCubes(cubes)
    fmt.Println(cubes)
    fmt.Println(cubes)
    for i := 0; i < 6; i++{
        min := minValues(cubes)
        max := maxValues(cubes)
        newCubes := make([][]int,0)
        fmt.Println("STATE")
        for w := min[3] - 1; w <= max[3] +1; w++{
            for z := min[2] - 1; z <= max[2] + 1; z++{
                for y := min[1] - 1; y <= max[1] + 1; y++{
                    for x := min[0] - 1; x <= max[0] + 1; x++{
                        test := []int{x,y,z,w}
                        active := isActive(cubes, test)
                        if active{
                            newCubes = append(newCubes,test)
                        }else{
                        }
                    }
                }  
            }
        }
        fmt.Println(newCubes)
        cubes = newCubes
    }
    fmt.Println(len(cubes))
}

func printCubes(cubes [][]int) {
    min := minValues(cubes)
    max := maxValues(cubes)
    fmt.Println(min,max)
    for w := min[3] - 1; w <= max[3] +1; w++{
        fmt.Println("W=")
        for z := min[2] - 1; z <= max[2] + 1; z++{
            fmt.Println("Z=",z)
            for y := min[1] - 1; y <= max[1] + 1; y++{
                for x := min[0] - 1; x <= max[0] + 1; x++{
                    if isInArray(cubes, []int{x,y,z}){
                        fmt.Print("#")
                    }else{
                        fmt.Print(".")
                    }
                }
                fmt.Println("")
            }
        }
    }
}

func isActive(cubes [][]int, c []int) bool {
    neighbors := countNeighbors(cubes, c)
    in := isInArray(cubes, c)
    result := false
    if in && (neighbors == 2 || neighbors == 3){
        result = true
    }
    if !in && neighbors == 3{
        result = true
    }
    return result
}

func countNeighbors(arr [][]int, x []int) int {
    result := 0
    for _,value := range(arr){
        if isNeighbor(value,x){
            result = result + 1
        }
    }
    return result
}

func isNeighbor(a,b []int) bool {
    result := false
    if !isSame(a,b) {
        neighbor := true
        for i,_ := range(a){
            if abs(a[i]-b[i]) > 1 {
                neighbor = false
            }
        }
        if neighbor {
            result = true
        }
    }
    return result
}

func abs(n int) int {
    result := n
    if n < 0 {
        result = -1 * n
    }
    return result
}

func isInArray(arr [][]int, x []int) bool {
    result := false
    for _,e := range(arr){
        if isSame(x,e){
            result = true
        }
    }
    return result
}

func isSame(a,b []int) bool {
    same := true
    if len(a) != len(b) {
        same = false
    }else{
        for i,_ := range(a){
            if a[i] != b[i] {
                same = false
            }
        }
    }
    return same
}

func minValues(cubes [][]int) []int{
    result := make([]int, len(cubes[0]))
    copy(result,cubes[0])
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
    result := make([]int, len(cubes[0]))
    copy(result,cubes[0])
    for _,cube := range(cubes){
        for i,value := range(result){
            if cube[i] > value {
                result[i] = cube[i]
            }
        }
    }
    return result
}
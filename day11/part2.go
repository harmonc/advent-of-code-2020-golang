package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    //"strconv"
    //"sort"
)

type bag struct {
    color string
    count int
}


func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    lines := strings.Split(string(data),"\n")
    grid := make([][]string,0)
    for _,value := range(lines){
        grid = append(grid, strings.Split(value,""))
    }
    
    changed := true
    for changed {
        fmt.Println("grid")
        printGrid(grid)
        grid, changed = run(grid)
        fmt.Println(changed)
    }
    fmt.Println(countOccupied(grid))
}

func run(grid [][]string) ([][]string,bool) {
    changed := false
    result := make([][]string,0)
    for i:=0; i < len(grid); i++{
        row := make([]string,0)
        for j:=0; j < len(grid[i]); j++{
            numOccupied := countOccupiedAdj(grid,j,i)
            if strings.Compare(grid[i][j],"L") == 0 && numOccupied == 0 {
                row = append(row,"#")
                changed = true
            }else if strings.Compare(grid[i][j],"#") == 0 && numOccupied >= 5 {
                row = append(row, "L")
                changed = true
            }else {
                row = append(row, grid[i][j])
            }
        }
        result = append(result,row)
    }
    return result, changed
}

func countOccupiedAdj(grid [][]string, x int, y int) int {
    a := [][]int{{1,0},{1,-1},{1,1},{0,1},{0,-1},{-1,0},{-1,1},{-1,-1}}
    result := 0
    for _,value := range(a){
        search := true
        l := 1
        for search {
            xPos := x + l*value[1]
            yPos := y + l*value[0]
            if onGrid(grid,xPos,yPos){
                if strings.Compare(grid[yPos][xPos],"#") == 0 {
                    result = result + 1
                    search = false
                } else if strings.Compare(grid[yPos][xPos],"L") == 0{
                    search = false
                }
            }else{
                search = false
            }
            l = l + 1
        }
    }
    return result
}

func countOccupied(grid [][]string) int {
    result := 0
    for _,value := range(grid){
        for _,seat := range(value) {
            if strings.Compare(seat, "#") == 0{
                result = result + 1
            }
        }
    }
    return result
}

func onGrid(grid [][]string, x int, y int) bool{
    result := true
    if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]){
        result = false
    }
    return result
}

func printGrid(grid [][]string){
    for _,value := range(grid){
        row := strings.Join(value,"")
        fmt.Println(row)
    }
}
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "sort"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    arr := strings.Split(string(data),"\n")
    fmt.Println(arr)
    seatIDs := make([]int,0)
    for _,value := range(arr){
        seatIDs = append(seatIDs,seatID(value))
    }
    fmt.Println(seatIDs)
    sort.Ints(seatIDs)
    fmt.Println(seatIDs)
    fmt.Println(seatIDs[len(seatIDs)-1])
    
    result := 0
    for i := 0; i < len(seatIDs)-1; i++ {
        if(seatIDs[i+1]-seatIDs[i]==2){
            result = seatIDs[i] + 1
                fmt.Println("GAP",result)

        }
	}
}

func seatID(value string) int {
    result := 0
    row := value[:7]
    col := value[7:]
    low := 0                   
    high := 127
    for _,value := range(strings.Split(row,"")){
        mid := (high-low)/2 + low
        if value == "B" {
            //upper
            low = mid + 1
        } else {
            //lower
            high = mid
        }
    }
    rowResult := high
    if(high != low){
        fmt.Println("HIGH != LOW, SEATID ROW ERROR")
    }
    low = 0
    high = 7
    for _,value := range(strings.Split(col,"")){
        mid := (high-low)/2 + low
        if value == "R" {
            //upper
            low = mid + 1
        } else {
            //lower
            high = mid
        }
    }
    if(high != low){
        fmt.Println("HIGH != LOW, SEATID COL ERROR")
    }
    colResult := high
    result = rowResult * 8 + colResult
    return result
}
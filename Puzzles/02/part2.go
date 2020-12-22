package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    fmt.Println("Hello, World!")
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
    arr := strings.Split(string(data),"\n")
    fmt.Println(arr)
    arr2 := make([][]string,0)
    ranges := make([]string,0)
    letter := make([]string,0)
    counts := make([]int,0)
    passwords := make([]string, 0)
    for _,value := range arr{
        strs := strings.Split(value," ")
        arr2 = append(arr2,strs)
        ranges = append(ranges,strs[0])
        l := strings.Split(strs[1],":")
        num := strings.Count(strs[2],l[0])
        passwords = append(passwords, strs[2])
        counts = append(counts,num)
        letter = append(letter,l[0])
    }
    ranges2 := make([][]int,0)
    for _,value := range ranges{
        r := strings.Split(value,"-")
        num1,_ := strconv.Atoi(r[0])
        num2,_ := strconv.Atoi(r[1])
        curr := []int{num1,num2}
        ranges2 = append(ranges2,curr)
    }
    
    valid := make([]bool,0)
    for i,_ := range(counts){
        b1 := strings.Compare(letter[i],string(passwords[i][ranges2[i][0]-1])) == 0
        b2 := strings.Compare(letter[i],string(passwords[i][ranges2[i][1]-1])) == 0
        v := (b1||b2) && !(b1&&b2)
        valid = append(valid,v)
    } 
    
    fmt.Println(ranges)
    fmt.Println(ranges2)
    fmt.Println(letter)
    fmt.Println(counts)
    fmt.Println(valid)
    
    result := 0
    for _,value := range(valid){
        if value {
            result++        
        }
    }
    fmt.Println(result)
}
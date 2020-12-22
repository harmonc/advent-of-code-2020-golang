package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    data, err := ioutil.ReadFile("data3.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    strs := strings.Split(string(data),"\n\n")
    fields := strings.Split(strs[0],"\n")
    ranges := make([][][]int,0)
    fieldNames := make([]string,0)
    for _,value := range(fields){
        r := strings.Split(value,": ")[1]
        fieldNames = append(fieldNames,strings.Split(value,": ")[0])
        fieldRange := make([][]int,0)
        for _,x := range(strings.Split(r," or ")) {
            nums := strings.Split(x,"-")
            num1,_ := strconv.Atoi(nums[0])
            num2,_ := strconv.Atoi(nums[1])
            range1 := []int{num1,num2}
            fieldRange = append(fieldRange,range1)
        }
        ranges = append(ranges,fieldRange)
    }
    fmt.Println("Field Names:\n",fieldNames,"\nRanges:\n",ranges)
    ticketStrings := strings.Split(strs[2],"\n")[1:]
    ticketData := make([][]int,0)
    for _,ticketString := range(ticketStrings){
        ticket := make([]int,0)
        for _,numString := range(strings.Split(ticketString,",")){
            numInt,_ := strconv.Atoi(numString)
            ticket = append(ticket,numInt)
        }
        ticketData = append(ticketData,ticket)
    }
    //fmt.Println(ticketData)
    validTickets := make([][]int, 0)
    
    //fmt.Println("Checking For Valid Tickets")
    for _,ticket := range(ticketData){
        validTicket := true
        for _,num := range(ticket){
            if !validNum(num,ranges){
                validTicket = false
            }
        }
        if validTicket {
            validTickets = append(validTickets,ticket)
        }
    }

    positions := make([][]int, 0)
    for _,r := range(ranges){
        possiblePositions := make([]int,0)
        for j,_ := range(ranges){
            allMatch := true
            for _,ticket := range(validTickets){
                valid := false
                for _,value := range(r){
                    if ticket[j] >= value[0] && ticket[j] <= value[1] {
                        valid = true
                    }
                }
                if !valid {
                    allMatch = false
                }
            }
            if allMatch{
                possiblePositions = append(possiblePositions,j)
            }
        }
        positions = append(positions,possiblePositions)
    }
    fmt.Println(positions)
    
    
}

func validNum(n int, ranges [][][]int) bool {
    result := false
    for _,pair := range(ranges){
        for _,r := range(pair){
            if n >= r[0] && n <= r[1] {
                result = true
            }
        }
    }
    return result
}
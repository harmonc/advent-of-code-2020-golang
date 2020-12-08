package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
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
    fmt.Println(lines)
    ops := make([]string,0)
    args := make([]int,0)
    for _,value := range(lines) {
        direction := strings.Split(value, " ")
        num,_ := strconv.Atoi(direction[1])
        ops = append(ops, direction[0])
        args = append(args, num)
    }
    fmt.Println(ops,args)
    
    //Brute Force Approach
    
    for i,_ := range(ops){
    
        temp := ops[i]

        if strings.Compare(temp,"acc") != 0{
            if strings.Compare(temp,"jmp") == 0 {
                ops[i] = "nop"
            }else{
                ops[i] = "jmp"
            }
            
            //Run program    
            visited := make([]bool,len(ops))
            curr := 0
            prevCurr := 0
            acc := 0
            terminate := false
            for !terminate {
                visited[curr] = true
                prevCurr = curr
                if strings.Compare(ops[curr],"acc") == 0{
                    acc = acc + args[curr]
                    curr = curr + 1
                }else if strings.Compare(ops[curr],"jmp") == 0{
                    curr = curr + args[curr]
                }else if strings.Compare(ops[curr],"nop") == 0{
                    curr = curr + 1
                }
                if curr >= len(visited){
                    terminate = true
                } else if visited[curr]{
                    terminate = true
                }
            }
            if curr == len(visited){
                fmt.Println(prevCurr,curr,acc)
            }
        }
        
        ops[i] = temp
    }
}


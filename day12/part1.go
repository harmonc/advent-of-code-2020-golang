package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "math"
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
    fmt.Println(part1(lines))
}

func part1(commands []string) int{
    x := 0
    y := 0
    a := float64(0)
    for _,value := range(commands){
        inst := string(value[0])
        num,_ := strconv.Atoi(value[1:])
        fmt.Println(string(inst), num)
        if strings.Compare(inst,"F") == 0{
            fmt.Println("Forward")
            x = x + int(float64(num) * math.Cos(a))
            y = y + int(float64(num) * math.Sin(a))
        }
        if strings.Compare(inst,"R") == 0{
            fmt.Println("Right")
            a = a - radians(num)
        }
        if strings.Compare(inst,"L") == 0{
            fmt.Println("Left")
            a = a + radians(num)
        }
        if strings.Compare(inst,"N") == 0{
            fmt.Println("North")
            y = y + num
        }
        if strings.Compare(inst,"E") == 0{
            fmt.Println("East")
            x = x + num
        }
        if strings.Compare(inst,"S") == 0{
            fmt.Println("South")
            y = y - num
        }
        if strings.Compare(inst,"W") == 0{
            fmt.Println("West")
            x = x - num
        }
        fmt.Println(x,y,"\n")
    }
    return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func radians(deg int) float64{
    return float64(deg) * math.Pi/180
}
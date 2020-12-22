package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "math"
    "strconv"
)

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
    dx := 10
    dy := 1
    for _,value := range(commands){
        inst := string(value[0])
        num,_ := strconv.Atoi(value[1:])
        fmt.Println(string(inst), num)
        if strings.Compare(inst,"F") == 0{
            fmt.Println("Forward")
            x = x + dx * num
            y = y + dy * num
        }
        if strings.Compare(inst,"R") == 0{
            fmt.Println("Right")
            cos := int(math.Cos(radians(-num)))
            sin := int(math.Sin(radians(-num)))
            dxPrev := dx
            dyPrev := dy
            dx = cos * dxPrev - sin * dyPrev
            dy = sin * dxPrev + cos * dyPrev
            
        }
        if strings.Compare(inst,"L") == 0{
            fmt.Println("Left")
            cos := int(math.Cos(radians(num)))
            sin := int(math.Sin(radians(num)))
            dxPrev := dx
            dyPrev := dy
            dx = cos * dxPrev - sin * dyPrev
            dy = sin * dxPrev + cos * dyPrev
        }
        if strings.Compare(inst,"N") == 0{
            fmt.Println("North")
            dy = dy + num
        }
        if strings.Compare(inst,"E") == 0{
            fmt.Println("East")
            dx = dx + num
        }
        if strings.Compare(inst,"S") == 0{
            fmt.Println("South")
            dy = dy - num
        }
        if strings.Compare(inst,"W") == 0{
            fmt.Println("West")
            dx = dx - num
        }
        fmt.Println(x,y,"\n")
    }
    return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func radians(deg int) float64{
    return float64(deg) * math.Pi/180
}
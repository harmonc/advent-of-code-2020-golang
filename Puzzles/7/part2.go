package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "regexp"
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
    
    rules := strings.Split(string(data),"\n")
    set := make(map[string][]bag)

    for i,_ := range(rules){
    rules[i] = strings.ReplaceAll(rules[i], " bags", "")
    rules[i] = strings.ReplaceAll(rules[i], " bag", "")
    
    arr := strings.Split(rules[i],"contain")
    key := strings.Trim(arr[0]," ")
    values := strings.Split(arr[1],",")
    
    bags := make([]bag,0)
    none := strings.Compare(values[0], " no other.")==0
    if !none {
        for _,value := range(values){
            r := regexp.MustCompile("(\\d+) ([\\w\\s]*)")
            matches := r.FindStringSubmatch(value)
            color := matches[2]
            num,_ := strconv.Atoi(matches[1])
            b := bag{color,num}
            bags = append(bags, b)
        }
    }
    set[key] = bags
    }
    
    lookingFor := "shiny gold"
    
    result := howManyBags(set, lookingFor)
    fmt.Println(result)
}

func howManyBags(set map[string][]bag, root string) int {
    result := 0
    for _,value := range(set[root]){
        result = result + value.count + value.count * howManyBags(set, value.color)
    }
    return result
}

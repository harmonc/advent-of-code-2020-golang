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
    
    result := 0
    for key,_ := range(set){
        if(strings.Compare(key,lookingFor)!=0){
            hasIt := canContain(set, key, lookingFor)
            if(hasIt){
                result = result + 1
            }
        }
    }
    fmt.Println(result)
}

func canContain(set map[string][]bag, root string, look string) bool {
    result := false
    if(strings.Compare(root,look)==0){
        result = true
    }else {
        for _,value := range(set[root]){
            result = result || canContain(set, value.color, look)
        }   
    }
    return result
}


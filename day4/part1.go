package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    //"strconv"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    /**
        byr (Birth Year)
        iyr (Issue Year)
        eyr (Expiration Year)
        hgt (Height)
        hcl (Hair Color)
        ecl (Eye Color)
        pid (Passport ID)
        cid (Country ID) can ignore
    **/
    
    creds := [7]string{"byr","iyr","eyr","hgt","hcl","ecl","pid"}
    m := make(map[string]bool)
    
    arr := strings.Split(string(data),"\n\n")
    
    var result = 0
    
    for i,_ := range(arr){
        for _,value := range(creds){
            m[value] = false
        }
        arr[i] = strings.ReplaceAll(arr[i],"\n"," ")
        fmt.Println(i,arr[i])
        test := strings.Split(arr[i]," ")
        for _,value := range(test){
            field := strings.Split(value,":")
            m[field[0]] = true
        }
        var valid = true
        for _,value := range(m){
            if(!value){
                valid = false
            }
        }
        if(valid){
            result = result + 1
        }
    }
    fmt.Println(result)
}
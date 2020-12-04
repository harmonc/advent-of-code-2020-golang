package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "regexp"
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
        test := strings.Split(arr[i]," ")
        for _,value := range(test){
            field := strings.Split(value,":")
            if validFunc(field[0],field[1]) {
                m[field[0]] = true
            } else {
                m[field[0]] = false
            }
        }
        var valid = true
        for _,value := range(m){
            if(!value){
                valid = false
            }
        }
        if(valid){
            fmt.Println(i,arr[i])
            result = result + 1
        }
    }
    fmt.Println(result)
    fmt.Println(validFunc("pid","23456789"))
}

func validFunc(field, value string) bool {
    result := true
    switch field {
        case "byr":
//        fmt.Println("byr",value)
            num,_ := strconv.Atoi(value)
            if num > 2002 || num < 1920 {
                result = false
            }
        case "iyr":
//        fmt.Println("iyr",value)
            num,_ := strconv.Atoi(value)
            if num < 2010 || num > 2020 {
                result = false
            }
        case "eyr":    
//        fmt.Println("eyr",value)
            num,_ := strconv.Atoi(value)
            if num < 2020 || num > 2030 {
                result = false
            }
        case "hgt":
//            fmt.Println("hgt:",value)
            unit := value[len(value)-2:len(value)]
            num,_ := strconv.Atoi(value[0:len(value)-2])
            if unit == "cm" && (num < 150 || num > 193) {
                result = false
            }
            if unit == "in" && (num < 59 || num > 76){
                result = false
            }
            if unit != "in" && unit != "cm" {
                result = false
            }
        case "hcl":
//            fmt.Println("hcl",value)
            matched, _ := regexp.MatchString(`#[a-f0-9]{6}`, value)
            if !matched {
                result = false
            }
        case "ecl":
//            fmt.Println("ecl:",value)
            result = value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth"
        case "pid":
//            fmt.Println("pid:",value)
            matched, _ := regexp.MatchString(`^[0-9]{9}$`, value)
            if !matched {
                result = false
            }
    }
//    fmt.Println(result)
    return result
}
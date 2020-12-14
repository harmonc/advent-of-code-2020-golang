package main
 
import (
    "fmt"
    "math/big"
    "strings"
    "io/ioutil"
    "strconv"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    
    lines := strings.Split(string(data),"\n")
    
    n := make([]*big.Int,0)
    a := make([]*big.Int,0)
    
    for i,value := range(strings.Split(lines[1],",")){
        fmt.Println(i,value)
        if strings.Compare(value,"x") == 0 {
        }else{
            num,_ := strconv.Atoi(value)
            if i == 0{
                a = append(a, big.NewInt(int64(i)))
            }else {
                a = append(a, big.NewInt(int64(num-i)))
            }
            n = append(n, big.NewInt(int64(num)))
        }
    }
        //big.NewInt(2),
    fmt.Println(a)
    fmt.Println(n)
    result,_ := crt(a, n)
    fmt.Println(result.Text(10))
}

//Chinese Remainder Theorem Code from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
//https://en.wikipedia.org/wiki/Chinese_remainder_theorem
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
    p := new(big.Int).Set(n[0])
    for _, n1 := range n[1:] {
        p.Mul(p, n1)
    }
    var x, q, s, z big.Int
    for i, n1 := range n {
        q.Div(p, n1)
        z.GCD(nil, &s, n1, &q)
        if z.Cmp(one) != 0 {
            return nil, fmt.Errorf("%d not coprime", n1)
        }
        x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
    }
    return x.Mod(&x, p), nil
}
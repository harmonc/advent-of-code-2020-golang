package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    data, err := ioutil.ReadFile("data2.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    players := strings.Split(string(data),"\n\n")
    player1 := make([]int,0)
    player2 := make([]int,0)
    loadPlayer(&player1,players[0])
    loadPlayer(&player2,players[1])
    i := 1
    for len(player1) > 0 && len(player2) > 0 {
        fmt.Println("Round",i)
        playRound(&player1,&player2)
        i = i + 1
    }
    fmt.Println("--Results--")
    fmt.Println("Player 1 Deck:",player1)
    fmt.Println("Player 2 Deck:",player2)
    fmt.Println("Score:",score(player1,player2))
}

func loadPlayer(player *[]int, data string){
    for i,value := range(strings.Split(data,"\n")){
        if i != 0 {
            num,_ := strconv.Atoi(value)
            *player = append(*player,num)
        }
    }
}

func score(p1,p2 []int) int{
    winningPlayer := p1
    if len(p1) == 0 {
        winningPlayer = p2
    }
    sum := 0
    for i,value := range(winningPlayer){
        sum = sum + value * (len(winningPlayer) - i)
    }
    return sum
}

func playRound(player1 *[]int, player2 *[]int){
    fmt.Println("Player 1 Deck:",*player1)
    fmt.Println("Player 2 Deck:",*player2)
    p1 := (*player1)[0]
    p2 := (*player2)[0]
    fmt.Println("Player 1 Plays",p1)
    fmt.Println("Player 2 Plays",p2)
    if p1 > p2 {
        fmt.Println("Player 1 Wins")
        *player2 = (*player2)[1:]
        *player1 = append((*player1)[1:],p1,p2)
    }else {
        fmt.Println("Player 2 Wins")
        *player1 = (*player1)[1:]
        *player2 = append((*player2)[1:],p2,p1)
    }
}
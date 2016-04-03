package main

import (
    . "github.com/Cepave/query/elus"
    "fmt"
)

func main(){
    enp := []string{"docker-agent"}
    con := []string{"cpu.idle"}
    x := Query(enp, con)
    fmt.Printf("%v", x)
}

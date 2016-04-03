package main

import (
    "encoding/json"
    "fmt"
)

type QueryString struct {
    Endpoints []string
    Counters []string
}

func main() {
    var q *QueryString
    decodeStr := "{\"endpoints\":[\"docker-agent\",\"docker-task\"],\"counters\":[\"cpu.idle\"]}"
    err := json.Unmarshal([]byte(decodeStr), &q)
    if err != nil {
        fmt.Printf("%v", err.Error())
    }
    fmt.Printf("%s", q.Endpoints)
}

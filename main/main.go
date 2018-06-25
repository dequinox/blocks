package main

import (
    "fmt"
    "blocks/api"
)

func main() {
    /*_, blocks := api.Collect("../data/test2", 51588663)
    err := api.Combine(blocks, "generated.mp4")
    if err != nil {
        fmt.Println(err)
    }*/
    nBlocks, err := api.Fragment("../data/test1")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(nBlocks)
}

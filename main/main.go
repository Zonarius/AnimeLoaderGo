package main

import(
    "fmt"
    "github.com/zonarius/AnimeLoaderGo"
)

func main() {
    items, err := AnimeLoaderGo.Fetch("baked")
    if err != nil {
        panic(err)
    }
    for _, item := range items {
        fmt.Printf("%+v\n", item)
    }
}
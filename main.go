package main

import "fmt"
import "math/rand"
import "sort"

func shuffle(list []int) {
    for i := len(list); i > 1; i-- {
        j := rand.Intn(i)
        list[i - 1], list[j] = list[j], list[i - 1]
    }
}

func main() {
    list := make([]int, 10, 10)
    for i := 0; i < 10; i++ {
        list[i] = i + 1
    }
    shuffle(list)
    fmt.Println(list)
    sort.Sort(list)
    fmt.Println(list)
}


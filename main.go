package main

import "fmt"
import "math/rand"
import "sort"
import "time"

func shuffle(list []int) {
    for i := len(list); i > 1; i-- {
        j := rand.Intn(i)
        list[i - 1], list[j] = list[j], list[i - 1]
    }
}

func make_list(n int) []int{
    list := make([]int, n, n)
    for i := 0; i < n; i++ {
        list[i] = i + 1
    }
    shuffle(list)
    return list
}

func main() {
    list := make_list(100)
    fmt.Println(list)

    start := time.Now()
    sort.Sort(sort.IntSlice(list))
    end := time.Now()

    fmt.Println(list)

    fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}


package util

import "math/rand"
import "time"

func Shuffle(list []int) {
    rand.Seed(time.Now().UnixNano())
    for i := len(list); i > 1; i-- {
        j := rand.Intn(i)
        list[i - 1], list[j] = list[j], list[i - 1]
    }
}

func MakeList(n int) []int{
    list := make([]int, n, n)
    for i := 0; i < n; i++ {
        list[i] = i + 1
    }
    Shuffle(list)
    return list
}


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

func bubble_sort(list []int) {
    length := len(list)
    for i := 0; i < length - 1; i++ {
        for j := 1; j < length - i; j++ {
            if (list[j] < list[j - 1]) {
                tmp := list[j - 1]
                list[j - 1] = list[j]
                list[j] = tmp
            }
        }
    }
}

func insertion_sort(list []int) {
    length := len(list)
    for i := 1; i < length; i++ {
        tmp := list[i]
        if (list[i - 1] > tmp) {
            j := i
            for ;j > 0 && list[j - 1] > tmp; {
                list[j] = list[j - 1]
                j--
            }
            list[j] = tmp
        }
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

func func_sort(list []int, fn func([]int)) {
    fmt.Println(list)

    start := time.Now()
    fn(list)
    end := time.Now()

    fmt.Println(list)
    fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func standard_sort(list []int) {
    sort.Sort(sort.IntSlice(list))
}

func main() {
    length := 10
    list := make_list(length)
    list_bk := make([]int, length)

    copy(list_bk, list)
    func_sort(list, bubble_sort)

    copy(list, list_bk)
    func_sort(list, insertion_sort)

    copy(list, list_bk)
    func_sort(list, standard_sort)
}


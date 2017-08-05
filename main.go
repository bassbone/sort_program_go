package main

import "fmt"
import "sort"
import "time"
import "./algorithm"
import "./util"

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
    list := util.MakeList(length)
    list_bk := make([]int, length)

    copy(list_bk, list)
    func_sort(list, algorithm.BubbleSort)

    copy(list, list_bk)
    func_sort(list, algorithm.InsertionSort)

    copy(list, list_bk)
    func_sort(list, standard_sort)
}


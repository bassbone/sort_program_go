package algorithm

func InsertionSort(list []int) {
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


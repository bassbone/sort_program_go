package algorithm

func BubbleSort(list []int) {
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


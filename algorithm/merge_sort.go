package algorithm

func MergeSort(list []int) {
    if (!len(list)) {
        return
    } elseif (len(list) == 1) {

    } elseif (len(list) == 2) {
        if (list[0] > list[1]) {
            tmp := list[1]
            list[1] = list[0]
            list[0] = tmp
        }
    } else {
        min_num := int(len(list) / 2)
        arr1 = MergeSort()
        arr2 = MergeSort()
        list = 
    }
}


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
        mid_num := int(len(list) / 2)
        arr1 := make([]int, mid_num + 1, mid_num + 1)
        arr2 := make([]int, len(list) - mid_num, len(list) - mid_num)
        arr1 = list[0:mid_num]
        arr2 = list[mid_num,len(list) - 1]
        MergeSort(arr1)
        MergeSort(arr2)
        list = merge_arr(arr1, arr2)
    }
}

func merge_arr(arr1 []int, arr2 []int) []int{
    list := make([]int, len(arr1) + len(arr2))
    while (len(arr1) || len(arr2)) {
        if (len(arr1) == 0) {
            

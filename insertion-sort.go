package main

import (
    "fmt"
)

// Sorts an array in increasing order.
func insertionSort(a []int) {
    for j := 1; j < len(a); j++ {
        key := a[j]
        i := j - 1
        for ; i >= 0 && a[i] > key; i-- {
            a[i + 1] = a[i]
        }
        a[i + 1] = key
    }
    return
}

func main() {
    a := []int{5, 2, 4, 6, 1, 3}
    insertionSort(a)
    fmt.Println(a)
}

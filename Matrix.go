package main

import (
    "fmt"
)

func matrixMultiplication(a [][]int, b [][]int) [][]int {
    if len(a) == 0 || len(b) == 0 {
        return nil
    }

    n, m, p := len(a), len(b), len(b[0])
    if len(a[0]) != m {
        return nil
    }

    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, p)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < p; j++ {
            for k := 0; k < m; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }

    return result
}

func main() {
    a := [][]int{{1, 2, 3}, {4, 5, 6}}
    b := [][]int{{7, 8}, {9, 10}, {11, 12}}

    result := matrixMultiplication(a, b)
    if result == nil {
        fmt.Println("Invalid matrices")
        return
    }

    fmt.Println(result)
}

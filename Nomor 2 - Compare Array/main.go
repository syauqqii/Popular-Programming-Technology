package main

import "fmt"

func compareArray(arrayA []string, arrayB []string, size int) bool {
        isEqual := true
        for i := 0; i < size; i++ {
                if arrayA[i] != arrayB[i] {
                        isEqual = false
                        break
                }
        }

        return isEqual
}

func main() {
    size := 0

    fmt.Print("\n > Masukkan ukuran array: ")
        fmt.Scanln(&size)

        arrayA := make([]string, size)
        arrayB := make([]string, size)

        fmt.Print("\n > Array 1: ")
        for i := 0; i < size; i++ {
                fmt.Scan(&arrayA[i])
        }

        fmt.Print(" > Array 2: ")
        for i := 0; i < size; i++ {
                fmt.Scan(&arrayB[i])
        }

        if compareArray(arrayA, arrayB, size) {
                fmt.Println("\n > Hasil: Kedua array memiliki isi yang sama!")
        } else {
                fmt.Println("\n > Hasil:")
                for i := 0; i < size; i++ {
                    if arrayA[i] != arrayB[i] {
                        fmt.Printf("   - Index ke-%d berbeda\n", i)
                    }
                }
        }
}

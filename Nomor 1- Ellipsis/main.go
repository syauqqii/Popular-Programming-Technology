package main

import "fmt"

func hitungTotal(items ...int) float64{
        total := 0

        for _, item := range items{
                total += item
        }

        return float64(total)
}

func main(){
        sabunMandi := 10000
        sabunMuka  := 30000

        jumlahBarang1 := 3
        jumlahBarang2 := 1

        totalHarga := hitungTotal(sabunMandi*jumlahBarang1, sabunMuka*jumlahBarang2)
        fmt.Printf("Total harga: %.2f\n", totalHarga)
}

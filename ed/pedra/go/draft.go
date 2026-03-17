package main
import (
    "fmt"
    "math"
)
func main() {
    var N int
    fmt.Scan(&N)

    indVencedor := -1
    menDiferenca := 101

    for i := 0; i < N; i++ {
        var a, b int
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10 {
            diferenca := int(math.Abs(float64(a - b)))

            if indVencedor == -1 || diferenca < menDiferenca{
                indVencedor = i
                menDiferenca = diferenca
            }
        }
    }

    if indVencedor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(indVencedor)
    }
}

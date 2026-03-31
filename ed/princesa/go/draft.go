package main
import "fmt"
func main() {
    var n, e int
    fmt.Scanln(&n, &e)

    pessoas := make([]int, n)
    for i := 0; i < n; i++ {
        pessoas[i] = i+1
    }

    tamanho := n
    indice := e-1

    for tamanho > 1 {
        fmt.Print("[ ")
        for i := 0; i < tamanho; i++ {
            if i == indice {
                fmt.Printf("%d> ", pessoas[i])
            } else {
                fmt.Printf("%d ", pessoas[i])
            }
        }
        fmt.Println("]")

        vitima := (indice+1)%tamanho

        for i := vitima; i < tamanho-1; i++ {
            pessoas[i] = pessoas[1+i]
        }

        tamanho--
        indice = vitima
        if indice >= tamanho {
            indice = 0
        }
    }

    fmt.Printf("[ %d> ]\n", pessoas[0])
}
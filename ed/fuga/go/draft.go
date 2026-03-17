package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)

    posicaoAtual := F

    for{
        posicaoAtual += D
        if posicaoAtual < 0 {
            posicaoAtual = 15
        } else if posicaoAtual > 15 {
            posicaoAtual = 0
        }

        if posicaoAtual == H {
            fmt.Println("S")
            break
        }
        if posicaoAtual == P {
            fmt.Println("N")
            break
        }
    }
}

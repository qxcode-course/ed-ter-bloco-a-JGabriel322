package main
import "fmt"

func num_resto(n int) {
    if n == 0 {
        return
    }

    r := n%2
    q := n/2

    num_resto(q)
    fmt.Printf("%d %d\n", q, r)

}
func main() {
    var n int
    fmt.Scan(&n)

    num_resto(n)
}

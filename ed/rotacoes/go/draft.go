package main
import "fmt"
func main() {
    var t, r int
    fmt.Scan(&t, &r)

    rotaciona := make([]int, t)
    for i := 0; i < t; i++ {
        fmt.Scan(&rotaciona[i])
    }

    result := make([]int, t)

    for i := 0; i < t; i++ {
        novaPos := (i+r)%t
        result[novaPos] = rotaciona[i]
    }

    fmt.Print("[ ")
    for i := 0; i < t; i++ {
        if i == t-1 {
			fmt.Print(result[i])
		} else {
			fmt.Print(result[i], " ")
		}
	}

	fmt.Println(" ]")
    
}

package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x <= 1 {
		return false
	}

	if div*div > x {
		return true
	}

	if x%div == 0 {
		return false
	}

	return eh_primo(x, div+1)
}

func gerar_primos(x int) []int {
    primos := []int{}
    num := 2

    for len(primos) < 2 {
        if eh_primo(num, 2){
            primos = append(primos, num)
        }

        num++
    }

    return primos
}

func main() {
	var x int
	fmt.Scan(&x)

    result := gerar_primos(x)
    fmt.Println(result)
}
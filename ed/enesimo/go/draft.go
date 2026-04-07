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

func e_nesimo(x int, atual int) int {
    if eh_primo(atual, 2){
        if x == 1 {
            return atual  
        }

        return e_nesimo(x-1, atual+1)
    }

    return e_nesimo(x, atual+1)
}

func main() {
	var x int
	fmt.Scan(&x)

    result := e_nesimo(x, 2)
    fmt.Println(result)
}

package main
import "fmt"

func matchingStrings (strings []string, consultas []string ) []int {
    cont := make(map[string]int)
    for i := 0; i < len(strings); i++ {
        cont[strings[i]] = cont[strings[i]]+1
    }

    result := make([]int, len(consultas))
    for i := 0; i < len(consultas); i++ {
        result[i] = cont[consultas[i]]
    }

    return result
}

func main() {
    var n int
    fmt.Scan(&n)

    strings := make([]string, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&strings[i])
    }

    var m int
    fmt.Scan(&m)

    consultas := make([]string, m)
    for i := 0; i < m; i++ {
        fmt.Scan(&consultas[i])
    }

    result := matchingStrings(strings, consultas)
    for i := 0; i < len(result); i++ {
        if i == len(result)-1 {
            fmt.Print(result[i])
        } else {
            fmt.Print(result[i], " ")
        }
    }

    fmt.Println()
}

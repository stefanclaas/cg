package main

import "fmt"

func main() {
    for i := 0; i < 26*26*26*26*26; i++ {
        fmt.Printf("%c%c%c%c%c\n", 'A'+i/26/26/26/26%26, 'A'+i/26/26/26%26, 'A'+i/26/26%26, 'A'+i/26%26, 'A'+i%26)
    }
}

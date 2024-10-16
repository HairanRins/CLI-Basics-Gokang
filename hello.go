package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <nom> <nombre>")
        return
    }

    name := os.Args[1]
    timesStr := os.Args[2]

    times, err := strconv.Atoi(timesStr)
    if err != nil {
        fmt.Println("Veuillez entrer un nombre valide.")
        return
    }

    for i := 0; i < times; i++ {
        fmt.Printf("Bonjour, %s! Ceci est le message numéro %d\n", name, i+1)
    }
}

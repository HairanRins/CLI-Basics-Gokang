package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run transforme.go <mode> <input>")
		fmt.Println("Modes disponibles: string (Transformation en majuscules) | hex (conversion en héxadécimal)")
		return
	}
	
	mode := os.Args[1]
	input := os.Args[2]

	switch mode {
		case "string":
			upperStr := strings.ToUpper(input)
			fmt.Printf("Chaine transformée en majuscules : %s\n", upperStr)
		case "hex":
			decimal, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Veuillez entrer un nombre valide pour la conversion hexadécimale")
				return
			}
			hexValue := strconv.FormatInt(int64(decimal), 16)
			fmt.Printf("Valeur hexadécimale de %d est: %s\n", decimal, hexValue)
		default:
			fmt.Println("Mode invalide. Utilisez 'string' ou 'hex'")
	}
}


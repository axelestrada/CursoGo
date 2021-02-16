package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Cual es su nombre")

	reader := bufio.NewReader(os.Stdin)

	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimRight(nombre, "\r\n")

	fmt.Println("Cual es su edad")

	edad, _ := reader.ReadString('\n')
	edad = strings.TrimRight(edad, "\r\n")

	fmt.Println("Cual es su peso en libras")

	peso, _ := reader.ReadString('\n')
	peso = strings.TrimRight(peso, "\r\n")

	fmt.Println("Nombre:", nombre, "Edad:", edad, "Peso:", peso)
}

// 9670-3730

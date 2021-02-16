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

	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimRight(entrada, "\r\n")

	fmt.Println("Nombre:", entrada)
}

// 9670-3730

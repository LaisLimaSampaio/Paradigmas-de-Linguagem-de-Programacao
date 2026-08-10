package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World");
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um número: ")
	input, _ := reader.ReadString('\n')

	var numero int
	fmt.Sscanf(input, "%d", &numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}

package main

import (
	"fmt"

	"github.com/goWorkspace/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1933)
	fmt.Println(estado)
	fmt.Println(texto)
}

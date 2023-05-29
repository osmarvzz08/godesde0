package main

import (
	"fmt"
	"github/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvierteaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}

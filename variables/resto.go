package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Osmar"
	Estado = true
	Sueldo = 1600.13
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvierteaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}

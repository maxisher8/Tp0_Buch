package main
import (
	"os"
	"bufio"
	"fmt"
	"tp0/ejercicios"
	"strconv"
)
const (
	RUTA_ARCHIVO_1 = "archivo1.in"
	RUTA_ARCHIVO_2 = "archivo2.in"
)

func lectura(ruta string) []int{
	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	vector := []int{}
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		contenido, _ := strconv.Atoi(lectura.Text()) 
		vector = append(vector, contenido)
	}
	return vector
}

func main() {
	vector1 := lectura(RUTA_ARCHIVO_1)
	vector2 := lectura(RUTA_ARCHIVO_2)

	igualdad := ejercicios.Comparar(vector1, vector2)
	
	if igualdad >= 0 {
		ejercicios.Seleccion(vector1)
		for i := 0; i < len(vector1); i++ {
			fmt.Println(vector1[i])
		}
	} else {
		ejercicios.Seleccion(vector2)
		for i := 0; i < len(vector2); i++ {
			fmt.Println(vector2[i])
		}
	}
}

package main

// Importar librerias necesarias
import (
    "fmt" // Libreria necesaria para usar funcion Print
)

//Comienza el programa
func main() {

    //LOOP 100 VESES
    
    //con la siguiente linea, el programa preguntara 100 veses
    //for numero_de_preguntas := 1; numero_de_preguntas <= 100; numero_de_preguntas++ {

    //Pero con esta linea, seran preguntas infinitas, a menos que 
    //el usuario escoja 0
    //comenta esta linea y descomenta el for-loop de arriba
    //para seleccionar cuantas preguntas serian
    for { //LOOP INFINITO
	fmt.Println()

	//si se escoge LOOP INFINITO de arriba, debemos comentar esta linea
	//fmt.Println("Loop Numero ", numero_de_preguntas)

	fmt.Println("Menu Principal")
	fmt.Println()
	fmt.Println("	0 Salir")
	fmt.Println("	1 Multiplicar")
	fmt.Println("	2 Dividir")
	fmt.Println("	3 Sumar")
	fmt.Println("	4 Restar")

	fmt.Println()
	fmt.Print("Escribe tu Opcion? ")

	//Definir variable para seleccion de opcion
	var opcion int

	//preguntale al usuario la opcion y 
	//asignala a mi variable "opcion"
	fmt.Scan(&opcion)

	//verifica si la opcion es valida
	if opcion == 0 {
		//es mayor de 0, es valida
		fmt.Println("Programa Terminado!")
		break
	}

	if opcion == 1 {
		//es 1, osea es multiplicacion
		fmt.Println("Multiplicar!")
		
		//preguntale al usuario por los 2 
		//valores a multiplicar
		var valor1 int
		var valor2 int

		//pregunta por primer valor
		fmt.Print("Dame el primer valor? ")
		fmt.Scan(&valor1)
		fmt.Println()

		//pregunta por segundo valor
		fmt.Print("Dame el segundo valor? ")
		fmt.Scan(&valor2)
		fmt.Println()

		//imprimimos el resultado al usuario
		fmt.Println("Tu respuesta de ", valor1, " X ", valor2, " = ",valor1*valor2)
	}

	if opcion == 2 {
		//es 2, osea es Divicion
		fmt.Println("Dividir!")


		//preguntale al usuario por los 2 
		//valores a Sumar
		var valor1 float64
		var valor2 float64

		//pregunta por primer valor
		fmt.Print("Dame el primer valor? ")
		fmt.Scan(&valor1)
		fmt.Println()

		//pregunta por segundo valor
		fmt.Print("Dame el segundo valor? ")
		fmt.Scan(&valor2)
		fmt.Println()

		//imprimimos el resultado al usuario
		fmt.Println("Tu division de ", valor1, " / ", valor2, " = ",valor1/valor2)
	}

	if opcion == 3 {
		//es 3, osea es Suma
		fmt.Println("Sumar!")

		//preguntale al usuario por los 2 
		//valores a Sumar
		var valor1 int
		var valor2 int

		//pregunta por primer valor
		fmt.Print("Dame el primer valor? ")
		fmt.Scan(&valor1)
		fmt.Println()

		//pregunta por segundo valor
		fmt.Print("Dame el segundo valor? ")
		fmt.Scan(&valor2)
		fmt.Println()

		//imprimimos el resultado al usuario
		fmt.Println("Tu suma de ", valor1, " + ", valor2, " = ",valor1+valor2)
	}

	if opcion == 4 {
		//es 4, osea es resta
		fmt.Println("Restar!")


		//preguntale al usuario por los 2 
		//valores a Sumar
		var valor1 int
		var valor2 int

		//pregunta por primer valor
		fmt.Print("Dame el primer valor? ")
		fmt.Scan(&valor1)
		fmt.Println()

		//pregunta por segundo valor
		fmt.Print("Dame el segundo valor? ")
		fmt.Scan(&valor2)
		fmt.Println()

		//imprimimos el resultado al usuario
		fmt.Println("Tu Resta de ", valor1, " - ", valor2, " = ",valor1-valor2)
	}

	if opcion < 0 {
		//es negativo, es invalido
		fmt.Println("Esta opcion no esta disponible, porque es negativo")
	}

	if opcion > 4 {
		//nomas tenemos hasta el 4 de opciones, osea es invalido
		fmt.Println("Esta opcion no esta disponible, escoge entre 0 y 4")
	}
     
     } //Fin de mi for-loop de numero_de_preguntas

} //Aqui termina programa

package main

import (
	"errors"
	"fmt"
)

func main() {
	//suma
	var food [3]string
	food[0] = "banana"
	food[1] = "pizza"
	food[2] = "piña"
	/*names := [6]string{"jose", "marco", "luis", "lucia", "bertha", "luciana"}
	men := names[:3]
	woman := names[3:]
	men[1] = "julio"


	fmt.Println(men)
	fmt.Println(woman)

	fruits := []string{"apple", "pine"}
	fruits = append(fruits, "peach", "watermelon", "carrot", "tomato")
	fmt.Println("fruits: ", (fruits))
	fmt.Println("fruits: ", len(fruits))
	fmt.Println("fruits: ", cap(fruits))*/

	//mapas
	/*animals := make(map[string]string)
	animals["cat"] = "cat"
	animals["dog"] = "dog"
	animals["duck"] = "duck"
	animals["elephant"] = "elephant"
	fmt.Println(animals)
	//mapas literales
	vegetables := map[string]string{
		"apple":  "apple",
		"banana": "banana",
	}

	delete(vegetables, "banana")
	_, ok := vegetables["carrot"]
	if !ok {
		vegetables["carrot"] = "carrot"
	}
	fmt.Println(vegetables)*/
	/*type countrys struct {
		Name      string
		latitude  float32
		altitude  float32
		continent string
	}
	db := countrys{
		Name:      "Belgium",
		altitude:  475521,
		latitude:  478557,
		continent: "EU",
	}
	peru := countrys{
		"Peru",
		345345,
		434123,
		"EU",
	}
	bolivia := countrys{
		Name: "Bol",
	}
	p := &peru
	(*p).latitude = 45465
	p.latitude = 234
	fmt.Printf("%+v \n", db.continent)
	fmt.Printf("%+v \n", peru.latitude)
	fmt.Printf("%+v \n", bolivia.Name)*/

	/*type circle struct {
		radio     float32
		diametro  float32
		perimetro float32
		area      float32
	}
	C1 := circle{
		radio:     5,
		diametro:  0,
		perimetro: 0,
		area:      0,
	}
	C1.diametro = C1.radio * 2
	C1.perimetro = 2 * 3.1416 * C1.radio
	C1.area = 2 * 3.1416 * C1.radio * C1.radio
	fmt.Println("Radio: ", C1.radio)
	fmt.Println("Perimetro: ", C1.perimetro)
	fmt.Println("Diametro: ", C1.diametro)
	fmt.Println("Area: ", C1.area)*/

	// FUNCIONES
	/*
		resultado, err := division(10, 3)
		if err != nil {
			fmt.Printf("Ocurrio un error: %v", err)
			return
		}
		fmt.Println(resultado)
	*/
	//ARCHIVOS Y DEFER
	//defer asiga la linea de codigo a una pila y la vuelve a ejecutar en FILO
	//despues de ejecutar todo el código
	/*file, err := os.Create("EDteam.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al Crear el archivo: %v", err)
		return
	}
	defer file.Close()
	bits, err := file.Write([]byte("Ya pertenezco a EDTeam"))
	if err != nil {
		fmt.Printf("Ocurrió un error al escribir: %v ", err)
		return
	}
	fmt.Printf("Bits escritos: %v", bits)*/
	//PANIC
	divi(10, 2)
	divi(10, 3)
	divi(10, 0)
	divi(10, 4)
}
func divi(dividendo, divisor int) {
	validar(divisor)
	fmt.Println(dividendo/divisor, ".", (dividendo % divisor))
}
func validar(divisor int) {
	if divisor == 0 {
		panic(" Divisor==0")
	}
}
func division(dividendo, divisor int) (result float32, err error) {
	if divisor == 0 {
		err = errors.New("No es posible dividir por cero")
		return
	}
	result = float32(dividendo) / float32(divisor)
	return
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//todo: Crear funciones paraprobar el codigo

//Mazo -> 52 cards
//Baraja -> 1+ mazos
//La carta esta conformado por [valor, figura]

//----------------------------------------------------
//Creamos un objeto, en este caso será la carta:
type carta struct {
	//Las propiedades de la carta serán:
	valor  string
	figura string
}

//El metodo del objeto card serán:
func (c carta) obtenerValor() int {
	switch c.valor {
	case "A":
		return 11
	case "K", "Q", "J":
		return 10
	default:
		valTemporal, _ := strconv.Atoi(c.valor)
		return valTemporal
	}
}

//----------------------------------------------------

//----------------------------------------------------
//Creamos al objeto: Jugador
type jugador struct {
	nombre string
	mano   []carta
}

func (j jugador) obtenerTotal() int {
	var total int

	for _, cartas := range j.mano {
		total = total + cartas.obtenerValor()
	}
	//todo: condicion de >=21 y ademas tener As en una mano
	return total
}

//----------------------------------------------------

func main() {

	baraja := crearBaraja(5)
	//fmt.Println(baraja)
	//fmt.Println(len(baraja))
	_ = baraja

	//Vamos a guardar el nombre del usuario
	fmt.Println("¿Cuál es tu nombre?")
	var nombreJugador string
	fmt.Scanf("%s", &nombreJugador)

	//Jugadores
	jugadorReal := jugador{
		nombre: nombreJugador,
	}
	jugadorVirtual := jugador{}

	//Repartir primeras 2 cartas
	baraja, jugadorReal.mano = repartirCartasAUnJugador(baraja, jugadorReal, 2)
	baraja, jugadorVirtual.mano = repartirCartasAUnJugador(baraja, jugadorVirtual, 2)

	//Sumar los puntos de cada jugador
	//totalJugadorReal := jugadorReal.obtenerTotal()
	//totalJugadorVirtual := jugadorVirtual.obtenerTotal()

	//fmt.Println(jugadorReal.mano, jugadorVirtual.mano)

	//Muestra las cartas al jugador
	fmt.Println("Tus cartas son:", jugadorReal.mano)
	fmt.Println("La carta del Dealer es:", jugadorVirtual.mano[0])

	var sobrevivientes []jugador

	for {
		//Debe de puntuacion actual
		totalJugadorReal := jugadorReal.obtenerTotal()
		fmt.Println("Tienes", totalJugadorReal, "puntos.")

		if totalJugadorReal > 21 {
			break
		}
		if totalJugadorReal == 21 {
			sobrevivientes = append(sobrevivientes, jugadorReal)
			break
		}
		//Como sus puntos son menores a 21, preguntamos si quiere otra carta
		//Pregunta al jugador si quiere otra carta
		fmt.Println("¿Quieres otra carta? s/n")
		//reader := bufio.NewReader(os.Stdin)
		//respuesta, _ := reader.ReadString('\n')
		var respuesta string
		fmt.Scanf("%s", &respuesta)
		//ToDo: si la funcion ReadString manda un error, debe terminar el programa
		//Sino quiere carta terminar ciclo
		if respuesta == "n" {
			sobrevivientes = append(sobrevivientes, jugadorReal)
			break
		}

		baraja, jugadorReal.mano = repartirCartasAUnJugador(baraja, jugadorReal, 1)
	}

	//Ahora realizamos el "for" para l dealer
	if len(sobrevivientes) == 0 {
		fmt.Println("¡La casa gana!")
		os.Exit(0)
	}

	dealerPerdio := false
	for {
		//Debe de puntuacion actual
		totalJugadorVirtual := jugadorVirtual.obtenerTotal()

		if totalJugadorVirtual > 21 {
			dealerPerdio = true
			break
		} else if totalJugadorVirtual <= 21 && totalJugadorVirtual >= 17 {
			break
		} else {
			baraja, jugadorVirtual.mano = repartirCartasAUnJugador(baraja, jugadorVirtual, 1)
		}

	}
	fmt.Println("La carta del Dealer es:", jugadorVirtual.mano)

	if dealerPerdio == true {
		fmt.Println("Los sobrevivientes ganan")
		os.Exit(0)
	}

	//El dealer no perdio y hay sobrevivientes, ent comparamos los puntos con c/u de 1en1
	for _, sobreviviente := range sobrevivientes {
		if sobreviviente.obtenerTotal() > jugadorVirtual.obtenerTotal() {
			fmt.Println("¡Jugador", sobreviviente.nombre, "gana!")
		} else if sobreviviente.obtenerTotal() == jugadorVirtual.obtenerTotal() {
			fmt.Println("Jugador", sobreviviente.nombre, "empata con Dealer")
		} else {
			fmt.Println("Dealer gana a", sobreviviente.nombre)
		}
	}
}

func crearBaraja(mazosDeseados int) []carta {
	// Cartas
	//Para crear unan variable vacia: var card [2]string
	//Para crear una variable con valores: card := [2]string("valor1", "valor 2")

	valores := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	figuras := [4]string{"Corazones", "Diamantes", "Espadas", "Treboles"}
	var baraja []carta //slice: es un arreglo sin num de elementos determinado

	//MAZOS
	//Creacion de los 5 mazos
	for i := 0; i < mazosDeseados; i++ {
		//Creacion de UN SOLO mazo
		for _, val := range valores {
			for _, fig := range figuras {
				//fmt.Println(val, fig)
				card := carta{
					valor:  val,
					figura: fig,
				}
				baraja = append(baraja, card)
			}
		}
	}
	//Para revolver la baraja:
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(baraja), func(i, j int) { baraja[i], baraja[j] = baraja[j], baraja[i] })

	return baraja
}

//Esta funcion nos va a regresar (baraja modificada,nueva mano del jugador)
func repartirCartasAUnJugador(b []carta, j jugador, numeroCartas int) ([]carta, []carta) {
	j.mano = append(j.mano, b[0:numeroCartas]...)
	b = b[numeroCartas:]
	return b, j.mano
}

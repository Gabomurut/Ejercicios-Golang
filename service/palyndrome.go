package service

import (

	"encoding/json"
	"strconv"
	"math/big"
	"net/http"
	"github.com/go-chi/chi"
)

func Palyndrome(w http.ResponseWriter, r *http.Request) {

	a := chi.URLParam(r, "number")  // Obtener number de los parámetros de la URL
	b, err := strconv.ParseInt(a, 0, 64) // Parsear a a Int

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Si hay error en el parseo devuelve Bad Request 
		return
	}

	for i := b; ; i++ {

		if isPalyndrome(i) && isPrime(i) {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(i) //Devuelve el número más próximo a b que sea a la vez Palindrome y Primo
			break

		}
	}

}


// Función que comprueba si el parámetro number es palíndrome o no
func isPalyndrome(number int64) bool {

	var reverseNumber int64 = 0

	tempNumber := number

	for {
		remainder := tempNumber % 10

		reverseNumber = reverseNumber*10 + remainder

		tempNumber = tempNumber / 10

		if tempNumber == 0 {

			break

		}

	}

	if reverseNumber == number {

		return true

	} else {

		return false

	}
}

// Función que comprueba si elparamtero number es número primo o no
func isPrime(number int64) bool {

	return (big.NewInt(number).ProbablyPrime(0))

}

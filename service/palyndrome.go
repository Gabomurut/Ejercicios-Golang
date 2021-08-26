package service

import (
	"encoding/json"

	"strconv"

	"math/big"

	"net/http"

	"github.com/go-chi/chi"
)

func Palyndrome(w http.ResponseWriter, r *http.Request) {

	a := chi.URLParam(r, "number")
	b, err := strconv.ParseInt(a, 0, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := b; ; i++ {

		if isPalyndrome(i) && isPrime(i) {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(i)
			break

		}
	}

}

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

func isPrime(number int64) bool {

	return (big.NewInt(number).ProbablyPrime(0))

}

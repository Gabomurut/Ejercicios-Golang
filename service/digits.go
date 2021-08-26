package service

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"github.com/go-chi/chi"
)

func Digits(w http.ResponseWriter, r *http.Request) {

	a := chi.URLParam(r, "number") // Obtener number de los parámetros de la URL
	b, err := strconv.ParseFloat(a, 64) // Parsear a a Float

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Si hay error en el parseo devuelve Bad Request 
		return
	}

	number := math.Pow((3 + math.Sqrt(5)), b) // Función solicitada en el ejercicio

	digits := int64(number) // Quitamos la coma

	c := strconv.FormatInt(digits, 10) // Parseo digits a String 

	chars := []rune(c)  // Obtenemos slice de carácteres de c

	var result string

	switch { // Switch para casos excepción 
	case len(chars) == 1:
		result = "00" + string(chars)

	case len(chars) == 2:
		result = "0" + string(chars)

	default:
		index := len(chars) - 3

		result = string(chars[index:]) // Sacamos los valores no necesarios del slice
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result) // Devuelve Result en formato String

}

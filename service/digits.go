package service

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Digits(w http.ResponseWriter, r *http.Request) {

	a := chi.URLParam(r, "number")
	b, err := strconv.ParseFloat(a, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	number := math.Pow((3 + math.Sqrt(5)), b)

	digits := int64(number)

	c := strconv.FormatInt(digits, 10)

	chars := []rune(c)

	var result string

	switch {
	case len(chars) == 1:
		result = "00" + string(chars)

	case len(chars) == 2:
		result = "0" + string(chars)

	default:
		index := len(chars) - 3

		result = string(chars[index:])
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

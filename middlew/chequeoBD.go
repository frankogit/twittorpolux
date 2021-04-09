package middlew

import (
	"net/http"

	"github.com/frankogit/twittorpolux/bd"
)

/* funcion de chequeo de bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdida con base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

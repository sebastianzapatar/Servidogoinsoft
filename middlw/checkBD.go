package middlw

import (
	"ingsoft/bd"
	"net/http"
)

/*
Los middlewares reciben algo y devuelven el mismo tipo
*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConection() == 0 {
			http.Error(w, "Conexion perdida con la bd", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

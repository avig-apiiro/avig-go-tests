package funcs

import (
	"io"
	"net/http"
)

func GetOtherHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Other Hello called!\n")
}

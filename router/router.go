package router

import(
	"net/http"

	"github.com/juanplagos/bubble/handler"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	//mux.HandleFunc("GET /entries/", GetEntries())
	//mux.HandleFunc("POST /entries/", CreateEntry())
	mux.HandleFunc("/authors", handler.GetAuthors)
	return mux
}
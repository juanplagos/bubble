package router

import(
	"net/http"

	"github.com/juanplagos/bubble/handler"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("GET /entries/", handler.GetEntries)
	// mux.HandleFunc("POST /entries/", handler.CreateEntry)
	mux.HandleFunc("GET /authors/", handler.GetAuthors)
	return mux
}
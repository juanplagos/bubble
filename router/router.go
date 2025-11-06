package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/juanplagos/bubble/handler"
	"github.com/juanplagos/bubble/repository"
	"github.com/juanplagos/bubble/usecase"
)

func RegisterRoutes(pool *pgxpool.Pool) *http.ServeMux {
	repo := repository.NewPostgresEntryRepo(pool)

	entryUseCase := usecase.NewEntryUseCase(*repo)
	
	entryHandler := handler.NewEntryHandler(entryUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /entries/", entryHandler.GetAll)
	//mux.HandleFunc("GET /authors/", entryHandler.GetAuthors)

	return mux
}
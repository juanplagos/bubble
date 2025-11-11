package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/juanplagos/bubble/handler"
	"github.com/juanplagos/bubble/repository"
	"github.com/juanplagos/bubble/usecase"
)

func RegisterRoutes(pool *pgxpool.Pool) *http.ServeMux {
	entryRepo := repository.NewPostgresEntryRepo(pool)
	authorRepo := repository.NewPostgresAuthorRepo(pool)

	entryUseCase := usecase.NewEntryUseCase(entryRepo)
	authorUseCase := usecase.NewAuthorUseCase(authorRepo)

	entryHandler := handler.NewEntryHandler(entryUseCase)
	authorHandler := handler.NewAuthorHandler(authorUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /entries/slug/", entryHandler.GetBySlug)
	mux.HandleFunc("GET /entries", entryHandler.GetAll)
	mux.HandleFunc("GET /entries/", entryHandler.GetByID)
	mux.HandleFunc("POST /entries", entryHandler.Create)
	mux.HandleFunc("PUT /entries/", entryHandler.Update)
	mux.HandleFunc("DELETE /entries/", entryHandler.Delete)

	mux.HandleFunc("GET /authors/email/", authorHandler.GetByEmail)
	mux.HandleFunc("GET /authors", authorHandler.GetAll)
	mux.HandleFunc("GET /authors/", authorHandler.GetByUsername)
	mux.HandleFunc("POST /authors", authorHandler.Create)
	mux.HandleFunc("PUT /authors/", authorHandler.Update)
	mux.HandleFunc("DELETE /authors/", authorHandler.Delete)

	return mux
}

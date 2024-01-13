package reader

import (
	"github.com/go-chi/chi/v5"
)

func NewReaderRouter() *chi.Mux {
	r := chi.NewRouter()

	r.With(GetDocumentID).Get("/{documentID}/{pageStart:[0-9]+}", HandleGetDocPage)

	return r
}

package reader

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetDocumentID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dID := chi.URLParam(r, "documentID")

		ctx := r.Context()
		ctx = context.WithValue(ctx, "documentID", dID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDocumentPage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pS := chi.URLParam(r, "pageStart")

		pageStart, err := strconv.ParseInt(pS, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "pageStart", pageStart)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

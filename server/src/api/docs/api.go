package docs

import (
	"encoding/json"
	"net/http"

	"github.com/Southclaws/qstring"
	"github.com/go-chi/chi"
	"github.com/golobby/container"

	"github.com/openmultiplayer/web/server/src/docsindex"
	"github.com/openmultiplayer/web/server/src/web"
)

type service struct {
	idx *docsindex.Index
}

func New() *chi.Mux {
	rtr := chi.NewRouter()
	svc := service{}
	container.Make(&svc.idx)

	fs := http.FileServer(http.Dir("docs/"))

	rtr.With(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/markdown")
			next.ServeHTTP(w, r)
		})
	}).Get("/*", http.StripPrefix("/docs/", fs).ServeHTTP)
	rtr.Get("/search", svc.search)

	return rtr
}

type query struct {
	Query string `qstring:"q"`
}

func (s *service) search(w http.ResponseWriter, r *http.Request) {
	var q query
	if err := qstring.Unmarshal(r.URL.Query(), &q); err != nil {
		web.StatusBadRequest(w, err)
		return
	}

	results, err := s.idx.Search(q.Query)
	if err != nil {
		web.StatusBadRequest(w, err)
		return
	}

	//nolint:errcheck
	json.NewEncoder(w).Encode(results)
}

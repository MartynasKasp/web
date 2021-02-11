package ppackage

import (
	"github.com/go-chi/chi"
	"github.com/golobby/container"

	"github.com/openmultiplayer/web/server/src/pkgstorage"
)

type service struct {
	store pkgstorage.Storer
}

func New() *chi.Mux {
	rtr := chi.NewMux()
	svc := service{}
	container.Make(&svc.store)

	rtr.Get("/", svc.list)
	rtr.Get("/package/{user}/{repo}", svc.get)
	rtr.Get("/package/{user}/{repo}/latest", svc.getLatest)

	return rtr
}

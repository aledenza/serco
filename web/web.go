package web

import (
	"net/http"

	"github.com/aledenza/serco"
	"github.com/aledenza/serco/utils"
	"github.com/danielgtaylor/huma/v2"
)

func register[I, O any](
	server *serco.Server,
	path string,
	method string,
	handler serco.HumaHandler[I, O],
	name string,
	description ...serco.HandlerDescriprion,
) {
	desc := utils.GetOptionalValue(description...)
	huma.Register(
		server.Api,
		huma.Operation{
			Method:      method,
			Path:        path,
			OperationID: name,
			Description: desc.Description,
			Tags:        desc.Tags,
		},
		handler,
	)
}

func Get[I, O any](
	server *serco.Server,
	path string,
	handler serco.HumaHandler[I, O],
	name string,
	description ...serco.HandlerDescriprion,
) {
	register(server, path, http.MethodGet, handler, name, description...)
}
func Post[I, O any](
	server *serco.Server,
	path string,
	handler serco.HumaHandler[I, O],
	name string,
	description ...serco.HandlerDescriprion) {
	register(server, path, http.MethodPost, handler, name, description...)
}
func Patch[I, O any](
	server *serco.Server,
	path string,
	handler serco.HumaHandler[I, O],
	name string,
	description ...serco.HandlerDescriprion) {
	register(server, path, http.MethodPatch, handler, name, description...)
}
func Delete[I, O any](
	server *serco.Server,
	path string,
	handler serco.HumaHandler[I, O],
	name string,
	description ...serco.HandlerDescriprion) {
	register(server, path, http.MethodDelete, handler, name, description...)
}
func Put[I, O any](
	server *serco.Server,
	path string,
	handler serco.HumaHandler[I, O],
	name string,
	description ...serco.HandlerDescriprion) {
	register(server, path, http.MethodPut, handler, name, description...)
}

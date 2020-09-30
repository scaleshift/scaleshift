// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/scaleshift/scaleshift/api/src/auth"
)

// GetRemoteRepositoriesHandlerFunc turns a function with the right signature into a get remote repositories handler
type GetRemoteRepositoriesHandlerFunc func(GetRemoteRepositoriesParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRemoteRepositoriesHandlerFunc) Handle(params GetRemoteRepositoriesParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRemoteRepositoriesHandler interface for that can handle valid get remote repositories params
type GetRemoteRepositoriesHandler interface {
	Handle(GetRemoteRepositoriesParams, *auth.Principal) middleware.Responder
}

// NewGetRemoteRepositories creates a new http.Handler for the get remote repositories operation
func NewGetRemoteRepositories(ctx *middleware.Context, handler GetRemoteRepositoriesHandler) *GetRemoteRepositories {
	return &GetRemoteRepositories{Context: ctx, Handler: handler}
}

/*GetRemoteRepositories swagger:route GET /repositories repository getRemoteRepositories

returns remote repositories


*/
type GetRemoteRepositories struct {
	Context *middleware.Context
	Handler GetRemoteRepositoriesHandler
}

func (o *GetRemoteRepositories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRemoteRepositoriesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *auth.Principal
	if uprinc != nil {
		principal = uprinc.(*auth.Principal) // this is really a auth.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

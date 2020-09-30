// Code generated by go-swagger; DO NOT EDIT.

package rescale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/scaleshift/scaleshift/api/src/auth"
)

// GetApplicationVersionHandlerFunc turns a function with the right signature into a get application version handler
type GetApplicationVersionHandlerFunc func(GetApplicationVersionParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApplicationVersionHandlerFunc) Handle(params GetApplicationVersionParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetApplicationVersionHandler interface for that can handle valid get application version params
type GetApplicationVersionHandler interface {
	Handle(GetApplicationVersionParams, *auth.Principal) middleware.Responder
}

// NewGetApplicationVersion creates a new http.Handler for the get application version operation
func NewGetApplicationVersion(ctx *middleware.Context, handler GetApplicationVersionHandler) *GetApplicationVersion {
	return &GetApplicationVersion{Context: ctx, Handler: handler}
}

/*GetApplicationVersion swagger:route GET /applications/{code}/{version}/ rescale getApplicationVersion

returns version information of a specified Rescale application


*/
type GetApplicationVersion struct {
	Context *middleware.Context
	Handler GetApplicationVersionHandler
}

func (o *GetApplicationVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetApplicationVersionParams()

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

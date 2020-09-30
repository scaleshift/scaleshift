// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/scaleshift/scaleshift/api/src/auth"
)

// GetJobLogsHandlerFunc turns a function with the right signature into a get job logs handler
type GetJobLogsHandlerFunc func(GetJobLogsParams, *auth.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetJobLogsHandlerFunc) Handle(params GetJobLogsParams, principal *auth.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetJobLogsHandler interface for that can handle valid get job logs params
type GetJobLogsHandler interface {
	Handle(GetJobLogsParams, *auth.Principal) middleware.Responder
}

// NewGetJobLogs creates a new http.Handler for the get job logs operation
func NewGetJobLogs(ctx *middleware.Context, handler GetJobLogsHandler) *GetJobLogs {
	return &GetJobLogs{Context: ctx, Handler: handler}
}

/*GetJobLogs swagger:route GET /jobs/{id}/logs job getJobLogs

returns the logs of a job


*/
type GetJobLogs struct {
	Context *middleware.Context
	Handler GetJobLogsHandler
}

func (o *GetJobLogs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetJobLogsParams()

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

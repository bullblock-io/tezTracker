// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFutureBakingRightsHandlerFunc turns a function with the right signature into a get future baking rights handler
type GetFutureBakingRightsHandlerFunc func(GetFutureBakingRightsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFutureBakingRightsHandlerFunc) Handle(params GetFutureBakingRightsParams) middleware.Responder {
	return fn(params)
}

// GetFutureBakingRightsHandler interface for that can handle valid get future baking rights params
type GetFutureBakingRightsHandler interface {
	Handle(GetFutureBakingRightsParams) middleware.Responder
}

// NewGetFutureBakingRights creates a new http.Handler for the get future baking rights operation
func NewGetFutureBakingRights(ctx *middleware.Context, handler GetFutureBakingRightsHandler) *GetFutureBakingRights {
	return &GetFutureBakingRights{Context: ctx, Handler: handler}
}

/*GetFutureBakingRights swagger:route GET /v2/data/{platform}/{network}/future_baking_rights Blocks getFutureBakingRights

GetFutureBakingRights get future baking rights API

*/
type GetFutureBakingRights struct {
	Context *middleware.Context
	Handler GetFutureBakingRightsHandler
}

func (o *GetFutureBakingRights) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFutureBakingRightsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAirportIATATypeLastHandlerFunc turns a function with the right signature into a get airport i a t a type last handler
type GetAirportIATATypeLastHandlerFunc func(GetAirportIATATypeLastParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAirportIATATypeLastHandlerFunc) Handle(params GetAirportIATATypeLastParams) middleware.Responder {
	return fn(params)
}

// GetAirportIATATypeLastHandler interface for that can handle valid get airport i a t a type last params
type GetAirportIATATypeLastHandler interface {
	Handle(GetAirportIATATypeLastParams) middleware.Responder
}

// NewGetAirportIATATypeLast creates a new http.Handler for the get airport i a t a type last operation
func NewGetAirportIATATypeLast(ctx *middleware.Context, handler GetAirportIATATypeLastHandler) *GetAirportIATATypeLast {
	return &GetAirportIATATypeLast{Context: ctx, Handler: handler}
}

/*GetAirportIATATypeLast swagger:route GET /airport/{IATA}/{type}/last getAirportIATATypeLast

GetAirportIATATypeLast get airport i a t a type last API

*/
type GetAirportIATATypeLast struct {
	Context *middleware.Context
	Handler GetAirportIATATypeLastHandler
}

func (o *GetAirportIATATypeLast) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAirportIATATypeLastParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

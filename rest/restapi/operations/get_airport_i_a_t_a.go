// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/COMA-tor/rtm/rest/models"
)

// GetAirportIATAHandlerFunc turns a function with the right signature into a get airport i a t a handler
type GetAirportIATAHandlerFunc func(GetAirportIATAParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAirportIATAHandlerFunc) Handle(params GetAirportIATAParams) middleware.Responder {
	return fn(params)
}

// GetAirportIATAHandler interface for that can handle valid get airport i a t a params
type GetAirportIATAHandler interface {
	Handle(GetAirportIATAParams) middleware.Responder
}

// NewGetAirportIATA creates a new http.Handler for the get airport i a t a operation
func NewGetAirportIATA(ctx *middleware.Context, handler GetAirportIATAHandler) *GetAirportIATA {
	return &GetAirportIATA{Context: ctx, Handler: handler}
}

/*GetAirportIATA swagger:route GET /airport/{IATA} getAirportIATA

GetAirportIATA get airport i a t a API

*/
type GetAirportIATA struct {
	Context *middleware.Context
	Handler GetAirportIATAHandler
}

func (o *GetAirportIATA) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAirportIATAParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetAirportIATAOKBody get airport i a t a o k body
//
// swagger:model GetAirportIATAOKBody
type GetAirportIATAOKBody struct {

	// pressure
	// Required: true
	Pressure *models.Measure `json:"pressure"`

	// temperature
	// Required: true
	Temperature *models.Measure `json:"temperature"`

	// wind speed
	// Required: true
	WindSpeed *models.Measure `json:"wind_speed"`
}

// Validate validates this get airport i a t a o k body
func (o *GetAirportIATAOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePressure(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemperature(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWindSpeed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAirportIATAOKBody) validatePressure(formats strfmt.Registry) error {

	if err := validate.Required("getAirportIATAOK"+"."+"pressure", "body", o.Pressure); err != nil {
		return err
	}

	if o.Pressure != nil {
		if err := o.Pressure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAirportIATAOK" + "." + "pressure")
			}
			return err
		}
	}

	return nil
}

func (o *GetAirportIATAOKBody) validateTemperature(formats strfmt.Registry) error {

	if err := validate.Required("getAirportIATAOK"+"."+"temperature", "body", o.Temperature); err != nil {
		return err
	}

	if o.Temperature != nil {
		if err := o.Temperature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAirportIATAOK" + "." + "temperature")
			}
			return err
		}
	}

	return nil
}

func (o *GetAirportIATAOKBody) validateWindSpeed(formats strfmt.Registry) error {

	if err := validate.Required("getAirportIATAOK"+"."+"wind_speed", "body", o.WindSpeed); err != nil {
		return err
	}

	if o.WindSpeed != nil {
		if err := o.WindSpeed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAirportIATAOK" + "." + "wind_speed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAirportIATAOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAirportIATAOKBody) UnmarshalBinary(b []byte) error {
	var res GetAirportIATAOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

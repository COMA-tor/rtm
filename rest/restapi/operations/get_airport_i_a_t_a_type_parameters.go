// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAirportIATATypeParams creates a new GetAirportIATATypeParams object
// with the default values initialized.
func NewGetAirportIATATypeParams() GetAirportIATATypeParams {

	var (
		// initialize parameters with default values

		countDefault = int64(50)
		stepDefault  = int64(1000)
	)

	return GetAirportIATATypeParams{
		Count: &countDefault,

		Step: &stepDefault,
	}
}

// GetAirportIATATypeParams contains all the bound params for the get airport i a t a type operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAirportIATAType
type GetAirportIATATypeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	IATA string
	/*Number of measures to return
	  In: query
	  Default: 50
	*/
	Count *int64
	/*Time step between measures for aggregation in ms
	  In: query
	  Default: 1000
	*/
	Step *int64
	/*
	  Required: true
	  In: path
	*/
	Type string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAirportIATATypeParams() beforehand.
func (o *GetAirportIATATypeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rIATA, rhkIATA, _ := route.Params.GetOK("IATA")
	if err := o.bindIATA(rIATA, rhkIATA, route.Formats); err != nil {
		res = append(res, err)
	}

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qStep, qhkStep, _ := qs.GetOK("step")
	if err := o.bindStep(qStep, qhkStep, route.Formats); err != nil {
		res = append(res, err)
	}

	rType, rhkType, _ := route.Params.GetOK("type")
	if err := o.bindType(rType, rhkType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIATA binds and validates parameter IATA from path.
func (o *GetAirportIATATypeParams) bindIATA(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.IATA = raw

	return nil
}

// bindCount binds and validates parameter Count from query.
func (o *GetAirportIATATypeParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAirportIATATypeParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("count", "query", "int64", raw)
	}
	o.Count = &value

	return nil
}

// bindStep binds and validates parameter Step from query.
func (o *GetAirportIATATypeParams) bindStep(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAirportIATATypeParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("step", "query", "int64", raw)
	}
	o.Step = &value

	return nil
}

// bindType binds and validates parameter Type from path.
func (o *GetAirportIATATypeParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Type = raw

	return nil
}

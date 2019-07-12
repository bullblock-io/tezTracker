// Code generated by go-swagger; DO NOT EDIT.

package operation_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams creates a new GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams object
// no default values defined in spec.
func NewGetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams() GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams {

	return GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams{}
}

// GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams contains all the bound params for the get v2 data platform network operation groups operation group ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV2DataPlatformNetworkOperationGroupsOperationGroupID
type GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	APIKey string
	/*
	  Required: true
	  In: path
	*/
	Network string
	/*
	  Required: true
	  In: path
	*/
	OperationGroupID string
	/*
	  Required: true
	  In: path
	*/
	Platform string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams() beforehand.
func (o *GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAPIKey(r.Header[http.CanonicalHeaderKey("apiKey")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	rOperationGroupID, rhkOperationGroupID, _ := route.Params.GetOK("operationGroupId")
	if err := o.bindOperationGroupID(rOperationGroupID, rhkOperationGroupID, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAPIKey binds and validates parameter APIKey from header.
func (o *GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams) bindAPIKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("apiKey", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("apiKey", "header", raw); err != nil {
		return err
	}

	o.APIKey = raw

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindOperationGroupID binds and validates parameter OperationGroupID from path.
func (o *GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams) bindOperationGroupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.OperationGroupID = raw

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}

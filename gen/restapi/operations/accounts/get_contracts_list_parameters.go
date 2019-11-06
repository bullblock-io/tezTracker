// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetContractsListParams creates a new GetContractsListParams object
// with the default values initialized.
func NewGetContractsListParams() GetContractsListParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)

		offsetDefault = int64(0)
	)

	return GetContractsListParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetContractsListParams contains all the bound params for the get contracts list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getContractsList
type GetContractsListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Not used
	  In: query
	  Collection Format: multi
	*/
	AccountDelegate []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	AccountID []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	AccountManager []string
	/*
	  In: query
	*/
	AfterID *string
	/*
	  Maximum: 500
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*Offset
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Not used
	  In: query
	*/
	Order *string
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
	/*Not used
	  In: query
	*/
	SortBy *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetContractsListParams() beforehand.
func (o *GetContractsListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccountDelegate, qhkAccountDelegate, _ := qs.GetOK("account_delegate")
	if err := o.bindAccountDelegate(qAccountDelegate, qhkAccountDelegate, route.Formats); err != nil {
		res = append(res, err)
	}

	qAccountID, qhkAccountID, _ := qs.GetOK("account_id")
	if err := o.bindAccountID(qAccountID, qhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	qAccountManager, qhkAccountManager, _ := qs.GetOK("account_manager")
	if err := o.bindAccountManager(qAccountManager, qhkAccountManager, route.Formats); err != nil {
		res = append(res, err)
	}

	qAfterID, qhkAfterID, _ := qs.GetOK("after_id")
	if err := o.bindAfterID(qAfterID, qhkAfterID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrder, qhkOrder, _ := qs.GetOK("order")
	if err := o.bindOrder(qOrder, qhkOrder, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortBy, qhkSortBy, _ := qs.GetOK("sort_by")
	if err := o.bindSortBy(qSortBy, qhkSortBy, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAccountDelegate binds and validates array parameter AccountDelegate from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetContractsListParams) bindAccountDelegate(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	accountDelegateIC := rawData

	if len(accountDelegateIC) == 0 {
		return nil
	}

	var accountDelegateIR []string
	for _, accountDelegateIV := range accountDelegateIC {
		accountDelegateI := accountDelegateIV

		accountDelegateIR = append(accountDelegateIR, accountDelegateI)
	}

	o.AccountDelegate = accountDelegateIR

	return nil
}

// bindAccountID binds and validates array parameter AccountID from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetContractsListParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	accountIDIC := rawData

	if len(accountIDIC) == 0 {
		return nil
	}

	var accountIDIR []string
	for _, accountIDIV := range accountIDIC {
		accountIDI := accountIDIV

		accountIDIR = append(accountIDIR, accountIDI)
	}

	o.AccountID = accountIDIR

	return nil
}

// bindAccountManager binds and validates array parameter AccountManager from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetContractsListParams) bindAccountManager(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	accountManagerIC := rawData

	if len(accountManagerIC) == 0 {
		return nil
	}

	var accountManagerIR []string
	for _, accountManagerIV := range accountManagerIC {
		accountManagerI := accountManagerIV

		accountManagerIR = append(accountManagerIR, accountManagerI)
	}

	o.AccountManager = accountManagerIR

	return nil
}

// bindAfterID binds and validates parameter AfterID from query.
func (o *GetContractsListParams) bindAfterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AfterID = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetContractsListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetContractsListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetContractsListParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 500, false); err != nil {
		return err
	}

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetContractsListParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetContractsListParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetContractsListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetContractsListParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", int64(*o.Offset), 0, false); err != nil {
		return err
	}

	return nil
}

// bindOrder binds and validates parameter Order from query.
func (o *GetContractsListParams) bindOrder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Order = &raw

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetContractsListParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}

// bindSortBy binds and validates parameter SortBy from query.
func (o *GetContractsListParams) bindSortBy(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SortBy = &raw

	return nil
}

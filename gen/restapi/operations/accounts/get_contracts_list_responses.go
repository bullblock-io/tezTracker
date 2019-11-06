// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetContractsListOKCode is the HTTP code returned for type GetContractsListOK
const GetContractsListOKCode int = 200

/*GetContractsListOK Query compatibility endpoint for contract accounts

swagger:response getContractsListOK
*/
type GetContractsListOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.AccountsRow `json:"body,omitempty"`
}

// NewGetContractsListOK creates GetContractsListOK with default headers values
func NewGetContractsListOK() *GetContractsListOK {

	return &GetContractsListOK{}
}

// WithXTotalCount adds the xTotalCount to the get contracts list o k response
func (o *GetContractsListOK) WithXTotalCount(xTotalCount int64) *GetContractsListOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get contracts list o k response
func (o *GetContractsListOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get contracts list o k response
func (o *GetContractsListOK) WithPayload(payload []*models.AccountsRow) *GetContractsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get contracts list o k response
func (o *GetContractsListOK) SetPayload(payload []*models.AccountsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContractsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AccountsRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetContractsListNotFoundCode is the HTTP code returned for type GetContractsListNotFound
const GetContractsListNotFoundCode int = 404

/*GetContractsListNotFound Not Found

swagger:response getContractsListNotFound
*/
type GetContractsListNotFound struct {
}

// NewGetContractsListNotFound creates GetContractsListNotFound with default headers values
func NewGetContractsListNotFound() *GetContractsListNotFound {

	return &GetContractsListNotFound{}
}

// WriteResponse to the client
func (o *GetContractsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

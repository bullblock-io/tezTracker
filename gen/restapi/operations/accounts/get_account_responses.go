// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetAccountOKCode is the HTTP code returned for type GetAccountOK
const GetAccountOKCode int = 200

/*GetAccountOK Query compatibility endpoint for account

swagger:response getAccountOK
*/
type GetAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccountsRow `json:"body,omitempty"`
}

// NewGetAccountOK creates GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {

	return &GetAccountOK{}
}

// WithPayload adds the payload to the get account o k response
func (o *GetAccountOK) WithPayload(payload *models.AccountsRow) *GetAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account o k response
func (o *GetAccountOK) SetPayload(payload *models.AccountsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountNotFoundCode is the HTTP code returned for type GetAccountNotFound
const GetAccountNotFoundCode int = 404

/*GetAccountNotFound Not Found

swagger:response getAccountNotFound
*/
type GetAccountNotFound struct {
}

// NewGetAccountNotFound creates GetAccountNotFound with default headers values
func NewGetAccountNotFound() *GetAccountNotFound {

	return &GetAccountNotFound{}
}

// WriteResponse to the client
func (o *GetAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetAccountInternalServerErrorCode is the HTTP code returned for type GetAccountInternalServerError
const GetAccountInternalServerErrorCode int = 500

/*GetAccountInternalServerError Internal error

swagger:response getAccountInternalServerError
*/
type GetAccountInternalServerError struct {
}

// NewGetAccountInternalServerError creates GetAccountInternalServerError with default headers values
func NewGetAccountInternalServerError() *GetAccountInternalServerError {

	return &GetAccountInternalServerError{}
}

// WriteResponse to the client
func (o *GetAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

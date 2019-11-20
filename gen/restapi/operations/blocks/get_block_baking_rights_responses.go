// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetBlockBakingRightsOKCode is the HTTP code returned for type GetBlockBakingRightsOK
const GetBlockBakingRightsOKCode int = 200

/*GetBlockBakingRightsOK Endpoint for baking rights of a block by hash or level

swagger:response getBlockBakingRightsOK
*/
type GetBlockBakingRightsOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.BakingRightsRow `json:"body,omitempty"`
}

// NewGetBlockBakingRightsOK creates GetBlockBakingRightsOK with default headers values
func NewGetBlockBakingRightsOK() *GetBlockBakingRightsOK {

	return &GetBlockBakingRightsOK{}
}

// WithXTotalCount adds the xTotalCount to the get block baking rights o k response
func (o *GetBlockBakingRightsOK) WithXTotalCount(xTotalCount int64) *GetBlockBakingRightsOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get block baking rights o k response
func (o *GetBlockBakingRightsOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get block baking rights o k response
func (o *GetBlockBakingRightsOK) WithPayload(payload []*models.BakingRightsRow) *GetBlockBakingRightsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get block baking rights o k response
func (o *GetBlockBakingRightsOK) SetPayload(payload []*models.BakingRightsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlockBakingRightsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BakingRightsRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBlockBakingRightsNotFoundCode is the HTTP code returned for type GetBlockBakingRightsNotFound
const GetBlockBakingRightsNotFoundCode int = 404

/*GetBlockBakingRightsNotFound Not Found

swagger:response getBlockBakingRightsNotFound
*/
type GetBlockBakingRightsNotFound struct {
}

// NewGetBlockBakingRightsNotFound creates GetBlockBakingRightsNotFound with default headers values
func NewGetBlockBakingRightsNotFound() *GetBlockBakingRightsNotFound {

	return &GetBlockBakingRightsNotFound{}
}

// WriteResponse to the client
func (o *GetBlockBakingRightsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBlockBakingRightsInternalServerErrorCode is the HTTP code returned for type GetBlockBakingRightsInternalServerError
const GetBlockBakingRightsInternalServerErrorCode int = 500

/*GetBlockBakingRightsInternalServerError Internal error

swagger:response getBlockBakingRightsInternalServerError
*/
type GetBlockBakingRightsInternalServerError struct {
}

// NewGetBlockBakingRightsInternalServerError creates GetBlockBakingRightsInternalServerError with default headers values
func NewGetBlockBakingRightsInternalServerError() *GetBlockBakingRightsInternalServerError {

	return &GetBlockBakingRightsInternalServerError{}
}

// WriteResponse to the client
func (o *GetBlockBakingRightsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
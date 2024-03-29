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

// GetBakingRightsOKCode is the HTTP code returned for type GetBakingRightsOK
const GetBakingRightsOKCode int = 200

/*GetBakingRightsOK Endpoint for baking rights

swagger:response getBakingRightsOK
*/
type GetBakingRightsOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.BakingRightsPerBlock `json:"body,omitempty"`
}

// NewGetBakingRightsOK creates GetBakingRightsOK with default headers values
func NewGetBakingRightsOK() *GetBakingRightsOK {

	return &GetBakingRightsOK{}
}

// WithXTotalCount adds the xTotalCount to the get baking rights o k response
func (o *GetBakingRightsOK) WithXTotalCount(xTotalCount int64) *GetBakingRightsOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get baking rights o k response
func (o *GetBakingRightsOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get baking rights o k response
func (o *GetBakingRightsOK) WithPayload(payload []*models.BakingRightsPerBlock) *GetBakingRightsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get baking rights o k response
func (o *GetBakingRightsOK) SetPayload(payload []*models.BakingRightsPerBlock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBakingRightsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BakingRightsPerBlock, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBakingRightsBadRequestCode is the HTTP code returned for type GetBakingRightsBadRequest
const GetBakingRightsBadRequestCode int = 400

/*GetBakingRightsBadRequest Bad request

swagger:response getBakingRightsBadRequest
*/
type GetBakingRightsBadRequest struct {
}

// NewGetBakingRightsBadRequest creates GetBakingRightsBadRequest with default headers values
func NewGetBakingRightsBadRequest() *GetBakingRightsBadRequest {

	return &GetBakingRightsBadRequest{}
}

// WriteResponse to the client
func (o *GetBakingRightsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBakingRightsNotFoundCode is the HTTP code returned for type GetBakingRightsNotFound
const GetBakingRightsNotFoundCode int = 404

/*GetBakingRightsNotFound Not Found

swagger:response getBakingRightsNotFound
*/
type GetBakingRightsNotFound struct {
}

// NewGetBakingRightsNotFound creates GetBakingRightsNotFound with default headers values
func NewGetBakingRightsNotFound() *GetBakingRightsNotFound {

	return &GetBakingRightsNotFound{}
}

// WriteResponse to the client
func (o *GetBakingRightsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

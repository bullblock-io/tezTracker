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

// GetFutureBakingRightsOKCode is the HTTP code returned for type GetFutureBakingRightsOK
const GetFutureBakingRightsOKCode int = 200

/*GetFutureBakingRightsOK Endpoint for future baking rights

swagger:response getFutureBakingRightsOK
*/
type GetFutureBakingRightsOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.FutureBakingRightsPerBlock `json:"body,omitempty"`
}

// NewGetFutureBakingRightsOK creates GetFutureBakingRightsOK with default headers values
func NewGetFutureBakingRightsOK() *GetFutureBakingRightsOK {

	return &GetFutureBakingRightsOK{}
}

// WithXTotalCount adds the xTotalCount to the get future baking rights o k response
func (o *GetFutureBakingRightsOK) WithXTotalCount(xTotalCount int64) *GetFutureBakingRightsOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get future baking rights o k response
func (o *GetFutureBakingRightsOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get future baking rights o k response
func (o *GetFutureBakingRightsOK) WithPayload(payload []*models.FutureBakingRightsPerBlock) *GetFutureBakingRightsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get future baking rights o k response
func (o *GetFutureBakingRightsOK) SetPayload(payload []*models.FutureBakingRightsPerBlock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFutureBakingRightsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.FutureBakingRightsPerBlock, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetFutureBakingRightsNotFoundCode is the HTTP code returned for type GetFutureBakingRightsNotFound
const GetFutureBakingRightsNotFoundCode int = 404

/*GetFutureBakingRightsNotFound Not Found

swagger:response getFutureBakingRightsNotFound
*/
type GetFutureBakingRightsNotFound struct {
}

// NewGetFutureBakingRightsNotFound creates GetFutureBakingRightsNotFound with default headers values
func NewGetFutureBakingRightsNotFound() *GetFutureBakingRightsNotFound {

	return &GetFutureBakingRightsNotFound{}
}

// WriteResponse to the client
func (o *GetFutureBakingRightsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
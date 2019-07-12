// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bullblock-io/tezTracker/gen/models"
)

// GetV2MetadataPlatformNetworksOKCode is the HTTP code returned for type GetV2MetadataPlatformNetworksOK
const GetV2MetadataPlatformNetworksOKCode int = 200

/*GetV2MetadataPlatformNetworksOK Metadata endpoint for listing available networks

swagger:response getV2MetadataPlatformNetworksOK
*/
type GetV2MetadataPlatformNetworksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.PlatformDiscoveryTypesNetwork `json:"body,omitempty"`
}

// NewGetV2MetadataPlatformNetworksOK creates GetV2MetadataPlatformNetworksOK with default headers values
func NewGetV2MetadataPlatformNetworksOK() *GetV2MetadataPlatformNetworksOK {

	return &GetV2MetadataPlatformNetworksOK{}
}

// WithPayload adds the payload to the get v2 metadata platform networks o k response
func (o *GetV2MetadataPlatformNetworksOK) WithPayload(payload []*models.PlatformDiscoveryTypesNetwork) *GetV2MetadataPlatformNetworksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v2 metadata platform networks o k response
func (o *GetV2MetadataPlatformNetworksOK) SetPayload(payload []*models.PlatformDiscoveryTypesNetwork) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV2MetadataPlatformNetworksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.PlatformDiscoveryTypesNetwork, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV2MetadataPlatformNetworksNotFoundCode is the HTTP code returned for type GetV2MetadataPlatformNetworksNotFound
const GetV2MetadataPlatformNetworksNotFoundCode int = 404

/*GetV2MetadataPlatformNetworksNotFound Not found

swagger:response getV2MetadataPlatformNetworksNotFound
*/
type GetV2MetadataPlatformNetworksNotFound struct {
}

// NewGetV2MetadataPlatformNetworksNotFound creates GetV2MetadataPlatformNetworksNotFound with default headers values
func NewGetV2MetadataPlatformNetworksNotFound() *GetV2MetadataPlatformNetworksNotFound {

	return &GetV2MetadataPlatformNetworksNotFound{}
}

// WriteResponse to the client
func (o *GetV2MetadataPlatformNetworksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// Code generated by go-swagger; DO NOT EDIT.

package baking_rights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bullblock-io/tezTracker/services/rpc_client/models"
)

// GetBakingRightsReader is a Reader for the GetBakingRights structure.
type GetBakingRightsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBakingRightsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBakingRightsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetBakingRightsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBakingRightsOK creates a GetBakingRightsOK with default headers values
func NewGetBakingRightsOK() *GetBakingRightsOK {
	return &GetBakingRightsOK{}
}

/*GetBakingRightsOK handles this case with default header values.

Query compatibility endpoint for account
*/
type GetBakingRightsOK struct {
	Payload []*models.BakingRight
}

func (o *GetBakingRightsOK) Error() string {
	return fmt.Sprintf("[GET /chains/{network}/blocks/{block}/helpers/baking_rights][%d] getBakingRightsOK  %+v", 200, o.Payload)
}

func (o *GetBakingRightsOK) GetPayload() []*models.BakingRight {
	return o.Payload
}

func (o *GetBakingRightsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBakingRightsInternalServerError creates a GetBakingRightsInternalServerError with default headers values
func NewGetBakingRightsInternalServerError() *GetBakingRightsInternalServerError {
	return &GetBakingRightsInternalServerError{}
}

/*GetBakingRightsInternalServerError handles this case with default header values.

Internal error
*/
type GetBakingRightsInternalServerError struct {
}

func (o *GetBakingRightsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chains/{network}/blocks/{block}/helpers/baking_rights][%d] getBakingRightsInternalServerError ", 500)
}

func (o *GetBakingRightsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

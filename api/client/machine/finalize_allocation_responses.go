// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/metal-go/api/models"
)

// FinalizeAllocationReader is a Reader for the FinalizeAllocation structure.
type FinalizeAllocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FinalizeAllocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFinalizeAllocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFinalizeAllocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFinalizeAllocationOK creates a FinalizeAllocationOK with default headers values
func NewFinalizeAllocationOK() *FinalizeAllocationOK {
	return &FinalizeAllocationOK{}
}

/*FinalizeAllocationOK handles this case with default header values.

OK
*/
type FinalizeAllocationOK struct {
	Payload *models.V1MachineResponse
}

func (o *FinalizeAllocationOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/finalize-allocation][%d] finalizeAllocationOK  %+v", 200, o.Payload)
}

func (o *FinalizeAllocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFinalizeAllocationDefault creates a FinalizeAllocationDefault with default headers values
func NewFinalizeAllocationDefault(code int) *FinalizeAllocationDefault {
	return &FinalizeAllocationDefault{
		_statusCode: code,
	}
}

/*FinalizeAllocationDefault handles this case with default header values.

Error
*/
type FinalizeAllocationDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the finalize allocation default response
func (o *FinalizeAllocationDefault) Code() int {
	return o._statusCode
}

func (o *FinalizeAllocationDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/finalize-allocation][%d] finalizeAllocation default  %+v", o._statusCode, o.Payload)
}

func (o *FinalizeAllocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

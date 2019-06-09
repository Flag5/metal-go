// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// MachineBiosReader is a Reader for the MachineBios structure.
type MachineBiosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineBiosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMachineBiosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewMachineBiosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMachineBiosOK creates a MachineBiosOK with default headers values
func NewMachineBiosOK() *MachineBiosOK {
	return &MachineBiosOK{}
}

/*MachineBiosOK handles this case with default header values.

OK
*/
type MachineBiosOK struct {
	Payload *models.V1MachineResponse
}

func (o *MachineBiosOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/bios][%d] machineBiosOK  %+v", 200, o.Payload)
}

func (o *MachineBiosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachineBiosDefault creates a MachineBiosDefault with default headers values
func NewMachineBiosDefault(code int) *MachineBiosDefault {
	return &MachineBiosDefault{
		_statusCode: code,
	}
}

/*MachineBiosDefault handles this case with default header values.

Error
*/
type MachineBiosDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the machine bios default response
func (o *MachineBiosDefault) Code() int {
	return o._statusCode
}

func (o *MachineBiosDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/bios][%d] machineBios default  %+v", o._statusCode, o.Payload)
}

func (o *MachineBiosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
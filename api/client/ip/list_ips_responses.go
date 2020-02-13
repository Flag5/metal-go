// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/metal-go/api/models"
)

// ListIpsReader is a Reader for the ListIps structure.
type ListIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListIpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListIpsOK creates a ListIpsOK with default headers values
func NewListIpsOK() *ListIpsOK {
	return &ListIpsOK{}
}

/*ListIpsOK handles this case with default header values.

OK
*/
type ListIpsOK struct {
	Payload []*models.V1IPResponse
}

func (o *ListIpsOK) Error() string {
	return fmt.Sprintf("[GET /v1/ip][%d] listIpsOK  %+v", 200, o.Payload)
}

func (o *ListIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIpsDefault creates a ListIpsDefault with default headers values
func NewListIpsDefault(code int) *ListIpsDefault {
	return &ListIpsDefault{
		_statusCode: code,
	}
}

/*ListIpsDefault handles this case with default header values.

Error
*/
type ListIpsDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the list ips default response
func (o *ListIpsDefault) Code() int {
	return o._statusCode
}

func (o *ListIpsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ip][%d] listIPs default  %+v", o._statusCode, o.Payload)
}

func (o *ListIpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

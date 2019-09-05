// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// NewAllocateIPParams creates a new AllocateIPParams object
// with the default values initialized.
func NewAllocateIPParams() *AllocateIPParams {
	var ()
	return &AllocateIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllocateIPParamsWithTimeout creates a new AllocateIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllocateIPParamsWithTimeout(timeout time.Duration) *AllocateIPParams {
	var ()
	return &AllocateIPParams{

		timeout: timeout,
	}
}

// NewAllocateIPParamsWithContext creates a new AllocateIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllocateIPParamsWithContext(ctx context.Context) *AllocateIPParams {
	var ()
	return &AllocateIPParams{

		Context: ctx,
	}
}

// NewAllocateIPParamsWithHTTPClient creates a new AllocateIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllocateIPParamsWithHTTPClient(client *http.Client) *AllocateIPParams {
	var ()
	return &AllocateIPParams{
		HTTPClient: client,
	}
}

/*AllocateIPParams contains all the parameters to send to the API endpoint
for the allocate IP operation typically these are written to a http.Request
*/
type AllocateIPParams struct {

	/*Body*/
	Body *models.V1IPAllocateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the allocate IP params
func (o *AllocateIPParams) WithTimeout(timeout time.Duration) *AllocateIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allocate IP params
func (o *AllocateIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allocate IP params
func (o *AllocateIPParams) WithContext(ctx context.Context) *AllocateIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allocate IP params
func (o *AllocateIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allocate IP params
func (o *AllocateIPParams) WithHTTPClient(client *http.Client) *AllocateIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allocate IP params
func (o *AllocateIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the allocate IP params
func (o *AllocateIPParams) WithBody(body *models.V1IPAllocateRequest) *AllocateIPParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the allocate IP params
func (o *AllocateIPParams) SetBody(body *models.V1IPAllocateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AllocateIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

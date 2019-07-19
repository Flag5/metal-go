// Code generated by go-swagger; DO NOT EDIT.

package health

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
)

// NewHealthParams creates a new HealthParams object
// with the default values initialized.
func NewHealthParams() *HealthParams {

	return &HealthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHealthParamsWithTimeout creates a new HealthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHealthParamsWithTimeout(timeout time.Duration) *HealthParams {

	return &HealthParams{

		timeout: timeout,
	}
}

// NewHealthParamsWithContext creates a new HealthParams object
// with the default values initialized, and the ability to set a context for a request
func NewHealthParamsWithContext(ctx context.Context) *HealthParams {

	return &HealthParams{

		Context: ctx,
	}
}

// NewHealthParamsWithHTTPClient creates a new HealthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHealthParamsWithHTTPClient(client *http.Client) *HealthParams {

	return &HealthParams{
		HTTPClient: client,
	}
}

/*HealthParams contains all the parameters to send to the API endpoint
for the health operation typically these are written to a http.Request
*/
type HealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the health params
func (o *HealthParams) WithTimeout(timeout time.Duration) *HealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the health params
func (o *HealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the health params
func (o *HealthParams) WithContext(ctx context.Context) *HealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the health params
func (o *HealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the health params
func (o *HealthParams) WithHTTPClient(client *http.Client) *HealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the health params
func (o *HealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

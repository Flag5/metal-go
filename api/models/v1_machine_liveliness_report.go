// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineLivelinessReport v1 machine liveliness report
// swagger:model v1.MachineLivelinessReport
type V1MachineLivelinessReport struct {

	// the number of machines alive
	// Required: true
	AliveCount *int32 `json:"alive_count"`

	// the number of dead machines
	// Required: true
	DeadCount *int32 `json:"dead_count"`

	// the number of machines with unknown liveliness
	// Required: true
	UnknownCount *int32 `json:"unknown_count"`
}

// Validate validates this v1 machine liveliness report
func (m *V1MachineLivelinessReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAliveCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnknownCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineLivelinessReport) validateAliveCount(formats strfmt.Registry) error {

	if err := validate.Required("alive_count", "body", m.AliveCount); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineLivelinessReport) validateDeadCount(formats strfmt.Registry) error {

	if err := validate.Required("dead_count", "body", m.DeadCount); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineLivelinessReport) validateUnknownCount(formats strfmt.Registry) error {

	if err := validate.Required("unknown_count", "body", m.UnknownCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineLivelinessReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineLivelinessReport) UnmarshalBinary(b []byte) error {
	var res V1MachineLivelinessReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

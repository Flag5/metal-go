// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1Quota v1 quota
// swagger:model v1.Quota
type V1Quota struct {

	// quota
	Quota *WrappersInt32Value `json:"quota,omitempty"`
}

// Validate validates this v1 quota
func (m *V1Quota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Quota) validateQuota(formats strfmt.Registry) error {

	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Quota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Quota) UnmarshalBinary(b []byte) error {
	var res V1Quota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

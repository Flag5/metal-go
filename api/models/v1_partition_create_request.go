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

// V1PartitionCreateRequest v1 partition create request
// swagger:model v1.PartitionCreateRequest
type V1PartitionCreateRequest struct {

	// the boot configuration of this partition
	// Required: true
	Bootconfig *V1PartitionBootConfiguration `json:"bootconfig"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	// Unique: true
	ID *string `json:"id"`

	// the address to the management service of this partition
	Mgmtserviceaddress string `json:"mgmtserviceaddress,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// the length of project networks for this partition, default 22
	// Maximum: 30
	// Minimum: 16
	Projectnetworkprefixlength int32 `json:"projectnetworkprefixlength,omitempty"`
}

// Validate validates this v1 partition create request
func (m *V1PartitionCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectnetworkprefixlength(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PartitionCreateRequest) validateBootconfig(formats strfmt.Registry) error {

	if err := validate.Required("bootconfig", "body", m.Bootconfig); err != nil {
		return err
	}

	if m.Bootconfig != nil {
		if err := m.Bootconfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootconfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1PartitionCreateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PartitionCreateRequest) validateProjectnetworkprefixlength(formats strfmt.Registry) error {

	if swag.IsZero(m.Projectnetworkprefixlength) { // not required
		return nil
	}

	if err := validate.MinimumInt("projectnetworkprefixlength", "body", int64(m.Projectnetworkprefixlength), 16, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("projectnetworkprefixlength", "body", int64(m.Projectnetworkprefixlength), 30, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PartitionCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PartitionCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1PartitionCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
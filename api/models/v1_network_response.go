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

// V1NetworkResponse v1 network response
// swagger:model v1.NetworkResponse
type V1NetworkResponse struct {

	// the last changed timestamp of this entity
	// Required: true
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed"`

	// the creation time of this entity
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the destination prefixes of this network
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes"`

	// the unique ID of this entity
	// Required: true
	// Unique: true
	ID *string `json:"id"`

	// free labels that you associate with this network.
	// Required: true
	Labels map[string]string `json:"labels"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// if set to true, packets leaving this network get masqueraded behind interface ip
	// Required: true
	Nat *bool `json:"nat"`

	// the id of the parent network
	// Required: true
	Parentnetworkid *string `json:"parentnetworkid"`

	// the partition this network belongs to
	Partitionid string `json:"partitionid,omitempty"`

	// the prefixes of this network
	// Required: true
	Prefixes []string `json:"prefixes"`

	// if set to true, this network will serve as a partition's super network for the internal machine networks,there can only be one privatesuper network per partition
	// Required: true
	Privatesuper *bool `json:"privatesuper"`

	// the project id this network belongs to, can be empty if globally available
	Projectid string `json:"projectid,omitempty"`

	// if set to true, this network can be used for underlay communication
	// Required: true
	Underlay *bool `json:"underlay"`

	// usage of ips and prefixes in this network
	// Required: true
	Usage *V1NetworkUsage `json:"usage"`

	// the vrf this network is associated with
	Vrf int64 `json:"vrf,omitempty"`

	// if set to true, given vrf can be used by multiple networks, which is sometimes useful for network partioning (default: false)
	Vrfshared bool `json:"vrfshared,omitempty"`
}

// Validate validates this v1 network response
func (m *V1NetworkResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationprefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentnetworkid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivatesuper(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderlay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkResponse) validateChanged(formats strfmt.Registry) error {

	if err := validate.Required("changed", "body", strfmt.DateTime(m.Changed)); err != nil {
		return err
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateDestinationprefixes(formats strfmt.Registry) error {

	if err := validate.Required("destinationprefixes", "body", m.Destinationprefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateLabels(formats strfmt.Registry) error {

	return nil
}

func (m *V1NetworkResponse) validateNat(formats strfmt.Registry) error {

	if err := validate.Required("nat", "body", m.Nat); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateParentnetworkid(formats strfmt.Registry) error {

	if err := validate.Required("parentnetworkid", "body", m.Parentnetworkid); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validatePrivatesuper(formats strfmt.Registry) error {

	if err := validate.Required("privatesuper", "body", m.Privatesuper); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateUnderlay(formats strfmt.Registry) error {

	if err := validate.Required("underlay", "body", m.Underlay); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkResponse) UnmarshalBinary(b []byte) error {
	var res V1NetworkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

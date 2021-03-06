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

// PrivateGetMarginsResponse private get margins response
// swagger:model private_get_margins_response
type PrivateGetMarginsResponse struct {

	// result
	// Required: true
	Result *PrivateGetMarginsResponseResult `json:"result"`
}

// Validate validates this private get margins response
func (m *PrivateGetMarginsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetMarginsResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetMarginsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetMarginsResponse) UnmarshalBinary(b []byte) error {
	var res PrivateGetMarginsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateGetMarginsResponseResult private get margins response result
// swagger:model PrivateGetMarginsResponseResult
type PrivateGetMarginsResponseResult struct {

	// Margin when buying
	// Required: true
	Buy *float64 `json:"buy"`

	// max price
	// Required: true
	MaxPrice MaxPrice `json:"max_price"`

	// min price
	// Required: true
	MinPrice MinPrice `json:"min_price"`

	// Margin when selling
	// Required: true
	Sell *float64 `json:"sell"`
}

// Validate validates this private get margins response result
func (m *PrivateGetMarginsResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSell(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateGetMarginsResponseResult) validateBuy(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"buy", "body", m.Buy); err != nil {
		return err
	}

	return nil
}

func (m *PrivateGetMarginsResponseResult) validateMaxPrice(formats strfmt.Registry) error {

	if err := m.MaxPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result" + "." + "max_price")
		}
		return err
	}

	return nil
}

func (m *PrivateGetMarginsResponseResult) validateMinPrice(formats strfmt.Registry) error {

	if err := m.MinPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result" + "." + "min_price")
		}
		return err
	}

	return nil
}

func (m *PrivateGetMarginsResponseResult) validateSell(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"sell", "body", m.Sell); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetMarginsResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetMarginsResponseResult) UnmarshalBinary(b []byte) error {
	var res PrivateGetMarginsResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

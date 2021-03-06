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

// PublicGetFundingChartDataResponse public get funding chart data response
// swagger:model public_get_funding_chart_data_response
type PublicGetFundingChartDataResponse struct {

	// result
	// Required: true
	Result *PublicGetFundingChartDataResponseResult `json:"result"`
}

// Validate validates this public get funding chart data response
func (m *PublicGetFundingChartDataResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicGetFundingChartDataResponse) validateResult(formats strfmt.Registry) error {

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
func (m *PublicGetFundingChartDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicGetFundingChartDataResponse) UnmarshalBinary(b []byte) error {
	var res PublicGetFundingChartDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicGetFundingChartDataResponseResult public get funding chart data response result
// swagger:model PublicGetFundingChartDataResponseResult
type PublicGetFundingChartDataResponseResult struct {

	// Current interest
	// Required: true
	CurrentInterest *float64 `json:"current_interest"`

	// data
	// Required: true
	Data []string `json:"data"`

	// Current index price
	// Required: true
	IndexPrice *float64 `json:"index_price"`

	// Current interest 8h
	// Required: true
	Interest8h *float64 `json:"interest_8h"`

	// maximal interest
	// Required: true
	Max *float64 `json:"max"`
}

// Validate validates this public get funding chart data response result
func (m *PublicGetFundingChartDataResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentInterest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterest8h(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicGetFundingChartDataResponseResult) validateCurrentInterest(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"current_interest", "body", m.CurrentInterest); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseResult) validateData(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseResult) validateIndexPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"index_price", "body", m.IndexPrice); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseResult) validateInterest8h(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"interest_8h", "body", m.Interest8h); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseResult) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicGetFundingChartDataResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicGetFundingChartDataResponseResult) UnmarshalBinary(b []byte) error {
	var res PublicGetFundingChartDataResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tuanito/go-deribit/v3/models"
)

// GetPublicGetHistoricalVolatilityReader is a Reader for the GetPublicGetHistoricalVolatility structure.
type GetPublicGetHistoricalVolatilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetHistoricalVolatilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetHistoricalVolatilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicGetHistoricalVolatilityOK creates a GetPublicGetHistoricalVolatilityOK with default headers values
func NewGetPublicGetHistoricalVolatilityOK() *GetPublicGetHistoricalVolatilityOK {
	return &GetPublicGetHistoricalVolatilityOK{}
}

/*GetPublicGetHistoricalVolatilityOK handles this case with default header values.

ok response
*/
type GetPublicGetHistoricalVolatilityOK struct {
	Payload *models.PublicGetHistoricalVolatilityResponse
}

func (o *GetPublicGetHistoricalVolatilityOK) Error() string {
	return fmt.Sprintf("[GET /public/get_historical_volatility][%d] getPublicGetHistoricalVolatilityOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetHistoricalVolatilityOK) GetPayload() *models.PublicGetHistoricalVolatilityResponse {
	return o.Payload
}

func (o *GetPublicGetHistoricalVolatilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGetHistoricalVolatilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

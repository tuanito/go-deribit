// Code generated by go-swagger; DO NOT EDIT.

package market_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tuanito/go-deribit/v3/models"
)

// GetPublicGetLastSettlementsByInstrumentReader is a Reader for the GetPublicGetLastSettlementsByInstrument structure.
type GetPublicGetLastSettlementsByInstrumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetLastSettlementsByInstrumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetLastSettlementsByInstrumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicGetLastSettlementsByInstrumentOK creates a GetPublicGetLastSettlementsByInstrumentOK with default headers values
func NewGetPublicGetLastSettlementsByInstrumentOK() *GetPublicGetLastSettlementsByInstrumentOK {
	return &GetPublicGetLastSettlementsByInstrumentOK{}
}

/*GetPublicGetLastSettlementsByInstrumentOK handles this case with default header values.

GetPublicGetLastSettlementsByInstrumentOK get public get last settlements by instrument o k
*/
type GetPublicGetLastSettlementsByInstrumentOK struct {
	Payload *models.PublicSettlementResponse
}

func (o *GetPublicGetLastSettlementsByInstrumentOK) Error() string {
	return fmt.Sprintf("[GET /public/get_last_settlements_by_instrument][%d] getPublicGetLastSettlementsByInstrumentOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetLastSettlementsByInstrumentOK) GetPayload() *models.PublicSettlementResponse {
	return o.Payload
}

func (o *GetPublicGetLastSettlementsByInstrumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicSettlementResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

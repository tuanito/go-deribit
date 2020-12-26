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

// GetPublicGetFundingChartDataReader is a Reader for the GetPublicGetFundingChartData structure.
type GetPublicGetFundingChartDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetFundingChartDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetFundingChartDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicGetFundingChartDataOK creates a GetPublicGetFundingChartDataOK with default headers values
func NewGetPublicGetFundingChartDataOK() *GetPublicGetFundingChartDataOK {
	return &GetPublicGetFundingChartDataOK{}
}

/*GetPublicGetFundingChartDataOK handles this case with default header values.

GetPublicGetFundingChartDataOK get public get funding chart data o k
*/
type GetPublicGetFundingChartDataOK struct {
	Payload *models.PublicGetFundingChartDataResponse
}

func (o *GetPublicGetFundingChartDataOK) Error() string {
	return fmt.Sprintf("[GET /public/get_funding_chart_data][%d] getPublicGetFundingChartDataOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetFundingChartDataOK) GetPayload() *models.PublicGetFundingChartDataResponse {
	return o.Payload
}

func (o *GetPublicGetFundingChartDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGetFundingChartDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

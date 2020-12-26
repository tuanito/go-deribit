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

// GetPublicGetContractSizeReader is a Reader for the GetPublicGetContractSize structure.
type GetPublicGetContractSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetContractSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetContractSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicGetContractSizeOK creates a GetPublicGetContractSizeOK with default headers values
func NewGetPublicGetContractSizeOK() *GetPublicGetContractSizeOK {
	return &GetPublicGetContractSizeOK{}
}

/*GetPublicGetContractSizeOK handles this case with default header values.

ok response
*/
type GetPublicGetContractSizeOK struct {
	Payload *models.PublicGetContractSizeResponse
}

func (o *GetPublicGetContractSizeOK) Error() string {
	return fmt.Sprintf("[GET /public/get_contract_size][%d] getPublicGetContractSizeOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetContractSizeOK) GetPayload() *models.PublicGetContractSizeResponse {
	return o.Payload
}

func (o *GetPublicGetContractSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGetContractSizeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

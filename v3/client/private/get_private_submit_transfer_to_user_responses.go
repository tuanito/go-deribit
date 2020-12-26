// Code generated by go-swagger; DO NOT EDIT.

package private

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tuanito/go-deribit/v3/models"
)

// GetPrivateSubmitTransferToUserReader is a Reader for the GetPrivateSubmitTransferToUser structure.
type GetPrivateSubmitTransferToUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateSubmitTransferToUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateSubmitTransferToUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSubmitTransferToUserOK creates a GetPrivateSubmitTransferToUserOK with default headers values
func NewGetPrivateSubmitTransferToUserOK() *GetPrivateSubmitTransferToUserOK {
	return &GetPrivateSubmitTransferToUserOK{}
}

/*GetPrivateSubmitTransferToUserOK handles this case with default header values.

GetPrivateSubmitTransferToUserOK get private submit transfer to user o k
*/
type GetPrivateSubmitTransferToUserOK struct {
	Payload *models.PrivateSubmitTransferResponse
}

func (o *GetPrivateSubmitTransferToUserOK) Error() string {
	return fmt.Sprintf("[GET /private/submit_transfer_to_user][%d] getPrivateSubmitTransferToUserOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSubmitTransferToUserOK) GetPayload() *models.PrivateSubmitTransferResponse {
	return o.Payload
}

func (o *GetPrivateSubmitTransferToUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSubmitTransferResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

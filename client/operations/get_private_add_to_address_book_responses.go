// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tuanito/go-deribit/models"
)

// GetPrivateAddToAddressBookReader is a Reader for the GetPrivateAddToAddressBook structure.
type GetPrivateAddToAddressBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateAddToAddressBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateAddToAddressBookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateAddToAddressBookOK creates a GetPrivateAddToAddressBookOK with default headers values
func NewGetPrivateAddToAddressBookOK() *GetPrivateAddToAddressBookOK {
	return &GetPrivateAddToAddressBookOK{}
}

/*GetPrivateAddToAddressBookOK handles this case with default header values.

foo
*/
type GetPrivateAddToAddressBookOK struct {
	Payload *models.PrivateAddToAddressBookResponse
}

func (o *GetPrivateAddToAddressBookOK) Error() string {
	return fmt.Sprintf("[GET /private/add_to_address_book][%d] getPrivateAddToAddressBookOK  %+v", 200, o.Payload)
}

func (o *GetPrivateAddToAddressBookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateAddToAddressBookResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

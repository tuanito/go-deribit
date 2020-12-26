// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tuanito/go-deribit/v3/models"
)

// GetPrivateGetNewAnnouncementsReader is a Reader for the GetPrivateGetNewAnnouncements structure.
type GetPrivateGetNewAnnouncementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetNewAnnouncementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetNewAnnouncementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetNewAnnouncementsOK creates a GetPrivateGetNewAnnouncementsOK with default headers values
func NewGetPrivateGetNewAnnouncementsOK() *GetPrivateGetNewAnnouncementsOK {
	return &GetPrivateGetNewAnnouncementsOK{}
}

/*GetPrivateGetNewAnnouncementsOK handles this case with default header values.

GetPrivateGetNewAnnouncementsOK get private get new announcements o k
*/
type GetPrivateGetNewAnnouncementsOK struct {
	Payload *models.PublicGetAnnouncementsResponse
}

func (o *GetPrivateGetNewAnnouncementsOK) Error() string {
	return fmt.Sprintf("[GET /private/get_new_announcements][%d] getPrivateGetNewAnnouncementsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetNewAnnouncementsOK) GetPayload() *models.PublicGetAnnouncementsResponse {
	return o.Payload
}

func (o *GetPrivateGetNewAnnouncementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGetAnnouncementsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

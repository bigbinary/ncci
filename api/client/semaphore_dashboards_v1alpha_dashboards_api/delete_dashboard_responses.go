// Code generated by go-swagger; DO NOT EDIT.

package semaphore_dashboards_v1alpha_dashboards_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semaphoreci/cli/api/models"
)

// DeleteDashboardReader is a Reader for the DeleteDashboard structure.
type DeleteDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDashboardOK creates a DeleteDashboardOK with default headers values
func NewDeleteDashboardOK() *DeleteDashboardOK {
	return &DeleteDashboardOK{}
}

/*DeleteDashboardOK handles this case with default header values.

(empty)
*/
type DeleteDashboardOK struct {
	Payload models.SemaphoreDashboardsV1alphaEmpty
}

func (o *DeleteDashboardOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha/dashboards/{id_or_name}][%d] deleteDashboardOK  %+v", 200, o.Payload)
}

func (o *DeleteDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

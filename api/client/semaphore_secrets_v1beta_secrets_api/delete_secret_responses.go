// Code generated by go-swagger; DO NOT EDIT.

package semaphore_secrets_v1beta_secrets_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semaphoreci/cli/api/models"
)

// DeleteSecretReader is a Reader for the DeleteSecret structure.
type DeleteSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSecretOK creates a DeleteSecretOK with default headers values
func NewDeleteSecretOK() *DeleteSecretOK {
	return &DeleteSecretOK{}
}

/*DeleteSecretOK handles this case with default header values.

(empty)
*/
type DeleteSecretOK struct {
	Payload models.SemaphoreSecretsV1betaEmpty
}

func (o *DeleteSecretOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1beta/secrets/{secret_id_or_name}][%d] deleteSecretOK  %+v", 200, o.Payload)
}

func (o *DeleteSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

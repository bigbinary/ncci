// Code generated by go-swagger; DO NOT EDIT.

package semaphore_secrets_v1beta_secrets_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semaphoreci/cli/api/models"
)

// NewCreateSecretParams creates a new CreateSecretParams object
// with the default values initialized.
func NewCreateSecretParams() *CreateSecretParams {
	var ()
	return &CreateSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecretParamsWithTimeout creates a new CreateSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSecretParamsWithTimeout(timeout time.Duration) *CreateSecretParams {
	var ()
	return &CreateSecretParams{

		timeout: timeout,
	}
}

// NewCreateSecretParamsWithContext creates a new CreateSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSecretParamsWithContext(ctx context.Context) *CreateSecretParams {
	var ()
	return &CreateSecretParams{

		Context: ctx,
	}
}

// NewCreateSecretParamsWithHTTPClient creates a new CreateSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSecretParamsWithHTTPClient(client *http.Client) *CreateSecretParams {
	var ()
	return &CreateSecretParams{
		HTTPClient: client,
	}
}

/*CreateSecretParams contains all the parameters to send to the API endpoint
for the create secret operation typically these are written to a http.Request
*/
type CreateSecretParams struct {

	/*Body*/
	Body *models.SemaphoreSecretsV1betaSecret

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create secret params
func (o *CreateSecretParams) WithTimeout(timeout time.Duration) *CreateSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create secret params
func (o *CreateSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create secret params
func (o *CreateSecretParams) WithContext(ctx context.Context) *CreateSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create secret params
func (o *CreateSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create secret params
func (o *CreateSecretParams) WithHTTPClient(client *http.Client) *CreateSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create secret params
func (o *CreateSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create secret params
func (o *CreateSecretParams) WithBody(body *models.SemaphoreSecretsV1betaSecret) *CreateSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create secret params
func (o *CreateSecretParams) SetBody(body *models.SemaphoreSecretsV1betaSecret) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

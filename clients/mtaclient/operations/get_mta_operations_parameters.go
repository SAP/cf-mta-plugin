// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMtaOperationsParams creates a new GetMtaOperationsParams object
// with the default values initialized.
func NewGetMtaOperationsParams() *GetMtaOperationsParams {
	var ()
	return &GetMtaOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMtaOperationsParamsWithTimeout creates a new GetMtaOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMtaOperationsParamsWithTimeout(timeout time.Duration) *GetMtaOperationsParams {
	var ()
	return &GetMtaOperationsParams{

		timeout: timeout,
	}
}

// NewGetMtaOperationsParamsWithContext creates a new GetMtaOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMtaOperationsParamsWithContext(ctx context.Context) *GetMtaOperationsParams {
	var ()
	return &GetMtaOperationsParams{

		Context: ctx,
	}
}

// NewGetMtaOperationsParamsWithHTTPClient creates a new GetMtaOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMtaOperationsParamsWithHTTPClient(client *http.Client) *GetMtaOperationsParams {
	var ()
	return &GetMtaOperationsParams{
		HTTPClient: client,
	}
}

/*GetMtaOperationsParams contains all the parameters to send to the API endpoint
for the get mta operations operation typically these are written to a http.Request
*/
type GetMtaOperationsParams struct {

	/*Last*/
	Last *int64
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mta operations params
func (o *GetMtaOperationsParams) WithTimeout(timeout time.Duration) *GetMtaOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mta operations params
func (o *GetMtaOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mta operations params
func (o *GetMtaOperationsParams) WithContext(ctx context.Context) *GetMtaOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mta operations params
func (o *GetMtaOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mta operations params
func (o *GetMtaOperationsParams) WithHTTPClient(client *http.Client) *GetMtaOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mta operations params
func (o *GetMtaOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLast adds the last to the get mta operations params
func (o *GetMtaOperationsParams) WithLast(last *int64) *GetMtaOperationsParams {
	o.SetLast(last)
	return o
}

// SetLast adds the last to the get mta operations params
func (o *GetMtaOperationsParams) SetLast(last *int64) {
	o.Last = last
}

// WithStatus adds the status to the get mta operations params
func (o *GetMtaOperationsParams) WithStatus(status []string) *GetMtaOperationsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get mta operations params
func (o *GetMtaOperationsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetMtaOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Last != nil {

		// query param last
		var qrLast int64
		if o.Last != nil {
			qrLast = *o.Last
		}
		qLast := swag.FormatInt64(qrLast)
		if qLast != "" {
			if err := r.SetQueryParam("last", qLast); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

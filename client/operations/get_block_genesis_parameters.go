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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBlockGenesisParams creates a new GetBlockGenesisParams object
// with the default values initialized.
func NewGetBlockGenesisParams() *GetBlockGenesisParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockGenesisParams{
		TxEncoding: &txEncodingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlockGenesisParamsWithTimeout creates a new GetBlockGenesisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlockGenesisParamsWithTimeout(timeout time.Duration) *GetBlockGenesisParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockGenesisParams{
		TxEncoding: &txEncodingDefault,

		timeout: timeout,
	}
}

// NewGetBlockGenesisParamsWithContext creates a new GetBlockGenesisParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlockGenesisParamsWithContext(ctx context.Context) *GetBlockGenesisParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockGenesisParams{
		TxEncoding: &txEncodingDefault,

		Context: ctx,
	}
}

// NewGetBlockGenesisParamsWithHTTPClient creates a new GetBlockGenesisParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlockGenesisParamsWithHTTPClient(client *http.Client) *GetBlockGenesisParams {
	var (
		txEncodingDefault = string("json")
	)
	return &GetBlockGenesisParams{
		TxEncoding: &txEncodingDefault,
		HTTPClient: client,
	}
}

/*GetBlockGenesisParams contains all the parameters to send to the API endpoint
for the get block genesis operation typically these are written to a http.Request
*/
type GetBlockGenesisParams struct {

	/*TxEncoding
	  Transactions encoding

	*/
	TxEncoding *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get block genesis params
func (o *GetBlockGenesisParams) WithTimeout(timeout time.Duration) *GetBlockGenesisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get block genesis params
func (o *GetBlockGenesisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get block genesis params
func (o *GetBlockGenesisParams) WithContext(ctx context.Context) *GetBlockGenesisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get block genesis params
func (o *GetBlockGenesisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get block genesis params
func (o *GetBlockGenesisParams) WithHTTPClient(client *http.Client) *GetBlockGenesisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get block genesis params
func (o *GetBlockGenesisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTxEncoding adds the txEncoding to the get block genesis params
func (o *GetBlockGenesisParams) WithTxEncoding(txEncoding *string) *GetBlockGenesisParams {
	o.SetTxEncoding(txEncoding)
	return o
}

// SetTxEncoding adds the txEncoding to the get block genesis params
func (o *GetBlockGenesisParams) SetTxEncoding(txEncoding *string) {
	o.TxEncoding = txEncoding
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlockGenesisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TxEncoding != nil {

		// query param tx_encoding
		var qrTxEncoding string
		if o.TxEncoding != nil {
			qrTxEncoding = *o.TxEncoding
		}
		qTxEncoding := qrTxEncoding
		if qTxEncoding != "" {
			if err := r.SetQueryParam("tx_encoding", qTxEncoding); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
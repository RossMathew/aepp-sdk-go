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

// NewGetTransactionFromBlockLatestParams creates a new GetTransactionFromBlockLatestParams object
// with the default values initialized.
func NewGetTransactionFromBlockLatestParams() *GetTransactionFromBlockLatestParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockLatestParams{
		TxEncoding: &txEncodingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionFromBlockLatestParamsWithTimeout creates a new GetTransactionFromBlockLatestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTransactionFromBlockLatestParamsWithTimeout(timeout time.Duration) *GetTransactionFromBlockLatestParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockLatestParams{
		TxEncoding: &txEncodingDefault,

		timeout: timeout,
	}
}

// NewGetTransactionFromBlockLatestParamsWithContext creates a new GetTransactionFromBlockLatestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTransactionFromBlockLatestParamsWithContext(ctx context.Context) *GetTransactionFromBlockLatestParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockLatestParams{
		TxEncoding: &txEncodingDefault,

		Context: ctx,
	}
}

// NewGetTransactionFromBlockLatestParamsWithHTTPClient creates a new GetTransactionFromBlockLatestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTransactionFromBlockLatestParamsWithHTTPClient(client *http.Client) *GetTransactionFromBlockLatestParams {
	var (
		txEncodingDefault = string("message_pack")
	)
	return &GetTransactionFromBlockLatestParams{
		TxEncoding: &txEncodingDefault,
		HTTPClient: client,
	}
}

/*GetTransactionFromBlockLatestParams contains all the parameters to send to the API endpoint
for the get transaction from block latest operation typically these are written to a http.Request
*/
type GetTransactionFromBlockLatestParams struct {

	/*TxEncoding
	  Transactions encoding

	*/
	TxEncoding *string
	/*TxIndex
	  Index of the transaction in the block

	*/
	TxIndex int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) WithTimeout(timeout time.Duration) *GetTransactionFromBlockLatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) WithContext(ctx context.Context) *GetTransactionFromBlockLatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) WithHTTPClient(client *http.Client) *GetTransactionFromBlockLatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTxEncoding adds the txEncoding to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) WithTxEncoding(txEncoding *string) *GetTransactionFromBlockLatestParams {
	o.SetTxEncoding(txEncoding)
	return o
}

// SetTxEncoding adds the txEncoding to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) SetTxEncoding(txEncoding *string) {
	o.TxEncoding = txEncoding
}

// WithTxIndex adds the txIndex to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) WithTxIndex(txIndex int64) *GetTransactionFromBlockLatestParams {
	o.SetTxIndex(txIndex)
	return o
}

// SetTxIndex adds the txIndex to the get transaction from block latest params
func (o *GetTransactionFromBlockLatestParams) SetTxIndex(txIndex int64) {
	o.TxIndex = txIndex
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionFromBlockLatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tx_index
	if err := r.SetPathParam("tx_index", swag.FormatInt64(o.TxIndex)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

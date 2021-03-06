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

// NewGetLatestBlockTxsCountParams creates a new GetLatestBlockTxsCountParams object
// with the default values initialized.
func NewGetLatestBlockTxsCountParams() *GetLatestBlockTxsCountParams {
	var ()
	return &GetLatestBlockTxsCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestBlockTxsCountParamsWithTimeout creates a new GetLatestBlockTxsCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLatestBlockTxsCountParamsWithTimeout(timeout time.Duration) *GetLatestBlockTxsCountParams {
	var ()
	return &GetLatestBlockTxsCountParams{

		timeout: timeout,
	}
}

// NewGetLatestBlockTxsCountParamsWithContext creates a new GetLatestBlockTxsCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLatestBlockTxsCountParamsWithContext(ctx context.Context) *GetLatestBlockTxsCountParams {
	var ()
	return &GetLatestBlockTxsCountParams{

		Context: ctx,
	}
}

// NewGetLatestBlockTxsCountParamsWithHTTPClient creates a new GetLatestBlockTxsCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLatestBlockTxsCountParamsWithHTTPClient(client *http.Client) *GetLatestBlockTxsCountParams {
	var ()
	return &GetLatestBlockTxsCountParams{
		HTTPClient: client,
	}
}

/*GetLatestBlockTxsCountParams contains all the parameters to send to the API endpoint
for the get latest block txs count operation typically these are written to a http.Request
*/
type GetLatestBlockTxsCountParams struct {

	/*ExcludeTxTypes
	  Transactions types not to show. Comma separated. If a tx type appears in both tx_types and exclude_tx_types, it is excluded.

	*/
	ExcludeTxTypes *string
	/*TxTypes
	  Transactions types to show. Comma separated

	*/
	TxTypes *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) WithTimeout(timeout time.Duration) *GetLatestBlockTxsCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) WithContext(ctx context.Context) *GetLatestBlockTxsCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) WithHTTPClient(client *http.Client) *GetLatestBlockTxsCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeTxTypes adds the excludeTxTypes to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) WithExcludeTxTypes(excludeTxTypes *string) *GetLatestBlockTxsCountParams {
	o.SetExcludeTxTypes(excludeTxTypes)
	return o
}

// SetExcludeTxTypes adds the excludeTxTypes to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) SetExcludeTxTypes(excludeTxTypes *string) {
	o.ExcludeTxTypes = excludeTxTypes
}

// WithTxTypes adds the txTypes to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) WithTxTypes(txTypes *string) *GetLatestBlockTxsCountParams {
	o.SetTxTypes(txTypes)
	return o
}

// SetTxTypes adds the txTypes to the get latest block txs count params
func (o *GetLatestBlockTxsCountParams) SetTxTypes(txTypes *string) {
	o.TxTypes = txTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestBlockTxsCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeTxTypes != nil {

		// query param exclude_tx_types
		var qrExcludeTxTypes string
		if o.ExcludeTxTypes != nil {
			qrExcludeTxTypes = *o.ExcludeTxTypes
		}
		qExcludeTxTypes := qrExcludeTxTypes
		if qExcludeTxTypes != "" {
			if err := r.SetQueryParam("exclude_tx_types", qExcludeTxTypes); err != nil {
				return err
			}
		}

	}

	if o.TxTypes != nil {

		// query param tx_types
		var qrTxTypes string
		if o.TxTypes != nil {
			qrTxTypes = *o.TxTypes
		}
		qTxTypes := qrTxTypes
		if qTxTypes != "" {
			if err := r.SetQueryParam("tx_types", qTxTypes); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

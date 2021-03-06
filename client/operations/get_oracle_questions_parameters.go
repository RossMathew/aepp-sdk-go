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

// NewGetOracleQuestionsParams creates a new GetOracleQuestionsParams object
// with the default values initialized.
func NewGetOracleQuestionsParams() *GetOracleQuestionsParams {
	var ()
	return &GetOracleQuestionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOracleQuestionsParamsWithTimeout creates a new GetOracleQuestionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOracleQuestionsParamsWithTimeout(timeout time.Duration) *GetOracleQuestionsParams {
	var ()
	return &GetOracleQuestionsParams{

		timeout: timeout,
	}
}

// NewGetOracleQuestionsParamsWithContext creates a new GetOracleQuestionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOracleQuestionsParamsWithContext(ctx context.Context) *GetOracleQuestionsParams {
	var ()
	return &GetOracleQuestionsParams{

		Context: ctx,
	}
}

// NewGetOracleQuestionsParamsWithHTTPClient creates a new GetOracleQuestionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOracleQuestionsParamsWithHTTPClient(client *http.Client) *GetOracleQuestionsParams {
	var ()
	return &GetOracleQuestionsParams{
		HTTPClient: client,
	}
}

/*GetOracleQuestionsParams contains all the parameters to send to the API endpoint
for the get oracle questions operation typically these are written to a http.Request
*/
type GetOracleQuestionsParams struct {

	/*From
	  Last query id in previous page

	*/
	From *string
	/*Max
	  Max number of oracle queries received

	*/
	Max *int64
	/*OraclePubKey
	  Oracle public key

	*/
	OraclePubKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oracle questions params
func (o *GetOracleQuestionsParams) WithTimeout(timeout time.Duration) *GetOracleQuestionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oracle questions params
func (o *GetOracleQuestionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oracle questions params
func (o *GetOracleQuestionsParams) WithContext(ctx context.Context) *GetOracleQuestionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oracle questions params
func (o *GetOracleQuestionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oracle questions params
func (o *GetOracleQuestionsParams) WithHTTPClient(client *http.Client) *GetOracleQuestionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oracle questions params
func (o *GetOracleQuestionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get oracle questions params
func (o *GetOracleQuestionsParams) WithFrom(from *string) *GetOracleQuestionsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get oracle questions params
func (o *GetOracleQuestionsParams) SetFrom(from *string) {
	o.From = from
}

// WithMax adds the max to the get oracle questions params
func (o *GetOracleQuestionsParams) WithMax(max *int64) *GetOracleQuestionsParams {
	o.SetMax(max)
	return o
}

// SetMax adds the max to the get oracle questions params
func (o *GetOracleQuestionsParams) SetMax(max *int64) {
	o.Max = max
}

// WithOraclePubKey adds the oraclePubKey to the get oracle questions params
func (o *GetOracleQuestionsParams) WithOraclePubKey(oraclePubKey string) *GetOracleQuestionsParams {
	o.SetOraclePubKey(oraclePubKey)
	return o
}

// SetOraclePubKey adds the oraclePubKey to the get oracle questions params
func (o *GetOracleQuestionsParams) SetOraclePubKey(oraclePubKey string) {
	o.OraclePubKey = oraclePubKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetOracleQuestionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Max != nil {

		// query param max
		var qrMax int64
		if o.Max != nil {
			qrMax = *o.Max
		}
		qMax := swag.FormatInt64(qrMax)
		if qMax != "" {
			if err := r.SetQueryParam("max", qMax); err != nil {
				return err
			}
		}

	}

	// query param oracle_pub_key
	qrOraclePubKey := o.OraclePubKey
	qOraclePubKey := qrOraclePubKey
	if qOraclePubKey != "" {
		if err := r.SetQueryParam("oracle_pub_key", qOraclePubKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

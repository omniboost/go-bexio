package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewManualEntriesPostRequest() ManualEntriesPostRequest {
	r := ManualEntriesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ManualEntriesPostRequest struct {
	client      *Client
	queryParams *ManualEntriesPostRequestQueryParams
	pathParams  *ManualEntriesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody ManualEntriesPostRequestBody
}

func (r ManualEntriesPostRequest) NewQueryParams() *ManualEntriesPostRequestQueryParams {
	return &ManualEntriesPostRequestQueryParams{}
}

type ManualEntriesPostRequestQueryParams struct{}

func (p ManualEntriesPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ManualEntriesPostRequest) QueryParams() *ManualEntriesPostRequestQueryParams {
	return r.queryParams
}

func (r ManualEntriesPostRequest) NewPathParams() *ManualEntriesPostRequestPathParams {
	return &ManualEntriesPostRequestPathParams{}
}

type ManualEntriesPostRequestPathParams struct {
}

func (p *ManualEntriesPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ManualEntriesPostRequest) PathParams() *ManualEntriesPostRequestPathParams {
	return r.pathParams
}

func (r *ManualEntriesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ManualEntriesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *ManualEntriesPostRequest) Method() string {
	return r.method
}

func (r ManualEntriesPostRequest) NewRequestBody() ManualEntriesPostRequestBody {
	return ManualEntriesPostRequestBody{}
}

type ManualEntriesPostRequestBody struct {
	Type        string `json:"type"`
	Date        Date   `json:"date"`
	ReferenceNr string `json:"reference_nr"`
	Entries     []struct {
		DebitAccountID  int     `json:"debit_account_id"`
		CreditAccountID int     `json:"credit_account_id"`
		TaxID           int     `json:"tax_id,omitempty"`
		TaxAccountID    int     `json:"tax_account_id,omitempty"`
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		CurrencyID      int     `json:"currency_id,omitempty"`
		CurrencyFactor  int     `json:"currency_factor,omitempty"`
	} `json:"entries"`
}

func (r *ManualEntriesPostRequest) RequestBody() *ManualEntriesPostRequestBody {
	return &r.requestBody
}

func (r *ManualEntriesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ManualEntriesPostRequest) SetRequestBody(body ManualEntriesPostRequestBody) {
	r.requestBody = body
}

func (r *ManualEntriesPostRequest) NewResponseBody() *ManualEntriesPostResponseBody {
	return &ManualEntriesPostResponseBody{}
}

type ManualEntriesPostResponseBody Accounts

func (r *ManualEntriesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/accounting/manual_entries", r.PathParams())
	return &u
}

func (r *ManualEntriesPostRequest) Do() (ManualEntriesPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

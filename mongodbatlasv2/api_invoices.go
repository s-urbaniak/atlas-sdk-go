/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type InvoicesApi interface {

	/*
	DownloadInvoiceCSV Return One Organization Invoice as CSV

	Returns one invoice that MongoDB issued to the specified organization in CSV format. A unique 24-hexadecimal digit string identifies the invoice. To use this resource, the requesting API Key must have the Organization Member role. If you have a cross-organization setup, you can query for a linked invoice if you have an Organization Billing Admin or Organization Owner Role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return InvoicesApiDownloadInvoiceCSVRequest
	*/
	DownloadInvoiceCSV(ctx context.Context, orgId string, invoiceId string) InvoicesApiDownloadInvoiceCSVRequest

	// DownloadInvoiceCSVExecute executes the request
	DownloadInvoiceCSVExecute(r InvoicesApiDownloadInvoiceCSVRequest) (*http.Response, error)

	/*
	GetInvoice Return One Organization Invoice

	Returns one invoice that MongoDB issued to the specified organization. A unique 24-hexadecimal digit string identifies the invoice. You can choose to receive this invoice in JSON or CSV format. To use this resource, the requesting API Key must have the Organization Member role. If you have a cross-organization setup, you can query for a linked invoice if you have an Organization Billing Admin or Organization Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return InvoicesApiGetInvoiceRequest
	*/
	GetInvoice(ctx context.Context, orgId string, invoiceId string) InvoicesApiGetInvoiceRequest

	// GetInvoiceExecute executes the request
	//  @return Invoice
	GetInvoiceExecute(r InvoicesApiGetInvoiceRequest) (*Invoice, *http.Response, error)

	/*
	ListInvoices Return All Invoices for One Organization

	Returns all invoices that MongoDB issued to the specified organization. This list includes all invoices regardless of invoice status. To use this resource, the requesting API Key must have the Organization Member role. If you have a cross-organization setup, to view linked invoices, you must have an Organization Billing Admin or Organization Owner role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return InvoicesApiListInvoicesRequest
	*/
	ListInvoices(ctx context.Context, orgId string) InvoicesApiListInvoicesRequest

	// ListInvoicesExecute executes the request
	//  @return PaginatedApiInvoice
	ListInvoicesExecute(r InvoicesApiListInvoicesRequest) (*PaginatedApiInvoice, *http.Response, error)

	/*
	ListPendingInvoices Return All Pending Invoices for One Organization

	Returns all invoices accruing charges for the current billing cycle for the specified organization. To use this resource, the requesting API Key must have the Organization Member role.  If you have a cross-organization setup, to view linked invoices, you must have an Organization Billing Admin or Organization Owner Role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return InvoicesApiListPendingInvoicesRequest
	*/
	ListPendingInvoices(ctx context.Context, orgId string) InvoicesApiListPendingInvoicesRequest

	// ListPendingInvoicesExecute executes the request
	//  @return PaginatedApiInvoice
	ListPendingInvoicesExecute(r InvoicesApiListPendingInvoicesRequest) (*PaginatedApiInvoice, *http.Response, error)
}

// InvoicesApiService InvoicesApi service
type InvoicesApiService service

type InvoicesApiDownloadInvoiceCSVRequest struct {
	ctx context.Context
	ApiService InvoicesApi
	orgId string
	invoiceId string
}

func (r InvoicesApiDownloadInvoiceCSVRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadInvoiceCSVExecute(r)
}

/*
DownloadInvoiceCSV Return One Organization Invoice as CSV

Returns one invoice that MongoDB issued to the specified organization in CSV format. A unique 24-hexadecimal digit string identifies the invoice. To use this resource, the requesting API Key must have the Organization Member role. If you have a cross-organization setup, you can query for a linked invoice if you have an Organization Billing Admin or Organization Owner Role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
 @param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
 @return InvoicesApiDownloadInvoiceCSVRequest
*/
func (a *InvoicesApiService) DownloadInvoiceCSV(ctx context.Context, orgId string, invoiceId string) InvoicesApiDownloadInvoiceCSVRequest {
	return InvoicesApiDownloadInvoiceCSVRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
func (a *InvoicesApiService) DownloadInvoiceCSVExecute(r InvoicesApiDownloadInvoiceCSVRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.DownloadInvoiceCSV")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return nil, reportError("orgId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+csv", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type InvoicesApiGetInvoiceRequest struct {
	ctx context.Context
	ApiService InvoicesApi
	orgId string
	invoiceId string
}

func (r InvoicesApiGetInvoiceRequest) Execute() (*Invoice, *http.Response, error) {
	return r.ApiService.GetInvoiceExecute(r)
}

/*
GetInvoice Return One Organization Invoice

Returns one invoice that MongoDB issued to the specified organization. A unique 24-hexadecimal digit string identifies the invoice. You can choose to receive this invoice in JSON or CSV format. To use this resource, the requesting API Key must have the Organization Member role. If you have a cross-organization setup, you can query for a linked invoice if you have an Organization Billing Admin or Organization Owner role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
 @param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
 @return InvoicesApiGetInvoiceRequest
*/
func (a *InvoicesApiService) GetInvoice(ctx context.Context, orgId string, invoiceId string) InvoicesApiGetInvoiceRequest {
	return InvoicesApiGetInvoiceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return Invoice
func (a *InvoicesApiService) GetInvoiceExecute(r InvoicesApiGetInvoiceRequest) (*Invoice, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Invoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.invoiceId) < 24 {
		return localVarReturnValue, nil, reportError("invoiceId must have at least 24 elements")
	}
	if strlen(r.invoiceId) > 24 {
		return localVarReturnValue, nil, reportError("invoiceId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type InvoicesApiListInvoicesRequest struct {
	ctx context.Context
	ApiService InvoicesApi
	orgId string
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r InvoicesApiListInvoicesRequest) IncludeCount(includeCount bool) InvoicesApiListInvoicesRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r InvoicesApiListInvoicesRequest) ItemsPerPage(itemsPerPage int32) InvoicesApiListInvoicesRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r InvoicesApiListInvoicesRequest) PageNum(pageNum int32) InvoicesApiListInvoicesRequest {
	r.pageNum = &pageNum
	return r
}

func (r InvoicesApiListInvoicesRequest) Execute() (*PaginatedApiInvoice, *http.Response, error) {
	return r.ApiService.ListInvoicesExecute(r)
}

/*
ListInvoices Return All Invoices for One Organization

Returns all invoices that MongoDB issued to the specified organization. This list includes all invoices regardless of invoice status. To use this resource, the requesting API Key must have the Organization Member role. If you have a cross-organization setup, to view linked invoices, you must have an Organization Billing Admin or Organization Owner role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
 @return InvoicesApiListInvoicesRequest
*/
func (a *InvoicesApiService) ListInvoices(ctx context.Context, orgId string) InvoicesApiListInvoicesRequest {
	return InvoicesApiListInvoicesRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return PaginatedApiInvoice
func (a *InvoicesApiService) ListInvoicesExecute(r InvoicesApiListInvoicesRequest) (*PaginatedApiInvoice, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedApiInvoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int32 = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int32 = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type InvoicesApiListPendingInvoicesRequest struct {
	ctx context.Context
	ApiService InvoicesApi
	orgId string
}

func (r InvoicesApiListPendingInvoicesRequest) Execute() (*PaginatedApiInvoice, *http.Response, error) {
	return r.ApiService.ListPendingInvoicesExecute(r)
}

/*
ListPendingInvoices Return All Pending Invoices for One Organization

Returns all invoices accruing charges for the current billing cycle for the specified organization. To use this resource, the requesting API Key must have the Organization Member role.  If you have a cross-organization setup, to view linked invoices, you must have an Organization Billing Admin or Organization Owner Role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
 @return InvoicesApiListPendingInvoicesRequest
*/
func (a *InvoicesApiService) ListPendingInvoices(ctx context.Context, orgId string) InvoicesApiListPendingInvoicesRequest {
	return InvoicesApiListPendingInvoicesRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return PaginatedApiInvoice
func (a *InvoicesApiService) ListPendingInvoicesExecute(r InvoicesApiListPendingInvoicesRequest) (*PaginatedApiInvoice, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedApiInvoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListPendingInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/pending"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

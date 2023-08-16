// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type X509AuthenticationApi interface {

	/*
		CreateDatabaseUserCertificate Create One X.509 Certificate for One MongoDB User

		Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Owner role.

	To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.

	If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param username Human-readable label that represents the MongoDB database user account for whom to create a certificate.
		@return CreateDatabaseUserCertificateApiRequest
	*/
	CreateDatabaseUserCertificate(ctx context.Context, groupId string, username string, userCert *UserCert) CreateDatabaseUserCertificateApiRequest
	/*
		CreateDatabaseUserCertificate Create One X.509 Certificate for One MongoDB User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateDatabaseUserCertificateApiParams - Parameters for the request
		@return CreateDatabaseUserCertificateApiRequest
	*/
	CreateDatabaseUserCertificateWithParams(ctx context.Context, args *CreateDatabaseUserCertificateApiParams) CreateDatabaseUserCertificateApiRequest

	// Interface only available internally
	createDatabaseUserCertificateExecute(r CreateDatabaseUserCertificateApiRequest) (*http.Response, error)

	/*
		DisableCustomerManagedX509 Disable Customer-Managed X.509

		Clears the customer-managed X.509 settings on a project, including the uploaded Certificate Authority, which disables self-managed X.509.

	 Updating this configuration triggers a rolling restart of the database. You must have the Project Owner role to use this endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DisableCustomerManagedX509ApiRequest
	*/
	DisableCustomerManagedX509(ctx context.Context, groupId string) DisableCustomerManagedX509ApiRequest
	/*
		DisableCustomerManagedX509 Disable Customer-Managed X.509


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DisableCustomerManagedX509ApiParams - Parameters for the request
		@return DisableCustomerManagedX509ApiRequest
	*/
	DisableCustomerManagedX509WithParams(ctx context.Context, args *DisableCustomerManagedX509ApiParams) DisableCustomerManagedX509ApiRequest

	// Interface only available internally
	disableCustomerManagedX509Execute(r DisableCustomerManagedX509ApiRequest) (*UserSecurity, *http.Response, error)

	/*
		ListDatabaseUserCertificates Return All X.509 Certificates Assigned to One MongoDB User

		Returns all unexpired X.509 certificates for the specified MongoDB user. This MongoDB user belongs to one project. Atlas manages these certificates and the MongoDB user. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param username Human-readable label that represents the MongoDB database user account whose certificates you want to return.
		@return ListDatabaseUserCertificatesApiRequest
	*/
	ListDatabaseUserCertificates(ctx context.Context, groupId string, username string) ListDatabaseUserCertificatesApiRequest
	/*
		ListDatabaseUserCertificates Return All X.509 Certificates Assigned to One MongoDB User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListDatabaseUserCertificatesApiParams - Parameters for the request
		@return ListDatabaseUserCertificatesApiRequest
	*/
	ListDatabaseUserCertificatesWithParams(ctx context.Context, args *ListDatabaseUserCertificatesApiParams) ListDatabaseUserCertificatesApiRequest

	// Interface only available internally
	listDatabaseUserCertificatesExecute(r ListDatabaseUserCertificatesApiRequest) (*PaginatedUserCert, *http.Response, error)
}

// X509AuthenticationApiService X509AuthenticationApi service
type X509AuthenticationApiService service

type CreateDatabaseUserCertificateApiRequest struct {
	ctx        context.Context
	ApiService X509AuthenticationApi
	groupId    string
	username   string
	userCert   *UserCert
}

type CreateDatabaseUserCertificateApiParams struct {
	GroupId  string
	Username string
	UserCert *UserCert
}

func (a *X509AuthenticationApiService) CreateDatabaseUserCertificateWithParams(ctx context.Context, args *CreateDatabaseUserCertificateApiParams) CreateDatabaseUserCertificateApiRequest {
	return CreateDatabaseUserCertificateApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		username:   args.Username,
		userCert:   args.UserCert,
	}
}

func (r CreateDatabaseUserCertificateApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.createDatabaseUserCertificateExecute(r)
}

/*
CreateDatabaseUserCertificate Create One X.509 Certificate for One MongoDB User

Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Owner role.

To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.

If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param username Human-readable label that represents the MongoDB database user account for whom to create a certificate.
	@return CreateDatabaseUserCertificateApiRequest
*/
func (a *X509AuthenticationApiService) CreateDatabaseUserCertificate(ctx context.Context, groupId string, username string, userCert *UserCert) CreateDatabaseUserCertificateApiRequest {
	return CreateDatabaseUserCertificateApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		username:   username,
		userCert:   userCert,
	}
}

// Execute executes the request
func (a *X509AuthenticationApiService) createDatabaseUserCertificateExecute(r CreateDatabaseUserCertificateApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "X509AuthenticationApiService.CreateDatabaseUserCertificate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return nil, reportError("groupId must have less than 24 elements")
	}
	if r.userCert == nil {
		return nil, reportError("userCert is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

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
	// body params
	localVarPostBody = r.userCert
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
		var v ApiError
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

type DisableCustomerManagedX509ApiRequest struct {
	ctx        context.Context
	ApiService X509AuthenticationApi
	groupId    string
}

type DisableCustomerManagedX509ApiParams struct {
	GroupId string
}

func (a *X509AuthenticationApiService) DisableCustomerManagedX509WithParams(ctx context.Context, args *DisableCustomerManagedX509ApiParams) DisableCustomerManagedX509ApiRequest {
	return DisableCustomerManagedX509ApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DisableCustomerManagedX509ApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.disableCustomerManagedX509Execute(r)
}

/*
DisableCustomerManagedX509 Disable Customer-Managed X.509

Clears the customer-managed X.509 settings on a project, including the uploaded Certificate Authority, which disables self-managed X.509.

	Updating this configuration triggers a rolling restart of the database. You must have the Project Owner role to use this endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DisableCustomerManagedX509ApiRequest
*/
func (a *X509AuthenticationApiService) DisableCustomerManagedX509(ctx context.Context, groupId string) DisableCustomerManagedX509ApiRequest {
	return DisableCustomerManagedX509ApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return UserSecurity
func (a *X509AuthenticationApiService) disableCustomerManagedX509Execute(r DisableCustomerManagedX509ApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "X509AuthenticationApiService.DisableCustomerManagedX509")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/customerX509"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
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
		var v ApiError
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

type ListDatabaseUserCertificatesApiRequest struct {
	ctx          context.Context
	ApiService   X509AuthenticationApi
	groupId      string
	username     string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListDatabaseUserCertificatesApiParams struct {
	GroupId      string
	Username     string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *X509AuthenticationApiService) ListDatabaseUserCertificatesWithParams(ctx context.Context, args *ListDatabaseUserCertificatesApiParams) ListDatabaseUserCertificatesApiRequest {
	return ListDatabaseUserCertificatesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		username:     args.Username,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListDatabaseUserCertificatesApiRequest) IncludeCount(includeCount bool) ListDatabaseUserCertificatesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListDatabaseUserCertificatesApiRequest) ItemsPerPage(itemsPerPage int) ListDatabaseUserCertificatesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListDatabaseUserCertificatesApiRequest) PageNum(pageNum int) ListDatabaseUserCertificatesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListDatabaseUserCertificatesApiRequest) Execute() (*PaginatedUserCert, *http.Response, error) {
	return r.ApiService.listDatabaseUserCertificatesExecute(r)
}

/*
ListDatabaseUserCertificates Return All X.509 Certificates Assigned to One MongoDB User

Returns all unexpired X.509 certificates for the specified MongoDB user. This MongoDB user belongs to one project. Atlas manages these certificates and the MongoDB user. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param username Human-readable label that represents the MongoDB database user account whose certificates you want to return.
	@return ListDatabaseUserCertificatesApiRequest
*/
func (a *X509AuthenticationApiService) ListDatabaseUserCertificates(ctx context.Context, groupId string, username string) ListDatabaseUserCertificatesApiRequest {
	return ListDatabaseUserCertificatesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		username:   username,
	}
}

// Execute executes the request
//
//	@return PaginatedUserCert
func (a *X509AuthenticationApiService) listDatabaseUserCertificatesExecute(r ListDatabaseUserCertificatesApiRequest) (*PaginatedUserCert, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedUserCert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "X509AuthenticationApiService.ListDatabaseUserCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
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
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
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
		var v ApiError
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

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


type AuditingApi interface {

	/*
	GetAuditingConfiguration Return the Auditing Configuration for One Project

	Returns the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AuditingApiGetAuditingConfigurationRequest
	*/
	GetAuditingConfiguration(ctx context.Context, groupId string) AuditingApiGetAuditingConfigurationRequest

	// GetAuditingConfigurationExecute executes the request
	//  @return AuditLog
	GetAuditingConfigurationExecute(r AuditingApiGetAuditingConfigurationRequest) (*AuditLog, *http.Response, error)

	/*
	UpdateAuditingConfiguration Update Auditing Configuration for One Project

	Updates the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AuditingApiUpdateAuditingConfigurationRequest
	*/
	UpdateAuditingConfiguration(ctx context.Context, groupId string) AuditingApiUpdateAuditingConfigurationRequest

	// UpdateAuditingConfigurationExecute executes the request
	//  @return AuditLog
	UpdateAuditingConfigurationExecute(r AuditingApiUpdateAuditingConfigurationRequest) (*AuditLog, *http.Response, error)
}

// AuditingApiService AuditingApi service
type AuditingApiService service

type AuditingApiGetAuditingConfigurationRequest struct {
	ctx context.Context
	ApiService AuditingApi
	groupId string
}

func (r AuditingApiGetAuditingConfigurationRequest) Execute() (*AuditLog, *http.Response, error) {
	return r.ApiService.GetAuditingConfigurationExecute(r)
}

/*
GetAuditingConfiguration Return the Auditing Configuration for One Project

Returns the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return AuditingApiGetAuditingConfigurationRequest
*/
func (a *AuditingApiService) GetAuditingConfiguration(ctx context.Context, groupId string) AuditingApiGetAuditingConfigurationRequest {
	return AuditingApiGetAuditingConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return AuditLog
func (a *AuditingApiService) GetAuditingConfigurationExecute(r AuditingApiGetAuditingConfigurationRequest) (*AuditLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditingApiService.GetAuditingConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/auditLog"
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

type AuditingApiUpdateAuditingConfigurationRequest struct {
	ctx context.Context
	ApiService AuditingApi
	groupId string
	auditLog *AuditLog
}

// Updated auditing configuration for the specified project.
func (r AuditingApiUpdateAuditingConfigurationRequest) AuditLog(auditLog AuditLog) AuditingApiUpdateAuditingConfigurationRequest {
	r.auditLog = &auditLog
	return r
}

func (r AuditingApiUpdateAuditingConfigurationRequest) Execute() (*AuditLog, *http.Response, error) {
	return r.ApiService.UpdateAuditingConfigurationExecute(r)
}

/*
UpdateAuditingConfiguration Update Auditing Configuration for One Project

Updates the auditing configuration for the specified project. The auditing configuration defines the events that MongoDB Cloud records in the audit log. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List. This feature isn't available for `M0`, `M2`, `M5`, or serverless clusters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return AuditingApiUpdateAuditingConfigurationRequest
*/
func (a *AuditingApiService) UpdateAuditingConfiguration(ctx context.Context, groupId string) AuditingApiUpdateAuditingConfigurationRequest {
	return AuditingApiUpdateAuditingConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return AuditLog
func (a *AuditingApiService) UpdateAuditingConfigurationExecute(r AuditingApiUpdateAuditingConfigurationRequest) (*AuditLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditingApiService.UpdateAuditingConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/auditLog"
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
	if r.auditLog == nil {
		return localVarReturnValue, nil, reportError("auditLog is required and must be specified")
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
	localVarPostBody = r.auditLog
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

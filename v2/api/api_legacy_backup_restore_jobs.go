/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type LegacyBackupRestoreJobsApi interface {

	/*
	CreateOneLegacyBackupRestoreJob Create One Legacy Backup Restore Job

	Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Atlas Admin role and an entry for the project access list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@return LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest

	Deprecated
	*/
	CreateOneLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string) LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest

	// CreateOneLegacyBackupRestoreJobExecute executes the request
	//  @return PaginatedRestoreJobView
	// Deprecated
	CreateOneLegacyBackupRestoreJobExecute(r LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest) (*PaginatedRestoreJobView, *http.Response, error)
}

// LegacyBackupRestoreJobsApiService LegacyBackupRestoreJobsApi service
type LegacyBackupRestoreJobsApiService service

type LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest struct {
	ctx context.Context
	ApiService LegacyBackupRestoreJobsApi
	groupId string
	clusterName string
	apiRestoreJobView *ApiRestoreJobView
	envelope *bool
	pretty *bool
}

// Legacy backup to restore to one cluster in the specified project.
func (r LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest) ApiRestoreJobView(apiRestoreJobView ApiRestoreJobView) LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest {
	r.apiRestoreJobView = &apiRestoreJobView
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest) Envelope(envelope bool) LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest) Pretty(pretty bool) LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest {
	r.pretty = &pretty
	return r
}

func (r LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest) Execute() (*PaginatedRestoreJobView, *http.Response, error) {
	return r.ApiService.CreateOneLegacyBackupRestoreJobExecute(r)
}

/*
CreateOneLegacyBackupRestoreJob Create One Legacy Backup Restore Job

Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Atlas Admin role and an entry for the project access list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
 @return LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest

Deprecated
*/
func (a *LegacyBackupRestoreJobsApiService) CreateOneLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string) LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest {
	return LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//  @return PaginatedRestoreJobView
// Deprecated
func (a *LegacyBackupRestoreJobsApiService) CreateOneLegacyBackupRestoreJobExecute(r LegacyBackupRestoreJobsApiCreateOneLegacyBackupRestoreJobRequest) (*PaginatedRestoreJobView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRestoreJobView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupRestoreJobsApiService.CreateOneLegacyBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.apiRestoreJobView == nil {
		return localVarReturnValue, nil, reportError("apiRestoreJobView is required and must be specified")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiRestoreJobView
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
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

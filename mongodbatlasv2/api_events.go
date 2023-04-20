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
	"time"
	"reflect"
)


type EventsApi interface {

	/*
	GetOrganizationEvent Return One Event from One Organization

	Returns one event for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-orgs-get-one/).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access.
	@return EventsApiGetOrganizationEventRequest
	*/
	GetOrganizationEvent(ctx context.Context, orgId string, eventId string) EventsApiGetOrganizationEventRequest

	// GetOrganizationEventExecute executes the request
	//  @return EventViewForOrg
	GetOrganizationEventExecute(r EventsApiGetOrganizationEventRequest) (*EventViewForOrg, *http.Response, error)

	/*
	GetProjectEvent Return One Event from One Project

	Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-one/).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access.
	@return EventsApiGetProjectEventRequest
	*/
	GetProjectEvent(ctx context.Context, groupId string, eventId string) EventsApiGetProjectEventRequest

	// GetProjectEventExecute executes the request
	//  @return EventViewForNdsGroup
	GetProjectEventExecute(r EventsApiGetProjectEventRequest) (*EventViewForNdsGroup, *http.Response, error)

	/*
	ListOrganizationEvents Return All Events from One Organization

	Returns all events for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-orgs-get-all/).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return EventsApiListOrganizationEventsRequest
	*/
	ListOrganizationEvents(ctx context.Context, orgId string) EventsApiListOrganizationEventsRequest

	// ListOrganizationEventsExecute executes the request
	//  @return OrgPaginatedEvent
	ListOrganizationEventsExecute(r EventsApiListOrganizationEventsRequest) (*OrgPaginatedEvent, *http.Response, error)

	/*
	ListProjectEvents Return All Events from One Project

	Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-all/).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return EventsApiListProjectEventsRequest
	*/
	ListProjectEvents(ctx context.Context, groupId string) EventsApiListProjectEventsRequest

	// ListProjectEventsExecute executes the request
	//  @return GroupPaginatedEvent
	ListProjectEventsExecute(r EventsApiListProjectEventsRequest) (*GroupPaginatedEvent, *http.Response, error)
}

// EventsApiService EventsApi service
type EventsApiService service

type EventsApiGetOrganizationEventRequest struct {
	ctx context.Context
	ApiService EventsApi
	orgId string
	eventId string
	includeRaw *bool
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r EventsApiGetOrganizationEventRequest) IncludeRaw(includeRaw bool) EventsApiGetOrganizationEventRequest {
	r.includeRaw = &includeRaw
	return r
}

func (r EventsApiGetOrganizationEventRequest) Execute() (*EventViewForOrg, *http.Response, error) {
	return r.ApiService.GetOrganizationEventExecute(r)
}

/*
GetOrganizationEvent Return One Event from One Organization

Returns one event for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-orgs-get-one/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
 @param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access.
 @return EventsApiGetOrganizationEventRequest
*/
func (a *EventsApiService) GetOrganizationEvent(ctx context.Context, orgId string, eventId string) EventsApiGetOrganizationEventRequest {
	return EventsApiGetOrganizationEventRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		eventId: eventId,
	}
}

// Execute executes the request
//  @return EventViewForOrg
func (a *EventsApiService) GetOrganizationEventExecute(r EventsApiGetOrganizationEventRequest) (*EventViewForOrg, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EventViewForOrg
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetOrganizationEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/events/{eventId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.eventId) < 24 {
		return localVarReturnValue, nil, reportError("eventId must have at least 24 elements")
	}
	if strlen(r.eventId) > 24 {
		return localVarReturnValue, nil, reportError("eventId must have less than 24 elements")
	}

	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
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

type EventsApiGetProjectEventRequest struct {
	ctx context.Context
	ApiService EventsApi
	groupId string
	eventId string
	includeRaw *bool
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r EventsApiGetProjectEventRequest) IncludeRaw(includeRaw bool) EventsApiGetProjectEventRequest {
	r.includeRaw = &includeRaw
	return r
}

func (r EventsApiGetProjectEventRequest) Execute() (*EventViewForNdsGroup, *http.Response, error) {
	return r.ApiService.GetProjectEventExecute(r)
}

/*
GetProjectEvent Return One Event from One Project

Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-one/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access.
 @return EventsApiGetProjectEventRequest
*/
func (a *EventsApiService) GetProjectEvent(ctx context.Context, groupId string, eventId string) EventsApiGetProjectEventRequest {
	return EventsApiGetProjectEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		eventId: eventId,
	}
}

// Execute executes the request
//  @return EventViewForNdsGroup
func (a *EventsApiService) GetProjectEventExecute(r EventsApiGetProjectEventRequest) (*EventViewForNdsGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EventViewForNdsGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetProjectEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/events/{eventId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.eventId) < 24 {
		return localVarReturnValue, nil, reportError("eventId must have at least 24 elements")
	}
	if strlen(r.eventId) > 24 {
		return localVarReturnValue, nil, reportError("eventId must have less than 24 elements")
	}

	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
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

type EventsApiListOrganizationEventsRequest struct {
	ctx context.Context
	ApiService EventsApi
	orgId string
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
	eventType *EventTypeForOrg
	includeRaw *bool
	maxDate *time.Time
	minDate *time.Time
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r EventsApiListOrganizationEventsRequest) IncludeCount(includeCount bool) EventsApiListOrganizationEventsRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r EventsApiListOrganizationEventsRequest) ItemsPerPage(itemsPerPage int32) EventsApiListOrganizationEventsRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r EventsApiListOrganizationEventsRequest) PageNum(pageNum int32) EventsApiListOrganizationEventsRequest {
	r.pageNum = &pageNum
	return r
}

// Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently.
func (r EventsApiListOrganizationEventsRequest) EventType(eventType EventTypeForOrg) EventsApiListOrganizationEventsRequest {
	r.eventType = &eventType
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r EventsApiListOrganizationEventsRequest) IncludeRaw(includeRaw bool) EventsApiListOrganizationEventsRequest {
	r.includeRaw = &includeRaw
	return r
}

// Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.
func (r EventsApiListOrganizationEventsRequest) MaxDate(maxDate time.Time) EventsApiListOrganizationEventsRequest {
	r.maxDate = &maxDate
	return r
}

// Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.
func (r EventsApiListOrganizationEventsRequest) MinDate(minDate time.Time) EventsApiListOrganizationEventsRequest {
	r.minDate = &minDate
	return r
}

func (r EventsApiListOrganizationEventsRequest) Execute() (*OrgPaginatedEvent, *http.Response, error) {
	return r.ApiService.ListOrganizationEventsExecute(r)
}

/*
ListOrganizationEvents Return All Events from One Organization

Returns all events for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-orgs-get-all/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
 @return EventsApiListOrganizationEventsRequest
*/
func (a *EventsApiService) ListOrganizationEvents(ctx context.Context, orgId string) EventsApiListOrganizationEventsRequest {
	return EventsApiListOrganizationEventsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return OrgPaginatedEvent
func (a *EventsApiService) ListOrganizationEventsExecute(r EventsApiListOrganizationEventsRequest) (*OrgPaginatedEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrgPaginatedEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListOrganizationEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/events"
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
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	}
	if r.maxDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxDate", r.maxDate, "")
	}
	if r.minDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minDate", r.minDate, "")
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

type EventsApiListProjectEventsRequest struct {
	ctx context.Context
	ApiService EventsApi
	groupId string
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
	clusterNames *[]string
	eventType *EventTypeForNdsGroup
	includeRaw *bool
	maxDate *time.Time
	minDate *time.Time
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r EventsApiListProjectEventsRequest) IncludeCount(includeCount bool) EventsApiListProjectEventsRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r EventsApiListProjectEventsRequest) ItemsPerPage(itemsPerPage int32) EventsApiListProjectEventsRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r EventsApiListProjectEventsRequest) PageNum(pageNum int32) EventsApiListProjectEventsRequest {
	r.pageNum = &pageNum
	return r
}

// Human-readable label that identifies the cluster.
func (r EventsApiListProjectEventsRequest) ClusterNames(clusterNames []string) EventsApiListProjectEventsRequest {
	r.clusterNames = &clusterNames
	return r
}

// Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently.
func (r EventsApiListProjectEventsRequest) EventType(eventType EventTypeForNdsGroup) EventsApiListProjectEventsRequest {
	r.eventType = &eventType
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r EventsApiListProjectEventsRequest) IncludeRaw(includeRaw bool) EventsApiListProjectEventsRequest {
	r.includeRaw = &includeRaw
	return r
}

// Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.
func (r EventsApiListProjectEventsRequest) MaxDate(maxDate time.Time) EventsApiListProjectEventsRequest {
	r.maxDate = &maxDate
	return r
}

// Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.
func (r EventsApiListProjectEventsRequest) MinDate(minDate time.Time) EventsApiListProjectEventsRequest {
	r.minDate = &minDate
	return r
}

func (r EventsApiListProjectEventsRequest) Execute() (*GroupPaginatedEvent, *http.Response, error) {
	return r.ApiService.ListProjectEventsExecute(r)
}

/*
ListProjectEvents Return All Events from One Project

Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-all/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return EventsApiListProjectEventsRequest
*/
func (a *EventsApiService) ListProjectEvents(ctx context.Context, groupId string) EventsApiListProjectEventsRequest {
	return EventsApiListProjectEventsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return GroupPaginatedEvent
func (a *EventsApiService) ListProjectEventsExecute(r EventsApiListProjectEventsRequest) (*GroupPaginatedEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupPaginatedEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ListProjectEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/events"
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
	if r.clusterNames != nil {
		t := *r.clusterNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "clusterNames", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "clusterNames", t, "multi")
		}
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.includeRaw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	} else {
		var defaultValue bool = false
		r.includeRaw = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeRaw", r.includeRaw, "")
	}
	if r.maxDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxDate", r.maxDate, "")
	}
	if r.minDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minDate", r.minDate, "")
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

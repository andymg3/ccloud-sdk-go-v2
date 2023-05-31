// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type StatementsSqlV1alpha1Api interface {

	/*
	CreateSqlV1alpha1Statement Create a Statement

	[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Make a request to create a statement.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param environmentId The unique identifier for the environment.
	 @return ApiCreateSqlV1alpha1StatementRequest
	*/
	CreateSqlV1alpha1Statement(ctx _context.Context, environmentId string) ApiCreateSqlV1alpha1StatementRequest

	// CreateSqlV1alpha1StatementExecute executes the request
	//  @return SqlV1alpha1Statement
	CreateSqlV1alpha1StatementExecute(r ApiCreateSqlV1alpha1StatementRequest) (SqlV1alpha1Statement, *_nethttp.Response, error)

	/*
	DeleteSqlV1alpha1Statement Delete a Statement

	[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Make a request to delete a statement.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param environmentId The unique identifier for the environment.
	 @param statementName The unique identifier for the statement.
	 @return ApiDeleteSqlV1alpha1StatementRequest
	*/
	DeleteSqlV1alpha1Statement(ctx _context.Context, environmentId string, statementName string) ApiDeleteSqlV1alpha1StatementRequest

	// DeleteSqlV1alpha1StatementExecute executes the request
	DeleteSqlV1alpha1StatementExecute(r ApiDeleteSqlV1alpha1StatementRequest) (*_nethttp.Response, error)

	/*
	GetSqlV1alpha1Statement Read a Statement

	[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Make a request to read a statement.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param environmentId The unique identifier for the environment.
	 @param statementName The unique identifier for the statement.
	 @return ApiGetSqlV1alpha1StatementRequest
	*/
	GetSqlV1alpha1Statement(ctx _context.Context, environmentId string, statementName string) ApiGetSqlV1alpha1StatementRequest

	// GetSqlV1alpha1StatementExecute executes the request
	//  @return SqlV1alpha1Statement
	GetSqlV1alpha1StatementExecute(r ApiGetSqlV1alpha1StatementRequest) (SqlV1alpha1Statement, *_nethttp.Response, error)

	/*
	ListSqlV1alpha1Statements List of Statements

	[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Retrieve a sorted, filtered, paginated list of all statements.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param environmentId The unique identifier for the environment.
	 @return ApiListSqlV1alpha1StatementsRequest
	*/
	ListSqlV1alpha1Statements(ctx _context.Context, environmentId string) ApiListSqlV1alpha1StatementsRequest

	// ListSqlV1alpha1StatementsExecute executes the request
	//  @return SqlV1alpha1StatementList
	ListSqlV1alpha1StatementsExecute(r ApiListSqlV1alpha1StatementsRequest) (SqlV1alpha1StatementList, *_nethttp.Response, error)
}

// StatementsSqlV1alpha1ApiService StatementsSqlV1alpha1Api service
type StatementsSqlV1alpha1ApiService service

type ApiCreateSqlV1alpha1StatementRequest struct {
	ctx _context.Context
	ApiService StatementsSqlV1alpha1Api
	orgId *string
	environmentId string
	sqlV1alpha1Statement *SqlV1alpha1Statement
}

// The unique identifier for the organization.
func (r ApiCreateSqlV1alpha1StatementRequest) OrgId(orgId string) ApiCreateSqlV1alpha1StatementRequest {
	r.orgId = &orgId
	return r
}
func (r ApiCreateSqlV1alpha1StatementRequest) SqlV1alpha1Statement(sqlV1alpha1Statement SqlV1alpha1Statement) ApiCreateSqlV1alpha1StatementRequest {
	r.sqlV1alpha1Statement = &sqlV1alpha1Statement
	return r
}

func (r ApiCreateSqlV1alpha1StatementRequest) Execute() (SqlV1alpha1Statement, *_nethttp.Response, error) {
	return r.ApiService.CreateSqlV1alpha1StatementExecute(r)
}

/*
CreateSqlV1alpha1Statement Create a Statement

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Make a request to create a statement.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The unique identifier for the environment.
 @return ApiCreateSqlV1alpha1StatementRequest
*/
func (a *StatementsSqlV1alpha1ApiService) CreateSqlV1alpha1Statement(ctx _context.Context, environmentId string) ApiCreateSqlV1alpha1StatementRequest {
	return ApiCreateSqlV1alpha1StatementRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return SqlV1alpha1Statement
func (a *StatementsSqlV1alpha1ApiService) CreateSqlV1alpha1StatementExecute(r ApiCreateSqlV1alpha1StatementRequest) (SqlV1alpha1Statement, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1alpha1Statement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1alpha1ApiService.CreateSqlV1alpha1Statement")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1alpha1/environments/{environment_id}/statements"
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Org-Id"] = parameterToString(*r.orgId, "")
	// body params
	localVarPostBody = r.sqlV1alpha1Statement
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteSqlV1alpha1StatementRequest struct {
	ctx _context.Context
	ApiService StatementsSqlV1alpha1Api
	orgId *string
	environmentId string
	statementName string
}

// The unique identifier for the organization.
func (r ApiDeleteSqlV1alpha1StatementRequest) OrgId(orgId string) ApiDeleteSqlV1alpha1StatementRequest {
	r.orgId = &orgId
	return r
}

func (r ApiDeleteSqlV1alpha1StatementRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteSqlV1alpha1StatementExecute(r)
}

/*
DeleteSqlV1alpha1Statement Delete a Statement

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Make a request to delete a statement.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The unique identifier for the environment.
 @param statementName The unique identifier for the statement.
 @return ApiDeleteSqlV1alpha1StatementRequest
*/
func (a *StatementsSqlV1alpha1ApiService) DeleteSqlV1alpha1Statement(ctx _context.Context, environmentId string, statementName string) ApiDeleteSqlV1alpha1StatementRequest {
	return ApiDeleteSqlV1alpha1StatementRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		statementName: statementName,
	}
}

// Execute executes the request
func (a *StatementsSqlV1alpha1ApiService) DeleteSqlV1alpha1StatementExecute(r ApiDeleteSqlV1alpha1StatementRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1alpha1ApiService.DeleteSqlV1alpha1Statement")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1alpha1/environments/{environment_id}/statements/{statement_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statement_name"+"}", _neturl.PathEscape(parameterToString(r.statementName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return nil, reportError("orgId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Org-Id"] = parameterToString(*r.orgId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSqlV1alpha1StatementRequest struct {
	ctx _context.Context
	ApiService StatementsSqlV1alpha1Api
	orgId *string
	environmentId string
	statementName string
}

// The unique identifier for the organization.
func (r ApiGetSqlV1alpha1StatementRequest) OrgId(orgId string) ApiGetSqlV1alpha1StatementRequest {
	r.orgId = &orgId
	return r
}

func (r ApiGetSqlV1alpha1StatementRequest) Execute() (SqlV1alpha1Statement, *_nethttp.Response, error) {
	return r.ApiService.GetSqlV1alpha1StatementExecute(r)
}

/*
GetSqlV1alpha1Statement Read a Statement

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Make a request to read a statement.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The unique identifier for the environment.
 @param statementName The unique identifier for the statement.
 @return ApiGetSqlV1alpha1StatementRequest
*/
func (a *StatementsSqlV1alpha1ApiService) GetSqlV1alpha1Statement(ctx _context.Context, environmentId string, statementName string) ApiGetSqlV1alpha1StatementRequest {
	return ApiGetSqlV1alpha1StatementRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		statementName: statementName,
	}
}

// Execute executes the request
//  @return SqlV1alpha1Statement
func (a *StatementsSqlV1alpha1ApiService) GetSqlV1alpha1StatementExecute(r ApiGetSqlV1alpha1StatementRequest) (SqlV1alpha1Statement, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1alpha1Statement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1alpha1ApiService.GetSqlV1alpha1Statement")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1alpha1/environments/{environment_id}/statements/{statement_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statement_name"+"}", _neturl.PathEscape(parameterToString(r.statementName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Org-Id"] = parameterToString(*r.orgId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListSqlV1alpha1StatementsRequest struct {
	ctx _context.Context
	ApiService StatementsSqlV1alpha1Api
	orgId *string
	environmentId string
	computePoolId *string
	pageSize *int32
	pageToken *string
}

// The unique identifier for the organization.
func (r ApiListSqlV1alpha1StatementsRequest) OrgId(orgId string) ApiListSqlV1alpha1StatementsRequest {
	r.orgId = &orgId
	return r
}
// Filter the results by exact match for spec.compute_pool_id.
func (r ApiListSqlV1alpha1StatementsRequest) ComputePoolId(computePoolId string) ApiListSqlV1alpha1StatementsRequest {
	r.computePoolId = &computePoolId
	return r
}
// A pagination size for collection requests.
func (r ApiListSqlV1alpha1StatementsRequest) PageSize(pageSize int32) ApiListSqlV1alpha1StatementsRequest {
	r.pageSize = &pageSize
	return r
}
// An opaque pagination token for collection requests.
func (r ApiListSqlV1alpha1StatementsRequest) PageToken(pageToken string) ApiListSqlV1alpha1StatementsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListSqlV1alpha1StatementsRequest) Execute() (SqlV1alpha1StatementList, *_nethttp.Response, error) {
	return r.ApiService.ListSqlV1alpha1StatementsExecute(r)
}

/*
ListSqlV1alpha1Statements List of Statements

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1alpha1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1alpha1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1alpha1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1alpha1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Retrieve a sorted, filtered, paginated list of all statements.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The unique identifier for the environment.
 @return ApiListSqlV1alpha1StatementsRequest
*/
func (a *StatementsSqlV1alpha1ApiService) ListSqlV1alpha1Statements(ctx _context.Context, environmentId string) ApiListSqlV1alpha1StatementsRequest {
	return ApiListSqlV1alpha1StatementsRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return SqlV1alpha1StatementList
func (a *StatementsSqlV1alpha1ApiService) ListSqlV1alpha1StatementsExecute(r ApiListSqlV1alpha1StatementsRequest) (SqlV1alpha1StatementList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1alpha1StatementList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1alpha1ApiService.ListSqlV1alpha1Statements")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1alpha1/environments/{environment_id}/statements"
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	if r.computePoolId != nil {
		localVarQueryParams.Add("compute_pool_id", parameterToString(*r.computePoolId, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Org-Id"] = parameterToString(*r.orgId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

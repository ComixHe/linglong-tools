/*
linglong仓库

玲珑仓库接口

API version: 1.0.0
Contact: wurongjie@deepin.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiserver

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// ClientAPIService ClientAPI service
type ClientAPIService service

type ApiFuzzySearchAppRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	data *RequestFuzzySearchReq
}

// app json数据
func (r ApiFuzzySearchAppRequest) Data(data RequestFuzzySearchReq) ApiFuzzySearchAppRequest {
	r.data = &data
	return r
}

func (r ApiFuzzySearchAppRequest) Execute() (*FuzzySearchApp200Response, *http.Response, error) {
	return r.ApiService.FuzzySearchAppExecute(r)
}

/*
FuzzySearchApp 模糊查找App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFuzzySearchAppRequest
*/
func (a *ClientAPIService) FuzzySearchApp(ctx context.Context) ApiFuzzySearchAppRequest {
	return ApiFuzzySearchAppRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FuzzySearchApp200Response
func (a *ClientAPIService) FuzzySearchAppExecute(r ApiFuzzySearchAppRequest) (*FuzzySearchApp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FuzzySearchApp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.FuzzySearchApp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apps/fuzzysearchapp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.data == nil {
		return localVarReturnValue, nil, reportError("data is required and must be specified")
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
	// body params
	localVarPostBody = r.data
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

type ApiGetRepoRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	repo string
}

func (r ApiGetRepoRequest) Execute() (*GetRepo200Response, *http.Response, error) {
	return r.ApiService.GetRepoExecute(r)
}

/*
GetRepo 查看仓库信息

returns repository mode and resolve all branches

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param repo 仓库名称
 @return ApiGetRepoRequest
*/
func (a *ClientAPIService) GetRepo(ctx context.Context, repo string) ApiGetRepoRequest {
	return ApiGetRepoRequest{
		ApiService: a,
		ctx: ctx,
		repo: repo,
	}
}

// Execute executes the request
//  @return GetRepo200Response
func (a *ClientAPIService) GetRepoExecute(r ApiGetRepoRequest) (*GetRepo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetRepo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.GetRepo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repos/{repo}"
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

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

type ApiNewUploadTaskIDRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	xToken *string
	req *SchemaNewUploadTaskReq
}

// 31a165ba1be6dec616b1f8f3207b4273
func (r ApiNewUploadTaskIDRequest) XToken(xToken string) ApiNewUploadTaskIDRequest {
	r.xToken = &xToken
	return r
}

// JSON数据
func (r ApiNewUploadTaskIDRequest) Req(req SchemaNewUploadTaskReq) ApiNewUploadTaskIDRequest {
	r.req = &req
	return r
}

func (r ApiNewUploadTaskIDRequest) Execute() (*NewUploadTaskID200Response, *http.Response, error) {
	return r.ApiService.NewUploadTaskIDExecute(r)
}

/*
NewUploadTaskID generate a new upload task id

generate a new upload task id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNewUploadTaskIDRequest
*/
func (a *ClientAPIService) NewUploadTaskID(ctx context.Context) ApiNewUploadTaskIDRequest {
	return ApiNewUploadTaskIDRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NewUploadTaskID200Response
func (a *ClientAPIService) NewUploadTaskIDExecute(r ApiNewUploadTaskIDRequest) (*NewUploadTaskID200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NewUploadTaskID200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.NewUploadTaskID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/upload-tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xToken == nil {
		return localVarReturnValue, nil, reportError("xToken is required and must be specified")
	}
	if r.req == nil {
		return localVarReturnValue, nil, reportError("req is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Token", r.xToken, "")
	// body params
	localVarPostBody = r.req
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

type ApiSignInRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	data *RequestAuth
}

// auth json数据
func (r ApiSignInRequest) Data(data RequestAuth) ApiSignInRequest {
	r.data = &data
	return r
}

func (r ApiSignInRequest) Execute() (*SignIn200Response, *http.Response, error) {
	return r.ApiService.SignInExecute(r)
}

/*
SignIn 登陆帐号

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSignInRequest
*/
func (a *ClientAPIService) SignIn(ctx context.Context) ApiSignInRequest {
	return ApiSignInRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SignIn200Response
func (a *ClientAPIService) SignInExecute(r ApiSignInRequest) (*SignIn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SignIn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.SignIn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sign-in"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.data == nil {
		return localVarReturnValue, nil, reportError("data is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.data
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

type ApiUploadTaskFileRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	xToken *string
	taskId string
	file *os.File
}

// 31a165ba1be6dec616b1f8f3207b4273
func (r ApiUploadTaskFileRequest) XToken(xToken string) ApiUploadTaskFileRequest {
	r.xToken = &xToken
	return r
}

// 文件路径
func (r ApiUploadTaskFileRequest) File(file *os.File) ApiUploadTaskFileRequest {
	r.file = file
	return r
}

func (r ApiUploadTaskFileRequest) Execute() (*UploadTaskLayerFile200Response, *http.Response, error) {
	return r.ApiService.UploadTaskFileExecute(r)
}

/*
UploadTaskFile upload tgz file to upload task

upload tgz file to upload task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId task id
 @return ApiUploadTaskFileRequest
*/
func (a *ClientAPIService) UploadTaskFile(ctx context.Context, taskId string) ApiUploadTaskFileRequest {
	return ApiUploadTaskFileRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return UploadTaskLayerFile200Response
func (a *ClientAPIService) UploadTaskFileExecute(r ApiUploadTaskFileRequest) (*UploadTaskLayerFile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadTaskLayerFile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.UploadTaskFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/upload-tasks/{task_id}/tar"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xToken == nil {
		return localVarReturnValue, nil, reportError("xToken is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Token", r.xToken, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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

type ApiUploadTaskInfoRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	xToken *string
	taskId string
}

// 31a165ba1be6dec616b1f8f3207b4273
func (r ApiUploadTaskInfoRequest) XToken(xToken string) ApiUploadTaskInfoRequest {
	r.xToken = &xToken
	return r
}

func (r ApiUploadTaskInfoRequest) Execute() (*UploadTaskInfo200Response, *http.Response, error) {
	return r.ApiService.UploadTaskInfoExecute(r)
}

/*
UploadTaskInfo get upload task status

get upload task status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId task id
 @return ApiUploadTaskInfoRequest
*/
func (a *ClientAPIService) UploadTaskInfo(ctx context.Context, taskId string) ApiUploadTaskInfoRequest {
	return ApiUploadTaskInfoRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return UploadTaskInfo200Response
func (a *ClientAPIService) UploadTaskInfoExecute(r ApiUploadTaskInfoRequest) (*UploadTaskInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadTaskInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.UploadTaskInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/upload-tasks/{task_id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xToken == nil {
		return localVarReturnValue, nil, reportError("xToken is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Token", r.xToken, "")
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

type ApiUploadTaskLayerFileRequest struct {
	ctx context.Context
	ApiService *ClientAPIService
	xToken *string
	taskId string
	file *os.File
}

// 31a165ba1be6dec616b1f8f3207b4273
func (r ApiUploadTaskLayerFileRequest) XToken(xToken string) ApiUploadTaskLayerFileRequest {
	r.xToken = &xToken
	return r
}

// 文件路径
func (r ApiUploadTaskLayerFileRequest) File(file *os.File) ApiUploadTaskLayerFileRequest {
	r.file = file
	return r
}

func (r ApiUploadTaskLayerFileRequest) Execute() (*UploadTaskLayerFile200Response, *http.Response, error) {
	return r.ApiService.UploadTaskLayerFileExecute(r)
}

/*
UploadTaskLayerFile upload layer file to upload task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId task id
 @return ApiUploadTaskLayerFileRequest
*/
func (a *ClientAPIService) UploadTaskLayerFile(ctx context.Context, taskId string) ApiUploadTaskLayerFileRequest {
	return ApiUploadTaskLayerFileRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return UploadTaskLayerFile200Response
func (a *ClientAPIService) UploadTaskLayerFileExecute(r ApiUploadTaskLayerFileRequest) (*UploadTaskLayerFile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadTaskLayerFile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAPIService.UploadTaskLayerFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/upload-tasks/{task_id}/layer"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xToken == nil {
		return localVarReturnValue, nil, reportError("xToken is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Token", r.xToken, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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

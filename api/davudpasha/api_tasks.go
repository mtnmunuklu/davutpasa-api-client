package davudpasha

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/mtnmunuklu/davudpasha-api-client-go/api/common"
)

// @title Davudpasha API
// @version 1.0.0
// @description Davudpasha API to demonstrate OpenAPI documentation for client-go
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @basePath /api
// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization

// TasksApi service type.
type TasksApi common.Service

// SearchTasksOptionalParameters holds optional parameters for SearchTasks.
type SearchTasksOptionalParameters struct {
	Body *TasksSearchRequest
}

// NewSearchTasksOptionalParameters creates an empty struct for parameters.
func NewSearchTasksOptionalParameters() *SearchTasksOptionalParameters {
	this := SearchTasksOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchTasksOptionalParameters) WithBody(body TasksSearchRequest) *SearchTasksOptionalParameters {
	r.Body = &body
	return r
}

// SearchTasks searches for tasks.
//
// @Summary Search Tasks
// @Description Retrieves tasks that match a search query.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param body body TasksSearchRequest true "Tasks Search Request"
// @Success 200 {array} TasksSearchResponse "Successful operation"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 429 {object} ErrorResponse "Too Many Requests"
// @Router /ITaskManagerAct/GetList [post]
// @Security ApiKeyAuth
func (a *TasksApi) SearchTasks(ctx _context.Context, o ...SearchTasksOptionalParameters) ([]TasksSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod    = _nethttp.MethodPost
		localVarPostBody      interface{}
		localVarReturnValue   []TasksSearchResponse
		optionalParams        SearchTasksOptionalParameters
		localVarInterfaceCode = "ITaskManagerAct"
		localVarMethodName    = "GetList"
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type SearchTasksOptionalParameters is allowed")
	}

	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "TasksApi.SearchTasks")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/" + localVarInterfaceCode + "/" + localVarMethodName

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	if optionalParams.Body != nil {
		localVarPostBody = &optionalParams.Body
	}

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "x-api-key"},
	)

	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewTasksApi returns TasksApi.
func NewTasksApi(client *common.APIClient) *TasksApi {
	return &TasksApi{
		Client: client,
	}
}

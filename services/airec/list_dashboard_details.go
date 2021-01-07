package airec

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListDashboardDetails invokes the airec.ListDashboardDetails API synchronously
func (client *Client) ListDashboardDetails(request *ListDashboardDetailsRequest) (response *ListDashboardDetailsResponse, err error) {
	response = CreateListDashboardDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDashboardDetailsWithChan invokes the airec.ListDashboardDetails API asynchronously
func (client *Client) ListDashboardDetailsWithChan(request *ListDashboardDetailsRequest) (<-chan *ListDashboardDetailsResponse, <-chan error) {
	responseChan := make(chan *ListDashboardDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDashboardDetails(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListDashboardDetailsWithCallback invokes the airec.ListDashboardDetails API asynchronously
func (client *Client) ListDashboardDetailsWithCallback(request *ListDashboardDetailsRequest, callback func(response *ListDashboardDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDashboardDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListDashboardDetails(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListDashboardDetailsRequest is the request struct for api ListDashboardDetails
type ListDashboardDetailsRequest struct {
	*requests.RoaRequest
	MetricType    string           `position:"Query" name:"metricType"`
	InstanceId    string           `position:"Path" name:"instanceId"`
	ExperimentIds string           `position:"Query" name:"experimentIds"`
	TraceIds      string           `position:"Query" name:"traceIds"`
	EndTime       requests.Integer `position:"Query" name:"endTime"`
	StartTime     requests.Integer `position:"Query" name:"startTime"`
	SceneIds      string           `position:"Query" name:"sceneIds"`
}

// ListDashboardDetailsResponse is the response struct for api ListDashboardDetails
type ListDashboardDetailsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"requestId" xml:"requestId"`
	Code      string       `json:"code" xml:"code"`
	Message   string       `json:"message" xml:"message"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListDashboardDetailsRequest creates a request to invoke ListDashboardDetails API
func CreateListDashboardDetailsRequest() (request *ListDashboardDetailsRequest) {
	request = &ListDashboardDetailsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ListDashboardDetails", "/v2/openapi/instances/[instanceId]/dashboard/details", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDashboardDetailsResponse creates a response to parse from ListDashboardDetails response
func CreateListDashboardDetailsResponse() (response *ListDashboardDetailsResponse) {
	response = &ListDashboardDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package arms

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

// UpdatePrometheusGlobalView invokes the arms.UpdatePrometheusGlobalView API synchronously
func (client *Client) UpdatePrometheusGlobalView(request *UpdatePrometheusGlobalViewRequest) (response *UpdatePrometheusGlobalViewResponse, err error) {
	response = CreateUpdatePrometheusGlobalViewResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrometheusGlobalViewWithChan invokes the arms.UpdatePrometheusGlobalView API asynchronously
func (client *Client) UpdatePrometheusGlobalViewWithChan(request *UpdatePrometheusGlobalViewRequest) (<-chan *UpdatePrometheusGlobalViewResponse, <-chan error) {
	responseChan := make(chan *UpdatePrometheusGlobalViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrometheusGlobalView(request)
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

// UpdatePrometheusGlobalViewWithCallback invokes the arms.UpdatePrometheusGlobalView API asynchronously
func (client *Client) UpdatePrometheusGlobalViewWithCallback(request *UpdatePrometheusGlobalViewRequest, callback func(response *UpdatePrometheusGlobalViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrometheusGlobalViewResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrometheusGlobalView(request)
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

// UpdatePrometheusGlobalViewRequest is the request struct for api UpdatePrometheusGlobalView
type UpdatePrometheusGlobalViewRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	SubClustersJson string `position:"Query" name:"SubClustersJson"`
}

// UpdatePrometheusGlobalViewResponse is the response struct for api UpdatePrometheusGlobalView
type UpdatePrometheusGlobalViewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdatePrometheusGlobalViewRequest creates a request to invoke UpdatePrometheusGlobalView API
func CreateUpdatePrometheusGlobalViewRequest() (request *UpdatePrometheusGlobalViewRequest) {
	request = &UpdatePrometheusGlobalViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdatePrometheusGlobalView", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePrometheusGlobalViewResponse creates a response to parse from UpdatePrometheusGlobalView response
func CreateUpdatePrometheusGlobalViewResponse() (response *UpdatePrometheusGlobalViewResponse) {
	response = &UpdatePrometheusGlobalViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// DeletePrometheusAlertRule invokes the arms.DeletePrometheusAlertRule API synchronously
func (client *Client) DeletePrometheusAlertRule(request *DeletePrometheusAlertRuleRequest) (response *DeletePrometheusAlertRuleResponse, err error) {
	response = CreateDeletePrometheusAlertRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePrometheusAlertRuleWithChan invokes the arms.DeletePrometheusAlertRule API asynchronously
func (client *Client) DeletePrometheusAlertRuleWithChan(request *DeletePrometheusAlertRuleRequest) (<-chan *DeletePrometheusAlertRuleResponse, <-chan error) {
	responseChan := make(chan *DeletePrometheusAlertRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePrometheusAlertRule(request)
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

// DeletePrometheusAlertRuleWithCallback invokes the arms.DeletePrometheusAlertRule API asynchronously
func (client *Client) DeletePrometheusAlertRuleWithCallback(request *DeletePrometheusAlertRuleRequest, callback func(response *DeletePrometheusAlertRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePrometheusAlertRuleResponse
		var err error
		defer close(result)
		response, err = client.DeletePrometheusAlertRule(request)
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

// DeletePrometheusAlertRuleRequest is the request struct for api DeletePrometheusAlertRule
type DeletePrometheusAlertRuleRequest struct {
	*requests.RpcRequest
	AlertId requests.Integer `position:"Query" name:"AlertId"`
}

// DeletePrometheusAlertRuleResponse is the response struct for api DeletePrometheusAlertRule
type DeletePrometheusAlertRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeletePrometheusAlertRuleRequest creates a request to invoke DeletePrometheusAlertRule API
func CreateDeletePrometheusAlertRuleRequest() (request *DeletePrometheusAlertRuleRequest) {
	request = &DeletePrometheusAlertRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeletePrometheusAlertRule", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePrometheusAlertRuleResponse creates a response to parse from DeletePrometheusAlertRule response
func CreateDeletePrometheusAlertRuleResponse() (response *DeletePrometheusAlertRuleResponse) {
	response = &DeletePrometheusAlertRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

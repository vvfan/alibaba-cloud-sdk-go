package config

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

// RevertAggregateEvaluationResults invokes the config.RevertAggregateEvaluationResults API synchronously
func (client *Client) RevertAggregateEvaluationResults(request *RevertAggregateEvaluationResultsRequest) (response *RevertAggregateEvaluationResultsResponse, err error) {
	response = CreateRevertAggregateEvaluationResultsResponse()
	err = client.DoAction(request, response)
	return
}

// RevertAggregateEvaluationResultsWithChan invokes the config.RevertAggregateEvaluationResults API asynchronously
func (client *Client) RevertAggregateEvaluationResultsWithChan(request *RevertAggregateEvaluationResultsRequest) (<-chan *RevertAggregateEvaluationResultsResponse, <-chan error) {
	responseChan := make(chan *RevertAggregateEvaluationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevertAggregateEvaluationResults(request)
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

// RevertAggregateEvaluationResultsWithCallback invokes the config.RevertAggregateEvaluationResults API asynchronously
func (client *Client) RevertAggregateEvaluationResultsWithCallback(request *RevertAggregateEvaluationResultsRequest, callback func(response *RevertAggregateEvaluationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevertAggregateEvaluationResultsResponse
		var err error
		defer close(result)
		response, err = client.RevertAggregateEvaluationResults(request)
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

// RevertAggregateEvaluationResultsRequest is the request struct for api RevertAggregateEvaluationResults
type RevertAggregateEvaluationResultsRequest struct {
	*requests.RpcRequest
	ConfigRuleId string                                       `position:"Body" name:"ConfigRuleId"`
	Resources    *[]RevertAggregateEvaluationResultsResources `position:"Body" name:"Resources"  type:"Json"`
	AggregatorId string                                       `position:"Body" name:"AggregatorId"`
}

// RevertAggregateEvaluationResultsResources is a repeated param struct in RevertAggregateEvaluationResultsRequest
type RevertAggregateEvaluationResultsResources struct {
	ResourceId        string `name:"ResourceId"`
	ResourceAccountId string `name:"ResourceAccountId"`
	Region            string `name:"Region"`
	ResourceType      string `name:"ResourceType"`
}

// RevertAggregateEvaluationResultsResponse is the response struct for api RevertAggregateEvaluationResults
type RevertAggregateEvaluationResultsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevertAggregateEvaluationResultsRequest creates a request to invoke RevertAggregateEvaluationResults API
func CreateRevertAggregateEvaluationResultsRequest() (request *RevertAggregateEvaluationResultsRequest) {
	request = &RevertAggregateEvaluationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "RevertAggregateEvaluationResults", "", "")
	request.Method = requests.POST
	return
}

// CreateRevertAggregateEvaluationResultsResponse creates a response to parse from RevertAggregateEvaluationResults response
func CreateRevertAggregateEvaluationResultsResponse() (response *RevertAggregateEvaluationResultsResponse) {
	response = &RevertAggregateEvaluationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
package rds

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

// DeleteDBNodes invokes the rds.DeleteDBNodes API synchronously
func (client *Client) DeleteDBNodes(request *DeleteDBNodesRequest) (response *DeleteDBNodesResponse, err error) {
	response = CreateDeleteDBNodesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDBNodesWithChan invokes the rds.DeleteDBNodes API asynchronously
func (client *Client) DeleteDBNodesWithChan(request *DeleteDBNodesRequest) (<-chan *DeleteDBNodesResponse, <-chan error) {
	responseChan := make(chan *DeleteDBNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDBNodes(request)
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

// DeleteDBNodesWithCallback invokes the rds.DeleteDBNodes API asynchronously
func (client *Client) DeleteDBNodesWithCallback(request *DeleteDBNodesRequest, callback func(response *DeleteDBNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDBNodesResponse
		var err error
		defer close(result)
		response, err = client.DeleteDBNodes(request)
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

// DeleteDBNodesRequest is the request struct for api DeleteDBNodes
type DeleteDBNodesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBNodeId             string           `position:"Query" name:"DBNodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DeleteDBNodesResponse is the response struct for api DeleteDBNodes
type DeleteDBNodesResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId      int64  `json:"OrderId" xml:"OrderId"`
}

// CreateDeleteDBNodesRequest creates a request to invoke DeleteDBNodes API
func CreateDeleteDBNodesRequest() (request *DeleteDBNodesRequest) {
	request = &DeleteDBNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteDBNodes", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDBNodesResponse creates a response to parse from DeleteDBNodes response
func CreateDeleteDBNodesResponse() (response *DeleteDBNodesResponse) {
	response = &DeleteDBNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

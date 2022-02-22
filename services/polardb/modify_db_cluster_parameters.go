package polardb

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

// ModifyDBClusterParameters invokes the polardb.ModifyDBClusterParameters API synchronously
func (client *Client) ModifyDBClusterParameters(request *ModifyDBClusterParametersRequest) (response *ModifyDBClusterParametersResponse, err error) {
	response = CreateModifyDBClusterParametersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterParametersWithChan invokes the polardb.ModifyDBClusterParameters API asynchronously
func (client *Client) ModifyDBClusterParametersWithChan(request *ModifyDBClusterParametersRequest) (<-chan *ModifyDBClusterParametersResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterParameters(request)
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

// ModifyDBClusterParametersWithCallback invokes the polardb.ModifyDBClusterParameters API asynchronously
func (client *Client) ModifyDBClusterParametersWithCallback(request *ModifyDBClusterParametersRequest, callback func(response *ModifyDBClusterParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterParametersResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterParameters(request)
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

// ModifyDBClusterParametersRequest is the request struct for api ModifyDBClusterParameters
type ModifyDBClusterParametersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ParameterGroupId     string           `position:"Query" name:"ParameterGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Parameters           string           `position:"Query" name:"Parameters"`
}

// ModifyDBClusterParametersResponse is the response struct for api ModifyDBClusterParameters
type ModifyDBClusterParametersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterParametersRequest creates a request to invoke ModifyDBClusterParameters API
func CreateModifyDBClusterParametersRequest() (request *ModifyDBClusterParametersRequest) {
	request = &ModifyDBClusterParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterParameters", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterParametersResponse creates a response to parse from ModifyDBClusterParameters response
func CreateModifyDBClusterParametersResponse() (response *ModifyDBClusterParametersResponse) {
	response = &ModifyDBClusterParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

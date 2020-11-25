package hitsdb

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

// DeleteHiTSDBInstance invokes the hitsdb.DeleteHiTSDBInstance API synchronously
func (client *Client) DeleteHiTSDBInstance(request *DeleteHiTSDBInstanceRequest) (response *DeleteHiTSDBInstanceResponse, err error) {
	response = CreateDeleteHiTSDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHiTSDBInstanceWithChan invokes the hitsdb.DeleteHiTSDBInstance API asynchronously
func (client *Client) DeleteHiTSDBInstanceWithChan(request *DeleteHiTSDBInstanceRequest) (<-chan *DeleteHiTSDBInstanceResponse, <-chan error) {
	responseChan := make(chan *DeleteHiTSDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHiTSDBInstance(request)
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

// DeleteHiTSDBInstanceWithCallback invokes the hitsdb.DeleteHiTSDBInstance API asynchronously
func (client *Client) DeleteHiTSDBInstanceWithCallback(request *DeleteHiTSDBInstanceRequest, callback func(response *DeleteHiTSDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHiTSDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.DeleteHiTSDBInstance(request)
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

// DeleteHiTSDBInstanceRequest is the request struct for api DeleteHiTSDBInstance
type DeleteHiTSDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DeleteHiTSDBInstanceResponse is the response struct for api DeleteHiTSDBInstance
type DeleteHiTSDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHiTSDBInstanceRequest creates a request to invoke DeleteHiTSDBInstance API
func CreateDeleteHiTSDBInstanceRequest() (request *DeleteHiTSDBInstanceRequest) {
	request = &DeleteHiTSDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2017-06-01", "DeleteHiTSDBInstance", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteHiTSDBInstanceResponse creates a response to parse from DeleteHiTSDBInstance response
func CreateDeleteHiTSDBInstanceResponse() (response *DeleteHiTSDBInstanceResponse) {
	response = &DeleteHiTSDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
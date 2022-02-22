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

// RemoveDBClusterFromGDN invokes the polardb.RemoveDBClusterFromGDN API synchronously
func (client *Client) RemoveDBClusterFromGDN(request *RemoveDBClusterFromGDNRequest) (response *RemoveDBClusterFromGDNResponse, err error) {
	response = CreateRemoveDBClusterFromGDNResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDBClusterFromGDNWithChan invokes the polardb.RemoveDBClusterFromGDN API asynchronously
func (client *Client) RemoveDBClusterFromGDNWithChan(request *RemoveDBClusterFromGDNRequest) (<-chan *RemoveDBClusterFromGDNResponse, <-chan error) {
	responseChan := make(chan *RemoveDBClusterFromGDNResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDBClusterFromGDN(request)
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

// RemoveDBClusterFromGDNWithCallback invokes the polardb.RemoveDBClusterFromGDN API asynchronously
func (client *Client) RemoveDBClusterFromGDNWithCallback(request *RemoveDBClusterFromGDNRequest, callback func(response *RemoveDBClusterFromGDNResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDBClusterFromGDNResponse
		var err error
		defer close(result)
		response, err = client.RemoveDBClusterFromGDN(request)
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

// RemoveDBClusterFromGDNRequest is the request struct for api RemoveDBClusterFromGDN
type RemoveDBClusterFromGDNRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	GDNId                string           `position:"Query" name:"GDNId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RemoveDBClusterFromGDNResponse is the response struct for api RemoveDBClusterFromGDN
type RemoveDBClusterFromGDNResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveDBClusterFromGDNRequest creates a request to invoke RemoveDBClusterFromGDN API
func CreateRemoveDBClusterFromGDNRequest() (request *RemoveDBClusterFromGDNRequest) {
	request = &RemoveDBClusterFromGDNRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "RemoveDBClusterFromGDN", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveDBClusterFromGDNResponse creates a response to parse from RemoveDBClusterFromGDN response
func CreateRemoveDBClusterFromGDNResponse() (response *RemoveDBClusterFromGDNResponse) {
	response = &RemoveDBClusterFromGDNResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

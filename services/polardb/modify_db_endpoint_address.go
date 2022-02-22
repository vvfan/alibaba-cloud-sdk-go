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

// ModifyDBEndpointAddress invokes the polardb.ModifyDBEndpointAddress API synchronously
func (client *Client) ModifyDBEndpointAddress(request *ModifyDBEndpointAddressRequest) (response *ModifyDBEndpointAddressResponse, err error) {
	response = CreateModifyDBEndpointAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBEndpointAddressWithChan invokes the polardb.ModifyDBEndpointAddress API asynchronously
func (client *Client) ModifyDBEndpointAddressWithChan(request *ModifyDBEndpointAddressRequest) (<-chan *ModifyDBEndpointAddressResponse, <-chan error) {
	responseChan := make(chan *ModifyDBEndpointAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBEndpointAddress(request)
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

// ModifyDBEndpointAddressWithCallback invokes the polardb.ModifyDBEndpointAddress API asynchronously
func (client *Client) ModifyDBEndpointAddressWithCallback(request *ModifyDBEndpointAddressRequest, callback func(response *ModifyDBEndpointAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBEndpointAddressResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBEndpointAddress(request)
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

// ModifyDBEndpointAddressRequest is the request struct for api ModifyDBEndpointAddress
type ModifyDBEndpointAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix   string           `position:"Query" name:"ConnectionStringPrefix"`
	DBEndpointId             string           `position:"Query" name:"DBEndpointId"`
	PrivateZoneName          string           `position:"Query" name:"PrivateZoneName"`
	PrivateZoneAddressPrefix string           `position:"Query" name:"PrivateZoneAddressPrefix"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId              string           `position:"Query" name:"DBClusterId"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	Port                     string           `position:"Query" name:"Port"`
	NetType                  string           `position:"Query" name:"NetType"`
}

// ModifyDBEndpointAddressResponse is the response struct for api ModifyDBEndpointAddress
type ModifyDBEndpointAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBEndpointAddressRequest creates a request to invoke ModifyDBEndpointAddress API
func CreateModifyDBEndpointAddressRequest() (request *ModifyDBEndpointAddressRequest) {
	request = &ModifyDBEndpointAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBEndpointAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBEndpointAddressResponse creates a response to parse from ModifyDBEndpointAddress response
func CreateModifyDBEndpointAddressResponse() (response *ModifyDBEndpointAddressResponse) {
	response = &ModifyDBEndpointAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

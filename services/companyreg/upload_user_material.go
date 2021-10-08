package companyreg

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

// UploadUserMaterial invokes the companyreg.UploadUserMaterial API synchronously
func (client *Client) UploadUserMaterial(request *UploadUserMaterialRequest) (response *UploadUserMaterialResponse, err error) {
	response = CreateUploadUserMaterialResponse()
	err = client.DoAction(request, response)
	return
}

// UploadUserMaterialWithChan invokes the companyreg.UploadUserMaterial API asynchronously
func (client *Client) UploadUserMaterialWithChan(request *UploadUserMaterialRequest) (<-chan *UploadUserMaterialResponse, <-chan error) {
	responseChan := make(chan *UploadUserMaterialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadUserMaterial(request)
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

// UploadUserMaterialWithCallback invokes the companyreg.UploadUserMaterial API asynchronously
func (client *Client) UploadUserMaterialWithCallback(request *UploadUserMaterialRequest, callback func(response *UploadUserMaterialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadUserMaterialResponse
		var err error
		defer close(result)
		response, err = client.UploadUserMaterial(request)
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

// UploadUserMaterialRequest is the request struct for api UploadUserMaterial
type UploadUserMaterialRequest struct {
	*requests.RpcRequest
	BizId     string                         `position:"Query" name:"BizId"`
	Attribute *[]UploadUserMaterialAttribute `position:"Query" name:"Attribute"  type:"Repeated"`
	Status    requests.Integer               `position:"Query" name:"Status"`
}

// UploadUserMaterialAttribute is a repeated param struct in UploadUserMaterialRequest
type UploadUserMaterialAttribute struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// UploadUserMaterialResponse is the response struct for api UploadUserMaterial
type UploadUserMaterialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUploadUserMaterialRequest creates a request to invoke UploadUserMaterial API
func CreateUploadUserMaterialRequest() (request *UploadUserMaterialRequest) {
	request = &UploadUserMaterialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "UploadUserMaterial", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUploadUserMaterialResponse creates a response to parse from UploadUserMaterial response
func CreateUploadUserMaterialResponse() (response *UploadUserMaterialResponse) {
	response = &UploadUserMaterialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
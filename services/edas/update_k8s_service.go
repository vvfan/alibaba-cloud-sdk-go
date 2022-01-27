package edas

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

// UpdateK8sService invokes the edas.UpdateK8sService API synchronously
func (client *Client) UpdateK8sService(request *UpdateK8sServiceRequest) (response *UpdateK8sServiceResponse, err error) {
	response = CreateUpdateK8sServiceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateK8sServiceWithChan invokes the edas.UpdateK8sService API asynchronously
func (client *Client) UpdateK8sServiceWithChan(request *UpdateK8sServiceRequest) (<-chan *UpdateK8sServiceResponse, <-chan error) {
	responseChan := make(chan *UpdateK8sServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateK8sService(request)
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

// UpdateK8sServiceWithCallback invokes the edas.UpdateK8sService API asynchronously
func (client *Client) UpdateK8sServiceWithCallback(request *UpdateK8sServiceRequest, callback func(response *UpdateK8sServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateK8sServiceResponse
		var err error
		defer close(result)
		response, err = client.UpdateK8sService(request)
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

// UpdateK8sServiceRequest is the request struct for api UpdateK8sService
type UpdateK8sServiceRequest struct {
	*requests.RoaRequest
	AppId        string `position:"Query" name:"AppId"`
	Name         string `position:"Query" name:"Name"`
	Type         string `position:"Query" name:"Type"`
	ServicePorts string `position:"Query" name:"ServicePorts"`
}

// UpdateK8sServiceResponse is the response struct for api UpdateK8sService
type UpdateK8sServiceResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateK8sServiceRequest creates a request to invoke UpdateK8sService API
func CreateUpdateK8sServiceRequest() (request *UpdateK8sServiceRequest) {
	request = &UpdateK8sServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateK8sService", "/pop/v5/k8s/acs/k8s_service", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateK8sServiceResponse creates a response to parse from UpdateK8sService response
func CreateUpdateK8sServiceResponse() (response *UpdateK8sServiceResponse) {
	response = &UpdateK8sServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// CreateOrUpdateIMRobot invokes the arms.CreateOrUpdateIMRobot API synchronously
func (client *Client) CreateOrUpdateIMRobot(request *CreateOrUpdateIMRobotRequest) (response *CreateOrUpdateIMRobotResponse, err error) {
	response = CreateCreateOrUpdateIMRobotResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateIMRobotWithChan invokes the arms.CreateOrUpdateIMRobot API asynchronously
func (client *Client) CreateOrUpdateIMRobotWithChan(request *CreateOrUpdateIMRobotRequest) (<-chan *CreateOrUpdateIMRobotResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateIMRobotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateIMRobot(request)
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

// CreateOrUpdateIMRobotWithCallback invokes the arms.CreateOrUpdateIMRobot API asynchronously
func (client *Client) CreateOrUpdateIMRobotWithCallback(request *CreateOrUpdateIMRobotRequest, callback func(response *CreateOrUpdateIMRobotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateIMRobotResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateIMRobot(request)
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

// CreateOrUpdateIMRobotRequest is the request struct for api CreateOrUpdateIMRobot
type CreateOrUpdateIMRobotRequest struct {
	*requests.RpcRequest
	DailyNoc     requests.Boolean `position:"Body" name:"DailyNoc"`
	RobotAddress string           `position:"Body" name:"RobotAddress"`
	RobotName    string           `position:"Body" name:"RobotName"`
	RobotId      requests.Integer `position:"Body" name:"RobotId"`
	Type         string           `position:"Body" name:"Type"`
	DailyNocTime string           `position:"Body" name:"DailyNocTime"`
}

// CreateOrUpdateIMRobotResponse is the response struct for api CreateOrUpdateIMRobot
type CreateOrUpdateIMRobotResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	AlertRobot AlertRobot `json:"AlertRobot" xml:"AlertRobot"`
}

// CreateCreateOrUpdateIMRobotRequest creates a request to invoke CreateOrUpdateIMRobot API
func CreateCreateOrUpdateIMRobotRequest() (request *CreateOrUpdateIMRobotRequest) {
	request = &CreateOrUpdateIMRobotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateOrUpdateIMRobot", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateOrUpdateIMRobotResponse creates a response to parse from CreateOrUpdateIMRobot response
func CreateCreateOrUpdateIMRobotResponse() (response *CreateOrUpdateIMRobotResponse) {
	response = &CreateOrUpdateIMRobotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

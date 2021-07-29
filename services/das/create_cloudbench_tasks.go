package das

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

// CreateCloudbenchTasks invokes the das.CreateCloudbenchTasks API synchronously
func (client *Client) CreateCloudbenchTasks(request *CreateCloudbenchTasksRequest) (response *CreateCloudbenchTasksResponse, err error) {
	response = CreateCreateCloudbenchTasksResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCloudbenchTasksWithChan invokes the das.CreateCloudbenchTasks API asynchronously
func (client *Client) CreateCloudbenchTasksWithChan(request *CreateCloudbenchTasksRequest) (<-chan *CreateCloudbenchTasksResponse, <-chan error) {
	responseChan := make(chan *CreateCloudbenchTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCloudbenchTasks(request)
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

// CreateCloudbenchTasksWithCallback invokes the das.CreateCloudbenchTasks API asynchronously
func (client *Client) CreateCloudbenchTasksWithCallback(request *CreateCloudbenchTasksRequest, callback func(response *CreateCloudbenchTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCloudbenchTasksResponse
		var err error
		defer close(result)
		response, err = client.CreateCloudbenchTasks(request)
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

// CreateCloudbenchTasksRequest is the request struct for api CreateCloudbenchTasks
type CreateCloudbenchTasksRequest struct {
	*requests.RpcRequest
	ClientType          string `position:"Query" name:"ClientType"`
	DstPort             string `position:"Query" name:"DstPort"`
	Description         string `position:"Query" name:"Description"`
	RequestStartTime    string `position:"Query" name:"RequestStartTime"`
	DstConnectionString string `position:"Query" name:"DstConnectionString"`
	DstSuperPassword    string `position:"Query" name:"DstSuperPassword"`
	DstSuperAccount     string `position:"Query" name:"DstSuperAccount"`
	DstInstanceId       string `position:"Query" name:"DstInstanceId"`
	Rate                string `position:"Query" name:"Rate"`
	RequestDuration     string `position:"Query" name:"RequestDuration"`
	DtsJobId            string `position:"Query" name:"DtsJobId"`
	RequestEndTime      string `position:"Query" name:"RequestEndTime"`
	Amount              string `position:"Query" name:"Amount"`
	TaskType            string `position:"Query" name:"TaskType"`
	EndState            string `position:"Query" name:"EndState"`
	BackupId            string `position:"Query" name:"BackupId"`
	SrcSuperPassword    string `position:"Query" name:"SrcSuperPassword"`
	BackupTime          string `position:"Query" name:"BackupTime"`
	GatewayVpcIp        string `position:"Query" name:"GatewayVpcIp"`
	WorkDir             string `position:"Query" name:"WorkDir"`
	DtsJobClass         string `position:"Query" name:"DtsJobClass"`
	SrcPublicIp         string `position:"Query" name:"SrcPublicIp"`
	SrcInstanceId       string `position:"Query" name:"SrcInstanceId"`
	DstType             string `position:"Query" name:"DstType"`
	SrcSuperAccount     string `position:"Query" name:"SrcSuperAccount"`
	GatewayVpcId        string `position:"Query" name:"GatewayVpcId"`
	SmartPressureTime   string `position:"Query" name:"SmartPressureTime"`
}

// CreateCloudbenchTasksResponse is the response struct for api CreateCloudbenchTasks
type CreateCloudbenchTasksResponse struct {
	*responses.BaseResponse
	Code      string                      `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Success   string                      `json:"Success" xml:"Success"`
	Data      DataInCreateCloudbenchTasks `json:"Data" xml:"Data"`
}

// CreateCreateCloudbenchTasksRequest creates a request to invoke CreateCloudbenchTasks API
func CreateCreateCloudbenchTasksRequest() (request *CreateCloudbenchTasksRequest) {
	request = &CreateCloudbenchTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "CreateCloudbenchTasks", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCloudbenchTasksResponse creates a response to parse from CreateCloudbenchTasks response
func CreateCreateCloudbenchTasksResponse() (response *CreateCloudbenchTasksResponse) {
	response = &CreateCloudbenchTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
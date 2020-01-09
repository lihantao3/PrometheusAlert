package sas

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

// ModifyOperateVul invokes the sas.ModifyOperateVul API synchronously
// api document: https://help.aliyun.com/api/sas/modifyoperatevul.html
func (client *Client) ModifyOperateVul(request *ModifyOperateVulRequest) (response *ModifyOperateVulResponse, err error) {
	response = CreateModifyOperateVulResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyOperateVulWithChan invokes the sas.ModifyOperateVul API asynchronously
// api document: https://help.aliyun.com/api/sas/modifyoperatevul.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyOperateVulWithChan(request *ModifyOperateVulRequest) (<-chan *ModifyOperateVulResponse, <-chan error) {
	responseChan := make(chan *ModifyOperateVulResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyOperateVul(request)
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

// ModifyOperateVulWithCallback invokes the sas.ModifyOperateVul API asynchronously
// api document: https://help.aliyun.com/api/sas/modifyoperatevul.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyOperateVulWithCallback(request *ModifyOperateVulRequest, callback func(response *ModifyOperateVulResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyOperateVulResponse
		var err error
		defer close(result)
		response, err = client.ModifyOperateVul(request)
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

// ModifyOperateVulRequest is the request struct for api ModifyOperateVul
type ModifyOperateVulRequest struct {
	*requests.RpcRequest
	Reason      string `position:"Query" name:"Reason"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	OperateType string `position:"Query" name:"OperateType"`
	Type        string `position:"Query" name:"Type"`
	Info        string `position:"Query" name:"Info"`
}

// ModifyOperateVulResponse is the response struct for api ModifyOperateVul
type ModifyOperateVulResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyOperateVulRequest creates a request to invoke ModifyOperateVul API
func CreateModifyOperateVulRequest() (request *ModifyOperateVulRequest) {
	request = &ModifyOperateVulRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyOperateVul", "sas", "openAPI")
	return
}

// CreateModifyOperateVulResponse creates a response to parse from ModifyOperateVul response
func CreateModifyOperateVulResponse() (response *ModifyOperateVulResponse) {
	response = &ModifyOperateVulResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
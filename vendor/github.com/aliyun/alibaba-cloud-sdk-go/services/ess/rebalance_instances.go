package ess

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

// RebalanceInstances invokes the ess.RebalanceInstances API synchronously
// api document: https://help.aliyun.com/api/ess/rebalanceinstances.html
func (client *Client) RebalanceInstances(request *RebalanceInstancesRequest) (response *RebalanceInstancesResponse, err error) {
	response = CreateRebalanceInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// RebalanceInstancesWithChan invokes the ess.RebalanceInstances API asynchronously
// api document: https://help.aliyun.com/api/ess/rebalanceinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RebalanceInstancesWithChan(request *RebalanceInstancesRequest) (<-chan *RebalanceInstancesResponse, <-chan error) {
	responseChan := make(chan *RebalanceInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebalanceInstances(request)
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

// RebalanceInstancesWithCallback invokes the ess.RebalanceInstances API asynchronously
// api document: https://help.aliyun.com/api/ess/rebalanceinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RebalanceInstancesWithCallback(request *RebalanceInstancesRequest, callback func(response *RebalanceInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebalanceInstancesResponse
		var err error
		defer close(result)
		response, err = client.RebalanceInstances(request)
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

// RebalanceInstancesRequest is the request struct for api RebalanceInstances
type RebalanceInstancesRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// RebalanceInstancesResponse is the response struct for api RebalanceInstances
type RebalanceInstancesResponse struct {
	*responses.BaseResponse
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

// CreateRebalanceInstancesRequest creates a request to invoke RebalanceInstances API
func CreateRebalanceInstancesRequest() (request *RebalanceInstancesRequest) {
	request = &RebalanceInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "RebalanceInstances", "ess", "openAPI")
	return
}

// CreateRebalanceInstancesResponse creates a response to parse from RebalanceInstances response
func CreateRebalanceInstancesResponse() (response *RebalanceInstancesResponse) {
	response = &RebalanceInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package cdn

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

// SetUserDomainBlackList invokes the cdn.SetUserDomainBlackList API synchronously
// api document: https://help.aliyun.com/api/cdn/setuserdomainblacklist.html
func (client *Client) SetUserDomainBlackList(request *SetUserDomainBlackListRequest) (response *SetUserDomainBlackListResponse, err error) {
	response = CreateSetUserDomainBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// SetUserDomainBlackListWithChan invokes the cdn.SetUserDomainBlackList API asynchronously
// api document: https://help.aliyun.com/api/cdn/setuserdomainblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserDomainBlackListWithChan(request *SetUserDomainBlackListRequest) (<-chan *SetUserDomainBlackListResponse, <-chan error) {
	responseChan := make(chan *SetUserDomainBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetUserDomainBlackList(request)
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

// SetUserDomainBlackListWithCallback invokes the cdn.SetUserDomainBlackList API asynchronously
// api document: https://help.aliyun.com/api/cdn/setuserdomainblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserDomainBlackListWithCallback(request *SetUserDomainBlackListRequest, callback func(response *SetUserDomainBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetUserDomainBlackListResponse
		var err error
		defer close(result)
		response, err = client.SetUserDomainBlackList(request)
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

// SetUserDomainBlackListRequest is the request struct for api SetUserDomainBlackList
type SetUserDomainBlackListRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
}

// SetUserDomainBlackListResponse is the response struct for api SetUserDomainBlackList
type SetUserDomainBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetUserDomainBlackListRequest creates a request to invoke SetUserDomainBlackList API
func CreateSetUserDomainBlackListRequest() (request *SetUserDomainBlackListRequest) {
	request = &SetUserDomainBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetUserDomainBlackList", "", "")
	return
}

// CreateSetUserDomainBlackListResponse creates a response to parse from SetUserDomainBlackList response
func CreateSetUserDomainBlackListResponse() (response *SetUserDomainBlackListResponse) {
	response = &SetUserDomainBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/models"
)

type CalculateExpansionPriceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    ClusterExpansion *jmr.ClusterExpansion `json:"clusterExpansion"`
}

/*
 * param regionId: 地域ID (Required)
 * param clusterExpansion:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCalculateExpansionPriceRequest(
    regionId string,
    clusterExpansion *jmr.ClusterExpansion,
) *CalculateExpansionPriceRequest {

	return &CalculateExpansionPriceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cluster/expansionPrice:calculate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterExpansion: clusterExpansion,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param clusterExpansion:  (Required)
 */
func NewCalculateExpansionPriceRequestWithAllParams(
    regionId string,
    clusterExpansion *jmr.ClusterExpansion,
) *CalculateExpansionPriceRequest {

    return &CalculateExpansionPriceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cluster/expansionPrice:calculate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterExpansion: clusterExpansion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCalculateExpansionPriceRequestWithoutParam() *CalculateExpansionPriceRequest {

    return &CalculateExpansionPriceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cluster/expansionPrice:calculate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CalculateExpansionPriceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterExpansion: (Required) */
func (r *CalculateExpansionPriceRequest) SetClusterExpansion(clusterExpansion *jmr.ClusterExpansion) {
    r.ClusterExpansion = clusterExpansion
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CalculateExpansionPriceRequest) GetRegionId() string {
    return r.RegionId
}

type CalculateExpansionPriceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CalculateExpansionPriceResult `json:"result"`
}

type CalculateExpansionPriceResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data int `json:"data"`
}
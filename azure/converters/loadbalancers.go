/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package converters

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-08-01/network"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
)

// SKUtoSDK converts infrav1.SKU into a network.LoadBalancerSkuName.
func SKUtoSDK(src infrav1.SKU) network.LoadBalancerSkuName {
	if src == infrav1.SKUStandard {
		return network.LoadBalancerSkuNameStandard
	}
	return ""
}
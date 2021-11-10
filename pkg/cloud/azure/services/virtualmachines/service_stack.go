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

package virtualmachines

import (
	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/compute/mgmt/compute"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

// StackHubService provides operations on resource groups
type StackHubService struct {
	Client compute.VirtualMachinesClient
	Scope  *actuators.MachineScope
}

// getVirtualNetworksClientStackHub creates a new groups client from subscriptionid.
func getVirtualMachinesClientStackHub(resourceManagerEndpoint, subscriptionID string, authorizer autorest.Authorizer) compute.VirtualMachinesClient {
	vmClient := compute.NewVirtualMachinesClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	vmClient.Authorizer = authorizer
	vmClient.AddToUserAgent(azure.UserAgent)
	return vmClient
}

// NewService creates a new groups service.
func NewServiceStackHub(scope *actuators.MachineScope) azure.Service {
	return &StackHubService{
		Client: getVirtualMachinesClientStackHub(scope.ResourceManagerEndpoint, scope.SubscriptionID, scope.Authorizer),
		Scope:  scope,
	}
}
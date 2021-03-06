// Copyright 2018 Authors of Cilium
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

package annotation

const (
	// Name is an optional annotation to the NetworkPolicy
	// resource which specifies the name of the policy node to which all
	// rules should be applied to.
	Name = "io.cilium.name"

	// V4CIDRName is the annotation name used to store the IPv4
	// pod CIDR in the node's annotations.
	V4CIDRName = "io.cilium.network.ipv4-pod-cidr"
	// V6CIDRName is the annotation name used to store the IPv6
	// pod CIDR in the node's annotations.
	V6CIDRName = "io.cilium.network.ipv6-pod-cidr"

	// V4HealthName is the annotation name used to store the IPv4
	// address of the cilium-health endpoint in the node's annotations.
	V4HealthName = "io.cilium.network.ipv4-health-ip"
	// V6HealthName is the annotation name used to store the IPv6
	// address of the cilium-health endpoint in the node's annotations.
	V6HealthName = "io.cilium.network.ipv6-health-ip"

	// CiliumHostIP is the annotation name used to store the IPv4 address
	// of the cilium host interface in the node's annotations.
	CiliumHostIP = "io.cilium.network.ipv4-cilium-host"
)

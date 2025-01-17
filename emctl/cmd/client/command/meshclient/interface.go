/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package meshclient

import (
	"context"

	"github.com/megaease/easemeshctl/cmd/client/resource"
)

// MeshClient is a client for accessing the EaseMesh control plane service
type MeshClient interface {
	V1Alpha1() V1Alpha1Interface
}

// V1Alpha1Interface is an interface that aggregates all resources accessor for the EaseMesh
type V1Alpha1Interface interface {
	TenantGetter
	ServiceGetter
	LoadbalanceGetter
	CanaryGetter
	ObservabilityGetter
	ResilienceGetter
	IngressGetter
}

// TenantGetter represents a Tenant resource accessor
type TenantGetter interface {
	Tenant() TenantInterface
}

// ServiceGetter represents a Service resource accessor
type ServiceGetter interface {
	Service() ServiceInterface
}

// LoadbalanceGetter represents a Loadbalance resource accessor
type LoadbalanceGetter interface {
	LoadBalance() LoadBalanceInterface
}

// CanaryGetter represents a Canary resource accessor
type CanaryGetter interface {
	Canary() CanaryInterface
}

// ObservabilityGetter represents an Observability resource accessor
type ObservabilityGetter interface {
	ObservabilityTracings() ObservabilityTracingInterface
	ObservabilityMetrics() ObservabilityMetricInterface
	ObservabilityOutputServer() ObservabilityOutputServerInterface
}

// ResilienceGetter represents a Resilience resource accessor
type ResilienceGetter interface {
	Resilience() ResilienceInterface
}

// IngressGetter represents an Ingress resource accessor
type IngressGetter interface {
	Ingress() IngressInterface
}

// TenantInterface captures the set of operations for interacting with the EaseMesh REST apis of the tenant resource.
type TenantInterface interface {
	Get(context.Context, string) (*resource.Tenant, error)
	Patch(context.Context, *resource.Tenant) error
	Create(context.Context, *resource.Tenant) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.Tenant, error)
}

// ServiceInterface captures the set of operations for interacting with the EaseMesh REST apis of the service resource.
type ServiceInterface interface {
	Get(context.Context, string) (*resource.Service, error)
	Patch(context.Context, *resource.Service) error
	Create(context.Context, *resource.Service) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.Service, error)
}

// LoadBalanceInterface captures the set of operations for interacting with the EaseMesh REST apis of the loadbalance resource.
type LoadBalanceInterface interface {
	Get(context.Context, string) (*resource.LoadBalance, error)
	Patch(context.Context, *resource.LoadBalance) error
	Create(context.Context, *resource.LoadBalance) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.LoadBalance, error)
}

// CanaryInterface captures the set of operations for interacting with the EaseMesh REST apis of the canaray resource.
type CanaryInterface interface {
	Get(context.Context, string) (*resource.Canary, error)
	Patch(context.Context, *resource.Canary) error
	Create(context.Context, *resource.Canary) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.Canary, error)
}

// ObservabilityOutputServerInterface captures the set of operations for interacting with the EaseMesh REST apis of the observability output server resource.
type ObservabilityOutputServerInterface interface {
	Get(context.Context, string) (*resource.ObservabilityOutputServer, error)
	Patch(context.Context, *resource.ObservabilityOutputServer) error
	Create(context.Context, *resource.ObservabilityOutputServer) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.ObservabilityOutputServer, error)
}

// ObservabilityMetricInterface captures the set of operations for interacting with the EaseMesh REST apis of the observability metric resource.
type ObservabilityMetricInterface interface {
	Get(context.Context, string) (*resource.ObservabilityMetrics, error)
	Patch(context.Context, *resource.ObservabilityMetrics) error
	Create(context.Context, *resource.ObservabilityMetrics) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.ObservabilityMetrics, error)
}

// ObservabilityTracingInterface captures the set of operations for interacting with the EaseMesh REST apis of the observability tracing resource.
type ObservabilityTracingInterface interface {
	Get(context.Context, string) (*resource.ObservabilityTracings, error)
	Patch(context.Context, *resource.ObservabilityTracings) error
	Create(context.Context, *resource.ObservabilityTracings) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.ObservabilityTracings, error)
}

// ResilienceInterface captures the set of operations for interacting with the EaseMesh REST apis of the resilience resource.
type ResilienceInterface interface {
	Get(context.Context, string) (*resource.Resilience, error)
	Patch(context.Context, *resource.Resilience) error
	Create(context.Context, *resource.Resilience) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.Resilience, error)
}

// IngressInterface captures the set of operations for interacting with the EaseMesh REST apis of the ingress resource.
type IngressInterface interface {
	Get(context.Context, string) (*resource.Ingress, error)
	Patch(context.Context, *resource.Ingress) error
	Create(context.Context, *resource.Ingress) error
	Delete(context.Context, string) error
	List(context.Context) ([]*resource.Ingress, error)
}

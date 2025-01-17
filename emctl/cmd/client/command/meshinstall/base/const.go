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

package installbase

const (
	ObjectsURL = "/apis/v1/objects"
	ObjectURL  = "/apis/v1/objects/%s"
	MemberList = "/apis/v1/status/members"
)

const (
	DefaultOperatorPath = "./manifests/easemesh-operator.yaml"

	DefaultMeshControlPlaneName                = "easemesh-control-plane"
	DefaultMeshClientPortName                  = "client-port"
	DefaultMeshPeerPortName                    = "peer-port"
	DefaultMeshAdminPortName                   = "admin-port"
	DefaultMeshControlPlanePlubicServiceName   = "easemesh-controlplane-public"
	DefaultMeshControlPlaneHeadlessServiceName = "easemesh-controlplane-hs"
	DefaultMeshControlPlaneServicePeerPort     = 2380
	DefaultMeshControlPlanelServiceAdminPort   = 2381

	DefaultMeshControlPlanePVName     = "easegress-control-plane-pv"
	DefaultMeshControlPlanePVHostPath = "/opt/easemesh"
	DefaultMeshControlPlaneConfig     = "easemesh-cluster-cm"

	DefaultMeshControllerName = "easemesh-controller"

	DefaultMeshOperatorName                         = "easemesh-operator"
	DefaultMeshOperatorControllerManagerServiceName = "mesh-operator-controller-manager-metrics-service"

	DefaultMeshIngressConfig         = "easemesh-ingress-config"
	DefaultMeshIngressService        = "easemesh-ingress-service"
	DefaultMeshIngressControllerName = "easemesh-ingress-easegress"

	// DefaultKubeDir represents default kubernetes client configuration directory
	DefaultKubeDir = ".kube"

	//
	DefaultKubernetesConfig = "config"
	WriterClusterRole       = "writer"
	ReaderClusterRole       = "reader"
)

type InstallPhase int

const (
	BeginPhase InstallPhase = iota
	EndPhase
	ErrorPhase
)

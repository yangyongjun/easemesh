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

package operator

import (
	"github.com/megaease/easemeshctl/cmd/client/command/flags"
	installbase "github.com/megaease/easemeshctl/cmd/client/command/meshinstall/base"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	roleVerbGet    = "get"
	roleVerbList   = "list"
	roleVerbWatch  = "watch"
	roleVerbCreate = "create"
	roleVerbUpdate = "update"
	roleVerbPatch  = "patch"
	roleVerbDelete = "delete"
)

func roleSpec(installFlags *flags.Install) installbase.InstallFunc {

	operatorLeaderElectionRole := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: installFlags.MeshNamespace,
			Name:      meshOperatorLeaderElectionRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"configmaps", "leases"},
				Verbs:     []string{roleVerbGet, roleVerbList, roleVerbWatch, roleVerbCreate, roleVerbUpdate, roleVerbPatch, roleVerbDelete},
			},
			{
				APIGroups: []string{"", "coordination.k8s.io"},
				Resources: []string{"events"},
				Verbs:     []string{roleVerbCreate, roleVerbPatch},
			},
		},
	}

	return func(cmd *cobra.Command, kubeClient *kubernetes.Clientset, installFlags *flags.Install) error {
		err := installbase.DeployRole(operatorLeaderElectionRole, kubeClient, installFlags.MeshNamespace)
		if err != nil {
			return err
		}
		return nil
	}
}

func clusterRoleSpec(installFlags *flags.Install) installbase.InstallFunc {
	operatorManagerClusterRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: meshOperatorManagerClusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"apps"},
				Resources: []string{"deployments"},
				Verbs:     []string{roleVerbGet, roleVerbList, roleVerbWatch, roleVerbCreate, roleVerbUpdate, roleVerbPatch, roleVerbDelete},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"pods"},
				Verbs:     []string{roleVerbGet, roleVerbList},
			},
			{
				APIGroups: []string{"mesh.megaease.com"},
				Resources: []string{"meshdeployments"},
				Verbs:     []string{roleVerbGet, roleVerbList, roleVerbWatch, roleVerbCreate, roleVerbUpdate, roleVerbPatch, roleVerbDelete},
			},
			{
				APIGroups: []string{"mesh.megaease.com"},
				Resources: []string{"meshdeployments/finalizers"},
				Verbs:     []string{roleVerbUpdate},
			},
			{
				APIGroups: []string{"mesh.megaease.com"},
				Resources: []string{"meshdeployments/status"},
				Verbs:     []string{roleVerbGet, roleVerbPatch, roleVerbUpdate},
			},
		},
	}

	metricsReaderClusterRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: meshOperatorMetricsReaderClusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				NonResourceURLs: []string{"/metrics"},
				Verbs:           []string{roleVerbGet},
			},
		},
	}

	operatorProxyClusterRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: meshOperatorProxyClusterRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{"authentication.k8s.io"},
				Resources: []string{"tokenreviews"},
				Verbs:     []string{roleVerbCreate},
			},
			{
				APIGroups: []string{"authentication.k8s.io"},
				Resources: []string{"subjectaccessreviews"},
				Verbs:     []string{roleVerbCreate},
			},
		},
	}

	return func(cmd *cobra.Command, kubeClient *kubernetes.Clientset, installFlags *flags.Install) error {
		for _, clusterRole := range []*rbacv1.ClusterRole{operatorManagerClusterRole, metricsReaderClusterRole, operatorProxyClusterRole} {
			err := installbase.DeployClusterRole(clusterRole, kubeClient)
			if err != nil {
				return errors.Wrapf(err, "createClusterRole role %s", clusterRole.Name)
			}

		}
		return nil
	}
}

func roleBindingSpec(installFlags *flags.Install) installbase.InstallFunc {
	operatorLeaderElectionRoleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      meshOperatorLeaderElectionRoleBinding,
			Namespace: installFlags.MeshNamespace,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     meshOperatorLeaderElectionRole,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: installFlags.MeshNamespace,
			},
		},
	}

	return func(cmd *cobra.Command, kubeClient *kubernetes.Clientset, installFlags *flags.Install) error {
		err := installbase.DeployRoleBinding(operatorLeaderElectionRoleBinding, kubeClient, installFlags.MeshNamespace)
		if err != nil {
			return err
		}
		return nil
	}
}

func clusterRoleBindingSpec(installFlags *flags.Install) installbase.InstallFunc {
	operatorManagerClusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: meshOperatorManagerClusterRoleBinding,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     meshOperatorManagerClusterRole,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: installFlags.MeshNamespace,
			},
		},
	}

	operatorProxyClusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: meshOperatorProxyClusterRoleBinding,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     meshOperatorProxyClusterRole,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: installFlags.MeshNamespace,
			},
		},
	}

	operatorMetricsReaderClusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: meshOperatorMetricsReaderClusterRoleBinding,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     meshOperatorMetricsReaderClusterRole,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: installFlags.MeshNamespace,
			},
		},
	}

	return func(cmd *cobra.Command, kubeClient *kubernetes.Clientset, installFlags *flags.Install) error {

		clusterRoleBindings := []*rbacv1.ClusterRoleBinding{
			operatorManagerClusterRoleBinding,
			operatorProxyClusterRoleBinding,
			operatorMetricsReaderClusterRoleBinding,
		}

		for _, clusterRoleBinding := range clusterRoleBindings {
			err := installbase.DeployClusterRoleBinding(clusterRoleBinding, kubeClient)
			if err != nil {
				return errors.Wrapf(err, "Create roleBinding %s", clusterRoleBinding.Name)
			}
		}
		return nil
	}
}

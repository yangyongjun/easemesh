package meshingress

import (
	"fmt"
	"time"

	installbase "github.com/megaease/easemeshctl/cmd/client/command/meshinstall/base"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Deploy deploy resources of mesh ingress controller
func Deploy(context *installbase.StageContext) error {
	err := installbase.BatchDeployResources(context.Cmd, context.Client, &context.Arguments, []installbase.InstallFunc{
		configMapSpec(&context.Arguments),
		serviceSpec(&context.Arguments),
		deploymentSpec(&context.Arguments),
	})
	if err != nil {
		return err
	}

	return checkMeshIngressStatus(context.Client, &context.Arguments)
}

// PreCheck check prerequisite for installing mesh ingress controller
func PreCheck(context *installbase.StageContext) error {
	return nil
}

// Clear will clear all installed resource about mesh ingress panel
func Clear(context *installbase.StageContext) error {
	appsV1Resources := [][]string{
		{"deployments", installbase.DefaultMeshIngressControllerName},
	}
	coreV1Resources := [][]string{
		{"services", installbase.DefaultMeshIngressService},
		{"configmap", installbase.DefaultMeshIngressConfig},
	}

	installbase.DeleteResources(context.Client, appsV1Resources, context.Arguments.MeshNameSpace, installbase.DeleteAppsV1Resource)
	installbase.DeleteResources(context.Client, coreV1Resources, context.Arguments.MeshNameSpace, installbase.DeleteCoreV1Resource)
	return nil
}

// Describe leverage human-readable text to describe different phase
// in the process of the mesh ingress controller
func Describe(context *installbase.StageContext, phase installbase.InstallPhase) string {
	switch phase {
	case installbase.BeginPhase:
		return fmt.Sprintf("Begin to install mesh ingress controller in the namespace:%s", context.Arguments.MeshNameSpace)
	case installbase.EndPhase:
		return fmt.Sprintf("\nMesh ingress controller deployed successfully, deployment:%s\n%s", installbase.DefaultMeshIngressControllerName,
			installbase.FormatPodStatus(context.Client, context.Arguments.MeshNameSpace,
				installbase.AdaptListPodFunc(meshIngressLabel())))
	}
	return ""
}

func checkMeshIngressStatus(client *kubernetes.Clientset, args *installbase.InstallArgs) error {
	i := 0
	for {
		time.Sleep(time.Millisecond * 100)
		i++
		if i > 600 {
			return errors.Errorf("easeMesh meshingress controller deploy failed, mesh ingress controller (EG deployment) not ready")
		}
		ready, err := installbase.CheckDeploymentResourceStatus(client, args.MeshNameSpace,
			installbase.DefaultMeshIngressControllerName,
			installbase.DeploymentReadyPredict)
		if ready {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
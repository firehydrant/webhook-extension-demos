package main

import (
	"context"
	"os"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// KubeClient holds the kubernetes clientset
type KubeClient struct {
	clientset *kubernetes.Clientset
}

// NewKubeClient creates a new Kubernetes client
func NewKubeClient() (*KubeClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return nil, err
	}

	// Create a new clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubeClient{clientset}, nil
}

// restartDeployment restarts a Kubernetes deployment
func (k *KubeClient) restartDeployment(namespace, deploymentName string) (*appsv1.Deployment, *appsv1.Deployment, error) {
	oldDeployment, err := k.clientset.AppsV1().Deployments(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	if oldDeployment.Spec.Template.ObjectMeta.Annotations == nil {
		oldDeployment.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	oldDeployment.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	newDeployment, err := k.clientset.AppsV1().Deployments(namespace).Update(context.Background(), oldDeployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, nil, err
	}

	return oldDeployment, newDeployment, nil
}

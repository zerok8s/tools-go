package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceRepository struct {
	kubeClient *kubernetes.Clientset
}

func NewNamesapceRepository(client *kubernetes.Clientset) *NamespaceRepository {
	return &NamespaceRepository{kubeClient: client}
}

func (r *NamespaceRepository) Apply(namespaceName string) (result *corev1.Namespace, err error) {
	return r.kubeClient.CoreV1().Namespaces().Apply(
		context.TODO(),
		applycorev1.Namespace(namespaceName),
		metav1.ApplyOptions{FieldManager: FieldManagerApplyPatch},
	)
}

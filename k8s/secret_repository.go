package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretRepository struct {
	kubeClient *kubernetes.Clientset
	secretList *ExtendedSecretList
}

func NewSecretRepository(client *kubernetes.Clientset) *SecretRepository {
	return &SecretRepository{kubeClient: client}
}

func (r *SecretRepository) GetAll() *ExtendedSecretList {
	if r.secretList == nil {
		r.secretList = r.fetchSecrets()
	}

	return r.secretList
}

func (r *SecretRepository) Apply(namespaceName string, secretConfiguration *applycorev1.SecretApplyConfiguration) (result *corev1.Secret, err error) {
	return r.kubeClient.CoreV1().Secrets(namespaceName).Apply(
		context.TODO(),
		secretConfiguration,
		metav1.ApplyOptions{FieldManager: FieldManagerApplyPatch},
	)
}

func (r *SecretRepository) Delete(namespaceName string, secretName string) error {
	return r.kubeClient.CoreV1().Secrets(namespaceName).Delete(
		context.TODO(),
		secretName,
		metav1.DeleteOptions{},
	)
}

func (r *SecretRepository) fetchSecrets() *ExtendedSecretList {
	secrets, _ := r.kubeClient.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})

	return (*ExtendedSecretList)(secrets)
}

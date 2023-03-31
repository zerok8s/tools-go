package k8s

import corev1 "k8s.io/api/core/v1"

type ExtendedSecretList corev1.SecretList

func (sl ExtendedSecretList) FindSecret(namespace string, name string) *corev1.Secret {
	for _, item := range sl.Items {
		if item.Namespace == namespace && item.Name == name {
			return &item
		}
	}

	return nil
}

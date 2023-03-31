package k8s

import corev1 "k8s.io/api/core/v1"

type ExtendedPodSpec corev1.PodSpec

func (ps ExtendedPodSpec) IsSecretUsed(name string) bool {
	for _, c := range ps.Containers {
		for _, e := range c.EnvFrom {
			if e.SecretRef == nil {
				continue
			}
			if e.SecretRef.Name == name {
				return true
			}
		}
	}
	for _, v := range ps.Volumes {
		if v.Secret == nil {
			continue
		}

		if v.Secret.SecretName == name {
			return true
		}
	}

	return false
}

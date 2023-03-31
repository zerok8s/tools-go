package k8s

import (
	"flag"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

const (
	FieldManagerApplyPatch string = "application/apply-patch"
)

func GetKubeConfig() *rest.Config {
	config, err := rest.InClusterConfig()
	if err == nil {
		return config
	}

	var kubeConfigPath *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfigPath = flag.String("config", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfigPath = flag.String("config", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err = clientcmd.BuildConfigFromFlags("", *kubeConfigPath)
	if err == nil {
		return config
	}

	panic(err)
}

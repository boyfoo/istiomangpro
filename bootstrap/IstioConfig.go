package bootstrap

import (
	"istio.io/client-go/pkg/informers/externalversions"
	"istiomang/pkg/vs"
	"k8s.io/client-go/dynamic"

	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type K8sConfig struct {
	VsHandler *vs.VsHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}
func (this *K8sConfig) IstioRestClient() *istio.Clientset {
	client, err := istio.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func (*K8sConfig) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", "./resources/config")
	//config.Insecure = true
	if err != nil {
		log.Fatal(err)
	}
	return config
}

//初始化client-go客户端
func (this *K8sConfig) InitClient() *kubernetes.Clientset {

	c, err := kubernetes.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (this *K8sConfig) InitDynamicClient() dynamic.Interface {
	c, err := dynamic.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//初始化Informer
func (this *K8sConfig) InitInformer() externalversions.SharedInformerFactory {
	fact := externalversions.NewSharedInformerFactoryWithOptions(this.IstioRestClient(), 0)
	//虚拟服务的监听
	fact.Networking().V1alpha3().VirtualServices().Informer().AddEventHandler(this.VsHandler)
	fact.Start(wait.NeverStop)
	return fact
}

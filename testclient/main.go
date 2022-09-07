package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	"istiomang/bootstrap"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	yaml2 "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/cache"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// 动态客户端新增
	config := bootstrap.NewK8sConfig()
	clientset := config.InitClient()

	depListWatcher := cache.NewListWatchFromClient(clientset.AppsV1().RESTClient(), "deployments", v1.NamespaceSystem, fields.Everything())

	// 自定义key名称 根据命名空间分索引 可以多个
	indexer := cache.Indexers{"test001": cache.MetaNamespaceIndexFunc}

	myIndexer, controller := cache.NewIndexerInformer(depListWatcher, &v12.DeploymentList{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			//fmt.Println(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			//fmt.Println(newObj)
		},
		DeleteFunc: func(obj interface{}) {
			//fmt.Println(obj)
		},
	}, indexer)

	stopChan := make(chan struct{})
	go controller.Run(stopChan)
	defer close(stopChan)
	// 等待同步后在运行后面的代码
	if !cache.WaitForCacheSync(stopChan, controller.HasSynced) {
		log.Fatalln("sync error")
	}
	fmt.Println(myIndexer.ListKeys())

	//fmt.Println(myIndexer.GetByKey("kube-system/metrics-server"))
	// 上面设置的key名称
	fmt.Println(myIndexer.IndexKeys("test001", v1.NamespaceSystem))
}

func main11() {
	// kustomization.yaml多个资源
	deployYaml := kustomize("testclient/ymals")
	decoder := yaml.NewYAMLOrJSONDecoder(strings.NewReader(deployYaml), len(deployYaml))

	config := bootstrap.NewK8sConfig()
	dc := config.InitDynamicClient()
	clientset := config.InitClient()

	for {
		var rawObj runtime.RawExtension
		err := decoder.Decode(&rawObj)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		obj, gvk, err := yaml2.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			log.Fatalln(err)
		}
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			log.Fatalln(err)
		}
		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}
		// 所有api-resoucse
		gr, err := restmapper.GetAPIGroupResources(clientset.Discovery())
		if err != nil {
			log.Fatalln(err)
		}
		mapping, err := restmapper.NewDiscoveryRESTMapper(gr).RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			log.Fatalln(err)
		}

		gvr := schema.GroupVersionResource{
			Group:    mapping.Resource.Group,
			Version:  mapping.Resource.Version,
			Resource: mapping.Resource.Resource,
		}

		_, err = dc.Resource(gvr).Namespace(unstructuredObj.GetNamespace()).Create(context.Background(), unstructuredObj, v1.CreateOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func kustomize(path string) string {
	command := exec.Command("kubectl", "kustomize", path)
	var ret bytes.Buffer
	command.Stdout = &ret
	if err := command.Run(); err != nil {
		log.Fatalln(err)
		return ""
	}

	return ret.String()
}

func main1() {
	// 动态客户端新增
	config := bootstrap.NewK8sConfig()
	client := config.InitDynamicClient()
	resource := client.Resource(schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1beta1",
		Resource: "gateways",
	})

	obj := &unstructured.Unstructured{}

	file, err := ioutil.ReadFile("./yamls/gateway.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(file, obj)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = resource.Namespace("myistio").Create(context.Background(), obj, v1.CreateOptions{})
	if err != nil {
		log.Fatalln(err)
	}
}

func main2() {
	// 动态客户端获取
	config := bootstrap.NewK8sConfig()
	client := config.InitDynamicClient()
	list, err := client.Resource(schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1beta1",
		Resource: "gateways",
	}).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, item := range list.Items {
		fmt.Println(item.GetName())
	}
	depUnstr, err := list.MarshalJSON()
	if err != nil {
		log.Fatalln(err)
	}

	depList := &v1beta1.GatewayList{}
	err = json.Unmarshal(depUnstr, depList)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range depList.Items {
		fmt.Println(item.Name)
	}
}

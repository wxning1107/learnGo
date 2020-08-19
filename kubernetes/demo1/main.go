package main

import (
	"fmt"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func InitClient() (clientset *kubernetes.Clientset, err error) {
	restConf, err := GetRestConf()
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func GetRestConf() (restConf *rest.Config, err error) {
	kubeConfig, err := ioutil.ReadFile("./admin.conf")
	if err != nil {
		return nil, err
	}

	restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return restConf, nil
}

func main() {
	clientset, err := InitClient()
	if err != nil {
		log.Panicln(err)
		return
	}
	podsList, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("pods list: %v\n", *podsList)
}

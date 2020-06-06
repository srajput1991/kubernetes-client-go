package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	//"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
//	"k8s.io/apimachinery/pkg/watch"
	"path/filepath"	
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

)

func main() {
	var ns,label,field string
//	kubeconfig := os.Getenv("KUBECONFIG")
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	fmt.Println(kubeconfig)
	flag.StringVar(&ns,"namespace","","provide the namespace")
	flag.StringVar(&field, "f", "", "Field selector")
	flag.StringVar(&label, "l","", "give the label")
	flag.StringVar(&kubeconfig, "kubeconfig",kubeconfig, "kubeconfig file")
	flag.Parse()
	
	// bootstrap config
	fmt.Println("Using kubeconfig:",kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	api := clientset.CoreV1()

	// initial list
	listOptions := metav1.ListOptions{LabelSelector: label, FieldSelector: field}
	pd, err := api.Pods(ns).List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	for _,pod := range pd.Items{
        fmt.Fprintf(os.Stdout, "pod name: %v\n", pod.Name)
         }
//	fmt.Println(pd)


}

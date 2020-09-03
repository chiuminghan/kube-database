package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/tools/clientcmd"
    examplecomclientset "github.com/chiuminghan/kube-database/pkg/client/clientset/versioned"
)

func main() {
    var kubeconfig *string
    if home := homeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    exampleClient, err := examplecomclientset.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    list, err := exampleClient.ExampleV1().Databases("default").List(metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }

    for _, db := range list.Items {
        fmt.Printf("database %s with user %q\n", db.Name, db.Spec.User)
    }
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}

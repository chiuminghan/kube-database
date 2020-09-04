package main

import (
    "flag"
    "os"
    "time"
    "path/filepath"

    "k8s.io/klog"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/tools/record"
    "k8s.io/client-go/util/workqueue"
    typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

    "k8s.io/apimachinery/pkg/util/wait"
    utilruntime "k8s.io/apimachinery/pkg/util/runtime"

    corev1 "k8s.io/api/core/v1"
    apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

    xscheme      "github.com/chiuminghan/kube-database/pkg/client/clientset/versioned/scheme"
    xclientset   "github.com/chiuminghan/kube-database/pkg/client/clientset/versioned"
    xlister      "github.com/chiuminghan/kube-database/pkg/client/listers/example.com/v1"
    xinformers   "github.com/chiuminghan/kube-database/pkg/client/informers/externalversions"
    xapi          "github.com/chiuminghan/kube-database/pkg/apis/example.com/v1"
)

type Controller struct {
    kubeclientset          kubernetes.Interface
    apiextensionsclientset apiextensionsclientset.Interface
    xclientset             xclientset.Interface
    informer               cache.SharedIndexInformer
    lister                 xlister.DatabaseLister
    recorder               record.EventRecorder
    workqueue              workqueue.RateLimitingInterface
}

func NewController() *Controller {

    var kubeconfig *string
    if home := homeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    kubeClient := kubernetes.NewForConfigOrDie(config)
    apiextensionsClient := apiextensionsclientset.NewForConfigOrDie(config)
    myClient := xclientset.NewForConfigOrDie(config)

    informerFactory := xinformers.NewSharedInformerFactory(myClient, time.Minute*1)
    informer := informerFactory.Example().V1().Databases()

    utilruntime.Must(xapi.AddToScheme(xscheme.Scheme))
    eventBroadcaster := record.NewBroadcaster()
    eventBroadcaster.StartLogging(klog.Infof)
    eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
    recorder := eventBroadcaster.NewRecorder(xscheme.Scheme, corev1.EventSource{Component: "database-controller"})

    workqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

    c := &Controller{
        kubeclientset:          kubeClient,
        apiextensionsclientset: apiextensionsClient,
        xclientset:             myClient,
        informer:               informer.Informer(),
        lister:                 informer.Lister(),
        recorder:               recorder,
        workqueue:              workqueue,
    }

    informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(newObject interface{}) {
            klog.Infof("Added: %v", newObject)
        },
        UpdateFunc: func(oldObject, newObject interface{}) {
            klog.Infof("Updated: %v", newObject)
        },
        DeleteFunc: func(object interface{}) {
            klog.Infof("Deleted: %v", object)
        },
    })

    informerFactory.Start(wait.NeverStop)

   return c
}

func (c *Controller) Run(){
    defer utilruntime.HandleCrash()
    defer c.workqueue.ShutDown()

    klog.Infoln("Waiting cache to be synced.")
    // Handle timeout for syncing.
    timeout := time.NewTimer(time.Second * 6000)
    timeoutCh := make(chan struct{})
    go func() {
        <-timeout.C
        timeoutCh <- struct{}{}
    }()

    if ok := cache.WaitForCacheSync(timeoutCh, c.informer.HasSynced); !ok {
        klog.Fatalln("Timeout expired during waiting for caches to sync.")
    }

    klog.Infoln("Starting custom controller.")
    select {}
}

func main() {
    controller := NewController()
    controller.Run()
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}

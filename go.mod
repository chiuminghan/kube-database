module github.com/chiuminghan/kube-database

go 1.13

require (
	github.com/imdario/mergo v0.3.11 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.19.0 // indirect
	k8s.io/apiextensions-apiserver v0.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200821003339-5e75c0163111 // indirect
)

replace (
	// https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-505627280
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/apiserver => k8s.io/apiserver v0.17.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
	k8s.io/component-base => k8s.io/component-base v0.17.6
	k8s.io/cri-api => k8s.io/cri-api v0.17.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.6
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.6
	k8s.io/kubectl => k8s.io/kubectl v0.17.6
	k8s.io/kubelet => k8s.io/kubelet v0.17.6
	k8s.io/kubernetes => k8s.io/kubernetes v1.17.6 // gitops-engine wants that
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.6
	k8s.io/metrics => k8s.io/metrics v0.17.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.6
)

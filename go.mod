module github.com/wonderflow/keda-api

go 1.14

require (
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	knative.dev/pkg v0.0.0-20201026015042-efadd36f9c63
	sigs.k8s.io/controller-runtime v0.6.1
)

replace k8s.io/client-go => k8s.io/client-go v0.18.6

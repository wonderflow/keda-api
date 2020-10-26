package main

import (
	"context"
	"log"

	kedaapi "github.com/wonderflow/keda-api/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var scheme = runtime.NewScheme()

func init() {
	_ = kedaapi.AddToScheme(scheme)
}

func main() {
	k8sClient, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.Create(context.Background(), &kedaapi.ScaledObject{
		ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
		Spec:       kedaapi.ScaledObjectSpec{},
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubean.io/api/generated/kubeanofflineversion/clientset/versioned/typed/kubeanofflineversion/v1alpha1"
)

type FakeKubeanV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKubeanV1alpha1) KuBeanOfflineVersions() v1alpha1.KuBeanOfflineVersionInterface {
	return &FakeKuBeanOfflineVersions{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubeanV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
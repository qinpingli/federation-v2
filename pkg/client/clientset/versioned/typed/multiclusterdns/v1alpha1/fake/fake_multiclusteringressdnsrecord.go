/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/multiclusterdns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressDNSRecords implements IngressDNSRecordInterface
type FakeIngressDNSRecords struct {
	Fake *FakeMulticlusterdnsV1alpha1
	ns   string
}

var ingressdnsrecordsResource = schema.GroupVersionResource{Group: "multiclusterdns.federation.k8s.io", Version: "v1alpha1", Resource: "ingressdnsrecords"}

var ingressdnsrecordsKind = schema.GroupVersionKind{Group: "multiclusterdns.federation.k8s.io", Version: "v1alpha1", Kind: "IngressDNSRecord"}

// Get takes name of the IngressDNSRecord, and returns the corresponding IngressDNSRecord object, and an error if there is any.
func (c *FakeIngressDNSRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.IngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingressdnsrecordsResource, c.ns, name), &v1alpha1.IngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressDNSRecord), err
}

// List takes label and field selectors, and returns the list of IngressDNSRecords that match those selectors.
func (c *FakeIngressDNSRecords) List(opts v1.ListOptions) (result *v1alpha1.IngressDNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingressdnsrecordsResource, ingressdnsrecordsKind, c.ns, opts), &v1alpha1.IngressDNSRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IngressDNSRecordList{ListMeta: obj.(*v1alpha1.IngressDNSRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.IngressDNSRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested IngressDNSRecords.
func (c *FakeIngressDNSRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingressdnsrecordsResource, c.ns, opts))

}

// Create takes the representation of a IngressDNSRecord and creates it.  Returns the server's representation of the IngressDNSRecord, and an error, if there is any.
func (c *FakeIngressDNSRecords) Create(IngressDNSRecord *v1alpha1.IngressDNSRecord) (result *v1alpha1.IngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingressdnsrecordsResource, c.ns, IngressDNSRecord), &v1alpha1.IngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressDNSRecord), err
}

// Update takes the representation of a IngressDNSRecord and updates it. Returns the server's representation of the IngressDNSRecord, and an error, if there is any.
func (c *FakeIngressDNSRecords) Update(IngressDNSRecord *v1alpha1.IngressDNSRecord) (result *v1alpha1.IngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingressdnsrecordsResource, c.ns, IngressDNSRecord), &v1alpha1.IngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressDNSRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIngressDNSRecords) UpdateStatus(IngressDNSRecord *v1alpha1.IngressDNSRecord) (*v1alpha1.IngressDNSRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ingressdnsrecordsResource, "status", c.ns, IngressDNSRecord), &v1alpha1.IngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressDNSRecord), err
}

// Delete takes name of the IngressDNSRecord and deletes it. Returns an error if one occurs.
func (c *FakeIngressDNSRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingressdnsrecordsResource, c.ns, name), &v1alpha1.IngressDNSRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressDNSRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingressdnsrecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IngressDNSRecordList{})
	return err
}

// Patch applies the patch and returns the patched IngressDNSRecord.
func (c *FakeIngressDNSRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IngressDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressdnsrecordsResource, c.ns, name, data, subresources...), &v1alpha1.IngressDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressDNSRecord), err
}

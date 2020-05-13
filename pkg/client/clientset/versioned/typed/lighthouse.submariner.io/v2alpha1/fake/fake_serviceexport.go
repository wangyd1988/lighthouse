/*
Copyright The Kubernetes Authors.

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
	v2alpha1 "github.com/submariner-io/lighthouse/pkg/apis/lighthouse.submariner.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceExports implements ServiceExportInterface
type FakeServiceExports struct {
	Fake *FakeLighthouseV2alpha1
	ns   string
}

var serviceexportsResource = schema.GroupVersionResource{Group: "lighthouse.submariner.io", Version: "v2alpha1", Resource: "serviceexports"}

var serviceexportsKind = schema.GroupVersionKind{Group: "lighthouse.submariner.io", Version: "v2alpha1", Kind: "ServiceExport"}

// Get takes name of the serviceExport, and returns the corresponding serviceExport object, and an error if there is any.
func (c *FakeServiceExports) Get(name string, options v1.GetOptions) (result *v2alpha1.ServiceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceexportsResource, c.ns, name), &v2alpha1.ServiceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ServiceExport), err
}

// List takes label and field selectors, and returns the list of ServiceExports that match those selectors.
func (c *FakeServiceExports) List(opts v1.ListOptions) (result *v2alpha1.ServiceExportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceexportsResource, serviceexportsKind, c.ns, opts), &v2alpha1.ServiceExportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.ServiceExportList{ListMeta: obj.(*v2alpha1.ServiceExportList).ListMeta}
	for _, item := range obj.(*v2alpha1.ServiceExportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceExports.
func (c *FakeServiceExports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceexportsResource, c.ns, opts))

}

// Create takes the representation of a serviceExport and creates it.  Returns the server's representation of the serviceExport, and an error, if there is any.
func (c *FakeServiceExports) Create(serviceExport *v2alpha1.ServiceExport) (result *v2alpha1.ServiceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceexportsResource, c.ns, serviceExport), &v2alpha1.ServiceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ServiceExport), err
}

// Update takes the representation of a serviceExport and updates it. Returns the server's representation of the serviceExport, and an error, if there is any.
func (c *FakeServiceExports) Update(serviceExport *v2alpha1.ServiceExport) (result *v2alpha1.ServiceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceexportsResource, c.ns, serviceExport), &v2alpha1.ServiceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ServiceExport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceExports) UpdateStatus(serviceExport *v2alpha1.ServiceExport) (*v2alpha1.ServiceExport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceexportsResource, "status", c.ns, serviceExport), &v2alpha1.ServiceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ServiceExport), err
}

// Delete takes name of the serviceExport and deletes it. Returns an error if one occurs.
func (c *FakeServiceExports) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceexportsResource, c.ns, name), &v2alpha1.ServiceExport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceExports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceexportsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2alpha1.ServiceExportList{})
	return err
}

// Patch applies the patch and returns the patched serviceExport.
func (c *FakeServiceExports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2alpha1.ServiceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceexportsResource, c.ns, name, pt, data, subresources...), &v2alpha1.ServiceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ServiceExport), err
}

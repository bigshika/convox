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

package v2

import (
	v2 "github.com/convox/convox/pkg/atom/pkg/apis/convox/v2"
	scheme "github.com/convox/convox/pkg/atom/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AtomsGetter has a method to return a AtomInterface.
// A group's client should implement this interface.
type AtomsGetter interface {
	Atoms(namespace string) AtomInterface
}

// AtomInterface has methods to work with Atom resources.
type AtomInterface interface {
	Create(*v2.Atom) (*v2.Atom, error)
	Update(*v2.Atom) (*v2.Atom, error)
	UpdateStatus(*v2.Atom) (*v2.Atom, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.Atom, error)
	List(opts v1.ListOptions) (*v2.AtomList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.Atom, err error)
	AtomExpansion
}

// atoms implements AtomInterface
type atoms struct {
	client rest.Interface
	ns     string
}

// newAtoms returns a Atoms
func newAtoms(c *ConvoxV2Client, namespace string) *atoms {
	return &atoms{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the atom, and returns the corresponding atom object, and an error if there is any.
func (c *atoms) Get(name string, options v1.GetOptions) (result *v2.Atom, err error) {
	result = &v2.Atom{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("atoms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Atoms that match those selectors.
func (c *atoms) List(opts v1.ListOptions) (result *v2.AtomList, err error) {
	result = &v2.AtomList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("atoms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested atoms.
func (c *atoms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("atoms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a atom and creates it.  Returns the server's representation of the atom, and an error, if there is any.
func (c *atoms) Create(atom *v2.Atom) (result *v2.Atom, err error) {
	result = &v2.Atom{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("atoms").
		Body(atom).
		Do().
		Into(result)
	return
}

// Update takes the representation of a atom and updates it. Returns the server's representation of the atom, and an error, if there is any.
func (c *atoms) Update(atom *v2.Atom) (result *v2.Atom, err error) {
	result = &v2.Atom{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("atoms").
		Name(atom.Name).
		Body(atom).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *atoms) UpdateStatus(atom *v2.Atom) (result *v2.Atom, err error) {
	result = &v2.Atom{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("atoms").
		Name(atom.Name).
		SubResource("status").
		Body(atom).
		Do().
		Into(result)
	return
}

// Delete takes name of the atom and deletes it. Returns an error if one occurs.
func (c *atoms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("atoms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *atoms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("atoms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched atom.
func (c *atoms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.Atom, err error) {
	result = &v2.Atom{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("atoms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

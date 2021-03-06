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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openservicemesh/osm/experimental/pkg/apis/policy/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackpressureLister helps list Backpressures.
// All objects returned here must be treated as read-only.
type BackpressureLister interface {
	// List lists all Backpressures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Backpressure, err error)
	// Backpressures returns an object that can list and get Backpressures.
	Backpressures(namespace string) BackpressureNamespaceLister
	BackpressureListerExpansion
}

// backpressureLister implements the BackpressureLister interface.
type backpressureLister struct {
	indexer cache.Indexer
}

// NewBackpressureLister returns a new BackpressureLister.
func NewBackpressureLister(indexer cache.Indexer) BackpressureLister {
	return &backpressureLister{indexer: indexer}
}

// List lists all Backpressures in the indexer.
func (s *backpressureLister) List(selector labels.Selector) (ret []*v1alpha1.Backpressure, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Backpressure))
	})
	return ret, err
}

// Backpressures returns an object that can list and get Backpressures.
func (s *backpressureLister) Backpressures(namespace string) BackpressureNamespaceLister {
	return backpressureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackpressureNamespaceLister helps list and get Backpressures.
// All objects returned here must be treated as read-only.
type BackpressureNamespaceLister interface {
	// List lists all Backpressures in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Backpressure, err error)
	// Get retrieves the Backpressure from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Backpressure, error)
	BackpressureNamespaceListerExpansion
}

// backpressureNamespaceLister implements the BackpressureNamespaceLister
// interface.
type backpressureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Backpressures in the indexer for a given namespace.
func (s backpressureNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Backpressure, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Backpressure))
	})
	return ret, err
}

// Get retrieves the Backpressure from the indexer for a given namespace and name.
func (s backpressureNamespaceLister) Get(name string) (*v1alpha1.Backpressure, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backpressure"), name)
	}
	return obj.(*v1alpha1.Backpressure), nil
}

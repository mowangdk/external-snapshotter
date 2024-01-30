/*
Copyright 2024 The Kubernetes Authors.

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
	v1alpha1 "github.com/kubernetes-csi/external-snapshotter/client/v7/apis/volumegroupsnapshot/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeGroupSnapshotLister helps list VolumeGroupSnapshots.
// All objects returned here must be treated as read-only.
type VolumeGroupSnapshotLister interface {
	// List lists all VolumeGroupSnapshots in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeGroupSnapshot, err error)
	// VolumeGroupSnapshots returns an object that can list and get VolumeGroupSnapshots.
	VolumeGroupSnapshots(namespace string) VolumeGroupSnapshotNamespaceLister
	VolumeGroupSnapshotListerExpansion
}

// volumeGroupSnapshotLister implements the VolumeGroupSnapshotLister interface.
type volumeGroupSnapshotLister struct {
	indexer cache.Indexer
}

// NewVolumeGroupSnapshotLister returns a new VolumeGroupSnapshotLister.
func NewVolumeGroupSnapshotLister(indexer cache.Indexer) VolumeGroupSnapshotLister {
	return &volumeGroupSnapshotLister{indexer: indexer}
}

// List lists all VolumeGroupSnapshots in the indexer.
func (s *volumeGroupSnapshotLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeGroupSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeGroupSnapshot))
	})
	return ret, err
}

// VolumeGroupSnapshots returns an object that can list and get VolumeGroupSnapshots.
func (s *volumeGroupSnapshotLister) VolumeGroupSnapshots(namespace string) VolumeGroupSnapshotNamespaceLister {
	return volumeGroupSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeGroupSnapshotNamespaceLister helps list and get VolumeGroupSnapshots.
// All objects returned here must be treated as read-only.
type VolumeGroupSnapshotNamespaceLister interface {
	// List lists all VolumeGroupSnapshots in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeGroupSnapshot, err error)
	// Get retrieves the VolumeGroupSnapshot from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VolumeGroupSnapshot, error)
	VolumeGroupSnapshotNamespaceListerExpansion
}

// volumeGroupSnapshotNamespaceLister implements the VolumeGroupSnapshotNamespaceLister
// interface.
type volumeGroupSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeGroupSnapshots in the indexer for a given namespace.
func (s volumeGroupSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeGroupSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeGroupSnapshot))
	})
	return ret, err
}

// Get retrieves the VolumeGroupSnapshot from the indexer for a given namespace and name.
func (s volumeGroupSnapshotNamespaceLister) Get(name string) (*v1alpha1.VolumeGroupSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("volumegroupsnapshot"), name)
	}
	return obj.(*v1alpha1.VolumeGroupSnapshot), nil
}

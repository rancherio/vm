/*
Copyright 2018 Rancher Labs, Inc.

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
	v1alpha1 "github.com/rancher/vm/pkg/apis/ranchervm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineImageLister helps list MachineImages.
type MachineImageLister interface {
	// List lists all MachineImages in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MachineImage, err error)
	// Get retrieves the MachineImage from the index for a given name.
	Get(name string) (*v1alpha1.MachineImage, error)
	MachineImageListerExpansion
}

// machineImageLister implements the MachineImageLister interface.
type machineImageLister struct {
	indexer cache.Indexer
}

// NewMachineImageLister returns a new MachineImageLister.
func NewMachineImageLister(indexer cache.Indexer) MachineImageLister {
	return &machineImageLister{indexer: indexer}
}

// List lists all MachineImages in the indexer.
func (s *machineImageLister) List(selector labels.Selector) (ret []*v1alpha1.MachineImage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineImage))
	})
	return ret, err
}

// Get retrieves the MachineImage from the index for a given name.
func (s *machineImageLister) Get(name string) (*v1alpha1.MachineImage, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("machineimage"), name)
	}
	return obj.(*v1alpha1.MachineImage), nil
}
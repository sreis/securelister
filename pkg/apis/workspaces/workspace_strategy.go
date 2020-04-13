/*
Copyright 2017 The Kubernetes Authors.

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

package workspaces

import (
	"context"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog"
)

// Validate checks that an instance of Workspace is well formed
func (WorkspaceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*Workspace)
	klog.V(5).Infof("Validating fields for Workspace %s", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// +k8s:deepcopy-gen=false
type WorkspaceREST struct {
}

func (r *WorkspaceREST) New() runtime.Object {
	log.Println("NewWorkspace")
	return &Workspace{}
}

func (r *WorkspaceREST) NamespaceScoped() bool {
	log.Println("NamespaceScoped")
	return false
}

// NewList implements rest.Lister interface
func (r *WorkspaceREST) NewList() runtime.Object {
	log.Println("NewList")
	return &WorkspaceList{}
}

// List implements rest.Lister interface
func (r *WorkspaceREST) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	log.Println("List")
	return nil, fmt.Errorf("List not supported")
}

// ConvertToTable implements rest.TableConverter interface
func (r *WorkspaceREST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	log.Println("ConvertToTable")
	return nil, fmt.Errorf("ConvertToTable not supported")
}

func NewWorkspaceREST(_ generic.RESTOptionsGetter) rest.Storage {
	log.Println("NewWorkspaceREST")
	return &WorkspaceREST{}
}

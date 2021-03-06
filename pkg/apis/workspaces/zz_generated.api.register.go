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

// Code generated by apiregister-gen. DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	WorkspacesWorkspaceStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalWorkspace,
		func() runtime.Object { return &Workspace{} },     // Register versioned resource
		func() runtime.Object { return &WorkspaceList{} }, // Register versioned resource list
		NewWorkspaceREST,
	)
	InternalWorkspace = builders.NewInternalResource(
		"workspaces",
		"Workspace",
		func() runtime.Object { return &Workspace{} },
		func() runtime.Object { return &WorkspaceList{} },
	)
	InternalWorkspaceStatus = builders.NewInternalResourceStatus(
		"workspaces",
		"WorkspaceStatus",
		func() runtime.Object { return &Workspace{} },
		func() runtime.Object { return &WorkspaceList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("workspaces.sreis.pt").WithKinds(
		InternalWorkspace,
		InternalWorkspaceStatus,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
	}).AddToScheme
	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Workspace struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   WorkspaceSpec
	Status WorkspaceStatus
}

type WorkspaceSpec struct {
}

type WorkspaceStatus struct {
}

//
// Workspace Functions and Structs
//
// +k8s:deepcopy-gen=false
type WorkspaceStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type WorkspaceStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WorkspaceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Workspace
}

func (Workspace) NewStatus() interface{} {
	return WorkspaceStatus{}
}

func (pc *Workspace) GetStatus() interface{} {
	return pc.Status
}

func (pc *Workspace) SetStatus(s interface{}) {
	pc.Status = s.(WorkspaceStatus)
}

func (pc *Workspace) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Workspace) SetSpec(s interface{}) {
	pc.Spec = s.(WorkspaceSpec)
}

func (pc *Workspace) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Workspace) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Workspace) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Workspace.
// +k8s:deepcopy-gen=false
type WorkspaceRegistry interface {
	ListWorkspaces(ctx context.Context, options *internalversion.ListOptions) (*WorkspaceList, error)
	GetWorkspace(ctx context.Context, id string, options *metav1.GetOptions) (*Workspace, error)
	CreateWorkspace(ctx context.Context, id *Workspace) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, id *Workspace) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewWorkspaceRegistry(sp builders.StandardStorageProvider) WorkspaceRegistry {
	return &storageWorkspace{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageWorkspace struct {
	builders.StandardStorageProvider
}

func (s *storageWorkspace) ListWorkspaces(ctx context.Context, options *internalversion.ListOptions) (*WorkspaceList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*WorkspaceList), err
}

func (s *storageWorkspace) GetWorkspace(ctx context.Context, id string, options *metav1.GetOptions) (*Workspace, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Workspace), nil
}

func (s *storageWorkspace) CreateWorkspace(ctx context.Context, object *Workspace) (*Workspace, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Workspace), nil
}

func (s *storageWorkspace) UpdateWorkspace(ctx context.Context, object *Workspace) (*Workspace, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Workspace), nil
}

func (s *storageWorkspace) DeleteWorkspace(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}


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


package workspaceadmission

import (
	"context"
	aggregatedadmission "github.com/sreis/securelister/plugin/admission"
	aggregatedinformerfactory "github.com/sreis/securelister/pkg/client/informers_generated/externalversions"
	aggregatedclientset "github.com/sreis/securelister/pkg/client/clientset_generated/clientset"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/admission"
)

var _ admission.Interface 											= &workspacePlugin{}
var _ admission.MutationInterface 									= &workspacePlugin{}
var _ admission.ValidationInterface 								= &workspacePlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &workspacePlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &workspacePlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &workspacePlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &workspacePlugin{}

func NewWorkspacePlugin() *workspacePlugin {
	return &workspacePlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type workspacePlugin struct {
	*admission.Handler
}

func (p *workspacePlugin) ValidateInitialization() error {
	return nil
}

func (p *workspacePlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *workspacePlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *workspacePlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *workspacePlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *workspacePlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *workspacePlugin) SetExternalKubeClientSet(kubernetes.Interface) {}

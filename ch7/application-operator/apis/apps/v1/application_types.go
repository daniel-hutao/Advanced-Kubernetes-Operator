/*
Copyright 2022 Daniel.Hu.

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

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	Deployment DeploymentTemplate `json:"deployment,omitempty"`
	Service    ServiceTemplate    `json:"service,omitempty"`
}

type DeploymentTemplate struct {
	appsv1.DeploymentSpec `json:",inline"`
}

type ServiceTemplate struct {
	corev1.ServiceSpec `json:",inline"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	Workflow appsv1.DeploymentStatus `json:"workflow"`
	Network  corev1.ServiceStatus    `json:"network"`
}

//+kubebuilder:subresource:status
//+kubebuilder:resource:path=applications,singular=application,scope=Namespaced,shortName=app
//+kubebuilder:storageversion
//+kubebuilder:object:root=true
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}

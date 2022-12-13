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

package v2

import (
	v12 "github.com/daniel-hutao/application-operator/apis/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Application to the Hub version (v1).
func (src *Application) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v12.Application)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Deployment = src.Spec.Workflow
	dst.Spec.Service = src.Spec.Service

	dst.Status.Workflow = src.Status.Workflow
	dst.Status.Network = src.Status.Network

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Application) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v12.Application)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Workflow = src.Spec.Deployment
	dst.Spec.Service = src.Spec.Service

	dst.Status.Workflow = src.Status.Workflow
	dst.Status.Network = src.Status.Network

	return nil
}

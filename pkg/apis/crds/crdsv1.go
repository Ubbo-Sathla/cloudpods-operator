// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crds

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "yunion.io/x/onecloud-operator/pkg/apis/onecloud/v1alpha1"
)

var (
	OnecloudClusterCRDv1 = &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: api.OnecloudClusterCRDName,
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group:    api.SchemeGroupVersion.Group,
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{},

			Scope: apiextensionsv1.NamespaceScoped,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Plural:     api.OnecloudClusterResourcePlural,
				Kind:       api.OnecloudClusterResourceKind,
				ShortNames: []string{"onecloud", "oc"},
			},
		},
	}
)

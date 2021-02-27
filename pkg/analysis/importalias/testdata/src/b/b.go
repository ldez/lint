// Copyright Project Contour Authors
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

package b

import (
	envoy_config "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/ext_authz/v2" // want `version "v2" not specified in alias "envoy_config" for import path "github/com/envoyproxy/gocontrolplane/envoy/config/filter/http/ext/authz/v2"`
	contour "github.com/projectcontour/contour/apis/projectcontour/v1"                          // want `"v1" not specified in alias "contour" for import path "github/com/projectcontour/contour/apis/projectcontour/v1"`
	contour_api_v1alpha1 "github.com/projectcontour/contour/apis/projectcontour/v1alpha1"
	v2 "gopkg.in/alecthomas/kingpin.v2"         // want `alias "v2" uses words that are not in path "in/alecthomas/kingpin/v2"`
	meta "k8s.io/apimachinery/pkg/apis/meta/v1" // want `version "v1" not specified in alias "meta" for import path "k8s/io/apimachinery/pkg/apis/meta/v1"`
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"   // want `alias "v1" uses words that are not in path "io/apimachinery/pkg/apis/meta/v1"`
	gateway_v1alpha1 "sigs.k8s.io/gateway-api/apis/v1alpha1"
)

func foo() {
	meta.Now()
	v1.Now()
	_ = envoy_config.AuthorizationRequest{}
	contour.AddKnownTypes(nil)
	_ = contour_api_v1alpha1.GroupVersion
	v2.Parse()
	_ = gateway_v1alpha1.GroupVersion
}

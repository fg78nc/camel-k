/*
Copyright 2018 The Knative Authors
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

package handler

import (
	"net/http"

	"github.com/knative/serving/pkg/activator"
)

// FilteringHandler will filter requests sent by the
// activator itself.
type FilteringHandler struct {
	NextHandler http.Handler
}

func (h *FilteringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If this header is set the request was sent by the activator itself, thus
	// we immediatly return a 503 to trigger a retry.
	if r.Header.Get(activator.RequestCountHTTPHeader) != "" {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	h.NextHandler.ServeHTTP(w, r)
}

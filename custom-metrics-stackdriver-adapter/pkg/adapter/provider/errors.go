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

package provider

import (
	"fmt"
	"net/http"

	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewOperationNotSupportedError returns a StatusError indicating that the invoked API call is not
// supported.
func NewOperationNotSupportedError(operation string) *apierr.StatusError {
	return &apierr.StatusError{metav1.Status{
		Status:  metav1.StatusFailure,
		Code:    int32(http.StatusNotImplemented),
		Reason:  metav1.StatusReasonBadRequest,
		Message: fmt.Sprintf("Operation: %q is not implemented", operation),
	}}
}

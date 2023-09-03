/*
Copyright 2023 yqc.

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

package v1alpha1

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var serviceingresscontrollerlog = logf.Log.WithName("serviceingresscontroller-resource")

func (r *ServiceIngressController) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-yqc-tech-github-com-v1alpha1-serviceingresscontroller,mutating=false,failurePolicy=fail,sideEffects=None,groups=yqc.tech.github.com,resources=serviceingresscontrollers,verbs=create;update,versions=v1alpha1,name=vserviceingresscontroller.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ServiceIngressController{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceIngressController) ValidateCreate() (admission.Warnings, error) {
	serviceingresscontrollerlog.Info("validate create", "name", r.Name)
	if r.Spec.EnableIngress && !r.Spec.EnableService {
		return admission.Warnings{"service switch should be true"}, errors.New("service switch should be true")
	}
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceIngressController) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	serviceingresscontrollerlog.Info("validate update", "name", r.Name)
	if r.Spec.EnableIngress && !r.Spec.EnableService {
		return admission.Warnings{"service switch should be true"}, errors.New("service switch should be true")
	}
	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceIngressController) ValidateDelete() (admission.Warnings, error) {
	serviceingresscontrollerlog.Info("validate delete", "name", r.Name)
	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}

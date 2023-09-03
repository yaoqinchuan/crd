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

package controller

import (
	"context"
	"github.com/internal/controller/utils"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	yqctechv1alpha1 "github.com/api/v1alpha1"
)

// ServiceIngressControllerReconciler reconciles a ServiceIngressController object
type ServiceIngressControllerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ingress.baiding.tech,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=yqc.tech.github.com,resources=serviceingresscontrollers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=yqc.tech.github.com,resources=serviceingresscontrollers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=yqc.tech.github.com,resources=serviceingresscontrollers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ServiceIngressController object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *ServiceIngressControllerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	serviceIngress := &yqctechv1alpha1.ServiceIngressController{}

	err := r.Get(ctx, req.NamespacedName, serviceIngress)
	if nil != err {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 更新deployment
	deployment := utils.NewDeployment(serviceIngress)
	err = controllerutil.SetControllerReference(serviceIngress, deployment, r.Scheme)
	if nil != err {
		return ctrl.Result{}, err
	}
	err = r.Get(ctx, req.NamespacedName, deployment)
	if nil != err && errors.IsNotFound(err) {
		err = r.Create(ctx, deployment)
		logger.Info("start to create deployment")
		if nil != err {
			logger.Error(err, "create deployment.yml failed")
			return ctrl.Result{}, err
		}
		logger.Info("finish to create deployment")
	} else {
		err = r.Update(ctx, deployment)
		logger.Info("start to update deployment")
		if nil != err {
			logger.Error(err, "update deployment failed")
		}
		logger.Info("finish to update deployment")
	}

	service := utils.NewService(serviceIngress)
	err = controllerutil.SetControllerReference(serviceIngress, service, r.Scheme)
	if nil != err {
		return ctrl.Result{}, err
	}
	if err = r.Get(ctx, types.NamespacedName{Name: serviceIngress.Name, Namespace: serviceIngress.Namespace}, service); err != nil {
		if errors.IsNotFound(err) && serviceIngress.Spec.EnableService {
			err = r.Create(ctx, service)
			if nil != err {
				logger.Error(err, "create service.yml failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if serviceIngress.Spec.EnableService {
			err = r.Update(ctx, service)
			if nil != err {
				logger.Error(err, "update service failed")
			}
		} else {
			err = r.Delete(ctx, service)
			if nil != err {
				logger.Error(err, "delete service failed")
			}
		}
	}

	ingress := utils.NewIngress(serviceIngress)
	err = controllerutil.SetControllerReference(serviceIngress, ingress, r.Scheme)
	if nil != err {
		return ctrl.Result{}, err
	}
	if err = r.Get(ctx, types.NamespacedName{Name: serviceIngress.Name, Namespace: serviceIngress.Namespace}, ingress); err != nil {
		if errors.IsNotFound(err) && serviceIngress.Spec.EnableIngress {
			err = r.Create(ctx, ingress)
			if nil != err {
				logger.Error(err, "create ingress.yml failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if serviceIngress.Spec.EnableIngress {
			err = r.Update(ctx, ingress)
			if nil != err {
				logger.Error(err, "update ingress failed")
			}
		} else {
			err = r.Delete(ctx, ingress)
			if nil != err {
				logger.Error(err, "delete ingress failed")
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceIngressControllerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&yqctechv1alpha1.ServiceIngressController{}).
		Owns(&v1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&netv1.Ingress{}).
		Complete(r)
}

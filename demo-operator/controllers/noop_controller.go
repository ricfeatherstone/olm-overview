/*
Copyright 2022.

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

package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1beta1 "olm-overview/api/v1beta1"
)

// NoopReconciler reconciles a Noop object
type NoopReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.ricfeatherstone.com,namespace=demo-operator-system,resources=noops,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.ricfeatherstone.com,namespace=demo-operator-system,resources=noops/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.ricfeatherstone.com,namespace=demo-operator-system,resources=noops/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Noop object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *NoopReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	noop := &demov1beta1.Noop{}
	if err := r.Get(ctx, req.NamespacedName, noop); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Failed to get noop, ignoring as must have been deleted.")
			return ctrl.Result{}, nil
		}

		logger.Error(err, "Failed to get noop.")
		return ctrl.Result{}, err
	}

	logger.Info("Reconciling Foo.", "Foo", noop.Spec.Foo, "Bar", noop.Spec.Bar)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NoopReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1beta1.Noop{}).
		Complete(r)
}

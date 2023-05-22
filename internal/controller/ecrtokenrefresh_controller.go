/*
Copyright 2023 @apanasiuk-el edenlabllc.
*/

package controller

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	ecrv1alpha1 "ecr-token-refresh.operators.infra/api/v1alpha1"
	"ecr-token-refresh.operators.infra/internal/secret"
)

// ECRTokenRefreshReconciler reconciles a ECRTokenRefresh object
type ECRTokenRefreshReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	SecretCreator secret.Creator
}

//+kubebuilder:rbac:groups=ecr.aws.edenlab.io,resources=ecrtokenrefreshes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ecr.aws.edenlab.io,resources=ecrtokenrefreshes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ecr.aws.edenlab.io,resources=ecrtokenrefreshes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ECRTokenRefresh object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *ECRTokenRefreshReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := log.FromContext(ctx)
	ecrTokenRefresh := &ecrv1alpha1.ECRTokenRefresh{}

	if err := r.Client.Get(ctx, req.NamespacedName, ecrTokenRefresh); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Error(nil, fmt.Sprintf("Can not find CRD by name: %s", req.Name))
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	// Create new Secret definition
	reqLogger.Info(fmt.Sprintf("Get token for ECR registry: %s", ecrTokenRefresh.Spec.ECRRegistry))
	newSecret, err := r.SecretCreator.CreateSecret(&secret.InputFromCRD{CRD: ecrTokenRefresh})
	if err != nil {
		reqLogger.Error(err, fmt.Sprintf("Can not get token for ECR registry: %s", ecrTokenRefresh.Spec.ECRRegistry))
		ecrTokenRefresh.Status.Phase = "Error"
		ecrTokenRefresh.Status.Error = err.Error()
		if err := r.Status().Update(ctx, ecrTokenRefresh); err != nil {
			reqLogger.Error(err, fmt.Sprintf("Unable to update status for CRD: %s", req.Name))
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, nil
	}

	// Used to ensure that the secret will be deleted when the custom resource object is removed
	if err := ctrl.SetControllerReference(ecrTokenRefresh, newSecret, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	defSecret := &v1.Secret{}
	// Create a new secret
	if err = r.Client.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, defSecret); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info(fmt.Sprintf("Create new secret: %s for namespace: %s", req.Name, req.Namespace))
			if err = r.Client.Create(ctx, newSecret); err != nil {
				return ctrl.Result{}, err
			}

			ecrTokenRefresh.Status.Phase = "Created"
		} else {
			return ctrl.Result{}, err
		}
	} else {
		reqLogger.Info(fmt.Sprintf("Update secret: %s for namespace: %s", req.Name, req.Namespace))
		defSecret.Data = newSecret.Data
		if err = r.Client.Update(ctx, defSecret); err != nil {
			return ctrl.Result{}, err
		}

		ecrTokenRefresh.Status.Phase = "Updated"
		ecrTokenRefresh.Status.LastUpdatedTime = &metav1.Time{Time: time.Now()}
	}

	// tokenRefresh.Status.Conditions.
	if err := r.Status().Update(ctx, ecrTokenRefresh); err != nil {
		reqLogger.Error(err, fmt.Sprintf("Unable to update status for CRD: %s", req.Name))
		return ctrl.Result{}, nil
	} else {
		reqLogger.Info(fmt.Sprintf("Update status for CRD: %s", req.Name))
	}

	return ctrl.Result{RequeueAfter: ecrTokenRefresh.Spec.Frequency.Duration}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ECRTokenRefreshReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ecrv1alpha1.ECRTokenRefresh{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}

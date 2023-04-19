package dnsmasq

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mcv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlfake "sigs.k8s.io/controller-runtime/pkg/client/fake"

	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	mock_dynamichelper "github.com/Azure/ARO-RP/pkg/util/mocks/dynamichelper"
)

func TestMachineConfigReconciler(t *testing.T) {
	fakeDh := func(controller *gomock.Controller) *mock_dynamichelper.MockInterface {
		return mock_dynamichelper.NewMockInterface(controller)
	}
	cluster := func(enabled bool) *arov1alpha1.Cluster {
		return &arov1alpha1.Cluster{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
			Status:     arov1alpha1.ClusterStatus{},
			Spec: arov1alpha1.ClusterSpec{
				OperatorFlags: arov1alpha1.OperatorFlags{
					controllerEnabled: strconv.FormatBool(enabled),
				},
			},
		}
	}

	t.Run("when no cluster resource is present, returns error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		client := ctrlfake.NewClientBuilder().Build()
		dh := fakeDh(controller)

		r := NewMachineConfigReconciler(
			logrus.NewEntry(logrus.StandardLogger()),
			client,
			dh,
		)

		request := ctrl.Request{}

		_, err := r.Reconcile(context.Background(), request)

		if !kerrors.IsNotFound(err) {
			t.Errorf("wanted error: cluster not found, got error: %v", err)
		}
	})

	t.Run("when controller is disabled, returns with no error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		client := ctrlfake.NewClientBuilder().WithObjects(cluster(false)).Build()
		dh := fakeDh(controller)

		r := NewMachineConfigReconciler(
			logrus.NewEntry(logrus.StandardLogger()),
			client,
			dh,
		)

		_, err := r.Reconcile(context.Background(), ctrl.Request{})

		if err != nil {
			t.Errorf("wanted no error, got error: %v", err)
		}
	})

	t.Run("when no corresponding MachineConfigPool for the requested MachineConfig is present, does nothing", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		client := ctrlfake.NewClientBuilder().WithObjects(cluster(true)).Build()
		dh := fakeDh(controller)

		request := ctrl.Request{
			NamespacedName: types.NamespacedName{
				Namespace: "",
				Name:      "99-custom-aro-dns",
			},
		}

		r := NewMachineConfigReconciler(
			logrus.NewEntry(logrus.StandardLogger()),
			client,
			dh,
		)

		_, err := r.Reconcile(context.Background(), request)

		if err != nil {
			t.Errorf("wanted no error, got error: %v", err)
		}
	})

	t.Run("when a valid MachineConfigPool exists for the requested MachineConfig, reconciles MachineConifg", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		client := ctrlfake.NewClientBuilder().
			WithObjects(
				cluster(true),
				&mcv1.MachineConfigPool{
					ObjectMeta: metav1.ObjectMeta{Name: "custom"},
					Status:     mcv1.MachineConfigPoolStatus{},
					Spec:       mcv1.MachineConfigPoolSpec{},
				},
			).
			Build()
		dh := fakeDh(controller)
		dh.EXPECT().Ensure(gomock.Any(), gomock.AssignableToTypeOf(&mcv1.MachineConfig{})).Times(1)

		request := ctrl.Request{
			NamespacedName: types.NamespacedName{
				Namespace: "",
				Name:      "99-custom-aro-dns",
			},
		}

		r := NewMachineConfigReconciler(
			logrus.NewEntry(logrus.StandardLogger()),
			client,
			dh,
		)

		_, err := r.Reconcile(context.Background(), request)

		if err != nil {
			t.Errorf("wanted no error, got error: %v", err)
		}
	})
}

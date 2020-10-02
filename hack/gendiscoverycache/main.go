package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	configv1 "github.com/openshift/api/config/v1"
	configclient "github.com/openshift/client-go/config/clientset/versioned"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	utillog "github.com/Azure/ARO-RP/pkg/util/log"
)

const discoveryCacheDir = "pkg/util/dynamichelper/discovery/cache"

func run(ctx context.Context, log *logrus.Entry) error {
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return err
	}

	err = os.RemoveAll(discoveryCacheDir)
	if err != nil {
		return err
	}

	err = genDiscoveryCache(restconfig)
	if err != nil {
		return err
	}

	err = genRBAC(restconfig)
	if err != nil {
		return err
	}

	return writeVersion(ctx, restconfig)
}

func writeVersion(ctx context.Context, restconfig *rest.Config) error {
	configcli, err := configclient.NewForConfig(restconfig)
	if err != nil {
		return err
	}

	clusterVersion, err := getClusterVersion(ctx, configcli)
	if err != nil {
		return err
	}

	versionPath := filepath.Join(discoveryCacheDir, "assets_version")
	return ioutil.WriteFile(versionPath, []byte(clusterVersion+"\n"), 0666)
}

func getClusterVersion(ctx context.Context, configcli configclient.Interface) (string, error) {
	cv, err := configcli.ConfigV1().ClusterVersions().Get(ctx, "version", metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	for _, history := range cv.Status.History {
		if history.State == configv1.CompletedUpdate {
			return history.Version, nil
		}
	}

	// Should never happen as a successfully created cluster
	// should have at least one completed update.
	return "", errors.New("could find actual cluster version")
}

func main() {
	log := utillog.GetLogger()

	if err := run(context.Background(), log); err != nil {
		log.Fatal(err)
	}
}

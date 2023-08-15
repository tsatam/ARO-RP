package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"os"

	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	failure "github.com/Azure/ARO-RP/pkg/hive/failure"
	utillog "github.com/Azure/ARO-RP/pkg/util/log"
)

const (
	hiveNamespaceName  = "hive"
	configMapName      = "additional-install-log-regexes"
	configMapPath      = "hack/hive-config/hive-additional-install-log-regexes.yaml"
	regexDataEntryName = "regexes"
)

type installLogRegex struct {
	Name                  string   `json:"name"`
	SearchRegexStrings    []string `json:"searchRegexStrings"`
	InstallFailingReason  string   `json:"installFailingReason"`
	InstallFailingMessage string   `json:"installFailingMessage"`
}

func run(ctx context.Context) error {
	ilrs := []installLogRegex{}

	for _, reason := range failure.Reasons {
		ilrs = append(ilrs, failureReasonToInstallLogRegex(reason))
	}

	ilrsRaw, err := yaml.Marshal(ilrs)
	if err != nil {
		return err
	}

	configmap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: corev1.SchemeGroupVersion.String(),
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: hiveNamespaceName,
			Name:      configMapName,
		},
		Data: map[string]string{
			regexDataEntryName: string(ilrsRaw),
		},
	}

	configmapRaw, err := yaml.Marshal(configmap)
	if err != nil {
		return err
	}
	return os.WriteFile(configMapPath, configmapRaw, 0666)
}

func failureReasonToInstallLogRegex(reason failure.InstallFailingReason) installLogRegex {
	ilr := installLogRegex{
		Name:                  reason.Name,
		InstallFailingReason:  reason.Reason,
		InstallFailingMessage: reason.Message,
		SearchRegexStrings:    []string{},
	}
	for _, regex := range reason.SearchRegexes {
		ilr.SearchRegexStrings = append(ilr.SearchRegexStrings, regex.String())
	}
	return ilr
}

func main() {
	log := utillog.GetLogger()

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

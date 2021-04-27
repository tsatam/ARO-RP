# Azure Red Hat OpenShift Resource Provider

## Welcome!

For information relating to the generally available Azure Red Hat OpenShift v4
service, please see the following links:

* https://azure.microsoft.com/en-us/services/openshift/
* https://www.openshift.com/products/azure-openshift
* https://docs.microsoft.com/en-us/azure/openshift/
* https://docs.openshift.com/aro/4/welcome/index.html


## Quickstarts

* If you are an end user and want to create an Azure Red Hat OpenShift 4
  cluster, follow [Create, access, and manage an Azure Red Hat OpenShift 4
  Cluster][1].

* If you want to deploy a development RP, follow [deploy development
  RP](docs/deploy-development-rp.md).

[1]: https://docs.microsoft.com/en-us/azure/openshift/howto-using-azure-redhat-openshift

## Contributing

This project welcomes contributions and suggestions. Most contributions require
you to agree to a Contributor License Agreement (CLA) declaring that you have
the right to, and actually do, grant us the rights to use your contribution. For
details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether
you need to provide a CLA and decorate the PR appropriately (e.g., label,
comment). Simply follow the instructions provided by the bot. You will only need
to do this once across all repositories using our CLA.

This project has adopted the [Microsoft Open Source Code of
Conduct](https://opensource.microsoft.com/codeofconduct/). For more information
see the [Code of Conduct
FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact
[opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional
questions or comments.


## Repository map

* .pipelines: CI workflows using Azure pipelines.

* cmd/aro: RP entrypoint.

* deploy: ARM templates to deploy RP in development and production.

* docs: Documentation.

* hack: Build scripts and utilities.

* pkg: RP source code:

  * pkg/api: RP internal and external API definitions.

  * pkg/backend: RP backend workers.

  * pkg/bootstraplogging: Bootstrap logging configuration

  * pkg/client: Autogenerated ARO service Go client.

  * pkg/cluster: Cluster create/update/delete operations wrapper for OCP installer.

  * pkg/database: RP CosmosDB wrapper layer.

  * pkg/deploy: /deploy ARM template generation code.

  * pkg/env: RP environment-specific shims for running in production,
    development or test

  * pkg/frontend: RP frontend webserver.

  * pkg/metrics: Handles RP metrics via statsd.

  * pkg/mirror: OpenShift release mirror tooling.

  * pkg/monitor: Monitors running clusters.

  * pkg/operator/controllers: A list of controllers instantiated by the operator
    component.

    * alertwebhook: Ensures that the receiver endpoint defined in the
      `alertmanager-main` secret matches the webserver endpoint at
      aro-operator-master.openshift-azure-operator:8080, to avoid the
      `AlertmanagerReceiversNotConfigured` warning.

    * checker: Watches the `Cluster` resource for changes and updates conditions
      of the resource based on checks mentioned below

      * internetchecker: validate outbound internet connectivity to the nodes

      * machinechecker: validate machine objects have the correct provider spec,
        vm type, vm image, disk size, three master nodes exist, and the number of worker nodes
        match the desired worker replicas

      * serviceprincipalchecker: validate cluster service principal has the
        correct role/permissions

    * dnsmasq: Ensures that a dnsmasq systemd service is defined as a machineconfig for all
      nodes to allow for api-int and *.apps domains resolve even if custom DNS on the VNET is set.

    * genevalogging: Ensures all the Geneva logging resources in the
      `openshift-azure-logging` namespace matches the pre-defined specification
      found in `pkg/operator/controllers/genevalogging/genevalogging.go`.

    * monitoring: Ensures that the OpenShift monitoring configuration in the `openshift-monitoring` namespace is consistent and immutable.

    * node: Force deletes pods when a node fails to drain for 1 hour.  It should clear up any pods that refuse to be evicted on a drain due to violating a pod disruption budget.

    * pullsecret: Ensures that the ACR credentials in the
      `openshift-config/pull-secret` secret match those in the
      `openshift/azure-operator/cluster` secret.

    * rbac: Ensures that the `aro-sre` clusterrole and clusterrolebinding exist and are consistent.

    * routefix: Ensures all the routefix resources in the namespace
      `openshift-azure-routefix` remain on the cluster.

    * workaround: Applies a set of temporay workarounds to the ARO cluster.

  * pkg/portal: Portal for running promql queries against a cluster or requesting a kubeconfig for a cluster.

  * pkg/proxy: Proxy service for portal kubeconfig cluster access.

  * pkg/swagger: Swagger specification generation code.

  * pkg/util: Utility libraries.

* python: Autogenerated ARO service Python client and `az aro` client extension.

* swagger: Autogenerated ARO service Swagger specification.

* test: End-to-end tests.

* vendor: Vendored Go libraries.


## Basic architecture

* pkg/frontend is intended to become a spec-compliant RP web server.  It is
  backed by CosmosDB.  Incoming PUT/DELETE requests are written to the database
  with an non-terminal (Updating/Deleting) provisioningState.

* pkg/backend reads documents with non-terminal provisioningStates,
  asynchronously updates them and finally updates document with a terminal
  provisioningState (Succeeded/Failed).  The backend updates the document with a
  heartbeat - if this fails, the document will be picked up by a different
  worker.

* As CosmosDB does not support document patch, care is taken to correctly pass
  through any fields in the internal model which the reader is unaware of (see
  `github.com/ugorji/go/codec.MissingFielder`).  This is intended to help in
  upgrade cases and (in the future) with multiple microservices reading from the
  database in parallel.

* Care is taken to correctly use optimistic concurrency to avoid document
  corruption through concurrent writes (see `RetryOnPreconditionFailed`).

* The pkg/api architecture differs somewhat from
  `github.com/openshift/openshift-azure`: the intention is to fix the broken
  merge semantics and try pushing validation into the versioned APIs to improve
  error reporting.

* Everything is intended to be crash/restart/upgrade-safe, horizontally
  scaleable, upgradeable...


## Useful links

* https://github.com/Azure/azure-resource-manager-rpc

* https://github.com/microsoft/api-guidelines

* https://docs.microsoft.com/en-gb/rest/api/cosmos-db

* https://github.com/jim-minter/go-cosmosdb

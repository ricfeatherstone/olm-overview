# Operator Framework Components

![](images/operator-framework-components.png)

## Custom Resource Definition (CRD)

Definition of an Object that extends the Kubernetes API and is reconciled by the Operator (the controller).

## Cluster Service Version (CSV)

Represents a single version of an Operator and defines the template that OLM uses to deploy the Operator.

Defines

* Deployment Template
* Permissions
  * cluster scoped
  * namespace scoped
* Capability level
  * Basic Install
  * Seamless Upgrades
  * Full Lifecycle
  * Deep Insights
  * Auto Pilot
* Supported Install Modes
  * Own Namespace
  * Single Namespace
  * Multi Namespace
  * All Namespaces
* Additional Metadata

## Operator Bundle Image

Packaging format for the associated manifests and metadata for a single version of an Operator.

* API definitions (CRDs)
* Cluster Service Version (CSV)
* Any additional manifests that are not part of the CSV
  * e.g. operator configuration configmap, metrics service

> :warning: When an Operator is installed, the bundle manifest are stored in a ConfigMap in the OLM namespace.
Kubernetes has a restriction of ~1MB for ConfigMaps. From OLM >= v0.19.0 the manifests are compressed, giving you
around ~4MB capacity for the uncompressed manifests. Only likely to affect large Operators, or Operators with many
CRD versions. 

## Operator Catalog Image

Also referred to as an Operator Index Image is a sqllite database containing the bundle images for every Operator in the
Catalog. It's served to OLM over a gRPC API. Is currently being migrating to a text based model.

## Catalog Source

Makes a catalog available so that the contained Operators can be installed. Polls for updates on a configurable 
interval.

When a CatalogSource is created, the ConfigMap for the Bundle is created in the OLM Namespace.

## Operator Group

Provides multi-tenant configuration capabilities allowing a Cluster Admin to control which Namespaces an Operator
will watch for CustomResources. Operators can only be installed in a Namespace if it's InstallMode is compatible
with the Namespaces configured in the OperatorGroup. There can only be one OperatorGroup in a specific Namespace.

It also allows the ServiceAccount that will be used to deploy the Operator to be specified, allowing Cluster Admins
to control the granted permissions and therefore what is installed i.e. installation will fail if the configured
ServiceAccount does not have permission to create a resource from the Bundle.

When the OperatorGroup is created the view, edit and admin ClusterRoles for the OperatorGroup are created along
with the view, edit, admin and crdview ClusterRoles for the CRDs and the Metrics Reader ClusterRole, see 
[OLM RBAC](olm-rbac.md).

## Subscription

When the Subscription is created the permissions in the CSV are transformed into Roles in the deployed and all 
watched namespaces and the CSV is copied into all these namespaces. The Namespace details (deployed and watched)
are removed from the CSV in the watched Namespaces.

Deleting a Subscription does not remove the Operator.

A Subscription also allows the configuration of specific fields on the deployment, including but not limited to
resource requests and limits and node selectors providing a way to customise specific deployments of the Operator.

## Install Plan

InstallPlans are created when the Subcription is created and when new versions of the Operator are available from
the Catalog (Catalogs usually reference the `latest` tag)

## Operator Development

By default, the Operator-SDK scaffolds a cluster-scoped operator. This is fine if AllNamespaces is the only 
InstallMade that is supported. 

In order to support all InstallModes, the Operator should be 
[Namespace Scoped](https://sdk.operatorframework.io/docs/building-operators/golang/operator-scope/) and should
dynamically set the Namespace(s) to watch, see [Restricting Watched Namespaces](restricting-watched-namespaces.md).
This will ensure RBAC permission are generated as a Role and not a ClusterRole and OLM will manage creating the 
Roles and RoleBindings based on the OperatorGroup's watched Namespaces.

# OLM RBAC

## Cluster Roles from CustomResources

* view
  * get, list, watch on CustomResources
  * e.g. `noops.demo.ricfeatherstone.com-v1beta1-view`
* edit
  * create, update, patch, delete on CustomResources
  * e.g. `noops.demo.ricfeatherstone.com-v1beta1-edit`
* admin
  * '*' on CustomResources
  * e.g. `noops.demo.ricfeatherstone.com-v1beta1-admin`
* crdview
  * get - restricted to named CustomResourceDefinitions
  * e.g. `noops.demo.ricfeatherstone.com-v1beta1-crdview`

## Cluster Roles from OperatorGroup

* _operatorgroup-name_-view
  * get, list, watch on CustomResources
  * get CustomResourceDefinitions
  * e.g. `demo-operator-system-view`
* _operatorgroup-name_-edit
  * create, update, patch, delete on CustomResources
  * e.g. `demo-operator-system-edit`
* _operatorgroup-name_-admin
  * '*' on CustomResources
  * e.g. `demo-operator-system-admin`

## Cluster Roles from Bundle

* _operator-name_._version_-_generated_
  * from clusterpermissions in csv
  * e.g. `demo-operator.v0.0.3-65d8ddc75d`
* _operator-name_-metrics-reader
  * Metrics reader permissions from bundle
  * e.g. `demo-operator-metrics-reader`

## New Role in OLM Namespace

* _generated-name_
  * create, get, update on configmap of same generated name which is the exploded bundle
  * e.g. `3169c1b59ea25a0dfb61f226ea6ee8c831c16f44269f5b13dd1c40450353cba`

## New Roles in Operator Namespace

* _operator-name_._version_
  * get, update, patch operatorconditions for deployed operator
  * e.g. `demo-operator.v0.0.3`
* _operator-name_._version_-_operator-name_-controller-manager-_generated_
  * from permissions in csv
  * also created in all watched namespaces
  * e.g. `demo-operator.v0.0.3-demo-operator-controller-manager-6fbcc96c5`

NAME ?= olm-overview
NAMESPACE ?= demo-operator-system

##@ General

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Kind

cluster-create: ## Create Cluster
	kind create cluster --name $(NAME)

cluster-delete: ## Delete Cluster
	kind delete cluster --name $(NAME)

##@ OLM

olm-install: ## Install OLM
	operator-sdk olm install

olm-uninstall: ## Uninstall OLM
	operator-sdk olm uninstall

olm-status: ## Display OLM Status
	operator-sdk olm status

olm-list-operators: ## List Available Operators
	kubectl -n olm get packagemanifest

olm-show-packagemanifest: ## Show PackageManifest for OPERATOR=xxx
	kubectl -n olm get packagemanifest $(OPERATOR) -oyaml

olm-dump: ## Get Manifests for OLM installed CatalogSource and OperatorGroups
	kubectl -n olm get catalogsource operatorhubio-catalog -oyaml > olm/catalog.yaml
	kubectl -n olm get operatorgroup olm-operators -oyaml > olm/olm-operators.yaml
	kubectl -n operators get operatorgroup global-operators -oyaml > olm/global-operators.yaml

##@ Demo Operator

operator-catalog-create: ## Create CatalogSource
	kubectl apply -f manifests/catalog-source.yaml

operator-catalog-delete: ## Delete CatalogSource
	kubectl delete -f manifests/catalog-source.yaml

operator-create-deployment-namespace: ## Create Namespace to deploy Operator into
	kubectl create ns $(NAMESPACE)

operator-create-watch-own-namespace: operator-create-deployment-namespace ## Deploy Demo Operator to watch own namespace
	kubectl apply -k manifests/watch-own-namespace

operator-create-watched-namespaces: ## Create namespaces to be watched by the Demo Operator
	-for i in 1 2 3; do \
		kubectl create ns watched-namespace-$$i; \
	done

operator-create-watch-multiple-namespaces: operator-create-deployment-namespace operator-create-watched-namespaces ## Deploy the Demo Operator to watch multiple namespaces
	kubectl apply -k manifests/watch-multiple-namespaces

operator-create-watch-all-namespaces: operator-create-deployment-namespace ## Deploy the Demo Operator to watch all namespaces
	kubectl apply -k manifests/watch-all-namespaces

operator-create-installplan-manual-approval: operator-create-deployment-namespace ## Deploy the Demo Operator with Manual Approval
	kubectl apply -k manifests/installplan-manual-approval

operator-list-installplans: ## List InstallPlans
	kubectl -n $(NAMESPACE) get installplans.operators.coreos.com

operator-approve-installplan: ## Approve an InstallPlan INSTALL_PLAN=xxx
	kubectl -n $(NAMESPACE) patch installplans.operators.coreos.com \
		$(INSTALL_PLAN) --type merge --patch-file manifests/installplan-patch.yaml

operator-create-configure-resources: operator-create-deployment-namespace ## Deploy the Demo Operator with resource requests and limits modified
	kubectl apply -k manifests/configure-resources

operator-delete: ## Delete the Demo Operator and all configuration
	-kubectl delete ns $(NAMESPACE)
	-for i in 1 2 3; do \
		kubectl delete ns watched-namespace-$$i; \
	done

operator-logs: ## Display the Demo Operator logs
	kubectl -n $(NAMESPACE) logs \
		-c manager \
		$$(kubectl -n $(NAMESPACE) get po -l=control-plane=controller-manager -oname)

operator-image: ## Display the current Demo Operator Image
	kubectl -n $(NAMESPACE) get \
		$$(kubectl -n $(NAMESPACE) get po -l=control-plane=controller-manager -oname) -ojson \
		| jq .spec.containers[1].image

##@ Demo CR

cr-create: ## Create CR - NAMESPACE=xxx
	-kubectl -n $(NAMESPACE) create -f demo-operator/config/samples/demo_v1beta1_noop.yaml

cr-delete: ## Delete CR - NAMESPACE=xxx
	-kubectl -n $(NAMESPACE) delete -f demo-operator/config/samples/demo_v1beta1_noop.yaml

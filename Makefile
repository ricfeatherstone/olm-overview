NAME = olm-overview

create-cluster:
	kind create cluster --name $(NAME)

delete-cluster:
	kind delete cluster --name $(NAME)

olm-install:
	operator-sdk olm install

olm-status:
	operator-sdk olm status

olm-catalog-source:
	kubectl -n olm get catalogsource operatorhubio-catalog -oyaml > olm/catalog.yaml

olm-operator-groups:
	kubectl -n olm get operatorgroup olm-operators -oyaml > olm/olm-operators.yaml
	kubectl -n operators get operatorgroup global-operators -oyaml > olm/global-operators.yaml

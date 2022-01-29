NAME = olm-overview

create-cluster:
	kind create cluster --name $(NAME)

delete-cluster:
	kind delete cluster --name $(NAME)

olm-install:
	operator-sdk olm install

olm-status:
	operator-sdk olm status

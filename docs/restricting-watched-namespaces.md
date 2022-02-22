# Restricting Watched Namespaces

Based on the OperatorGroup configuration, OLM annotates the PodSpec in the Deployment to specify the watched namespaces
e.g. `olm.targetNamespaces: watched-namespace-1,watched-namespace-2,watched-namespace-3`. This can be retrieved in the 
code to ensure that the Operator is only watching namespaces where it has the correct permissions.

[manager.yaml](../demo-operator/config/manager/manager.yaml)

```yaml
env:
  - name: TARGET_NAMESPACES
    valueFrom:
      fieldRef:
        fieldPath: metadata.annotations['olm.targetNamespaces']
```
[main.go](../demo-operator/main.go)

```golang
targetNamespacesString := os.Getenv("TARGET_NAMESPACES")
	if targetNamespacesString == "" {
		options.Namespace = targetNamespacesString
	} else {
		targetNamespaces := strings.Split(targetNamespacesString, ",")
		options.NewCache = cache.MultiNamespacedCacheBuilder(targetNamespaces)
	}
```

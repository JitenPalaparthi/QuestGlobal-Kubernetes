### Kubernetes api resources

```kubectl api-resources```

### Explain kubernetes resource

```kubectl explain ns.spec``` 

### To create a namespace through a file

```kubectl apply -f test-namespace.yaml```


### Exec into a pod

- exec into pod if it has only one container

```kubectl exec -it  nginx bash -n test```

- exec into container when pod has more than one container

```kubectl exec -it nginx -c pingapp /bin/sh -n test```

### To know more about a pod

```kubectl describe pod nginx -n test```

### Delete a pod

```kubectl delete pods nginx -n test```
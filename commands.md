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

### To get logs

- followup logs

```kubectl logs  -f  nginx -c nginx -n test```

- tailered logs[number of logs]

```kubectl logs  --tail=10  nginx -c nginx -n test```

## Service

```kubectl expose pod nginx --type=NodePort  --name=nginx -n test```

### Scaling

- scale up the replication controller
  
```kubectl scale --replicas rc/nginx-rc -n test```

### ADD Labels to the node

```kubectl label nodes minikube-m04 colour=blue priority=high```

- add label to replicaset

```kubectl label rs nginx-rs deployment=replicas -n test```

# AWS Cli

- ```aws configure```

- ```aws eks --region us-east-1 update-kubeconfig --name demo_cluster```

- ```aws sts get-caller-identity```
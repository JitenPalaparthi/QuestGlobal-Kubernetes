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

- ```aws eks create-cluster --name mycluster --role-arn arn:aws:iam::196139220645:role/demo --region us-east-1 --resources-vpc-config subnetIds=subnet-0ff0222db0a8255ff,subnet-018efc0f5769151db,securityGroupIds=sg-05424846ba0def0c4```

```aws eks update-kubeconfig --region us-east-1 --name mycluster```

- ```eksctl create nodegroup --cluster mycluster --region us-east-1 --name mynodegroup --node-type m3.medium --nodes 3 --nodes-min 2 --nodes-max 4```
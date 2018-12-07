
* Choose Kubernetes for AWS (Bitnami) AMI from AWS marketplace.

* Use below deployment yaml file to launch your pods.

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: userlogin-deployment1
spec:
  selector:
    matchLabels:
      app: userlogin
  replicas: 2
  template:
    metadata:
      labels:
        app: userlogin
    spec:
      containers:
      - name: userlogin
        image: kingslayerr/teamproject:version2
        ports:
        - containerPort: 5000
```

* Use below service yaml file to start the NodePort service.

```
apiVersion: v1
kind: Service
metadata:
  name: userlogin-service1
spec:
  ports:
  - port: 8000
    targetPort: 5000
    protocol: TCP
  selector:
    name: userlogin
  type: NodePort
```
* NodePort: Exposes the service on each Node’s IP at a static port (the NodePort). A ClusterIP service, to which the NodePort service will route, is automatically created. You’ll be able to contact the NodePort service, from outside the cluster, by requesting <NodeIP>:<NodePort>. (taken from kubernetes documentation)

* Use below steps to get the ip and port:

```
kubectl run user-login --replicas=2 --labels="run=user-login" --image=kingslayerr/teamproject:version2 --port=5000

kubectl expose deployment user-login --type=NodePort --name=user-login-service

kubectl describe services user-login-service (Note down the port)

kubect cluster-info (IP-> Get The IP where master is running)

Your service is accessible at (IP):(port)

```

* You have a working kubernetes service in the AWS cloud. :)

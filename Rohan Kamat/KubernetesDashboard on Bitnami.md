### Kubernetes Dashboard on Bitnami

Kubernetes dashboard will already be setup on Bitnami Instance to expose port edit the yaml file by executing the following command

```sh
kubectl -n kube-system edit service kubernetes-dashboard
```
Change  ```sh type: ClusterIP to type: NodePort ```

```sh
$ kubectl -n kube-system get service kubernetes-dashboard
NAME                   CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
kubernetes-dashboard   10.100.124.90   <nodes>       443:31707/TCP   21h
```

To get port details 
```sh
sudo kubectl -n kube-system get service kubernetes-dashboard
```
Should look something like this

```sh
 NAME                   TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
kubernetes-dashboard   NodePort   10.106.190.241   <none>        80:30009/TCP   12h
```
### Note the port number
### Change the security group for the instance, by opening 30009

Hit your instanceIP:Port  
vola! Dashboard service works!!

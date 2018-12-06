# Payment GoAPI

Payment GoAPI is deployed on internet-facing ELB with 2 docker private instances connected to the private mongodb cluster, nodes of which are on different VPCs.
* Application instances are on Ohio region and is accessable through ElB public DNS: PaymentAPI-705203207.us-east-2.elb.amazonaws.com
* Mongodb cluster has 3 nodes on N.California and 2 nodes on Oregon. They're not publicly accessable
## What is VPC peering?

1) Amazon VPC (Virtual Private Cloud) enables you to launch AWS resources into a virtual network that you've defined.
2) A VPC peering connection is a networking connection between two VPCs that enables you to route traffic between them using private IPv4 addresses or IPv6 addresses.
3) Instances in either VPC can communicate with each other as if they are within the same network.
4) VPC peering can be done between:
* Two VPCs of the same account  
* Two VPCs of different accounts  
* Two VPCs of different regions (inter-region VPC peering connection)  

## To create a VPC peering connection with a VPC in a different region:

1) Open the Amazon VPC console at https://console.aws.amazon.com/vpc/.
2) In the navigation pane, choose Peering Connections, Create Peering Connection.
3) Configure the following information, and choose Create Peering Connection when you are done:
4) Peering connection name tag = CA-OR (for N. California and Oregon)
5) VPC (Requester)* = cmpe281 (N. California)
6) Account = My account
7) Region = Another Region (Oregon)
8) VPC (Accepter): ID of the accepter VPC.
9) OK 
10) In the region selector, select the region of the accepter VPC (Oregon).  
11) In the navigation pane, choose Peering Connections. 
12) Select the VPC peering connection that you've created, and choose Actions, Accept Request.  
13) Accept peering connection request. 
14) Once your VPC peering connection is active, you must add an entry to your VPC route tables to enable traffic to be directed between the peered VPCs.    
15) Edit route the route tables by adding the CIDR of the other VPC.

## Copy AMI from one region to another:

* Since we will be creating the MongoDB cluster between two regions, we will need the MongoDB installed AMI in both the regions. The steps to copy the AMI are given below.

1) Open the Amazon EC2 console.
2) In the navigation pane, choose AMIs.
3) Select the AMI to be copied.
4) Select the appropriate destination region (Oregon).
5) Copy AMI. 

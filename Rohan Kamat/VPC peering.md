### VPC Peering Solution
It is necessary for instances in different regions to communicate with each other, for that purpose there must be a connection between 2 VPC. This is taken care by VPC peering.
VPC peering can be done under 3 types of scenarios:
* Between accounts
* Between regions
* In the same region
### Steps:
1.	Go to VPC launch Wizard.
2.	Select VPC with private and public.
3.	Keep the CIDR default.
 a. Assign a VPC name.
 b. Choose NAT instance to avoid charges
 c. Select key pair 
4.	Launch the VPC.
5.	Launch 3 instance from the AMI(mongoCluster) into the VPC.
Select any other region:
The problem encountered is AMI in previous region is not visible in the new region, so go to previous region>> AMI
                 Click on desired AMI
Copy image>>Select region
6.	Now the AMI is visible in the new region.
7.	Follow the steps from 1 to 5(Ensure the new VPC CIDR blocks are not overlapping)
8.	Go to VPC peering 
a.	Select Requestor as current 
b.	Select another region
c.	Add ID of VPC in new region
d.	Launch the peer connection
9.	Go to the new region to accept the pending peer request.
Editing the route table to ensure, instances can communicate.
10.Go to original region>> VPC>>Route Tables.
      a. Edit the public subnet route table.
      b. Add the CIDR of new region VPC, set target to VPC peer.(Same procedure in the new region)
Follow from [Mongo Documentation](https://github.com/nguyensjsu/fa18-281-code-monks/blob/master/Rohan%20Kamat/Mongo%20Cluster%20setup.md) for cluster setup.


## Scrum Value: Focus:

#### Iterated product delivery:
The stand-up meetings should be conduted with the focus of discussing the hindrances faced, about what percentage of the user story is completed, and what is planned to be completed for the next day.

#### Time factor:
In a week’s time, the team should create and assign user stories, code to develop the product, and carry out the white box and black box tests.

###### Week 1:
+ Basic idea of how and what application is to be developed were discussed.
+ Various research items were distributed amongst team members.
+ Decision of developing aaplication in GO API and using MongoDB cluster as backend was made.

###### Week 2:
+ Post research performed individually by all the team members, a clearer picture of application to be developed was defined.
+ A final decision of working on a Starbucks application was made.
+ Higher lever GO APIs where found out

###### Week 3:
+ A final architecture of the application was articulated.
+ Tasks such as developing GO API (6 GO APIs, 1 for each team member), deploying API on 2 dockers, adding a load balancer to balance the load of 2 dockers and using MongoDB clutser from personal project as backend DB for each API was made 
+ Additionally, 2 tasks were identified:
1) Would integrating APIs using MongoDB clusters in different private networks (and different VPCs) into a single application cause any   problems?
2) How can we make databases that lie within different networks exchange data with each other. For example, how can we pass information of items added in a cart (Cart API) to 'Payment' API?

###### Week 5:
###### Team met and helped each other out with the setup for:
+ Dockers that'd contain GO APIs for individual API written by every member
+ VPC peering to connect machines that are setup accross AZ's and regions
+ Setting up public facing loadbalancer to connect to MongoDB cluster

+ Plan was made to meet next time with finsihed frontend and integrate the code 
  

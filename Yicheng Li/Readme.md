# Week 1
### Project Research
- [x] 1. Understanding the project requirement and architecture from various articles and books.
- [x] 2. Research on AWS lambda was done.

## Week 2
### Initial GoAPI and MongoDB Replicaset
- [x] 1. Research on MongoDB database CP properties.
- [x] 2. Configuring a simple Go api to connect to mongoDB cluster.
- [x] 3. Setup an AMI for EC2 having mongoDB installed.
- [x] 4. Configure a VPC and create 5 different EC2 instances in 2 separate AZs.
- [x] 5. Test the mongoDB replicaset by inserting data in primary node.

## Week 3
### Develop Product GO API for Burger Online Purchase Website
- [x] 1. Define all the api calls(names, and methods)
- [x] 2. Write Ping to test the connection.
- [x] 3. Develop add, get methods

## Week 4
### Finish GoAPI
- [x] 1. Finish Product GoAPI includes getProductById, getAllProduct, add or delete a product by ID.
- [x] 2. Configure VPC Peering structure

## Week 5
### Set up peering network
- [x] Bring up three vpc in three regions. Two for MongoDB(Cal, Oregon), one for GoAPI(Ohio).
- [x] Set up MongoDB cluster in peering network.
- [x] Deploy ProductAPI on Region Ohio.
- [x] Setup ELB for two instances of ProductAPI.
- [-] Upload product data

## Week 1
### Project Research
- [x] 1. Understanding the project requirement and architecture from various articles and books.
- [x] 2. Research on AWS lambda and Amazon API Gatway was done.
- [X] 3. Understood AKF scale cube, partition tolerance and the quorum consensus.

## Week 2
### Initial GoAPI and MongoDB Replicaset
- [x] 1. Research on MongoDB database CP properties.
- [x] 2. Research on how to use redis cache along with mongodb.
- [X] 3. Configuring a simple Go api to connect to mongoDB cluster and redisCache.
- [X] 4. Setup an AMI for EC2 having mongoDB installed.
- [X] 5. Configure a VPC and create 5 different EC2 instances in 2 separate AZs.
- [X] 6. Test the mongoDB replicaset by inserting data in primary node.

## Week 3
### Test GoAPI and move MongoDB replicaset to private subnets.
- [x] 1. Tested the GoAPI for various corner cases using postman.
- [x] 2. Researched on VPC peering.
- [x] 3. Used VPC peering to connect two VPCs which are in different regions.
- [x] 4. Created a new MongoDb replicaSet in private subnets and in different regions using VPC peering.
- [x] 5. Tested the new setup using GoAPI.

## Week 4
### Deploy the GoAPI and start working on fronted NodeJs app
- [x] 1. Deploy the GOAPI on ec2 instance using docker compose.
- [x] 2. Setup load balancer for GoAPIs.
- [x] 3. Start working on NodeJS frontend app.
- [x] 4. Discuss the frontend dependencies with other teammates.

## Week 5
### Start integrating the NodeJS app and start testing
- [x] 1. Integrate nodeJS app with other teammates and start basic sanity testing.
- [x] 2. Try various corner cases where app might fail.
- [x] 3. Fix those issues and get ready with final application. 

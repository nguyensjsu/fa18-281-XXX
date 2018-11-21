# What is MongoDB?  
MongoDB is a document database with the scalability and flexibility that you want with the querying and indexing that you need.  

## About:
* MongoDB stores data in flexible, JSON-like documents, meaning fields can vary from document to document and data structure can be changed over time.
* The document model maps to the objects in your application code, making data easy to work with.
* Ad hoc queries, indexing, and real time aggregation provide powerful ways to access and analyze your data.
* MongoDB is a distributed database at its core, so high availability, horizontal scaling, and geographic distribution are built in and easy to use.
* MongoDB is free and open-source.

## Makes development easy
MongoDB’s document model is simple for developers to learn and use, while still providing all the capabilities needed to meet the most complex requirements at any scale. 

## Features:
* High availability through built-in replication and failover
* Horizontal scalability with native sharding
* End-to-end security
* Native document validation and schema exploration with Compass
* Always available global support
* Management tooling for automation, monitoring, and backup
* Fully elastic database as a service with built-in best practices

## MongoDB on the AWS Cloud:
The MongoDB cluster (version 2.6 or 3.0) makes use of Amazon Elastic Compute Cloud (EC2) and Amazon Virtual Private Cloud, and is launched via a AWS CloudFormation template. You can use the template directly or you can copy and then customize it as needed.  The template creates the following resources:
* VPC with private and public subnets (you can also launch the cluster into an existing VPC).
* A NAT instance in the public subnet to support SSH access to the cluster and outbound Internet connectivity.
* An IAM instance role with fine-grained permissions.
* Security groups.
* A fully customized MongoDB cluster with replica sets, shards, and config servers, along with customized EBS storage, all running in the private subnet.
* The document examines scaling, replication, and performance tradeoffs in depth, and provides guidance to help you to choose appropriate types of EC2 instances and EBS volumes.


## Reference:
https://docs.mongodb.com
https://aws.amazon.com/blogs/aws/mongodb-on-the-aws-cloud-new-quick-start-reference-deployment/
